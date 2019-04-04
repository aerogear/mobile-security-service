import {
  APPS_SORT,
  APPS_REQUEST,
  APPS_SUCCESS,
  APPS_FAILURE
} from '../actions/types.js';

import { SortByDirection } from '@patternfly/react-table';

const initialState = {
  data: [],
  sortBy: { direction: SortByDirection.asc, index: 0 },
  isAppsRequestFailed: false
};

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

/**
 * Selector to convert app objects to array of app values
 *
 * @param {Array} apps - Array of app objects
 */
export const getAppsTableRows = (apps) => {
  return apps.map(app => {
    return [
      app.appName,
      app.appId,
      app.numOfDeployedVersions,
      app.numOfCurrentInstalls,
      app.numOfAppLaunches
    ];
  });
};

export default (state = initialState, action) => {
  switch (action.type) {
    case APPS_SORT:
      return {
        ...state,
        sortBy: {
          direction: action.payload.direction,
          index: action.payload.index
        }
      };
    case APPS_REQUEST:
      return {
        ...state
      };
    case APPS_SUCCESS:
      return {
        ...state,
        data: action.result
      };
    case APPS_FAILURE:
      return {
        ...state,
        isAppsRequestFailed: true
      };
    default:
      return state;
  }
};
