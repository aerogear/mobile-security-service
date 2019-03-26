import {
  APPS_FAILURE,
  APPS_SUCCESS,
  APPS_REQUEST,
  APPS_SORT,
  TOGGLE_HEADER_DROPDOWN,
  TOGGLE_NAVIGATION_MODAL,
  TOGGLE_APP_DETAILED_IS_DIRTY,
  APP_SUCCESS,
  APP_REQUEST,
  APP_FAILURE,
  APP_VERSIONS_SORT,
  UPDATE_DISABLED_APP,
  UPDATE_VERSION_CUSTOM_MESSAGE
} from '../actions/types.js';

import { SortByDirection, sortable, cellWidth } from '@patternfly/react-table';

const initialState = {
  apps: { rows: [], data: {} },
  sortBy: { direction: SortByDirection.asc, index: 0 },
  appVersionsSortDirection: { direction: SortByDirection.asc, index: 0 },
  columns: [
    { title: 'APP NAME', transforms: [sortable] },
    { title: 'APP ID', transforms: [sortable] },
    { title: 'DEPLOYED VERSIONS', transforms: [sortable] },
    { title: 'CURRENT INSTALLS', transforms: [sortable] },
    { title: 'LAUNCHES', transforms: [sortable] }
  ],
  appVersionsColumns: [
    { title: 'APP VERSION', transforms: [sortable, cellWidth(10)] },
    { title: 'CURRENT INSTALLS', transforms: [sortable, cellWidth(10)] },
    { title: 'LAUNCHES', transforms: [sortable, cellWidth(10)] },
    { title: 'LAST LAUNCHED', transforms: [sortable, cellWidth(15)] },
    { title: 'DISABLE ON STARTUP', transforms: [sortable, cellWidth(10)] },
    { title: 'CUSTOM DISABLE MESSAGE', transforms: [sortable, cellWidth('max')] }
  ],
  isAppsRequestFailed: false,
  currentUser: 'currentUser',
  isUserDropdownOpen: false,
  navigationModal: {
    isOpen: false,
    targetLocation: undefined
  },
  isAppDetailedDirty: false,
  app: {
    data: {},
    versionsRows: []
  },
  isAppRequestFailed: false
};

// returns a new array sorted in preferred direction
const sortRows = (rows, index, direction) => {
  // sort in ascending direction
  const sortedRows = [ ...rows ].sort((a, b) => (a[index] < b[index] ? -1 : a[index] > b[index] ? 1 : 0));

  if (areColumnValuesEqual(rows, index)) {
    return rows;
  }

  // reverse if descending direction is preferred
  if (direction !== SortByDirection.asc) {
    sortedRows.reverse();
  }

  return sortedRows;
};

/**
 * Check if all of the table column values are the same
 *
 * @param {Array} rows The table rows in which we will compare values
 * @param {*} index The index of the table column to compare values
 *
 * @returns {Boolean} Return whether every value is the same or not
 */
const areColumnValuesEqual = (rows, index) => {
  if (!rows || !index) {
    return false;
  }

  return rows.every((r, i) => {
    if (i === 0) {
      return true;
    }

    if (r[index] !== rows[i - 1][index]) {
      return false;
    }

    return true;
  });
};

export default (state = initialState, action) => {
  switch (action.type) {
    case APPS_SORT:
      const direction = action.payload.direction;
      const index = action.payload.index;
      const sortedRows = sortRows(state.apps.rows, index, direction);
      return {
        ...state,
        sortBy: {
          direction: direction,
          index: index
        },
        apps: {
          rows: sortedRows,
          data: state.apps.data
        }
      };
    case APP_VERSIONS_SORT:
      const versionDirection = action.payload.direction;
      const versionIndex = action.payload.index;
      const sortedAppVersions = sortRows(state.app.versionsRows, versionIndex, versionDirection);
      const newState = {
        ...state,
        appVersionsSortDirection: {
          direction: versionDirection,
          index: versionIndex
        }
      };

      newState.app.versionsRows = sortedAppVersions;

      return newState;

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
    case APPS_FAILURE:
      return {
        ...state,
        isAppsRequestFailed: true
      };
    case APP_REQUEST:
      return {
        ...state
      };
    case APP_SUCCESS:
      const fetchedVersions = [];
      action.result.deployedVersions.forEach((version) => {
        const temp = [];
        temp[0] = version['version'];
        temp[1] = version['numOfCurrentInstalls'] || 0;
        temp[2] = version['numOfAppLaunches'] || 0;
        temp[3] = version['lastLaunchedAt'] || 'Never Launched';
        temp[4] = version['disabled'];
        temp[5] = version['disabledMessage'] || '';

        fetchedVersions.push(temp);
      });

      return {
        ...state,
        app: {
          data: action.result,
          versionsRows: fetchedVersions
        }
      };
    case APP_FAILURE:
      return {
        ...state,
        isAppRequestFailed: true
      };
    case TOGGLE_HEADER_DROPDOWN:
      const isUserDropdownOpen = state.isUserDropdownOpen;
      return {
        ...state,
        isUserDropdownOpen: !isUserDropdownOpen
      };
    case TOGGLE_NAVIGATION_MODAL:
      const targetLocation = action.payload.targetLocation || undefined;
      return {
        ...state,
        navigationModal: {
          isOpen: action.payload.isOpen,
          targetLocation: targetLocation
        }
      };
    case TOGGLE_APP_DETAILED_IS_DIRTY:
      return {
        ...state,
        isAppDetailedDirty: !state.isAppDetailedDirty
      };
    case UPDATE_DISABLED_APP:
      const id = action.payload.id;
      const isDisabled = action.payload.isDisabled;
      const updatedVersions = state.app.versionsRows.map((version) => {
        if (version[0] === id) {
          version[4] = isDisabled;
        }
        return version;
      });
      return {
        ...state,
        app: {
          data: state.app.data,
          versionsRows: [ ...updatedVersions ]
        }
      };
    case UPDATE_VERSION_CUSTOM_MESSAGE:
      const versionId = action.payload.id;
      const value = action.payload.value;
      const updatedVersionsRows = state.app.versionsRows.map((version) => {
        if (version[0] === versionId) {
          version[5] = value;
        }
        return version;
      });
      return {
        ...state,
        app: {
          data: state.app.data,
          versionsRows: [ ...updatedVersionsRows ]
        }
      };
    default:
      return state;
  }
};
