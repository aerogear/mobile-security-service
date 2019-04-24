import userReducer from './user';
import {
  USER_REQUEST,
  USER_SUCCESS,
  USER_FAILURE,
  USER_AUTHENTICATE_REQUEST,
  USER_AUTHENTICATE_SUCCESS,
  USER_AUTHENTICATE_FAILURE
} from '../actions/types.js';

describe('userReducer', () => {
  const initialState = {
    data: {
      username: '',
      email: ''
    },
    isUserRequestFailed: false,
    isLoggedIn: false,
    isLoading: true
  };

  const resultUser = {
    username: 'abcdef',
    email: 'abc@def.com'
  };

  it('should return the initial state', () => {
    const newInitialState = JSON.stringify(userReducer(undefined, {}), null, 2);
    const expectedInitialState = JSON.stringify(initialState, null, 2);
    expect(newInitialState).toEqual(expectedInitialState);
  });

  it('should handle USER_REQUEST', () => {
    const newState = userReducer(initialState, { type: USER_REQUEST });
    expect(newState).toEqual(initialState);
  });

  it('should handle USER_SUCCESS', () => {
    const newState = userReducer(initialState, { type: USER_SUCCESS, result: resultUser });
    expect(newState.isUserRequestFailed).toEqual(false);
    expect(newState.data).toEqual(resultUser);
  });

  it('should handle USER_FAILURE', () => {
    const newState = userReducer(initialState, { type: USER_FAILURE });
    expect(newState.isUserRequestFailed).toEqual(true);
  });

  it('should handle USER_AUTHENTICATE_REQUEST', () => {
    const newState = userReducer(initialState, { type: USER_AUTHENTICATE_REQUEST });
    expect(newState).toEqual(initialState);
  });

  it('should handle USER_AUTHENTICATE_SUCCESS', () => {
    const newState = userReducer(initialState, { type: USER_AUTHENTICATE_SUCCESS });
    expect(newState.isLoggedIn).toEqual(true);
    expect(newState.isLoading).toEqual(false);
  });

  it('should handle USER_AUTHENTICATE_FAILURE', () => {
    const newState = userReducer(initialState, { type: USER_AUTHENTICATE_FAILURE });
    expect(newState.isLoggedIn).toEqual(false);
    expect(newState.isLoading).toEqual(false);
  });
});
