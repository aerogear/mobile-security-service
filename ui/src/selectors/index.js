import { createSelector } from 'reselect';
import { SortByDirection } from '@patternfly/react-table';

/**
 * Sorts the rows for a UI table by column index and direction
 * @param {Array} rows Array of arrays for Patternfly table data
 * @param {Number} index Table column to sort by
 * @param {String} direction Either ASC or DESC
 */
export const getSortedTableRows = (rows, index, direction) => {
  // sort in ascending direction
  const sortedRows = [ ...rows ].sort((a, b) => (a[index] < b[index] ? -1 : a[index] > b[index] ? 1 : 0));

  if (areColumnValuesEqual(rows, index)) {
    return rows;
  }

  // reverse if descending direction is selected
  if (direction !== SortByDirection.asc) {
    sortedRows.reverse();
  }

  return sortedRows;
};

/**
 * Check if all of the table column values are the same
 *
 * @param {Array} rows The table rows in which we will compare values
 * @param {Number} index The index of the table column to compare values
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

/**
 * Gets the sort index from a sortBy object
 * @param {Object} _state Redux state
 * @param {Object} sortBy sortBy includes index and direction
 */
export const getSortIndex = (_state, sortBy) => sortBy.index;

/**
 * Gets the sort direction from a sortBy object
 * @param {Object} _state Redux state
 * @param {Object} sortBy sortBy includes index and direction
 */
export const getSortDirection = (_state, sortBy) => sortBy.direction;

/**
 * Selector to convert app objects to array of app values
 * @param {Object} state Redux state
 * @param {Object} _sortBy sortBy includes index and direction
 */
export const getAppsTableRows = (state, _sortBy) => {
  return state.apps.data.map(app => {
    return [
      app.appName,
      app.appId,
      app.numOfDeployedVersions,
      app.numOfCurrentInstalls,
      app.numOfAppLaunches
    ];
  });
};

/**
 * Selector to convert app version objects to array of app version values
 * @param {Object} state Redux state
 * @param {Object} _sortBy sortBy includes index and direction
 */
export const getAppVersionTableRows = (state, _sortBy) => {
  return state.app.data.deployedVersions.map(version => {
    return [
      version.version,
      version.numOfCurrentInstalls || 0,
      version.numOfAppLaunches || 0,
      version.lastLaunchedAt || 'Never Launched',
      version.disabled,
      version.disabledMessage || '',
      version.id
    ];
  });
};

/**
 * Memoised selector to retrieve the sorted rows for the apps table
 */
export const getSortedAppsTableRows = createSelector(
  [ getAppsTableRows, getSortIndex, getSortDirection ],
  (appsRows, sortIndex, sortDirection) => {
    return getSortedTableRows(appsRows, sortIndex, sortDirection);
  }
);

/**
 * Memoised selector to retrieve the sorted rows for the app versions table
 */
export const getSortedAppVersionTableRows = createSelector(
  [ getAppVersionTableRows, getSortIndex, getSortDirection ],
  (appVersionRows, sortIndex, sortDirection) => {
    return getSortedTableRows(appVersionRows, sortIndex, sortDirection);
  }
);
