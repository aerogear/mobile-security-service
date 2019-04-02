import React, { useEffect } from 'react';
import AppsTable from '../components/AppsTable';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';
import { connect } from 'react-redux';
import { getApps, appsTableSort } from '../actions/actions-ui';
import './TableContainer.css';

const AppsTableContainer = ({ apps, sortBy, columns, isAppsRequestFailed, getApps, appsTableSort, history, className }) => {
  useEffect(() => {
    getApps();
  }, []);

  const onRowClick = (_event, rowId) => {
    var app = apps.data.filter((app) => {
      return app.appName === rowId[0];
    });
    const id = app[0].id;
    const path = '/apps/' + id;
    history.push(path);
  };

  const onSort = (_event, index, direction) => {
    appsTableSort(index, direction);
  };

  const getTable = () => {
    return (
      <div className={className}>
        <AppsTable
          columns={columns}
          rows={apps.rows}
          sortBy={sortBy}
          onSort={onSort}
          onRowClick={onRowClick}
        />
      </div>
    );
  };

  if (isAppsRequestFailed) {
    return (
      <div className="no-apps">
        <p>Unable to fetch any apps :/</p>
      </div>
    );
  }
  return getTable();
};

AppsTableContainer.propTypes = {
  apps: PropTypes.object.isRequired,
  sortBy: PropTypes.object.isRequired,
  columns: PropTypes.array.isRequired,
  isAppsRequestFailed: PropTypes.bool.isRequired,
  getApps: PropTypes.func.isRequired,
  className: PropTypes.string.isRequired
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
