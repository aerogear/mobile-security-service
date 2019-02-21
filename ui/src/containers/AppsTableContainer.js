import React from 'react';
import { sortable, SortByDirection } from '@patternfly/react-table';
import AppsTable from '../components/AppsTable';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { getApps } from '../actions/actions';

class AppsTableContainer extends React.Component {
  constructor (props) {
    super(props);
    this.sortBy = {};
    this.onSort = this.onSort.bind(this);
    this.onRowClick = this.onRowClick.bind(this);
  }
  componentDidMount () {
    this.props.getApps();
  }
  onRowClick (event, rowId, props) {
    this.setState({ redirect: true });
  }
  onSort (_event, index, direction) {
    const sortedRows = this.props.apps.sort((a, b) => (a[index] < b[index] ? -1 : a[index] > b[index] ? 1 : 0));
    this.setState({
      sortBy: {
        index,
        direction
      },
      apps: direction === SortByDirection.asc ? sortedRows : sortedRows.reverse()
    });
  }

  render () {
    const columns = [
      { title: '', transforms: [ sortable ] },
      { title: 'Deployed Versions', transforms: [ sortable ] },
      { title: 'Current Installs', transforms: [ sortable ] },
      { title: 'Launches', transforms: [ sortable ] }
    ];
    const rows = this.props.apps;
    return (
      <AppsTable columns={columns} rows={rows} sortBy={this.sortBy} onSort= {this.onSort} onRowClick={this.onRowClick}/>
    );
  }
}

AppsTableContainer.propTypes = {
  getApps: PropTypes.func.isRequired,
  apps: PropTypes.array.isRequired
};

function mapStateToProps (state) {
  return {
    apps: state.apps
  };
};

export default connect(mapStateToProps, { getApps })(AppsTableContainer);
