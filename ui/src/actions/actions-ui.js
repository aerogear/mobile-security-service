import { APP_REQUEST, APP_SUCCESS, APP_FAILURE, APPS_REQUEST, APPS_SUCCESS, APPS_FAILURE, REVERSE_SORT, TOGGLE_HEADER_DROPDOWN } from '../actions/types.js';
import DataService from '../DataService';
import fetchAction from './fetch';

export const reverseAppsTableSort = (index) => {
  return {
    type: REVERSE_SORT,
    payload: {
      index: index
    }
  };
};

export const toggleHeaderDropdown = () => {
  return {
    type: TOGGLE_HEADER_DROPDOWN
  };
};

export const getApps = fetchAction([APPS_REQUEST, APPS_SUCCESS, APPS_FAILURE], DataService.fetchApps);

export const getAppById = appId => fetchAction([APP_REQUEST, APP_SUCCESS, APP_FAILURE], async () => DataService.getAppById(appId))();
