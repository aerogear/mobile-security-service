import reducer from '../index';
import { APPS_SUCCESS, REVERSE_SORT, APPS_FAILURE, TOGGLE_HEADER_DROPDOWN } from '../../actions/types.js';
import { SortByDirection, sortable } from '@patternfly/react-table';

describe('reducer', () => {
  const columns = [
    { title: 'App Name', transforms: [ sortable ] },
    { title: 'App ID', transforms: [ sortable ] },
    { title: 'Deployed Versions', transforms: [ sortable ] },
    { title: 'Current Installs', transforms: [ sortable ] },
    { title: 'Launches', transforms: [ sortable ] }
  ];
  const apps = [];

  const sortedRows = [
    [ 'Test App', 'com.aerogear.testapp', 2, 3, 6000 ],
    [ 'Foobar', 'com.aerogear.foobar', 0, 0, 0 ]
  ];

  const sortBy = { direction: SortByDirection.asc, index: 0 };
  const initialState = { apps: { rows: apps, data: {} }, sortBy: sortBy, columns: columns, isAppsRequestFailed: false, currentUser: 'currentUser', isUserDropdownOpen: false };

  const resultApps =
    [
      {
        'id': '1b9e7a5f-af7c-4055-b488-72f2b5f72266',
        'appId': 'com.aerogear.foobar',
        'appName': 'Foobar',
        'numOfDeployedVersions': 0,
        'numOfCurrentInstalls': 0,
        'numOfAppLaunches': 0
      },
      {
        'id': '0890506c-3dd1-43ad-8a09-21a4111a65a6',
        'appId': 'com.aerogear.testapp',
        'appName': 'Test App',
        'numOfDeployedVersions': 2,
        'numOfCurrentInstalls': 3,
        'numOfAppLaunches': 6000
      }
    ];
  const rows = [
    [ 'Foobar', 'com.aerogear.foobar', 0, 0, 0 ],
    [ 'Test App', 'com.aerogear.testapp', 2, 3, 6000 ]
  ];

  it('should return the initial state', () => {
    expect(reducer(undefined, {})).toEqual(initialState);
  });

  it('should handle REVERSE_SORT', () => {
    const appsState = reducer(initialState, { type: APPS_SUCCESS, result: resultApps });
    const newState = reducer(appsState, { type: REVERSE_SORT, payload: { index: 1 } });
    expect(newState).toEqual({ ...initialState, sortBy: { direction: SortByDirection.desc, index: 1 }, apps: { rows: sortedRows } });
  });

  it('should handle APPS_SUCCESS', () => {
    const newState = reducer(initialState, { type: APPS_SUCCESS, result: resultApps });
    expect(newState.isAppsRequestFailed).toEqual(false);
    expect(newState.apps).toEqual({ rows: rows, data: resultApps });
  });

  it('should handle APPS_FAILURE', () => {
    const newState = reducer(initialState, { type: APPS_FAILURE });
    expect(newState.isAppsRequestFailed).toEqual(true);
  });

  it('should toggle header dropdown state', () => {
    const dropdownBeforeToggle = initialState.isUserDropdownOpen;
    const newState = reducer(initialState, { type: TOGGLE_HEADER_DROPDOWN });
    expect(newState.isUserDropdownOpen).toEqual(!dropdownBeforeToggle);
  });
});
