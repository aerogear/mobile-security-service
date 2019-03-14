import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Checkbox } from '@patternfly/react-core';
import { getApps, appDetailsSort } from '../actions/actions-ui';
import AppsTable from '../components/AppsTable';
import './TableContainer.css';

export class AppTableContainer extends React.Component {
  componentWillMount () {}

  handleChange = (e) => {
    console.log('checkbox clicked');
  };

  onSort = (_event, index, direction) => {
    this.props.appDetailsSort(index, direction);
  };

  getTable = () => {
    const renderedRows = [];
    for (let i = 0; i < this.props.app.length; i++) {
      const tempRow = [];
      tempRow[0] = this.props.app[i][0];
      tempRow[1] = this.props.app[i][1];
      tempRow[2] = this.props.app[i][2];
      tempRow[3] = this.props.app[i][3];
      const cb = (
        <React.Fragment>
          <Checkbox
            label=""
            isChecked={this.props.app[i][4]}
            onChange={this.handleChange}
            aria-label="disable app checkbox"
            id={i.toString()}
          />
        </React.Fragment>
      );
      tempRow[4] = cb;
      tempRow[5] = this.props.app[i][5];
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
    if (this.props.isAppsRequestFailed) {
      return (
        <div className="no-apps">
          <p>Unable to fetch any apps</p>
        </div>
      );
    }
    return this.getTable();
  }
}

AppTableContainer.propTypes = {
  app: PropTypes.array.isRequired,
  sortBy: PropTypes.object.isRequired,
  columns: PropTypes.array.isRequired,
  isAppsRequestFailed: PropTypes.bool.isRequired
};

function mapStateToProps (state) {
  return {
    app: state.appDetailRows,
    sortBy: state.appDetailsSortDirection,
    columns: state.appDetailColumns,
    isAppsRequestFailed: state.isAppsRequestFailed
  };
}

export default connect(mapStateToProps, { appDetailsSort, getApps })(AppTableContainer);
