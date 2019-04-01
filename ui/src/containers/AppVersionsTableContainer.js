import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Checkbox, TextInput } from '@patternfly/react-core';
import moment from 'moment';
import { getApps, appDetailsSort } from '../actions/actions-ui';
import AppsTable from '../components/AppsTable';
import './TableContainer.css';
import config from '../config/config';

export class AppVersionsTableContainer extends React.Component {
  shouldComponentUpdate = () => {
    return this.state.shouldAppUpdate;
  }
  componentWillMount = () => {
    this.setState({
      ...this.state,
      shouldAppUpdate: false
    });
  }
  componentDidMount = () => {
    this.setState({
      ...this.state,
      updatedVersions: this.props.appVersions
    });
  }
  handleDisableAppVersionChange = (_event, e) => {
    const id = e.target.id;
    const isDisabled = e.target.checked;
    this.state.updatedVersions.forEach(version => {
      if (version.versionNum === id) {
        version.isDisable = isDisabled;
      }
    });
    this.setState({
      ...this.state,
      shouldAppUpdate: false,
      updatedVersions: this.state.updatedVersions
    });
  };

  handleCustomMessageInputChange = (_event, e) => {
    const id = e.target.id;
    const value = e.target.value;
    this.state.updatedVersions.map(version => {
      if (version.versionNum === id) {
        version.disabledMessage = value;
      }
      return version;
    });
    this.setState({
      ...this.state,
      shouldAppUpdate: false
    });
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
          defaultValue={text === null ? undefined : text}
          onChange={this.handleCustomMessageInputChange}
          aria-label="Custom Disable Message"
        />
      </React.Fragment>
    );
  };

  getTable = (versions = []) => {
    const renderedRows = [];
    versions.map(version => {
      const tempRow = [];
      const versionNum = version.versionNum;
      tempRow[0] = versionNum;
      tempRow[1] = version.numOfAppLaunches;
      tempRow[2] = version.currentInstalls;
      const lastLaunched = version.lastLaunchedAt;
      if (lastLaunched.isNullOrUndefined || lastLaunched === 'Never Launched') {
        tempRow[3] = 'Never Launched';
      } else {
        tempRow[3] = moment(lastLaunched).format(config.dateTimeFormat);
      }
      tempRow[4] = this.createCheckbox(versionNum, version.isDisabled);
      tempRow[5] = this.createTextInput(versionNum, version.disabledMessage);
      renderedRows.push(tempRow);
      return version;
    });
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
    appVersions: state.app.versionsRows,
    app: state.app.data
  };
}

const mapDispatchToProps = {
  appDetailsSort,
  getApps
};

export default connect(mapStateToProps, mapDispatchToProps)(AppVersionsTableContainer);
