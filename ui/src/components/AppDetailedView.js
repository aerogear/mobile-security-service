import React from 'react';
import { withRouter } from 'react-router-dom';
import NavigationModalContainer from '../containers/NavigationModalContainer';
import { Title, Grid, GridItem } from '@patternfly/react-core';
import Header from './common/Header';
import AppVersionsTableContainer from '../containers/AppVersionsTableContainer';
import AppOverviewContainer from '../containers/AppOverviewContainer';
import './AppDetailedView.css';
import { getAppById, toggleNavigationModal } from '../actions/actions-ui';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';

class AppDetailedView extends React.Component {
  componentWillMount () {
    this.props.getAppById(this.props.match.params.id);
  }

  componentDidMount () {
    this.unblock = this.props.history.block(targetLocation => {
      // If the view has a dirty state, display the popup
      if (this.props.isDirty) {
        this.props.toggleNavigationModal(true);
        return false;
      }
    });
  }

  componentWillUnmount () {
    this.unblock();
  }

  render () {
    return (
      <div className="app-detailed-view">
        <Header />
        <Grid gutter="md" className="container">
          <GridItem span={1} />
          <GridItem span={10}>
            <AppOverviewContainer app={this.props.app} className='app-overview-container' />
            <Title className="table-title" size="2xl">
              Deployed Versions
            </Title>
            <AppVersionsTableContainer className='table-scroll-x' />
            <NavigationModalContainer text="You still have unsaved changes." title="Are you sure you want to leave this page?" />
          </GridItem>
          <GridItem span={1} />
        </Grid>
      </div>
    );
  }
}

AppDetailedView.propTypes = {
  app: PropTypes.object.isRequired,
  getAppById: PropTypes.func.isRequired,
  isDirty: PropTypes.bool
};

function mapStateToProps (state) {
  return {
    app: state.app.data,
    getAppById: PropTypes.func.isRequired,
    isDirty: state.isAppDetailedDirty
  };
};

const mapDispatchToProps = {
  getAppById,
  toggleNavigationModal
};

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(AppDetailedView));
