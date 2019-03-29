import React from 'react';
import PropTypes from 'prop-types';
import { Table, TableHeader, TableBody } from '@patternfly/react-table';
import { withRouter } from 'react-router-dom';

export const AppsTable = ({ columns, rows, sortBy, onSort, onRowClick }) => (
  <Table aria-label="Mobile Apps" sortBy={sortBy} onSort={onSort} cells={columns} rows={rows}>
    <TableHeader />
    <TableBody onRowClick={onRowClick}/>
  </Table>
);

AppsTable.propTypes = {
  sortBy: PropTypes.object.isRequired,
  onSort: PropTypes.func.isRequired,
  columns: PropTypes.array.isRequired,
  rows: PropTypes.array.isRequired,
  onRowClick: PropTypes.func.isRequired
};

export default withRouter(AppsTable);
