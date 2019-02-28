import { REVERSE_SORT } from '../actions/types.js';
import { SortByDirection, sortable } from '@patternfly/react-table';

const columns = [
  { title: 'App ID', transforms: [ sortable ] },
  { title: 'Deployed Versions', transforms: [ sortable ] },
  { title: 'Current Installs', transforms: [ sortable ] },
  { title: 'Launches', transforms: [sortable] }
];

const apps = [
  [ 'App-F', 3, 245, 873 ],
  [ 'App-G', 4, 655, 435 ],
  [ 'App-H', 1, 970, 98 ],
  [ 'App-I', 6, 255, 3000 ],
  [ 'App-J', 5, 120, 765 ]
];

const sortBy = { direction: SortByDirection.asc, index: 0 };

const initialState = { apps: apps, sortBy: sortBy, columns: columns };

export default (state = initialState, action) => {
  switch (action.type) {
    case REVERSE_SORT:
      const reversedOrder = state.sortBy.direction === SortByDirection.asc ? SortByDirection.desc : SortByDirection.asc;
      const index = action.payload.index;
      const sortedRows = state.apps.sort((a, b) => (a[index] < b[index] ? -1 : a[index] > b[index] ? 1 : 0));
      const sortedApps = reversedOrder === SortByDirection.asc ? sortedRows : sortedRows.reverse();
      return {
        ...state,
        sortBy: {
          direction: reversedOrder,
          index: index
        },
        apps: sortedApps
      };
    default:
      return state;
  }
};
