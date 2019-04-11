import React, { useEffect } from 'react';
import TableView from '../components/common/TableView';
import PropTypes from 'prop-types';
import { withRouter } from 'react-router-dom';
import { connect } from 'react-redux';
import { sortable } from '@patternfly/react-table';
import { getApps, appsTableSort } from '../actions/actions-ui';
import { getSortedAppsTableRows } from '../selectors/index';

/**
 * Stateful container component to manage state and display the Apps table for the landing page
 *
 * @param {object} props Component props
 * @param {array} props.apps Array of app object to display
 * @param {array} props.appRows Transformed app array of arrays contains just the values needed for the table
 * @param {object} props.sortBy Contains column index and direction for sorting
 * @param {boolean} props.isAppsRequestFailed Boolean on if the apps request has failed
 * @param {func} props.getApps Retrieves apps from the server
 * @param {func} props.appsTableSort Triggers redux action creator for the table sort
 * @param {object} props.history Contains functions to modify the react-router-dom
 * @param {string} props.className Optionally provide a custom class for the component
 */
export const AppsTableContainer = ({
  apps,
  appRows,
  sortBy,
  isAppsRequestFailed,
  getApps,
  appsTableSort,
  history,
  className
}) => {
  const columns = [
    { title: 'APP NAME', transforms: [ sortable ] },
    { title: 'APP ID', transforms: [ sortable ] },
    { title: 'DEPLOYED VERSIONS', transforms: [ sortable ] },
    { title: 'CURRENT INSTALLS', transforms: [ sortable ] },
    { title: 'LAUNCHES', transforms: [ sortable ] }
  ];

  useEffect(() => {
    getApps();
  }, []);

  const onRowClick = (_event, rowId) => {
    var app = apps.filter((app) => {
      return app.appName === rowId[0];
    });
    const id = app[0].id;
    const path = '/apps/' + id;
    history.push(path);
  };

  const getTable = () => {
    return (
      <div className={className}>
        <TableView columns={columns} rows={appRows} sortBy={sortBy} onSort={appsTableSort} onRowClick={onRowClick} />
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
  apps: PropTypes.arrayOf(
    PropTypes.shape({
      appId: PropTypes.string.isRequired,
      appName: PropTypes.string.isRequired,
      id: PropTypes.string.isRequired,
      numOfAppLaunches: PropTypes.number.isRequired,
      numOfCurrentInstalls: PropTypes.number.isRequired,
      numOfDeployedVersions: PropTypes.number.isRequired
    })
  ),
  appRows: PropTypes.array.isRequired,
  sortBy: PropTypes.object.isRequired,
  isAppsRequestFailed: PropTypes.bool.isRequired,
  getApps: PropTypes.func.isRequired,
  appsTableSort: PropTypes.func.isRequired,
  className: PropTypes.string.isRequired,
  history: PropTypes.shape({
    push: PropTypes.func.isRequired
  }).isRequired
};

function mapStateToProps (state, props) {
  return {
    apps: state.apps.data,
    appRows: getSortedAppsTableRows(state, state.apps.sortBy),
    sortBy: state.apps.sortBy,
    isAppsRequestFailed: state.apps.isAppsRequestFailed
  };
}

export default withRouter(connect(mapStateToProps, { appsTableSort, getApps })(AppsTableContainer));
