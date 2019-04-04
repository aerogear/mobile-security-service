import React from 'react';
import PropTypes from 'prop-types';
import { Table, TableHeader, TableBody } from '@patternfly/react-table';
import './TableView.css';

export const TableView = ({ columns, rows, sortBy, onSort, onRowClick }) => (
  <Table aria-label="Mobile Apps" sortBy={sortBy} onSort={onSort} cells={columns} rows={rows}>
    <TableHeader />
    <TableBody onRowClick={onRowClick}/>
  </Table>
);

TableView.propTypes = {
  columns: PropTypes.array.isRequired,
  rows: PropTypes.array.isRequired,
  sortBy: PropTypes.object.isRequired,
  onSort: PropTypes.func.isRequired,
  onRowClick: PropTypes.func
};

export default TableView;
