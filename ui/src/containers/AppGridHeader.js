import React, { Component } from 'react';
import { Table } from 'patternfly-react';
import * as resolve from 'table-resolver';
import { connect } from 'react-redux';

class AppGridHeader extends Component {
  render() {
    const columns = this.props.columns;
    return <Table.Header headerRows={resolve.headerRows({ columns })} />;
  }
}

const mapStateToProps = (state) => {
  return {
    columns: state.appGrid.columns
  };
};

const mapDispatchToProps = (dispatch) => {
  return {};
};

export default connect(mapStateToProps, mapDispatchToProps)(AppGridHeader);
