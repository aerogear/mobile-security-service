import * as selectors from './index';
import { SortByDirection } from '@patternfly/react-table';

describe('selectors', () => {
  const mockState = {
    apps: {
      data: [
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
      ],
      sortBy: {
        index: 1,
        direction: SortByDirection.asc
      }
    },
    app: {
      data: {
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
      }
    }
  };

  it('should return correct getSortIndex', () => {
    const sortIndex = selectors.getSortIndex(null, mockState.apps.sortBy);
    expect(sortIndex).toEqual(1);
  });

  it('should return correct getSortDirection', () => {
    const sortIndex = selectors.getSortDirection(null, mockState.apps.sortBy);
    expect(sortIndex).toEqual(SortByDirection.asc);
  });

  it('should return getAppsTableRows in the correct format', () => {
    const appsTableRows = selectors.getAppsTableRows(mockState);
    expect(appsTableRows).toEqual([
      ['Foobar', 'com.aerogear.foobar', 0, 0, 0],
      ['Test App', 'com.aerogear.testapp', 2, 3, 6000]
    ]);
  });

  it('should return getAppVersionTableRows in the correct format', () => {
    const appVersionTableRows = selectors.getAppVersionTableRows(mockState);
    expect(appVersionTableRows).toEqual([
      ['v1.0', 100, 100, '2019-03-14T16:06:09.256498Z', true, 'Deprecated. Please upgrade to latest version', '23d334ef-e200-4639-8a22-c5aee389dd22'],
      ['v1.2', 75, 921, '2019-01-20 12:12:12.256498Z', false, 'LTS', 'a7ab467a-e719-49f3-9ec0-200898703583']
    ]);
  });

  it('should return getSortedTableRows for integer columns in the correct format', () => {
    const appsTableRows = selectors.getAppsTableRows(mockState);
    const sortedAppsTableRows = selectors.getSortedTableRows(appsTableRows, 2, SortByDirection.asc);
    expect(sortedAppsTableRows).toEqual([
      ['Foobar', 'com.aerogear.foobar', 0, 0, 0],
      ['Test App', 'com.aerogear.testapp', 2, 3, 6000]
    ]);
  });

  it('should return getSortedTableRows for string columns in the correct format', () => {
    const appsTableRows = selectors.getAppsTableRows(mockState);
    const sortedAppsTableRows = selectors.getSortedTableRows(appsTableRows, 1, SortByDirection.desc);
    expect(sortedAppsTableRows).toEqual([
      ['Test App', 'com.aerogear.testapp', 2, 3, 6000],
      ['Foobar', 'com.aerogear.foobar', 0, 0, 0]
    ]);
  });

  it('should ensure getSortedAppsTableRows selector does not recompute', () => {
    selectors.getSortedAppsTableRows(mockState, mockState.apps.sortBy);
    expect(selectors.getSortedAppsTableRows.recomputations()).toEqual(1);
    selectors.getSortedAppsTableRows(mockState, mockState.apps.sortBy);
    expect(selectors.getSortedAppsTableRows.recomputations()).toEqual(1);
    const sortedAppsTableRows = selectors.getSortedAppsTableRows(mockState, { index: 2, direction: SortByDirection.asc });
    expect(selectors.getSortedAppsTableRows.recomputations()).toEqual(2);
    expect(sortedAppsTableRows).toEqual([
      ['Foobar', 'com.aerogear.foobar', 0, 0, 0],
      ['Test App', 'com.aerogear.testapp', 2, 3, 6000]
    ]);
  });

  it('should ensure getSortedAppVersionTableRows selector does not recompute', () => {
    selectors.getSortedAppVersionTableRows(mockState, mockState.apps.sortBy);
    expect(selectors.getSortedAppVersionTableRows.recomputations()).toEqual(1);
    selectors.getSortedAppVersionTableRows(mockState, mockState.apps.sortBy);
    expect(selectors.getSortedAppVersionTableRows.recomputations()).toEqual(1);
    selectors.getSortedAppVersionTableRows(mockState, { index: 1, direction: SortByDirection.desc });
    expect(selectors.getSortedAppVersionTableRows.recomputations()).toEqual(2);
  });
});
