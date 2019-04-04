import { createSelector } from 'reselect';
import { SortByDirection } from '@patternfly/react-table';

// returns a new array sorted in preferred direction
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

const getSortIndex = (state, sortBy) => sortBy.index;
const getSortDirection = (state, sortBy) => sortBy.direction;

/**
 * Selector to convert app objects to array of app values
 *
 * @param {Array} apps - Array of app objects
 */
export const getAppsTableRows = (state) => {
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

export const getAppVersionTableRows = (state) => {
  return state.app.data.deployedVersions.map(version => {
    return [
      version['version'],
      version['numOfCurrentInstalls'] || 0,
      version['numOfAppLaunches'] || 0,
      version['lastLaunchedAt'] || 'Never Launched',
      version['disabled'],
      version['disabledMessage'] || '',
      version['id']
    ];
  });
};

export const getSortedAppsTableRows = createSelector(
  [ getAppsTableRows, getSortIndex, getSortDirection ],
  (appsRows, sortIndex, sortDirection) => {
    return getSortedTableRows(appsRows, sortIndex, sortDirection);
  }
);

export const getSortedAppVersionTableRows = createSelector(
  [ getAppVersionTableRows, getSortIndex, getSortDirection ],
  (appVersionRows, sortIndex, sortDirection) => {
    return getSortedTableRows(appVersionRows, sortIndex, sortDirection);
  }
);
