import React from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { Checkbox, TextInput } from '@patternfly/react-core';
import { sortable, cellWidth } from '@patternfly/react-table';
import moment from 'moment';
import { appDetailsSort, updateDisabledAppVersion, updateVersionCustomMessage } from '../actions/actions-ui';
import AppsTable from '../components/AppsTable';
import { getSortedAppVersionTableRows } from '../selectors/index';
import './TableContainer.css';
import config from '../config/config';

const AppVersionsTableContainer = ({ className, sortBy, appVersionRows, appDetailsSort, updateDisabledAppVersion, updateVersionCustomMessage }) => {
  const columns = [
    { title: 'APP VERSION', transforms: [sortable, cellWidth(10)] },
    { title: 'CURRENT INSTALLS', transforms: [sortable, cellWidth(10)] },
    { title: 'LAUNCHES', transforms: [sortable, cellWidth(10)] },
    { title: 'LAST LAUNCHED', transforms: [sortable, cellWidth(15)] },
    { title: 'DISABLE ON STARTUP', transforms: [sortable, cellWidth(10)] },
    { title: 'CUSTOM DISABLE MESSAGE', transforms: [sortable, cellWidth('max')] }
  ];

  const handleDisableAppVersionChange = (value, e) => {
    const id = e.target.id;
    updateDisabledAppVersion(id, value);
  };

  const handleCustomMessageInputChange = (value, e) => {
    const id = e.target.id;
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
      tempRow[4] = createCheckbox(versions[i][6].toString(), versions[i][4]);
      tempRow[5] = createTextInput(versions[i][6], versions[i][5]);
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

  if (!appVersionRows || !appVersionRows.length) {
    return (
      <div className="no-versions">
        <p>This app has no versions</p>
      </div>
    );
  }

  return getTable(appVersionRows);
};

AppVersionsTableContainer.propTypes = {
  className: PropTypes.string.isRequired,
  sortBy: PropTypes.shape({
    direction: PropTypes.string.isRequired,
    index: PropTypes.number.isRequired
  }).isRequired,
  appVersionRows: PropTypes.array.isRequired,
  appDetailsSort: PropTypes.func.isRequired,
  updateDisabledAppVersion: PropTypes.func.isRequired,
  updateVersionCustomMessage: PropTypes.func.isRequired
};

function mapStateToProps (state) {
  return {
    sortBy: state.app.sortBy,
    appVersionRows: getSortedAppVersionTableRows(state, state.app.sortBy, state.app.sortBy)
  };
}

const mapDispatchToProps = {
  appDetailsSort,
  updateDisabledAppVersion,
  updateVersionCustomMessage
};

export default connect(mapStateToProps, mapDispatchToProps)(AppVersionsTableContainer);
