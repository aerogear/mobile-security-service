import React from 'react';
import { withRouter } from 'react-router-dom';
import NavigationModalContainer from '../containers/NavigationModalContainer';
import SaveAppModalContainer from '../containers/SaveAppModalContainer';
import { Title } from '@patternfly/react-core';
import Header from './common/Header';
import AppVersionsTableContainer from '../containers/AppVersionsTableContainer';
import AppOverview from './AppOverview';
import './AppDetailedView.css';
import { getAppById, toggleNavigationModal, toggleSaveAppModal } from '../actions/actions-ui';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import Content from './common/Content';
import AppDetailedToolbar from './AppDetailedToolbar';

class AppDetailedView extends React.Component {
  componentWillMount () {
    this.props.getAppById(this.props.match.params.id);

    this.unblockHistory = this.props.history.block(targetLocation => {
      // If the view has a dirty state, display the popup
      if (this.props.isDirty) {
        this.props.toggleNavigationModal(true, targetLocation.pathname);
        return false;
      }
    });
  }

  componentWillUnmount () {
    this.unblockHistory();
  }

  onSaveApp = () => {
    this.props.toggleSaveAppModal(true);
  }

  onDisableApp () {
    // TODO: Show modal
  }

  onConfirmSaveApp () {
    this.props.toggleSaveAppModal(false);
    // TODO: Make a PUT request to API
    // to update the App versions
  }

  render () {
    return (
      <div className="app-detailed-view">
        <Header />
        <AppDetailedToolbar app={this.props.app} onSaveApp={() => this.onSaveApp()} onDisableApp={() => this.onDisableApp()}/>
        <Content className="container">
          <AppOverview app={this.props.app} className='app-overview' />
          <Title className="table-title" size="2xl">
            Deployed Versions
          </Title>
          <AppVersionsTableContainer className='table-scroll-x' />
          <NavigationModalContainer title="Are you sure you want to leave this page?" unblockHistory={this.unblockHistory}>
          You still have unsaved changes.
          </NavigationModalContainer>
          <SaveAppModalContainer title="Save Changes" onConfirm={() => this.onConfirmSaveApp()}>
          Are you sure you want to save your changes to this app?
          </SaveAppModalContainer>
        </Content>
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
  toggleNavigationModal,
  toggleSaveAppModal
};

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(AppDetailedView));
