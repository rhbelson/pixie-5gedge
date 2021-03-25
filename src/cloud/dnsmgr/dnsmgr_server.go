package main

import (
	"net/http"
	_ "net/http/pprof"

	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"pixielabs.ai/pixielabs/src/cloud/dnsmgr/controller"
	"pixielabs.ai/pixielabs/src/cloud/dnsmgr/dnsmgrenv"
	dnsmgrpb "pixielabs.ai/pixielabs/src/cloud/dnsmgr/dnsmgrpb"
	"pixielabs.ai/pixielabs/src/cloud/dnsmgr/schema"
	"pixielabs.ai/pixielabs/src/cloud/shared/pgmigrate"
	"pixielabs.ai/pixielabs/src/shared/services"
	"pixielabs.ai/pixielabs/src/shared/services/healthz"
	"pixielabs.ai/pixielabs/src/shared/services/pg"
	"pixielabs.ai/pixielabs/src/shared/services/server"
)

func init() {
	pflag.String("dns_zone", "cluster-dev-withpixie-dev", "The zone to use for cloud DNS")
	pflag.String("dns_project", "pl-dev-infra", "The project to use for cloud DNS")
	pflag.String("domain_name", "withpixie.ai", "The domain name")
	pflag.Bool("use_default_dns_cert", false, "Whether to use the default DNS ssl cert")
}

func main() {
	services.SetupService("dnsmgr-service", 51900)
	services.PostFlagSetupAndParse()
	services.CheckServiceFlags()
	services.SetupServiceLogging()

	mux := http.NewServeMux()
	// This handles all the pprof endpoints.
	mux.Handle("/debug/", http.DefaultServeMux)
	healthz.RegisterDefaultChecks(mux)

	db := pg.MustConnectDefaultPostgresDB()
	err := pgmigrate.PerformMigrationsUsingBindata(db, "dnsmgr_service_migrations",
		bindata.Resource(schema.AssetNames(), schema.Asset))
	if err != nil {
		log.WithError(err).Fatal("Failed to apply migrations")
	}

	env := dnsmgrenv.New()
	dnsService, err := controller.NewCloudDNSService(
		viper.GetString("dns_zone"),
		viper.GetString("dns_project"),
		"/secrets/clouddns/dns_service_account.json",
	)

	if err != nil {
		log.WithError(err).Fatal("Failed to connect to Cloud DNS service")
	}

	svr := controller.NewServer(env, dnsService, db)

	s := server.NewPLServer(env, mux)
	dnsmgrpb.RegisterDNSMgrServiceServer(s.GRPCServer(), svr)
	s.Start()
	s.StopOnInterrupt()
}
