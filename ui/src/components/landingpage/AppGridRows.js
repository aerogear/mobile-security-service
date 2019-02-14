import React, { Component } from 'react';
import { connect } from 'react-redux';

import { Table } from 'patternfly-react';
import { compose } from 'recompose';
import * as sort from 'sortabular';
import { orderBy } from 'lodash';

class AppGridRows extends Component {
  render() {
    const { rows, sortingColumns, columns } = this.props.appGrid;

    const sortedRows = compose(
      sort.sorter({
        columns: columns,
        sortingColumns,
        sort: orderBy,
        strategy: sort.strategies.byProperty
      })
    )(rows);

    return (
      <Table.Body
        rows={sortedRows}
        rowKey="id"
        onRow={() => {
          return {
            role: 'row'
          };
        }}
      />
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

export default connect(mapStateToProps, mapDispatchToProps)(AppGridRows);
