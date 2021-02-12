import { ClusterContext } from 'common/cluster-context';
import UserContext from 'common/user-context';
import * as storage from 'common/storage';
import { ClusterStatus, VizierGRPCClientProvider, CLUSTER_STATUS_UNKNOWN } from 'common/vizier-grpc-client-context';
import { useSnackbar } from 'pixie-components';
import AdminView from 'pages/admin/admin';
import CreditsView from 'pages/credits/credits';
import { SCRATCH_SCRIPT, ScriptsContextProvider } from 'containers/App/scripts-context';
import LiveView from 'pages/live/live';
import * as React from 'react';
import { Redirect, Route, Switch } from 'react-router-dom';
import { useParams, useLocation } from 'react-router';
import * as QueryString from 'query-string';

import { useQuery } from '@apollo/client';
import { createStyles, makeStyles } from '@material-ui/core/styles';
import { useLDClient } from 'launchdarkly-react-client-sdk';
import { CLUSTER_QUERIES, USER_QUERIES } from 'pixie-api';
import { scriptToEntityURL } from 'containers/live-widgets/utils/live-view-params';
import { LIVE_VIEW_SCRIPT_ID_KEY, useSessionStorage } from 'common/storage';
import { DeployInstructions } from './deploy-instructions';
import { selectCluster } from './cluster-info';

const useStyles = makeStyles(() => createStyles({
  banner: {
    position: 'absolute',
    width: '100%',
    textAlign: 'center',
    top: 0,
    zIndex: 1500, // TopBar has a z-index of 1300.
    color: 'white',
    background: 'rgba(220,0,0,0.5)',
  },
}));

const ClusterBanner = () => {
  const classes = useStyles();
  const { loading, error, data } = useQuery(USER_QUERIES.GET_USER_INFO, { fetchPolicy: 'network-only' });

  if (loading || error) {
    return null;
  }

  if (data.user.email.split('@')[1] === 'pixie.support') {
    return (
      <div className={classes.banner}>
        {
          `You are viewing clusters for an external org: ${data.user.orgName}`
        }
      </div>
    );
  }

  return null;
};

const useSelectedCluster = () => {
  const { selectedCluster } = React.useContext(ClusterContext);
  const { loading, error, data } = useQuery(
    CLUSTER_QUERIES.LIST_CLUSTERS,
    { pollInterval: 2500, fetchPolicy: 'network-only' },
  );
  const clusters = data?.clusters || [];
  const cluster = clusters.find((c) => c.id === selectedCluster);
  return {
    loading, cluster, numClusters: clusters.length, error,
  };
};

// Convenience routes: sends `/scratch`, `/script/http_data`, and others to the appropriate Live url.
// eslint-disable-next-line react/require-default-props
const ScriptShortcut = ({ toNamespace = '', toScript = '' }: { toNamespace?: string; toScript?: string }) => {
  const { namespace, scriptId } = useParams();
  let fullId = `${toNamespace || namespace || 'px'}/`;
  const normalized = (toScript || scriptId || '').trim().toLowerCase().replace(/[^a-z0-9]+/g, '-');
  if (['scratch', 'scratchpad', 'scratch-pad', 'scratch-script'].includes(normalized)) {
    fullId = SCRATCH_SCRIPT.id;
  } else {
    fullId += toScript || scriptId;
  }

  const { loading, cluster, error } = useSelectedCluster();

  const queryParams = QueryString.parse(useLocation().search);

  // The default /live route points to <Vizier /> where the actual error handling is.
  const destination = error || !cluster
    ? '/live'
    : scriptToEntityURL(fullId, cluster?.clusterName, queryParams);

  // ScriptContext reads session storage to find the currently selected script, and then redirects to it. This happens
  // AFTER the page finishes loading and rendering routes, which is a bug (PC-537) with no currently clear solution.
  // The workaround here, for now, is to simply overwrite that session storage before ScriptContext reads it.
  const [id, setId] = useSessionStorage(LIVE_VIEW_SCRIPT_ID_KEY, fullId);
  if (!error && !loading && cluster && fullId !== id) {
    setId(fullId);
  }

  return loading ? null : <Redirect to={destination} />;
};

// Selects a default cluster if one hasn't already been selected by the user.

