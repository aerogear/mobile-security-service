import { APPS_REQUEST, APPS_SUCCESS, APPS_FAILURE, REVERSE_SORT } from '../actions/types.js';
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

export const getApps = fetchAction([APPS_REQUEST, APPS_SUCCESS, APPS_FAILURE], DataService.fetchApps);
