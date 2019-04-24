import appReducer from './app';
import {
  APP_VERSIONS_SORT,
  APP_SUCCESS,
  APP_FAILURE,
  SAVE_APP_VERSIONS_SUCCESS,
  SAVE_APP_VERSIONS_FAILURE,
  SET_APP_DETAILED_DIRTY,
  UPDATE_DISABLED_APP,
  UPDATE_VERSION_CUSTOM_MESSAGE,
  DISABLE_APP_SUCCESS,
  DISABLE_APP_FAILURE
} from '../actions/types.js';
import { SortByDirection } from '@patternfly/react-table';

describe('appReducer', () => {
  const initialState = {
    savedData: {
      deployedVersions: []
    },
    data: {
      deployedVersions: []
    },
    sortBy: { direction: SortByDirection.asc, index: 0 },
    isAppRequestFailed: false,
    isSaveAppRequestFailed: false,
    isSaveAppRequestSuccess: false,
    isDisableAppRequestFailed: false,
    isDirty: false
  };

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
      }
    ]
  };

  it('should return the initial state', () => {
    const newInitialState = JSON.stringify(appReducer(undefined, {}), null, 2);
    const expectedInitialState = JSON.stringify(initialState, null, 2);
    expect(newInitialState).toEqual(expectedInitialState);
  });

  it('should handle APP_VERSIONS_SORT', () => {
    const newState = appReducer(initialState, {
      type: APP_VERSIONS_SORT,
      payload: { index: 0, direction: SortByDirection.asc }
    });

    expect(newState.sortBy).toEqual({ direction: SortByDirection.asc, index: 0 });
  });

  it('should handle APP_SUCCESS', () => {
    const newState = appReducer(initialState, { type: APP_SUCCESS, result: resultApp });
    expect(newState.isAppRequestFailed).toEqual(false);
    expect(newState.data).toEqual(resultApp);
  });

  it('should handle APP_FAILURE', () => {
    const newState = appReducer(initialState, { type: APP_FAILURE });
    expect(newState.isAppRequestFailed).toEqual(true);
  });

  it('should handle SAVE_APP_VERSIONS_SUCCESS', () => {
    const initialSaveAppVersionsState = {
      ...initialState,
      data: resultApp
    };

    const newState = appReducer(initialSaveAppVersionsState, { type: SAVE_APP_VERSIONS_SUCCESS });

    expect(newState.isSaveAppRequestFailed).toEqual(false);
    expect(newState.isDirty).toEqual(false);
    expect(newState.savedData).toEqual(resultApp);
  });

  it('should handle SAVE_APP_VERSIONS_FAILURE', () => {
    const newState = appReducer(initialState, { type: SAVE_APP_VERSIONS_FAILURE });
    expect(newState.isSaveAppRequestFailed).toEqual(true);
  });

  it('should handle DISABLE_APP_SUCCESS', () => {
    const initialDisableAppVersionsState = {
      ...initialState,
      data: resultApp
    };

    const newState = appReducer(initialDisableAppVersionsState, { type: DISABLE_APP_SUCCESS, result: 'custom message' });

    const newAppData = {
      ...initialDisableAppVersionsState.data,
      deployedVersions: initialDisableAppVersionsState.data.deployedVersions.map((version) => {
        return {
          ...version,
          disabledMessage: 'custom message',
          disabled: true
        };
      })
    };

    expect(newState.isDisableAppRequestFailed).toEqual(false);
    expect(newState.data).toEqual(newAppData);
    expect(newState.savedData).toEqual(newAppData);
  });

  it('should handle DISABLE_APP_FAILURE', () => {
    const newState = appReducer(initialState, { type: DISABLE_APP_FAILURE });
    expect(newState.isDisableAppRequestFailed).toEqual(true);
  });

  it('should handle SET_APP_DETAILED_DIRTY', () => {
    const newState = appReducer(initialState, {
      type: SET_APP_DETAILED_DIRTY,
      payload: {
        isDirty: true
      }
    });
    expect(newState.isDirty).toEqual(true);
  });

  it('should handle UPDATE_DISABLED_APP', () => {
    const appState = appReducer(initialState, { type: APP_SUCCESS, result: resultApp });
    const versionId = appState.data.deployedVersions[0].id;

    const updatedState = appReducer(appState, {
      type: UPDATE_DISABLED_APP,
      payload: { id: versionId, isDisabled: false }
    });

    expect(updatedState.data.deployedVersions[0].disabled).toEqual(false);
    expect(updatedState.data.deployedVersions).toEqual(appState.data.deployedVersions);
  });

  it('should handle UPDATE_VERSION_CUSTOM_MESSAGE', () => {
    const appState = appReducer(initialState, { type: APP_SUCCESS, result: resultApp });
    const versionId = appState.data.deployedVersions[0].id;
    const updatedMessage = appState.data.deployedVersions[0].disabledMessage + '-newText';

    const updatedState = appReducer(appState, {
      type: UPDATE_VERSION_CUSTOM_MESSAGE,
      payload: { id: versionId, value: updatedMessage }
    });

    expect(updatedState.data.deployedVersions[0].disabledMessage).toEqual(updatedMessage);
    expect(updatedState.data.deployedVersions).toEqual(appState.data.deployedVersions);
  });
});
