import {
  USER_REQUEST,
  USER_SUCCESS,
  USER_FAILURE,
  USER_AUTHENTICATE_SUCCESS,
  USER_AUTHENTICATE_FAILURE
} from '../actions/types.js';

const initialState = {
  data: {
    username: '',
    email: ''
  },
  isUserRequestFailed: false,
  isLoggedIn: false,
  isLoading: true
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
        data: action.result
      };
    }
    case USER_FAILURE: {
      return {
        ...state,
        isUserRequestFailed: true
      };
    }
    case USER_AUTHENTICATE_SUCCESS: {
      return {
        ...state,
        isLoggedIn: true,
        isLoading: false
      };
    }
    case USER_AUTHENTICATE_FAILURE: {
      return {
        ...state,
        isLoggedIn: false,
        isLoading: false
      };
    }
    default: {
      return state;
    }
  }
};
