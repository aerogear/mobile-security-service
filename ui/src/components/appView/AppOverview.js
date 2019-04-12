import React from 'react';
import PropTypes from 'prop-types';
import { Grid, GridItem } from '@patternfly/react-core';
import './AppOverview.css';

/**
 * Stateless presentational component to display app level information.
 * @param {object} props Component props
 * @param {string} props.className CSS className for this component
 * @param {object} props.app Contains all app properties
 */
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
  className: PropTypes.string.isRequired,
  app: PropTypes.shape({
    appId: PropTypes.string
  }).isRequired
};

export default AppOverview;
