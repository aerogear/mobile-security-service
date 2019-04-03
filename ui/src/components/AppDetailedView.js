import React from 'react';
import { withRouter } from 'react-router-dom';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import { Title } from '@patternfly/react-core';
import HeaderContainer from '../containers/common/HeaderContainer';
import AppOverview from './AppOverview';
import Content from './common/Content';
import AppDetailedToolbar from './AppDetailedToolbar';
import AppVersionsTableContainer from '../containers/AppVersionsTableContainer';
import DisableAppModalContainer from '../containers/DisableAppModalContainer';
import NavigationModalContainer from '../containers/NavigationModalContainer';
import SaveAppModalContainer from '../containers/SaveAppModalContainer';
import './AppDetailedView.css';
import { getAppById, toggleNavigationModal, toggleSaveAppModal, toggleDisableAppModal } from '../actions/actions-ui';

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

  onConfirmSaveApp () {
    this.props.toggleSaveAppModal();
    // TODO: Make a PUT request to API
    // to update the App versions
  }

  render () {
    return (
      <div className="app-detailed-view">
        <HeaderContainer />
        <AppDetailedToolbar app={this.props.app} onSaveAppClick={this.props.toggleSaveAppModal} onDisableAppClick={this.props.toggleDisableAppModal}/>
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
          <DisableAppModalContainer />
        </Content>
      </div>
    );
  }
}

AppDetailedView.propTypes = {
  app: PropTypes.shape({
    data: PropTypes.shape({
      id: PropTypes.string,
      appId: PropTypes.string,
      appName: PropTypes.string,
      deployedVersions: PropTypes.arrayOf(PropTypes.object)
    }),
    versionsRows: PropTypes.array
  }),
  isDirty: PropTypes.bool,
  getAppById: PropTypes.func.isRequired,
  toggleNavigationModal: PropTypes.func.isRequired,
  toggleSaveAppModal: PropTypes.func.isRequired,
  toggleDisableAppModal: PropTypes.func.isRequired
};

function mapStateToProps (state) {
  return {
    app: state.app.data,
    isDirty: state.isAppDetailedDirty
  };
};

const mapDispatchToProps = {
  getAppById,
  toggleNavigationModal,
  toggleSaveAppModal,
  toggleDisableAppModal
};

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(AppDetailedView));
