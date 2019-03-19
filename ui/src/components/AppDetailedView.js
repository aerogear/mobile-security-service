import React from 'react';
import { Title, Grid, GridItem } from '@patternfly/react-core';
import Header from './common/Header';
import AppVersionsTableContainer from '../containers/AppVersionsTableContainer';
import AppOverviewContainer from '../containers/AppOverviewContainer';
import './AppDetailedView.css';
import { getAppById } from '../actions/actions-ui';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';

class AppDetailedView extends React.Component {
  componentWillMount () {
    this.props.getAppById(this.props.match.params.id);
  }

  render () {
    return (
      <div className="app-detailed-view">
        <Header />
        <Grid gutter="md" className="container">
          <GridItem span={1} />
          <GridItem span={10}>
            <AppOverviewContainer app={this.props.app} className='app-overview-container'/>
            <Title className="table-title" size="2xl">
              Deployed Versions
            </Title>
            <AppVersionsTableContainer className='table-scroll-x'/>
          </GridItem>
          <GridItem span={1} />
        </Grid>
      </div>
    );
  }
}

AppDetailedView.propTypes = {
  app: PropTypes.object.isRequired,
  getAppById: PropTypes.func.isRequired
};

function mapStateToProps (state) {
  return {
    app: state.app.data,
    getAppById: PropTypes.func.isRequired
  };
};

const mapDispatchToProps = {
  getAppById
};

export default connect(mapStateToProps, mapDispatchToProps)(AppDetailedView);
