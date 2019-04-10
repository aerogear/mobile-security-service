import * as actions from '../actions/types.js';
import DataService from '../DataService';
import fetchAction from './fetch';

export const appsTableSort = (_event, index, direction) => {
  return {
    type: actions.APPS_SORT,
    payload: {
      index: index,
      direction: direction
    }
  };
};

export const appDetailsSort = (_event, index, direction) => {
  return {
    type: actions.APP_VERSIONS_SORT,
    payload: {
      index: index,
      direction: direction
    }
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

export const toggleSaveAppModal = () => {
  return {
    type: actions.TOGGLE_SAVE_APP_MODAL
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

export const setAppDetailedDirtyState = (isDirty) => {
  return {
    type: actions.APP_DETAILED_IS_DIRTY,
    payload: {
      isDirty: isDirty
    }
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

export const setAllAppVersionsDisabled = (customMessage) => {
  return {
    type: actions.SET_ALL_APP_VERSIONS_DISABLED,
    payload: {
      customMessage
    }
  };
};

export const getApps = fetchAction([actions.APPS_REQUEST, actions.APPS_SUCCESS, actions.APPS_FAILURE], DataService.fetchApps);

export const getAppById = (appId) =>
  fetchAction([actions.APP_REQUEST, actions.APP_SUCCESS, actions.APP_FAILURE], async () => DataService.getAppById(appId))();

export const saveAppVersions = (id, versions) =>
  fetchAction([actions.SAVE_APP_VERSIONS,
    actions.SAVE_APP_VERSIONS_SUCCESS,
    actions.SAVE_APP_VERSIONS_FAILURE], async () => DataService.updateAppVersions(id, versions))();

export const disableAppVersions = (id, customMessage) =>
  fetchAction([actions.DISABLE_APP,
    actions.DISABLE_APP_SUCCESS,
    actions.DISABLE_APP_FAILURE], async () => DataService.disableAppVersions(id, customMessage))();
