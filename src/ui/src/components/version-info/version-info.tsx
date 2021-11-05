/*
 * Copyright 2018- The Pixie Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import * as React from 'react';
import { Theme } from '@mui/material/styles';
import { createStyles, makeStyles } from '@mui/styles';

const useStyles = makeStyles(({ palette }: Theme) => createStyles({
  root: {
    position: 'absolute',
    top: 0,
    right: 0,
    color: palette.grey['800'],
    fontSize: '12px',
  },
}), { name: 'VersionInfo' });

interface VersionInfoProps {
  cloudVersion: string;
}

export const VersionInfo: React.FC<VersionInfoProps> = React.memo(function VersionInfo({ cloudVersion }) {
  const classes = useStyles();
  return <div className={classes.root}>{cloudVersion}</div>;
});
