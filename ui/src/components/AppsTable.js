import React from 'react';
import { Table, TableHeader, TableBody } from '@patternfly/react-table';
import { withRouter } from 'react-router-dom';
import './Table.css';

export const AppsTable = ({ columns, rows, sortBy, onSort, onRowClick }) => (
  <Table aria-label="Mobile Apps" sortBy={sortBy} onSort={onSort} cells={columns} rows={rows}>
    <TableHeader />
    <TableBody onRowClick={onRowClick}/>
  </Table>
);

export default withRouter(AppsTable);
