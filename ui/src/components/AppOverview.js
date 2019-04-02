import React from 'react';
import PropTypes from 'prop-types';
import { Grid, GridItem } from '@patternfly/react-core';
import './AppOverview.css';

const AppOverview = ({ className, app }) => {
  return (
    <div className={className}>
      <Grid>
        <GridItem span={3}>
          <div className="app-property">App ID</div>
        </GridItem>
        <GridItem span={6}>
          <div className="app-property">{app.appId}</div>
        </GridItem>
      </Grid>
    </div>
  );
};

AppOverview.propTypes = {
  app: PropTypes.object.isRequired
};

export default AppOverview;
