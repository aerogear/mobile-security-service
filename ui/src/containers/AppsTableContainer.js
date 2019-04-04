import React, { useEffect } from 'react';
import AppsTable from '../components/AppsTable';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';
import { connect } from 'react-redux';
import { sortable } from '@patternfly/react-table';
import { getApps, appsTableSort } from '../actions/actions-ui';
import { getAppsTableRows, getSortedTableRows } from '../reducers/apps';
import './TableContainer.css';

const AppsTableContainer = ({ apps, appRows, sortBy, isAppsRequestFailed, getApps, appsTableSort, history, className }) => {
  const columns = [
    { title: 'APP NAME', transforms: [sortable] },
    { title: 'APP ID', transforms: [sortable] },
    { title: 'DEPLOYED VERSIONS', transforms: [sortable] },
    { title: 'CURRENT INSTALLS', transforms: [sortable] },
    { title: 'LAUNCHES', transforms: [sortable] }
  ];

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
          rows={appRows}
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
  apps: PropTypes.shape({
    data: PropTypes.arrayOf(PropTypes.shape({
      appId: PropTypes.string.isRequired,
      appName: PropTypes.string.isRequired,
      id: PropTypes.string.isRequired,
      numOfAppLaunches: PropTypes.number.isRequired,
      numOfCurrentInstalls: PropTypes.number.isRequired,
      numOfDeployedVersions: PropTypes.number.isRequired
    }))
  }),
  sortBy: PropTypes.object.isRequired,
  isAppsRequestFailed: PropTypes.bool.isRequired,
  getApps: PropTypes.func.isRequired,
  className: PropTypes.string.isRequired,
  history: PropTypes.shape({
    push: PropTypes.func.isRequired
  }).isRequired
};

function mapStateToProps (state) {
  return {
    apps: state.apps,
    appRows: getSortedTableRows(getAppsTableRows(state.apps.data), state.apps.sortBy.index, state.apps.sortBy.direction),
    sortBy: state.apps.sortBy,
    isAppsRequestFailed: state.apps.isAppsRequestFailed
  };
}

export default withRouter(connect(mapStateToProps, { appsTableSort, getApps })(AppsTableContainer));
