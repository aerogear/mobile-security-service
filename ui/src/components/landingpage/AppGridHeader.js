import React, { Component } from 'react';
import { Table } from 'patternfly-react';
import * as resolve from 'table-resolver';
import { connect } from 'react-redux';

class AppGridHeader extends Component {
  render() {
    const columns = this.props.appGrid.columns;
    console.log('columns', columns);
    return <Table.Header headerRows={resolve.headerRows({ columns })} />;
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

export default connect(mapStateToProps, mapDispatchToProps)(AppGridHeader);
