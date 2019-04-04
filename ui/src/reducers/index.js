import * as actions from '../actions/types.js';

import { SortByDirection } from '@patternfly/react-table';

const initialState = {
  header: {
    currentUser: 'currentUser'
  },
  apps: {
    data: [],
    sortBy: { direction: SortByDirection.asc, index: 0 },
    isAppsRequestFailed: false
  },
  app: {
    savedData: {
      deployedVersions: []
    },
    data: {
      deployedVersions: []
    },
    sortBy: { direction: SortByDirection.asc, index: 0 },
    isAppRequestFailed: false,
    isDirty: false
  },
  modals: {
    disableApp: {
      isOpen: false
    },
    saveApp: {
      isOpen: false
    },
    navigationModal: {
      isOpen: false,
      targetLocation: undefined
    }
  }
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
    case actions.APPS_SORT:
      return {
        ...state,
        apps: {
          ...state.apps,
          sortBy: {
            direction: action.payload.direction,
            index: action.payload.index
          }
        }
      };
    case actions.APP_VERSIONS_SORT:
      return {
        ...state,
        app: {
          ...state.app,
          sortBy: {
            direction: action.payload.direction,
            index: action.payload.index
          }
        }
      };
    case actions.APPS_REQUEST:
      return {
        ...state
      };
    case actions.APPS_SUCCESS:
      return {
        ...state,
        apps: {
          ...state.apps,
          data: action.result
        }
      };
    case actions.APPS_FAILURE:
      return {
        ...state,
        apps: {
          ...state.apps,
          isAppsRequestFailed: true
        }
      };
    case actions.APP_REQUEST:
      return {
        ...state
      };
    case actions.APP_SUCCESS:
      return {
        ...state,
        app: {
          ...state.app,
          data: action.result,
          savedData: action.result
        }
      };
    case actions.APP_FAILURE:
      return {
        ...state,
        app: {
          ...state.app,
          isAppRequestFailed: true
        }
      };
    case actions.TOGGLE_NAVIGATION_MODAL:
      const targetLocation = action.payload.targetLocation || undefined;
      return {
        ...state,
        modals: {
          ...state.modals,
          navigationModal: {
            isOpen: action.payload.isOpen,
            targetLocation: targetLocation
          }
        }
      };
    case actions.TOGGLE_SAVE_APP_MODAL:
      return {
        ...state,
        modals: {
          ...state.modals,
          saveApp: {
            isOpen: !state.modals.saveApp.isOpen
          }
        }
      };
    case actions.TOGGLE_DISABLE_APP_MODAL:
      return {
        ...state,
        modals: {
          ...state.modals,
          disableApp: {
            isOpen: !state.modals.disableApp.isOpen
          }
        }
      };
    case actions.TOGGLE_APP_DETAILED_IS_DIRTY:
      return {
        ...state,
        app: {
          ...state.app,
          isDirty: !state.app.isDirty
        }
      };
    case actions.UPDATE_DISABLED_APP:
      const updatedVersions = state.app.data.deployedVersions.map((version) => {
        if (version.id === action.payload.id) {
          version.disabled = action.payload.isDisabled;
        }
        return version;
      });

      return {
        ...state,
        app: {
          ...state.app,
          data: {
            ...state.app.data,
            deployedVersions: updatedVersions
          }
        }
      };
    case actions.UPDATE_VERSION_CUSTOM_MESSAGE:
      const updatedVersions2 = state.app.data.deployedVersions.map((version) => {
        if (version.id === action.payload.id) {
          version.disabledMessage = action.payload.value;
        }
        return version;
      });

      return {
        ...state,
        app: {
          ...state.app,
          data: {
            ...state.app.data,
            deployedVersions: updatedVersions2
          }
        }
      };
    default:
      return state;
  }
};
