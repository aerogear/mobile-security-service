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

export default (state = initialState, action) => {
  switch (action.type) {
    case APPS_SORT: {
      return {
        ...state,
        sortBy: {
          direction: action.payload.direction,
          index: action.payload.index
        }
      };
    }
    case APPS_REQUEST: {
      return {
        ...state
      };
    }
    case APPS_SUCCESS: {
      return {
        ...state,
        data: action.result
      };
    }
    case APPS_FAILURE: {
      return {
        ...state,
        isAppsRequestFailed: true
      };
    }
    default: {
      return state;
    }
  }
};
