import appsReducer from './apps';
import {
  APPS_SORT,
  APPS_SUCCESS,
  APPS_FAILURE
} from '../actions/types.js';
import { SortByDirection } from '@patternfly/react-table';

describe('appsReducer', () => {
  const initialState = {
    data: [],
    sortBy: { direction: SortByDirection.asc, index: 0 },
    isAppsRequestFailed: false
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

  it('should return the initial state', () => {
    const newInitialState = JSON.stringify(appsReducer(undefined, {}), null, 2);
    const expectedInitialState = JSON.stringify(initialState, null, 2);
    expect(newInitialState).toEqual(expectedInitialState);
  });

  it('should handle APPS_SORT', () => {
    const newState = appsReducer(initialState, { type: APPS_SORT, payload: { index: 1, direction: SortByDirection.desc } });
    expect(newState.sortBy).toEqual({ direction: SortByDirection.desc, index: 1 });
  });

  it('should handle APPS_SUCCESS', () => {
    const newState = appsReducer(initialState, { type: APPS_SUCCESS, result: resultApps });
    expect(newState.isAppsRequestFailed).toEqual(false);
    expect(newState.data).toEqual(resultApps);
  });

  it('should handle APPS_FAILURE', () => {
    const newState = appsReducer(initialState, { type: APPS_FAILURE });
    expect(newState.isAppsRequestFailed).toEqual(true);
  });
});
