import { REVERSE_SORT } from '../actions/types.js';

export const reverseAppsTableSort = (index) => {
  return {
    type: REVERSE_SORT,
    payload: {
      index: index
    }
  };
};
