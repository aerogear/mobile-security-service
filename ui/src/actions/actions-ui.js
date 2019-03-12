import { APPS_REQUEST, APPS_SUCCESS, APPS_FAILURE, APPS_SORT, TOGGLE_HEADER_DROPDOWN } from '../actions/types.js';
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

export const toggleHeaderDropdown = () => {
  return {
    type: TOGGLE_HEADER_DROPDOWN
  };
};

export const getApps = fetchAction([ APPS_REQUEST, APPS_SUCCESS, APPS_FAILURE ], DataService.fetchApps);
