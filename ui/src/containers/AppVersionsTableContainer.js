import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Checkbox, TextInput } from '@patternfly/react-core';
import moment from 'moment';
import { appDetailsSort, updateDisabledAppVersion, updateVersionCustomMessage } from '../actions/actions-ui';
import AppsTable from '../components/AppsTable';
import './TableContainer.css';
import config from '../config/config';

const AppVersionsTableContainer = ({ className, sortBy, columns, appVersions, appDetailsSort, updateDisabledAppVersion, updateVersionCustomMessage }) => {
  const handleDisableAppVersionChange = (_event, e) => {
    const id = e.target.id;
    const isDisabled = e.target.checked;
    updateDisabledAppVersion(id, isDisabled);
  };

  const handleCustomMessageInputChange = (_event, e) => {
    const id = e.target.id;
    const value = e.target.value;
    updateVersionCustomMessage(id, value);
  };

  const onSort = (_event, index, direction) => {
    appDetailsSort(index, direction);
  };

  const createCheckbox = (id, checked) => {
    return (
      <React.Fragment>
        <Checkbox
          label=""
          isChecked={checked}
          onChange={handleDisableAppVersionChange}
          aria-label="disable app checkbox"
          id={id}
        />
      </React.Fragment>
    );
  };

  const createTextInput = (id, text) => {
    return (
      <React.Fragment>
        <TextInput
          id={id}
          type="text"
          placeholder="Add a custom message.."
          value={text}
          onChange={handleCustomMessageInputChange}
          aria-label="Custom Disable Message"
        />
      </React.Fragment>
    );
  };

  const getTable = (versions = []) => {
    const renderedRows = [];
    for (let i = 0; i < versions.length; i++) {
      const tempRow = [];
      tempRow[0] = versions[i][0];
      tempRow[1] = versions[i][1];
      tempRow[2] = versions[i][2];
      if (versions[i][3].isNullOrUndefined || versions[i][3] === 'Never Launched') {
        tempRow[3] = 'Never Launched';
      } else {
        tempRow[3] = moment(versions[i][3]).format(config.dateTimeFormat);
      }
      tempRow[4] = createCheckbox(versions[i][0].toString(), versions[i][4]);
      tempRow[5] = createTextInput(versions[i][0], versions[i][5]);
      renderedRows.push(tempRow);
    }

    return (

      <div className={className}>
        <AppsTable
          columns={columns}
          rows={renderedRows}
          sortBy={sortBy}
          onSort={onSort}
        />
      </div>
    );
  };

  if (!appVersions || !appVersions.length) {
    return (
      <div className="no-versions">
        <p>This app has no versions</p>
      </div>
    );
  }

  return getTable(appVersions);
};

AppVersionsTableContainer.propTypes = {
  className: PropTypes.string.isRequired,
  sortBy: PropTypes.shape({
    direction: PropTypes.string.isRequired,
    index: PropTypes.number.isRequired
  }).isRequired,
  columns: PropTypes.array.isRequired,
  appVersions: PropTypes.array.isRequired,
  appDetailsSort: PropTypes.func.isRequired,
  updateDisabledAppVersion: PropTypes.func.isRequired,
  updateVersionCustomMessage: PropTypes.func.isRequired
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
  updateDisabledAppVersion,
  updateVersionCustomMessage
};

export default connect(mapStateToProps, mapDispatchToProps)(AppVersionsTableContainer);
