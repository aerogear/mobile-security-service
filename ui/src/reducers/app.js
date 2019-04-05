import {
  APP_VERSIONS_SORT,
  APP_REQUEST,
  APP_SUCCESS,
  APP_FAILURE,
  APP_DETAILED_IS_DIRTY,
  UPDATE_DISABLED_APP,
  UPDATE_VERSION_CUSTOM_MESSAGE,
  SAVE_APP_VERSIONS,
  SAVE_APP_VERSIONS_SUCCESS,
  SAVE_APP_VERSIONS_FAILURE
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
  isSaveAppRequestFailed: false,
  isDirty: false
};

export const cloneAppData = (appData) => {
  return {
    ...appData,
    deployedVersions: [
      ...appData.deployedVersions
    ].map(version => ({ ...version }))
  };
};

export default (state = initialState, action) => {
  switch (action.type) {
    case APP_VERSIONS_SORT: {
      return {
        ...state,
        sortBy: {
          direction: action.payload.direction,
          index: action.payload.index
        }
      };
    }
    case APP_REQUEST: {
      return {
        ...state
      };
    }
    case APP_SUCCESS: {
      return {
        ...state,
        data: cloneAppData(action.result),
        savedData: cloneAppData(action.result)
      };
    }
    case APP_FAILURE: {
      return {
        ...state,
        isAppRequestFailed: true
      };
    }
    case APP_DETAILED_IS_DIRTY: {
      return {
        ...state,
        isDirty: action.payload.isDirty
      };
    }
    case UPDATE_DISABLED_APP: {
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
    }
    case UPDATE_VERSION_CUSTOM_MESSAGE: {
      const updatedVersions2 = state.data.deployedVersions.map((version) => {
        if (version.id === action.payload.id) {
          version.disabledMessage = action.payload.value;
        }
        return version;
      });

      return {
        ...state,
        direction: action.payload.direct,
        data: {
          ...state.data,
          deployedVersions: updatedVersions2
        }
      };
    }
    case SAVE_APP_VERSIONS: {
      return {
        ...state
      };
    }
    case SAVE_APP_VERSIONS_SUCCESS: {
      return {
        ...state,
        isSaveAppRequestFailed: false,
        savedData: cloneAppData(state.data),
        isDirty: false
      };
    }
    case SAVE_APP_VERSIONS_FAILURE: {
      return {
        ...state,
        isSaveAppRequestFailed: true
      };
    }
    default: {
      return state;
    }
  }
};
