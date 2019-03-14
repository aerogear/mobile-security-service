import reducer from '../index';
import {
  APPS_SUCCESS,
  APPS_SORT,
  APPS_FAILURE,
  TOGGLE_HEADER_DROPDOWN,
  APP_SUCCESS, APP_FAILURE,
  APP_VERSIONS_SORT
} from '../../actions/types.js';
import { SortByDirection, sortable } from '@patternfly/react-table';

describe('reducer', () => {
  const initialState = {
    apps: { rows: [], data: {} },
    sortBy: { direction: SortByDirection.asc, index: 0 },
    appVersionsSortDirection: { direction: SortByDirection.asc, index: 0 },
    columns: [
      { title: 'App Name', transforms: [sortable] },
      { title: 'App ID', transforms: [sortable] },
      { title: 'Deployed Versions', transforms: [sortable] },
      { title: 'Current Installs', transforms: [sortable] },
      { title: 'Launches', transforms: [sortable] }
    ],
    appVersionsColumns: [
      { title: 'App Version', transforms: [sortable] },
      { title: 'Current Installs', transforms: [sortable] },
      { title: 'Launches', transforms: [sortable] },
      { title: 'Last Launched', transforms: [sortable] },
      { title: 'Disable on Startup', transforms: [sortable] },
      { title: 'Custom Disable Message', transforms: [sortable] }
    ],
    isAppsRequestFailed: false,
    currentUser: 'currentUser',
    isUserDropdownOpen: false,
    app: {
      deployedVersions: { rows: [], data: {} }
    },
    isAppRequestFailed: false
  };

  const resultApps = [
    {
      id: '1b9e7a5f-af7c-4055-b488-72f2b5f72266',
      appId: 'com.aerogear.foobar',
      appName: 'Foobar',
      numOfDeployedVersions: 0,
      numOfCurrentInstalls: 0,
      numOfAppLaunches: 0
    },
    {
      id: '0890506c-3dd1-43ad-8a09-21a4111a65a6',
      appId: 'com.aerogear.testapp',
      appName: 'Test App',
      numOfDeployedVersions: 2,
      numOfCurrentInstalls: 3,
      numOfAppLaunches: 6000
    }
  ];

  const sortedRows = [
    ['Test App', 'com.aerogear.testapp', 2, 3, 6000],
    ['Foobar', 'com.aerogear.foobar', 0, 0, 0]
  ];

  const resultAppDetails = {
    deployedVersions: [
      {
        id: '23d334ef-e200-4639-8a22-c5aee389dd22',
        version: 'v1.0',
        appId: 'com.aerogear.testapp1',
        disabled: true,
        disabledMessage: 'Deprecated. Please upgrade to latest version',
        numOfCurrentInstalls: 100,
        numOfAppLaunches: 100,
        lastLaunchedAt: '2019-03-14T16:06:09.256498Z'
      }
    ]
  };

  // const resultAppDetails = [
  //   ['v1.0', 55, 621, '2019-01-11 10:45:03', true, 'Deprecated. Please upgrade to latest version'],
  //   ['v1.1', 55, 621, '2019-01-11 10:45:03', true, 'Deprecated. Please upgrade to latest version'],
  //   ['v1.2', 75, 921, '2019-01-20 12:12:12', false, 'LTS'],
  //   ['v1.3', 125, 1221, '2019-01-31 11:05:43', false, 'Curent version'],
  //   ['v1.4', 40, 120, '2019-02-15 10:02:50', false, 'Beta version']
  // ];

  const sortedAppVersions = [
    ['v1.0', 100, 100, '2019-03-14T16:06:09.256498Z', true, 'Deprecated. Please upgrade to latest version']
    // ['v1.1', 55, 621, '2019-01-11 10:45:03', true, 'Deprecated. Please upgrade to latest version'],
    // ['v1.2', 75, 921, '2019-01-20 12:12:12', false, 'LTS'],
    // ['v1.3', 125, 1221, '2019-01-31 11:05:43', false, 'Curent version'],
    // ['v1.4', 40, 120, '2019-02-15 10:02:50', false, 'Beta version']
  ];

  const resultApp = {
    'id': '0890506c-3dd1-43ad-8a09-21a4111a65a6',
    'appId': 'com.aerogear.testapp',
    'appName': 'Test App',
    'numOfDeployedVersions': 2,
    'numOfCurrentInstalls': 3,
    'numOfAppLaunches': 6000
  };

  const rows = [
    ['Foobar', 'com.aerogear.foobar', 0, 0, 0],
    ['Test App', 'com.aerogear.testapp', 2, 3, 6000]
  ];

  it('should return the initial state', () => {
    expect(reducer(undefined, {})).toEqual(initialState);
  });

  it('should handle APPS_SORT', () => {
    const appsState = reducer(initialState, { type: APPS_SUCCESS, result: resultApps });
    const newState = reducer(appsState, { type: APPS_SORT, payload: { index: 1, direction: SortByDirection.desc } });
    expect(newState).toEqual({
      ...initialState,
      sortBy: { direction: SortByDirection.desc, index: 1 },
      apps: { data: resultApps, rows: sortedRows }
    });
  });

  it('should handle APP_VERSIONS_SORT', () => {
    const appsState = reducer(initialState, { type: APP_SUCCESS, result: resultAppDetails });
    const newState = reducer(appsState, {
      type: APP_VERSIONS_SORT,
      payload: { index: 0, direction: SortByDirection.asc }
    });
    expect(newState).toEqual({
      ...initialState,
      appVersionsSortDirection: { direction: SortByDirection.asc, index: 0 },
      appVersions: sortedAppVersions
    });
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

  it('should handle APP_SUCCESS', () => {
    const newState = reducer(initialState, { type: APP_SUCCESS, result: resultApp });
    expect(newState.isAppRequestFailed).toEqual(false);
    expect(newState.app).toEqual(resultApp);
  });

  it('should handle APP_FAILURE', () => {
    const newState = reducer(initialState, { type: APP_FAILURE });
    expect(newState.isAppRequestFailed).toEqual(true);
  });

  it('should toggle header dropdown state', () => {
    const dropdownBeforeToggle = initialState.isUserDropdownOpen;
    const newState = reducer(initialState, { type: TOGGLE_HEADER_DROPDOWN });
    expect(newState.isUserDropdownOpen).toEqual(!dropdownBeforeToggle);
  });
});
