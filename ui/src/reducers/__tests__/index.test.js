import reducer from '../index';
import { REVERSE_SORT } from '../../actions/types.js';
import { SortByDirection, sortable } from '@patternfly/react-table';

describe('reducer', () => {
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

  const sortedApps = [
    ['App-J', 5, 120, 765],
    ['App-I', 6, 255, 3000],
    ['App-H', 1, 970, 98],
    ['App-G', 4, 655, 435],
    ['App-F', 3, 245, 873]
  ];

  const sortBy = { direction: SortByDirection.asc, index: 0 };
  const state = { apps: apps, sortBy: sortBy, columns: columns };

  it('should return the initial state', () => {
    expect(reducer(undefined, {})).toEqual(
      {
        apps: apps,
        sortBy: sortBy,
        columns: columns
      }
    );
  });

  it('should handle REVERSE_SORT', () => {
    const newState = reducer(state, { type: REVERSE_SORT, payload: { index: 0 } });
    expect(newState).toEqual({ ...state, sortBy: { direction: SortByDirection.desc, index: 0 }, apps: sortedApps });
  });
});
