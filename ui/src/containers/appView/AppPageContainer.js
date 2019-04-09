import React from 'react';
import { withRouter } from 'react-router-dom';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import { Title } from '@patternfly/react-core';
import AppOverview from '../../components/appView/AppOverview';
import Content from '../../components/common/Content';
import AppToolbar from '../../components/appView/AppToolbar';
import HeaderContainer from '../HeaderContainer';
import AppVersionsTableContainer from './AppVersionsTableContainer';
import DisableAppModalContainer from './DisableAppModalContainer';
import NavigationModalContainer from './NavigationModalContainer';
import SaveAppModalContainer from './SaveAppModalContainer';
import { getAppById, toggleNavigationModal, toggleSaveAppModal, toggleDisableAppModal, setAppDetailedDirtyState, saveAppVersions } from '../../actions/actions-ui';
import AppService from '../../services/appService';

/**
 * Redux container component for the AppPage
 *
 * @class AppPageContainer
 * @extends {React.Component}
 */
export class AppPageContainer extends React.Component {
  componentWillMount () {
    this.props.getAppById(this.props.match.params.id);

    this.props.setAppDetailedDirtyState(this.isAppVersionsDirty());

    this.unblockHistory = this.props.history.block(targetLocation => {
      // If the view has a dirty state, display the popup
      if (this.props.isDirty) {
        this.props.toggleNavigationModal(true, targetLocation.pathname);
        return false;
      }
    });
  }

  componentDidUpdate () {
    this.props.setAppDetailedDirtyState(this.isAppVersionsDirty());
  }

  componentWillUnmount () {
    this.unblockHistory();
  }

  /**
   * Handles when user clicks Confirm in the Save App Model.
   * This function closes the model, then it saves the changed versions to the server.
   * When that is complete, it fetches fresh app data from the server.
   *
   * @memberof AppPageContainer
   */
  onConfirmSaveApp () {
    const { app: { id, deployedVersions: currentVersions }, savedData: { deployedVersions: savedVersions } } = this.props;

    this.props.toggleSaveAppModal();

    const dirtyVersions = AppService.getDirtyVersions(savedVersions, currentVersions);

    this.props.saveAppVersions(id, dirtyVersions);
  }

  /**
   * Checks if the app versions form is dirty.
   *
   * @returns {Boolean} is the form dirty
   * @memberof AppPageContainer
   */
  isAppVersionsDirty () {
    const { app: { deployedVersions: currentVersions },
      savedData: { deployedVersions: savedVersions } } = this.props;

    const dirtyItems = AppService.getDirtyVersions(savedVersions, currentVersions);

    return !!dirtyItems.length;
  };

  render () {
    return (
      <div className="app-detailed-view">
        <HeaderContainer />
        <AppToolbar app={this.props.app} onSaveAppClick={this.props.toggleSaveAppModal} onDisableAppClick={this.props.toggleDisableAppModal} isViewDirty={this.props.isDirty} />
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

AppPageContainer.propTypes = {
  app: PropTypes.shape({
    data: PropTypes.shape({
      id: PropTypes.string,
      appId: PropTypes.string,
      appName: PropTypes.string,
      deployedVersions: PropTypes.arrayOf(PropTypes.object)
    })
  }),
  isDirty: PropTypes.bool,
  getAppById: PropTypes.func.isRequired,
  toggleNavigationModal: PropTypes.func.isRequired,
  toggleSaveAppModal: PropTypes.func.isRequired,
  toggleDisableAppModal: PropTypes.func.isRequired,
  saveAppVersions: PropTypes.func.isRequired
};

const mapStateToProps = (state) => {
  return {
    app: state.app.data,
    savedData: state.app.savedData,
    isDirty: state.app.isDirty
  };
};

const mapDispatchToProps = {
  getAppById,
  toggleNavigationModal,
  toggleSaveAppModal,
  toggleDisableAppModal,
  setAppDetailedDirtyState,
  saveAppVersions
};

export default withRouter(connect(mapStateToProps, mapDispatchToProps)(AppPageContainer));
