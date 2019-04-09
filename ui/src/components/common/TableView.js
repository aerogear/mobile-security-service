import React from 'react';
import PropTypes from 'prop-types';
import { Table, TableHeader, TableBody } from '@patternfly/react-table';
import './TableView.css';

/**
 * Stateless presentational component to display a UI table
 *
 * @param {object} props Component props
 * @param {array} props.columns Array of objects containing column title
 * @param {array} props.rows Array of arrays containing table row data
 * @param {object} props.sortBy sort information for the table, with the column index and sort direction
 * @param {func} props.onSort logic to determine how the table should sort
 * @param {func} props.onRowClick logic to execute on row click
 */
export const TableView = ({ columns, rows, sortBy, onSort, onRowClick }) => (
  <Table aria-label="Mobile Apps" sortBy={sortBy} onSort={onSort} cells={columns} rows={rows}>
    <TableHeader />
    <TableBody onRowClick={onRowClick}/>
  </Table>
);

TableView.propTypes = {
  columns: PropTypes.arrayOf(
    PropTypes.shape({
      title: PropTypes.string.isRequired,
      transforms: PropTypes.array
    }).isRequired
  ).isRequired,
  rows: PropTypes.arrayOf(
    PropTypes.array
  ).isRequired,
  sortBy: PropTypes.shape({
    index: PropTypes.number.isRequired,
    direction: PropTypes.string.isRequired
  }).isRequired,
  onSort: PropTypes.func.isRequired,
  onRowClick: PropTypes.func
};

export default TableView;
