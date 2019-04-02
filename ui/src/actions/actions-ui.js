import * as actions from '../actions/types.js';
import DataService from '../DataService';
import fetchAction from './fetch';

export const appsTableSort = (index, direction) => {
  return {
    type: actions.APPS_SORT,
    payload: {
      index: index,
      direction: direction
    }
  };
};

export const appDetailsSort = (index, direction) => {
  return {
    type: actions.APP_VERSIONS_SORT,
    payload: {
      index: index,
      direction: direction
    }
  };
};

export const toggleHeaderDropdown = () => {
  return {
    type: actions.TOGGLE_HEADER_DROPDOWN
  };
};

export const toggleNavigationModal = (isOpen, targetLocation) => {
  return {
    type: actions.TOGGLE_NAVIGATION_MODAL,
    payload: {
      isOpen: isOpen,
      targetLocation: targetLocation
    }
  };
};

export const toggleSaveAppModal = isSaveAppModalOpen => {
  return {
    type: actions.TOGGLE_SAVE_APP_MODAL,
    payload: { isSaveAppModalOpen }
  };
};

export const updateDisabledAppVersion = (id, isDisabled) => {
  return {
    type: actions.UPDATE_DISABLED_APP,
    payload: {
      id: id,
      isDisabled: isDisabled
    }
  };
};

export const toggleAppDetailedIsDirty = () => {
  return {
    type: actions.TOGGLE_APP_DETAILED_IS_DIRTY
  };
};

export const toggleDisableAppModal = () => {
  return {
    type: actions.TOGGLE_DISABLE_APP_MODAL
  };
};

export const updateVersionCustomMessage = (id, value) => {
  return {
    type: actions.UPDATE_VERSION_CUSTOM_MESSAGE,
    payload: {
      id: id,
      value: value
    }
  };
};

export const getApps = fetchAction([ actions.APPS_REQUEST, actions.APPS_SUCCESS, actions.APPS_FAILURE ], DataService.fetchApps);

export const getAppById = (appId) =>
  fetchAction([ actions.APP_REQUEST, actions.APP_SUCCESS, actions.APP_FAILURE ], async () => DataService.getAppById(appId))();
