import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import { AppVersionsTableContainer } from './AppVersionsTableContainer';
import TableView from '../../components/common/TableView';
import { Checkbox, TextInput } from '@patternfly/react-core';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      className: '',
      sortBy: {
        direction: 'asc',
        index: 1
      },
      appVersionRows: [],
      appVersionsTableSort: jest.fn(),
      updateDisabledAppVersion: jest.fn(),
      updateVersionCustomMessage: jest.fn()
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<AppVersionsTableContainer {...props} />);
  const output = renderer.getRenderOutput();

  return {
    props: props,
    output: output
  };
};

describe('components', () => {
  describe('TableView', () => {
    it('should not render when no app data', () => {
      const { output } = setup();
      const nonTableView = output.props.children;
      expect(nonTableView.type).not.toBe(TableView);
    });

    it('should render when app data', () => {
      const { output } = setup({
        appVersionRows: [
          [
            'v1.0',
            100,
            100,
            '2019-03-14T16:06:09.256498Z',
            true,
            'Deprecated. Please upgrade to latest version',
            '1234'
          ]
        ]
      });
      const tableView = output.props.children;
      expect(tableView.type).toBe(TableView);
    });

    describe('Table row data', () => {
      it('should render checkbox', () => {
        const { output } = setup({
          appVersionRows: [
            [
              'v1.0',
              100,
              100,
              '2019-03-14T16:06:09.256498Z',
              true,
              'Deprecated. Please upgrade to latest version',
              '1234'
            ]
          ]
        });
        const tableView = output.props.children;
        const tableRows = tableView.props.rows;
        const checkbox = tableRows[0][4];
        expect(JSON.stringify(checkbox.type)).toBe(JSON.stringify(Checkbox));
      });

      it('should render text input', () => {
        const { output } = setup({
          appVersionRows: [
            [
              'v1.0',
              100,
              100,
              '2019-03-14T16:06:09.256498Z',
              true,
              'Deprecated. Please upgrade to latest version',
              '1234'
            ]
          ]
        });
        const tableView = output.props.children;
        const tableRows = tableView.props.rows;
        const textInput = tableRows[0][5];
        expect(JSON.stringify(textInput.type)).toBe(JSON.stringify(TextInput));
      });
    });
  });
});

describe('events', () => {
  describe('Disable checkbox handles changes', () => {
    it('onChange should call updateDisabledAppVersion', () => {
      const { output, props } = setup({
        appVersionRows: [
          [
            'v1.0',
            100,
            100,
            '2019-03-14T16:06:09.256498Z',
            true,
            'Deprecated. Please upgrade to latest version',
            '1234'
          ]
        ]
      });
      const tableView = output.props.children;
      const tableRows = tableView.props.rows;
      const checkboxCell = tableRows[0][4];
      const checkbox = checkboxCell.props.children;
      checkbox.props.onChange(false, {
        target: {
          id: '1234'
        }
      });
      expect(props.updateDisabledAppVersion).toBeCalledWith('1234', false);
    });
  });

  describe('Custom message field handles onBlur', () => {
    it('onBlur should call updateVersionCustomMessage', () => {
      const { output, props } = setup({
        appVersionRows: [
          [
            'v1.0',
            100,
            100,
            '2019-03-14T16:06:09.256498Z',
            true,
            'Deprecated. Please upgrade to latest version',
            '1234'
          ]
        ]
      });
      const tableView = output.props.children;
      const tableRows = tableView.props.rows;
      const textInputCell = tableRows[0][5];
      const textInput = textInputCell.props.children;
      textInput.props.onBlur({
        target: {
          id: '1234',
          value: 'Deprecated. Please upgrade to latest version'
        }
      });
      expect(props.updateVersionCustomMessage).toBeCalledWith('1234', 'Deprecated. Please upgrade to latest version');
    });
  });
});
