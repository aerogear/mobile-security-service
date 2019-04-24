import {
  APP_VERSIONS_SORT,
  APP_REQUEST,
  APP_SUCCESS,
  APP_FAILURE,
  SET_APP_DETAILED_DIRTY,
  UPDATE_DISABLED_APP,
  UPDATE_VERSION_CUSTOM_MESSAGE,
  SAVE_APP_VERSIONS,
  SAVE_APP_VERSIONS_SUCCESS,
  SAVE_APP_VERSIONS_FAILURE,
  DISABLE_APP_SUCCESS,
  DISABLE_APP_FAILURE
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
  isSaveAppRequestSuccess: false,
  isDisableAppRequestFailed: false,
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
    case SET_APP_DETAILED_DIRTY: {
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
      const updatedVersions = state.data.deployedVersions.map((version) => {
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
          deployedVersions: updatedVersions
        }
      };
    }
    case SAVE_APP_VERSIONS: {
      return {
        ...state,
        isSaveAppRequestSuccess: false,
        isSaveAppRequestFailed: false
      };
    }
    case SAVE_APP_VERSIONS_SUCCESS: {
      return {
        ...state,
        isSaveAppRequestFailed: false,
        isSaveAppRequestSuccess: true,
        savedData: cloneAppData(state.data),
        isDirty: false
      };
    }
    case SAVE_APP_VERSIONS_FAILURE: {
      return {
        ...state,
        isSaveAppRequestFailed: true,
        isSaveAppRequestSuccess: false
      };
    }
    case DISABLE_APP_SUCCESS: {
      const updatedVersions = state.data.deployedVersions.map((version) => {
        return {
          ...version,
          disabledMessage: action.result || version.disabledMessage,
          disabled: true
        };
      });

      const appData = { ...state.data, deployedVersions: updatedVersions };

      return {
        ...state,
        data: appData,
        savedData: cloneAppData(appData),
        isDisableAppRequestFailed: false
      };
    }
    case DISABLE_APP_FAILURE: {
      return {
        ...state,
        isDisableAppRequestFailed: true
      };
    }
    default: {
      return state;
    }
  }
};
