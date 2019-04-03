import reducer from '../index';
import * as actions from '../../actions/types.js';
import { SortByDirection, sortable, cellWidth } from '@patternfly/react-table';

describe('reducer', () => {
  const initialState = {
    apps: { rows: [], data: [] },
    sortBy: { direction: SortByDirection.asc, index: 0 },
    appVersionsSortDirection: { direction: SortByDirection.asc, index: 0 },
    columns: [
      { title: 'APP NAME', transforms: [ sortable ] },
      { title: 'APP ID', transforms: [ sortable ] },
      { title: 'DEPLOYED VERSIONS', transforms: [ sortable ] },
      { title: 'CURRENT INSTALLS', transforms: [ sortable ] },
      { title: 'LAUNCHES', transforms: [ sortable ] }
    ],
    appVersionsColumns: [
      { title: 'APP VERSION', transforms: [ sortable, cellWidth(10) ] },
      { title: 'CURRENT INSTALLS', transforms: [ sortable, cellWidth(10) ] },
      { title: 'LAUNCHES', transforms: [ sortable, cellWidth(10) ] },
      { title: 'LAST LAUNCHED', transforms: [ sortable, cellWidth(15) ] },
      { title: 'DISABLE ON STARTUP', transforms: [ sortable, cellWidth(10) ] },
      { title: 'CUSTOM DISABLE MESSAGE', transforms: [ sortable, cellWidth('max') ] }
    ],
    isAppsRequestFailed: false,
    currentUser: 'currentUser',
    navigationModal: {
      isOpen: false,
      targetLocation: undefined
    },
    isSaveAppModalOpen: false,
    isAppDetailedDirty: false,
    app: {
      data: {},
      versionsRows: []
    },
    isAppRequestFailed: false,
    modals: {
      disableApp: {
        isOpen: false
      }
    }
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
    [ 'Test App', 'com.aerogear.testapp', 2, 3, 6000 ],
    [ 'Foobar', 'com.aerogear.foobar', 0, 0, 0 ]
  ];

  const sortedAppVersions = [
    [ 'v1.0', 100, 100, '2019-03-14T16:06:09.256498Z', true, 'Deprecated. Please upgrade to latest version' ],
    [ 'v1.1', 55, 621, '2019-01-11 10:45:03.256498Z', true, 'Deprecated. Please upgrade to latest version' ],
    [ 'v1.2', 75, 921, '2019-01-20 12:12:12.256498Z', false, 'LTS' ],
    [ 'v1.3', 125, 1221, '2017-01-31 11:05:43.256498Z', false, '' ],
    [ 'v1.4', 40, 120, '2018-02-15 10:02:50.256498Z', false, '' ]
  ];

  const resultApp = {
    id: '1b9e7a5f-af7c-4055-b488-72f2b5f72266',
    appId: 'com.aerogear.testapp1',
    appName: 'Foobar',
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
      },
      {
        id: 'e23bcfd4-0d0a-48ee-96a8-db79141226da',
        version: 'v1.1',
        appId: 'com.aerogear.testapp1',
        disabled: true,
        disabledMessage: 'Deprecated. Please upgrade to latest version',
        numOfCurrentInstalls: 55,
        numOfAppLaunches: 621,
        lastLaunchedAt: '2019-01-11 10:45:03.256498Z'
      },
      {
        id: 'a7ab467a-e719-49f3-9ec0-200898703583',
        version: 'v1.2',
        appId: 'com.aerogear.testapp1',
        disabled: false,
        disabledMessage: 'LTS',
        numOfCurrentInstalls: 75,
        numOfAppLaunches: 921,
        lastLaunchedAt: '2019-01-20 12:12:12.256498Z'
      },
      {
        id: '6c656492-62ef-406f-a7ec-866b112488f5',
        version: 'v1.3',
        appId: 'com.aerogear.testapp1',
        disabled: false,
        disabledMessage: '',
        numOfCurrentInstalls: 125,
        numOfAppLaunches: 1221,
        lastLaunchedAt: '2017-01-31 11:05:43.256498Z'
      },
      {
        id: '2019-02-15 10:02:50.256498Z',
        version: 'v1.4',
        appId: 'com.aerogear.testapp1',
        disabled: false,
        disabledMessage: '',
        numOfCurrentInstalls: 40,
        numOfAppLaunches: 120,
        lastLaunchedAt: '2018-02-15 10:02:50.256498Z'
      }
    ]
  };

  const rows = [ [ 'Foobar', 'com.aerogear.foobar', 0, 0, 0 ], [ 'Test App', 'com.aerogear.testapp', 2, 3, 6000 ] ];

  it('should return the initial state', () => {
    const newInitialState = JSON.stringify(reducer(undefined, {}), null, 2);
    const expectedInitialState = JSON.stringify(initialState, null, 2);
    expect(newInitialState).toEqual(expectedInitialState);
  });

  it('should handle APPS_SORT', () => {
    const appsState = reducer(initialState, { type: actions.APPS_SUCCESS, result: resultApps });
    const newState = reducer(appsState, { type: actions.APPS_SORT, payload: { index: 1, direction: SortByDirection.desc } });
    expect(newState).toEqual({
      ...initialState,
      sortBy: { direction: SortByDirection.desc, index: 1 },
      apps: { data: resultApps, rows: sortedRows }
    });
  });

  it('should handle APP_VERSIONS_SORT', () => {
    const appsState = reducer(initialState, { type: actions.APP_SUCCESS, result: resultApp });
    const newState = reducer(appsState, {
      type: actions.APP_VERSIONS_SORT,
      payload: { index: 0, direction: SortByDirection.asc }
    });
    expect(newState).toEqual({
      ...initialState,
      appVersionsSortDirection: { direction: SortByDirection.asc, index: 0 },
      app: {
        data: resultApp,
        versionsRows: sortedAppVersions
      }
    });
  });

  it('should handle APPS_SUCCESS', () => {
    const newState = reducer(initialState, { type: actions.APPS_SUCCESS, result: resultApps });
    expect(newState.isAppsRequestFailed).toEqual(false);
    expect(newState.apps).toEqual({ rows: rows, data: resultApps });
  });

  it('should handle APPS_FAILURE', () => {
    const newState = reducer(initialState, { type: actions.APPS_FAILURE });
    expect(newState.isAppsRequestFailed).toEqual(true);
  });

  it('should handle APP_SUCCESS', () => {
    const newState = reducer(initialState, { type: actions.APP_SUCCESS, result: resultApp });
    expect(newState.isAppRequestFailed).toEqual(false);
    expect(newState.app).toEqual({
      data: resultApp,
      versionsRows: sortedAppVersions
    });
  });

  it('should handle APP_FAILURE', () => {
    const newState = reducer(initialState, { type: actions.APP_FAILURE });
    expect(newState.isAppRequestFailed).toEqual(true);
  });

  it('should handle open TOGGLE_NAVIGATION_MODAL', () => {
    const newState = reducer(initialState, {
      type: actions.TOGGLE_NAVIGATION_MODAL,
      payload: {
        isOpen: true,
        targetLocation: '/'
      }
    });
    expect(newState.navigationModal.isOpen).toEqual(true);
    expect(newState.navigationModal.targetLocation).toEqual('/');
  });

  it('should handle close TOGGLE_NAVIGATION_MODAL', () => {
    const newState = reducer(initialState, {
      type: actions.TOGGLE_NAVIGATION_MODAL,
      payload: {
        isOpen: false
      }
    });
    expect(newState.navigationModal.isOpen).toEqual(false);
    expect(newState.navigationModal.targetLocation).toEqual(undefined);
  });

  it('should handle close TOGGLE_SAVE_APP_MODAL', () => {
    const newState = reducer(initialState, {
      type: actions.TOGGLE_SAVE_APP_MODAL,
      payload: { isSaveAppModalOpen: false }
    });
    expect(newState.isSaveAppModalOpen).toEqual(false);
  });

  it('should handle open TOGGLE_SAVE_APP_MODAL', () => {
    const newState = reducer(initialState, {
      type: actions.TOGGLE_SAVE_APP_MODAL,
      payload: { isSaveAppModalOpen: true }
    });
    expect(newState.isSaveAppModalOpen).toEqual(true);
  });

  it('should toggle TOGGLE_DISABLE_APP_MODAL', () => {
    const isOpen = initialState.modals.disableApp.isOpen;
    const newState = reducer(initialState, { type: actions.TOGGLE_DISABLE_APP_MODAL });
    expect(newState.modals.disableApp.isOpen).toEqual(!isOpen);
  });

  it('should handle TOGGLE_APP_DETAILED_IS_DIRTY', () => {
    const appDetailedDirtyBeforeToggle = initialState.isAppDetailedDirty;
    const newState = reducer(initialState, { type: actions.TOGGLE_APP_DETAILED_IS_DIRTY });
    expect(newState.isAppDetailedDirty).toEqual(!appDetailedDirtyBeforeToggle);
  });

  it('should update the checkbox disabled state', () => {
    const appState = reducer(initialState, { type: actions.APP_SUCCESS, result: resultApp });
    const isDisabled = appState.app.versionsRows[0][4];
    const updatedState = reducer(appState, {
      type: actions.UPDATE_DISABLED_APP,
      payload: { id: 'v1.0', isDisabled: !isDisabled }
    });
    expect(updatedState.app.versionsRows[0][4]).toEqual(!isDisabled);
  });

  it('should update the custom text state', () => {
    const appState = reducer(initialState, { type: actions.APP_SUCCESS, result: resultApp });
    const updatedMessage = appState.app.versionsRows[0][5] + '-newText';
    const updatedState = reducer(appState, {
      type: actions.UPDATE_VERSION_CUSTOM_MESSAGE,
      payload: { id: 'v1.0', value: updatedMessage }
    });
    expect(updatedState.app.versionsRows[0][5]).toEqual(updatedMessage);
  });
});
