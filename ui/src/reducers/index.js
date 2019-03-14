import {
  APPS_FAILURE,
  APPS_SUCCESS,
  APP_DETAILS_SUCCESS,
  APPS_REQUEST,
  APPS_SORT,
  TOGGLE_HEADER_DROPDOWN,
  APP_DETAILS_SORT
} from '../actions/types.js';

import { SortByDirection, sortable } from '@patternfly/react-table';

const initialState = {
  apps: { rows: [], data: {} },
  sortBy: { direction: SortByDirection.asc, index: 0 },
  appDetailsSortDirection: { direction: SortByDirection.asc, index: 0 },
  columns: [
    { title: 'App Name', transforms: [ sortable ] },
    { title: 'App ID', transforms: [ sortable ] },
    { title: 'Deployed Versions', transforms: [ sortable ] },
    { title: 'Current Installs', transforms: [ sortable ] },
    { title: 'Launches', transforms: [ sortable ] }
  ],
  appDetailRows: [],
  appDetailColumns: [
    { title: 'App Version', transforms: [ sortable ] },
    { title: 'Current Installs', transforms: [ sortable ] },
    { title: 'Launches', transforms: [ sortable ] },
    { title: 'Last Launched', transforms: [ sortable ] },
    { title: 'Disable on Startup', transforms: [ sortable ] },
    { title: 'Custom Disable Message', transforms: [ sortable ] }
  ],
  isAppsRequestFailed: false,
  currentUser: 'currentUser',
  isUserDropdownOpen: false
};

// returns a new array sorted in preferred order
const sortRows = (rows, index, order) => {
  // sort in ascending order
  const sortedRows = rows.sort((a, b) => (a[index] < b[index] ? -1 : a[index] > b[index] ? 1 : 0));

  // reverse if descending order is preferred
  if (order !== SortByDirection.asc) {
    sortedRows.reverse();
  }

  return sortedRows;
};

export default (state = initialState, action) => {
  switch (action.type) {
    case APPS_SORT:
      const order = action.payload.direction;
      const index = action.payload.index;
      const sortedRows = sortRows(state.apps.rows, index, order);
      return {
        ...state,
        sortBy: {
          direction: action.payload.direction,
          index: index
        },
        apps: {
          rows: sortedRows,
          data: state.apps.data
        }
      };
    case APP_DETAILS_SORT:
      const newOrder = action.payload.direction;
      const colIndex = action.payload.index;
      const sortedAppDetails = sortRows(state.appDetailRows, colIndex, newOrder);
      return {
        ...state,
        appDetailsSortDirection: {
          direction: newOrder,
          index: colIndex
        },
        appDetailRows: sortedAppDetails
      };
    case APPS_REQUEST:
      return {
        ...state
      };
    case APPS_SUCCESS:
      const fetchedApps = [];
      action.result.forEach((app) => {
        const temp = [];
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
    case APP_DETAILS_SUCCESS:
      const fetchedAppDetails = [];
      action.result.forEach((appDetail) => {
        fetchedAppDetails.push(appDetail);
      });
      return {
        ...state,
        appDetailRows: fetchedAppDetails
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
