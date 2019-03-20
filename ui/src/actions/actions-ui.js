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
  TOGGLE_APP_DETAILED_IS_DIRTY,
  APP_VERSIONS_SORT
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

export const toggleNavigationModal = (isNavigationModalOpen) => {
  return {
    type: TOGGLE_NAVIGATION_MODAL,
    payload: {
      isNavigationModalOpen: isNavigationModalOpen
    }
  };
};

export const toggleAppDetailedIsDirty = () => {
  return {
    type: TOGGLE_APP_DETAILED_IS_DIRTY
  };
};

export const getApps = fetchAction([APPS_REQUEST, APPS_SUCCESS, APPS_FAILURE], DataService.fetchApps);

export const getAppById = appId => fetchAction([APP_REQUEST, APP_SUCCESS, APP_FAILURE], async () => DataService.getAppById(appId))();
