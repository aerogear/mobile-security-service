import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Checkbox, TextInput } from '@patternfly/react-core';
import moment from 'moment';
import { getApps, appDetailsSort, updateDisabledAppVersion, updateVersionCustomMessage } from '../actions/actions-ui';
import AppsTable from '../components/AppsTable';
import './TableContainer.css';

export class AppVersionsTableContainer extends React.Component {
  handleDisableAppVersionChange = (_event, e) => {
    const id = e.target.id;
    const isDisabled = e.target.checked;
    this.props.updateDisabledAppVersion(id, isDisabled);
  };

  handleCustomMessageInputChange = (_event, e) => {
    const id = e.target.id;
    const value = e.target.value;
    this.props.updateVersionCustomMessage(id, value);
  };

  onSort = (_event, index, direction) => {
    this.props.appDetailsSort(index, direction);
  };

  createCheckbox = (id, checked) => {
    return (
      <React.Fragment>
        <Checkbox
          label=""
          isChecked={checked}
          onChange={this.handleDisableAppVersionChange}
          aria-label="disable app checkbox"
          id={id}
        />
      </React.Fragment>
    );
  };

  createTextInput = (id, text) => {
    return (
      <React.Fragment>
        <TextInput
          id={id}
          type="text"
          placeholder="Add a custom message.."
          value={text}
          onChange={this.handleCustomMessageInputChange}
          aria-label="Custom Disable Message"
        />
      </React.Fragment>
    );
  };

  getTable = (versions = []) => {
    const renderedRows = [];
    for (let i = 0; i < versions.length; i++) {
      const tempRow = [];
      tempRow[0] = versions[i][0];
      tempRow[1] = versions[i][1];
      tempRow[2] = versions[i][2];
      if (versions[i][3].isNullOrUndefined || versions[i][3] === 'Never Launched') {
        tempRow[3] = 'Never Launched';
      } else {
        tempRow[3] = moment(versions[i][3]).format('YYYY-MM-DD HH:mm:ss');
      }
      tempRow[4] = this.createCheckbox(versions[i][0].toString(), versions[i][4]);
      tempRow[5] = this.createTextInput(versions[i][0], versions[i][5]);
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

const mapDispatchToProps = {
  appDetailsSort,
  getApps,
  updateDisabledAppVersion,
  updateVersionCustomMessage
};

export default connect(mapStateToProps, mapDispatchToProps)(AppVersionsTableContainer);
