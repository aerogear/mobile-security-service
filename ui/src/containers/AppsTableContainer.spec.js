import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import { AppsTableContainer } from './AppsTableContainer';
import TableView from '../components/common/TableView';
import { SortByDirection } from '@patternfly/react-table';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      apps: [{
        appId: 'someAppId',
        appName: 'myAppName',
        id: 'myID',
        numOfAppLaunches: 1,
        numOfCurrentInstalls: 2,
        numOfDeployedVersions: 4
      }],
      appRows: [['someAppId', 'myAppName', 'myID', 1, 2, 4]],
      sortBy: {
        index: 0,
        direction: SortByDirection.desc
      },
      isAppsRequestFailed: false,
      getApps: jest.fn(),
      appsTableSort: jest.fn(),
      className: 'myClassName',
      history: {
        push: jest.fn()
      }
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<AppsTableContainer {...props} />);
  const output = renderer.getRenderOutput();

  return {
    props: props,
    output: output
  };
};

describe('components', () => {
  describe('HeaderContainer', () => {
    it('should render', () => {
      const { output } = setup();
      expect(output.type).toBe('div');
      expect(output.props.children.type).toBe(TableView);
    });

    it('should render no table when no apps exist', () => {
      const { output } = setup({ apps: [] });
      expect(output.type).toBe('div');
      expect(output.props.children.props.children).toMatch('Unable to fetch any apps');
    });

    it('onRowClick should call history push', () => {
      const { output, props } = setup();

      output.props.children.props.onRowClick(undefined, ['myAppName']);
      expect(props.history.push).toBeCalled();
    });
  });
});
