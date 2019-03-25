import {
  APP_REQUEST,
  APP_SUCCESS,
  APP_FAILURE,
  APPS_REQUEST,
  APPS_SUCCESS,
  APPS_FAILURE,
  APPS_SORT,
  TOGGLE_HEADER_DROPDOWN,
  TOGGLE_NAVIGATION_MODAL,
  TOGGLE_SAVE_APP_MODAL,
  TOGGLE_APP_DETAILED_IS_DIRTY,
  APP_VERSIONS_SORT,
  UPDATE_DISABLED_APP,
  UPDATE_VERSION_CUSTOM_MESSAGE
} from '../actions/types.js';
import DataService from '../DataService';
import fetchAction from './fetch';

export const appsTableSort = (index, direction) => {
  return {
    type: APPS_SORT,
    payload: {
      index: index,
      direction: direction
    }
  };
};

export const appDetailsSort = (index, direction) => {
  return {
    type: APP_VERSIONS_SORT,
    payload: {
      index: index,
      direction: direction
    }
  };
};

export const toggleHeaderDropdown = () => {
  return {
    type: TOGGLE_HEADER_DROPDOWN
  };
};

export const toggleNavigationModal = (isOpen, targetLocation) => {
  return {
    type: TOGGLE_NAVIGATION_MODAL,
    payload: {
      isOpen: isOpen,
      targetLocation: targetLocation
    }
  };
};

export const toggleSaveAppModal = isSaveAppModalOpen => {
  return {
    type: TOGGLE_SAVE_APP_MODAL,
    payload: { isSaveAppModalOpen }
  };
};

export const updateDisabledAppVersion = (id, isDisabled) => {
  return {
    type: UPDATE_DISABLED_APP,
    payload: {
      id: id,
      isDisabled: isDisabled
    }
  };
};

export const toggleAppDetailedIsDirty = () => {
  return {
    type: TOGGLE_APP_DETAILED_IS_DIRTY
  };
};

export const updateVersionCustomMessage = (id, value) => {
  return {
    type: UPDATE_VERSION_CUSTOM_MESSAGE,
    payload: {
      id: id,
      value: value
    }
  };
};

export const getApps = fetchAction([ APPS_REQUEST, APPS_SUCCESS, APPS_FAILURE ], DataService.fetchApps);

export const getAppById = (appId) =>
  fetchAction([ APP_REQUEST, APP_SUCCESS, APP_FAILURE ], async () => DataService.getAppById(appId))();
