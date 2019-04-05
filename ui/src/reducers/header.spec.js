import headerReducer from './header';
import { } from '../actions/types.js';

describe('headerReducer', () => {
  const initialState = {
    currentUser: 'currentUser'
  };

  it('should return the initial state', () => {
    const newInitialState = JSON.stringify(headerReducer(undefined, {}), null, 2);
    const expectedInitialState = JSON.stringify(initialState, null, 2);
    expect(newInitialState).toEqual(expectedInitialState);
  });
});
