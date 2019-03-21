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
            <div className="app-property">App ID</div>
          </GridItem>
          <GridItem span={6}>
            <div className="app-property">{app.appId}</div>
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
