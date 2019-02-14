import React, { Component } from 'react';
import { connect } from 'react-redux';

import AppGridHeader from './AppGridHeader';
import AppGridRows from './AppGridRows';
import { customHeaderFormattersDefinition, Table } from 'patternfly-react';

class AppGrid extends Component {
  constructor(props) {
    super(props);

    // enables our custom header formatters extensions to reactabular
    this.customHeaderFormatters = customHeaderFormattersDefinition;
  }

  render() {
    const { sortingColumns, columns } = this.props.appGrid;

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
          <AppGridRows />
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
