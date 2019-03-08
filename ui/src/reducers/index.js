import { APPS_FAILURE, APPS_SUCCESS, APPS_REQUEST, REVERSE_SORT, TOGGLE_HEADER_DROPDOWN } from '../actions/types.js';

import { SortByDirection, sortable } from '@patternfly/react-table';

const columns = [
  { title: 'App Name', transforms: [ sortable ] },
  { title: 'App ID', transforms: [ sortable ] },
  { title: 'Deployed Versions', transforms: [ sortable ] },
  { title: 'Current Installs', transforms: [ sortable ] },
  { title: 'Launches', transforms: [ sortable ] }
];

const apps = { rows: [], data: {} };

const sortBy = { direction: SortByDirection.asc, index: 0 };

const initialState = {
  apps: apps,
  sortBy: sortBy,
  columns: columns,
  isAppsRequestFailed: false,
  currentUser: 'currentUser',
  isUserDropdownOpen: false
};

export default (state = initialState, action) => {
  switch (action.type) {
    case REVERSE_SORT:
      const reversedOrder = state.sortBy.direction === SortByDirection.asc ? SortByDirection.desc : SortByDirection.asc;
      const index = action.payload.index;
      const sortedRows = state.apps.rows.sort((a, b) => (a[index] < b[index] ? -1 : a[index] > b[index] ? 1 : 0));
      const sortedApps = reversedOrder === SortByDirection.asc ? sortedRows : sortedRows.reverse();
      return {
        ...state,
        sortBy: {
          direction: reversedOrder,
          index: index
        },
        apps: {
          rows: sortedApps
        }
      };
    case APPS_REQUEST:
      return {
        ...state
      };
    case APPS_SUCCESS:
      var fetchedApps = [];
      action.result.forEach((app) => {
        var temp = [];
        temp[0] = app.appName;
        temp[1] = app.appId;
        temp[2] = app.numOfDeployedVersions;
        temp[3] = app.numOfCurrentInstalls;
        temp[4] = app.numOfAppLaunches;
        fetchedApps.push(temp);
      });
      return {
        ...state,
        apps: {
          rows: fetchedApps,
          data: action.result
        }
      };
    case APPS_FAILURE:
      return {
        ...state,
        isAppsRequestFailed: true
      };
    case TOGGLE_HEADER_DROPDOWN:
      const isUserDropdownOpen = state.isUserDropdownOpen;
      return {
        ...state,
        isUserDropdownOpen: !isUserDropdownOpen
      };
    default:
      return state;
  }
};
