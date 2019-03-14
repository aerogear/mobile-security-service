import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Checkbox } from '@patternfly/react-core';
import { getApps, appDetailsSort } from '../actions/actions-ui';
import AppsTable from '../components/AppsTable';
import './TableContainer.css';

export class AppTableContainer extends React.Component {
  onSort = (_event, index, direction) => {
    this.props.appDetailsSort(index, direction);
  };

  getTable = (versions = []) => {
    const renderedRows = [];
    for (let i = 0; i < versions.length; i++) {
      const tempRow = [];

      tempRow[0] = versions[i]['version'];
      tempRow[1] = versions[i]['numOfCurrentInstalls'];
      tempRow[2] = versions[i]['numOfAppLaunches'];
      tempRow[3] = versions[i]['lastLaunched'] || 'Never Launched';

      const cb = (
        <React.Fragment>
          <Checkbox
            label=""
            isChecked={versions[i]['disabled']}
            onChange={this.handleChange}
            aria-label="disable app checkbox"
            id={i.toString()}
          />
        </React.Fragment>
      );
      tempRow[4] = cb;
      tempRow[5] = versions[i]['disabledMessage'] || '';

      renderedRows.push(tempRow);
    }

    return (
      <div className="apps-table">
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
    if (!this.props.appVersions.length) {
      return (
        <div className="no-versions">
          <p>This app has no versions</p>
        </div>
      );
    }

    return this.getTable(this.props.appVersions);
  }
}

AppTableContainer.propTypes = {
  sortBy: PropTypes.object.isRequired,
  columns: PropTypes.array.isRequired
};

function mapStateToProps (state) {
  return {
    sortBy: state.appDetailsSortDirection,
    columns: state.appVersionsColumns,
    appVersions: state.app.deployedVersions
  };
}

export default connect(mapStateToProps, { appDetailsSort, getApps })(AppTableContainer);
