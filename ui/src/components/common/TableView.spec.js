import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import TableView from './TableView';
import { Table, TableHeader, TableBody } from '@patternfly/react-table';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      columns: [{ title: 'column1' }, { title: 'column2' }],
      rows: [[123, 'row1'], [456, 'row2']],
      sortBy: { index: 0, direction: 'asc' },
      onSort: jest.fn(),
      onRowClick: jest.fn()
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<TableView {...props} />);
  const output = renderer.getRenderOutput();

  return {
    props: props,
    output: output
  };
};

describe('components', () => {
  describe('TableView', () => {
    it('should render', () => {
      const { output } = setup();
      expect(output.type).toBe(Table);
      expect(output.props.cells).toEqual([{ title: 'column1' }, { title: 'column2' }]);
      expect(output.props.rows).toEqual([[123, 'row1'], [456, 'row2']]);

      const [ TableHeaderChild, TableBodyChild ] = output.props.children;

      expect(TableHeaderChild.type).toBe(TableHeader);
      expect(TableBodyChild.type).toBe(TableBody);
    });
  });
});