const Vizier = () => {
  const showSnackbar = useSnackbar();

  const {
    loading, error, cluster, numClusters,
  } = useSelectedCluster();

  if (loading) { return <div>Loading...</div>; }

  const errMsg = error?.message;
  if (errMsg) {
    // This is an error with pixie cloud, it is probably not relevant to the user.
    // Show a generic error message instead.
    showSnackbar({ message: 'There was a problem connecting to Pixie', autoHideDuration: 5000 });
    // eslint-disable-next-line no-console
    console.error(errMsg);
  }

  if (numClusters === 0) {
    return <DeployInstructions />;
  }

  const status: ClusterStatus = cluster?.status || CLUSTER_STATUS_UNKNOWN;

  return (
    <VizierGRPCClientProvider
      clusterID={cluster.id}
      passthroughEnabled={cluster.vizierConfig.passthroughEnabled}
      clusterStatus={errMsg ? CLUSTER_STATUS_UNKNOWN : status}
    >
      <ScriptsContextProvider>
        <LiveView />
      </ScriptsContextProvider>
    </VizierGRPCClientProvider>
  );
};

export default function WithClusterBanner() {
  const showSnackbar = useSnackbar();

  const [clusterId, setClusterId] = storage.useSessionStorage(storage.CLUSTER_ID_KEY, '');
  const { loading, error, data } = useQuery(
    CLUSTER_QUERIES.LIST_CLUSTERS,
    { pollInterval: 2500, fetchPolicy: 'network-only' },
  );
  const userQuery = useQuery(USER_QUERIES.GET_USER_INFO, { fetchPolicy: 'network-only' });
  const ldClient = useLDClient();

  const clusters = data?.clusters || [];
  const cluster = (clusterId && clusters.find((c) => c.id === clusterId)) || selectCluster(clusters);

  const context = React.useMemo(() => ({
    selectedCluster: clusterId,
    selectedClusterName: cluster?.clusterName,
    selectedClusterPrettyName: cluster?.prettyClusterName,
    selectedClusterUID: cluster?.clusterUID,
    setCluster: setClusterId,
    setClusterByName: (name: string) => {
      const foundId = name && clusters.find((c) => c.clusterName === name)?.id;
      const newClusterId = foundId || clusterId;
      setClusterId(newClusterId);
    },
  }), [clusterId, cluster?.clusterName, cluster?.prettyClusterName, cluster?.clusterUID, setClusterId, clusters]);

  const userEmail = userQuery.data?.user.email;
  const userOrg = userQuery.data?.user.orgName;

  const userContext = React.useMemo(() => ({
    user: {
      email: userEmail,
      orgName: userOrg,
    },
  }), [userEmail, userOrg]);

  React.useEffect(() => {
    if (ldClient != null && userEmail != null) {
      ldClient.identify({
        key: userEmail,
        email: userEmail,
        custom: {
          orgName: userOrg,
        },
      });
    }
  }, [ldClient, userEmail, userOrg]);

  if (loading || userQuery.loading) { return <div>Loading...</div>; }

  const errMsg = error?.message || userQuery.error?.message;
  if (errMsg) {
    // This is an error with pixie cloud, it is probably not relevant to the user.
    // Show a generic error message instead.
    showSnackbar({ message: 'There was a problem connecting to Pixie', autoHideDuration: 5000 });
    // eslint-disable-next-line no-console
    console.error(errMsg);
  }

  if (clusters.length > 0) {
    if (cluster?.id && clusterId !== cluster?.id) {
      setClusterId(cluster.id);
    }
  }

  return (
    <ClusterContext.Provider value={context}>
      <UserContext.Provider value={userContext}>
        <ClusterBanner />
        <Switch>
          <Route path='/admin' component={AdminView} />
          <Route path='/credits' component={CreditsView} />
          <Route path='/live' component={Vizier} />
          <Route
            path={[
              '/script/:namespace/:scriptId',
              '/scripts/:namespace/:scriptId',
              '/s/:namespace/:scriptId',
              '/script/:scriptId',
              '/scripts/:scriptId',
              '/s/:scriptId',
            ]}
            component={ScriptShortcut}
          />
          <Route path={['/scratch', '/scratchpad']} render={() => <ScriptShortcut toScript={SCRATCH_SCRIPT.id} />} />
          <Redirect from='/*' to='/live' />
        </Switch>
      </UserContext.Provider>
    </ClusterContext.Provider>
  );
}
