import {
  APP_VERSIONS_SORT,
  APP_REQUEST,
  APP_SUCCESS,
  APP_FAILURE,
  TOGGLE_APP_DETAILED_IS_DIRTY,
  UPDATE_DISABLED_APP,
  UPDATE_VERSION_CUSTOM_MESSAGE
} from '../actions/types.js';

import { SortByDirection } from '@patternfly/react-table';

const initialState = {
  savedData: {
    deployedVersions: []
  },
  data: {
    deployedVersions: []
  },
  sortBy: { direction: SortByDirection.asc, index: 0 },
  isAppRequestFailed: false,
  isDirty: false
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

export const getAppVersionTableRows = (versions) => {
  return versions.map(version => {
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

export default (state = initialState, action) => {
  switch (action.type) {
    case APP_VERSIONS_SORT:
      return {
        ...state,
        sortBy: {
          direction: action.payload.direction,
          index: action.payload.index
        }
      };
    case APP_REQUEST:
      return {
        ...state
      };
    case APP_SUCCESS:
      return {
        ...state,
        data: action.result,
        savedData: action.result
      };
    case APP_FAILURE:
      return {
        ...state,
        isAppRequestFailed: true
      };
    case TOGGLE_APP_DETAILED_IS_DIRTY:
      return {
        ...state,
        isDirty: !state.isDirty
      };
    case UPDATE_DISABLED_APP:
      const updatedVersions = state.data.deployedVersions.map((version) => {
        if (version.id === action.payload.id) {
          version.disabled = action.payload.isDisabled;
        }
        return version;
      });

      return {
        ...state,
        data: {
          ...state.data,
          deployedVersions: updatedVersions
        }
      };
    case UPDATE_VERSION_CUSTOM_MESSAGE:
      const updatedVersions2 = state.data.deployedVersions.map((version) => {
        if (version.id === action.payload.id) {
          version.disabledMessage = action.payload.value;
        }
        return version;
      });

      return {
        ...state,
        data: {
          ...state.data,
          deployedVersions: updatedVersions2
        }
      };
    default:
      return state;
  }
};
