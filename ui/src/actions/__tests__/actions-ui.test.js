import * as actions from '../actions-ui';
import { REVERSE_SORT } from '../types';

describe('actions', () => {
  it('should reverse the app Table Sort', () => {
    const index = 0;
    const expectedAction = {
      type: REVERSE_SORT,
      payload: {
        index: index
      }
    };
    expect(actions.reverseAppsTableSort(index)).toEqual(expectedAction);
  });
});
