import {
  USER_REQUEST,
  USER_SUCCESS,
  USER_FAILURE
} from '../actions/types.js';

const initialState = {
  user: {
    username: '',
    email: ''
  },
  isUserRequestFailed: false
};

export default (state = initialState, action) => {
  switch (action.type) {
    case USER_REQUEST: {
      return {
        ...state
      };
    }
    case USER_SUCCESS: {
      return {
        ...state,
        user: action.result
      };
    }
    case USER_FAILURE: {
      return {
        ...state,
        isUserRequestFailed: true
      };
    }
    default: {
      return state;
    }
  }
};
