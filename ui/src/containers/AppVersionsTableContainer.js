import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Checkbox } from '@patternfly/react-core';
import { getApps, appDetailsSort } from '../actions/actions-ui';
import AppsTable from '../components/AppsTable';
import './TableContainer.css';

export class AppVersionsTableContainer extends React.Component {
  onSort = (_event, index, direction) => {
    this.props.appDetailsSort(index, direction);
  };

  getTable = (versions = []) => {
    const renderedRows = [];
    for (let i = 0; i < versions.length; i++) {
      const tempRow = [];

      tempRow[0] = versions[i][0];
      tempRow[1] = versions[i][1];
      tempRow[2] = versions[i][2];
      tempRow[3] = versions[i][3];

      const cb = (
        <React.Fragment>
          <Checkbox
            label=""
            isChecked={versions[i][4]}
            onChange={this.handleChange}
            aria-label="disable app checkbox"
            id={i.toString()}
          />
        </React.Fragment>
      );
      tempRow[4] = cb;
      tempRow[5] = versions[i][5];

      renderedRows.push(tempRow);
    }

    return (

      <div className={this.props.className}>
        <AppsTable
          columns={this.props.columns}
          rows={renderedRows}
          sortBy={this.props.sortBy}
          onSort={this.onSort}
          onRowClick={this.onRowClick}
        />
      </div>
    );
  };

  render () {
    if (!this.props.appVersions || !this.props.appVersions.length) {
      return (
        <div className="no-versions">
          <p>This app has no versions</p>
        </div>
      );
    }

    return this.getTable(this.props.appVersions);
  }
}

AppVersionsTableContainer.propTypes = {
  sortBy: PropTypes.object.isRequired,
  columns: PropTypes.array.isRequired,
  appVersions: PropTypes.array.isRequired
};

function mapStateToProps (state) {
  return {
    sortBy: state.appVersionsSortDirection,
    columns: state.appVersionsColumns,
    appVersions: state.app.versionsRows
  };
}

export default connect(mapStateToProps, { appDetailsSort, getApps })(AppVersionsTableContainer);
