import React, { Component } from 'react';
import { connect } from 'react-redux';

import AppGridHeader from './AppGridHeader';
import AppGridRows from './AppGridRows';
import { orderBy } from 'lodash';
import * as resolve from 'table-resolver';
import * as sort from 'sortabular';
import {
  actionHeaderCellFormatter,
  customHeaderFormattersDefinition,
  Table,
  TABLE_SORT_DIRECTION
} from 'patternfly-react';
import { compose } from 'recompose';

class AppGrid extends Component {
  constructor(props) {
    super(props);

    // enables our custom header formatters extensions to reactabular
    this.customHeaderFormatters = customHeaderFormattersDefinition;
  }

  render() {
    const { rows, sortingColumns, columns } = this.props.appGrid;
    console.log('this.state', this.state);

    const sortedRows = compose(
      sort.sorter({
        columns: columns,
        sortingColumns,
        sort: orderBy,
        strategy: sort.strategies.byProperty
      })
    )(rows);

    return (
      <div>
        <Table.PfProvider
          striped
          bordered
          hover
          dataTable
          columns={columns}
          components={{
            header: {
              cell: (cellProps) => {
                return this.customHeaderFormatters({
                  cellProps,
                  columns,
                  sortingColumns
                });
              }
            }
          }}
        >
          <AppGridHeader />
          <Table.Body
            rows={sortedRows}
            rowKey="id"
            onRow={() => {
              return {
                role: 'row'
              };
            }}
          />
        </Table.PfProvider>
      </div>
    );
  }
}

const mapStateToProps = (state) => {
  return {
    appGrid: state.appGrid
  };
};

const mapDispatchToProps = (dispatch) => {
  return {};
};

export default connect(mapStateToProps, mapDispatchToProps)(AppGrid);
