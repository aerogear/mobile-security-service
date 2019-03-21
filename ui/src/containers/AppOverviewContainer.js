import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Grid, GridItem } from '@patternfly/react-core';
import './AppOverviewContainer.css';

export class AppOverviewContainer extends React.Component {
  render () {
    const { className, app } = this.props;

    return (
      <div className={className}>
        <Grid>
          <GridItem span={3}>
            <span className="app-property">App ID</span>
          </GridItem>
          <GridItem span={6}>
            <span className="app-property">{app.appId}</span>
          </GridItem>
        </Grid>
      </div>
    );
  }
}

AppOverviewContainer.propTypes = {
  app: PropTypes.object.isRequired
};

function mapStateToProps (state) {
  return {
    app: state.app.data
  };
}

export default connect(mapStateToProps)(AppOverviewContainer);
