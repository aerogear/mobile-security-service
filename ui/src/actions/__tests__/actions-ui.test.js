import * as actions from '../actions-ui';
import { APPS_SORT } from '../types';
import { SortByDirection } from '@patternfly/react-table';

describe('actions', () => {
  it('should reverse the app Table Sort', () => {
    const index = 0;
    const direction = SortByDirection.asc;
    const expectedAction = {
      type: APPS_SORT,
      payload: {
        index: index,
        direction: direction
      }
    };
    expect(actions.appsTableSort(index, direction)).toEqual(expectedAction);
  });
});
