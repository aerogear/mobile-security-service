import React from 'react';
import AppsTable from '../components/AppsTable';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';
import { connect } from 'react-redux';
import { getApps, appsTableSort } from '../actions/actions-ui';
import './TableContainer.css';

export class AppsTableContainer extends React.Component {
  constructor (props) {
    super(props);
    this.onSort = this.onSort.bind(this);
    this.onRowClick = this.onRowClick.bind(this);
    this.getTable = this.getTable.bind(this);
  }
  componentWillMount () {
    this.props.getApps();
  }
  onRowClick (_event, rowId) {
    var app = this.props.apps.data.filter((app) => {
      return app.appName === rowId[0];
    });
    const id = app[0].id;
    const path = '/apps/' + id;
    this.props.history.push(path);
  }
  onSort (_event, index, direction) {
    this.props.appsTableSort(index, direction);
  }

  getTable () {
    return (
      <div className={this.props.className}>
        <AppsTable
          columns={this.props.columns}
          rows={this.props.apps.rows}
          sortBy={this.props.sortBy}
          onSort={this.onSort}
          onRowClick={this.onRowClick}
        />
      </div>
    );
  }

  render () {
    if (this.props.isAppsRequestFailed) {
      return (
        <div className="no-apps">
          <p>Unable to fetch any apps :/</p>
        </div>
      );
    }
    return this.getTable();
  }
}

AppsTableContainer.propTypes = {
  apps: PropTypes.object.isRequired,
  sortBy: PropTypes.object.isRequired,
  columns: PropTypes.array.isRequired,
  isAppsRequestFailed: PropTypes.bool.isRequired,
  getApps: PropTypes.func.isRequired
};

function mapStateToProps (state) {
  return {
    apps: state.apps,
    sortBy: state.sortBy,
    columns: state.columns,
    isAppsRequestFailed: state.isAppsRequestFailed
  };
}

export default withRouter(connect(mapStateToProps, { appsTableSort, getApps })(AppsTableContainer));
