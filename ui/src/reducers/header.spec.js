import headerReducer from './header';
import {
  USER_REQUEST,
  USER_SUCCESS,
  USER_FAILURE
} from '../actions/types.js';

describe('headerReducer', () => {
  const initialState = {
    user: {
      username: '',
      email: ''
    },
    isUserRequestFailed: false
  };

  const resultUser = {
    username: 'abcdef',
    email: 'abc@def.com'
  };

  it('should return the initial state', () => {
    const newInitialState = JSON.stringify(headerReducer(undefined, {}), null, 2);
    const expectedInitialState = JSON.stringify(initialState, null, 2);
    expect(newInitialState).toEqual(expectedInitialState);
  });

  it('should handle USER_REQUEST', () => {
    const newState = headerReducer(initialState, { type: USER_REQUEST });
    expect(newState).toEqual(initialState);
  });

  it('should handle USER_SUCCESS', () => {
    const newState = headerReducer(initialState, { type: USER_SUCCESS, result: resultUser });
    expect(newState.isUserRequestFailed).toEqual(false);
    expect(newState.user).toEqual(resultUser);
  });

  it('should handle USER_FAILURE', () => {
    const newState = headerReducer(initialState, { type: USER_FAILURE });
    expect(newState.isUserRequestFailed).toEqual(true);
  });
});
