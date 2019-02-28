import React from 'react';
import AppsTable from '../components/AppsTable';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { reverseAppsTableSort } from '../actions/actions-ui';

export class AppsTableContainer extends React.Component {
  constructor (props) {
    super(props);
    this.onSort = this.onSort.bind(this);
    this.onRowClick = this.onRowClick.bind(this);
  }
  onRowClick (event, rowId, props) {
    console.log('On row click called');
    this.setState({ redirect: true });
  }
  onSort (_event, index) {
    this.props.reverseAppsTableSort(index);
  }

  render () {
    return (
      <div className="apps-table">
        <AppsTable columns={this.props.columns} rows={this.props.apps} sortBy={this.sortBy} onSort= {this.onSort} onRowClick={this.onRowClick}/>
      </div>
    );
  }
}

AppsTableContainer.propTypes = {
  apps: PropTypes.array.isRequired,
  sortBy: PropTypes.object.isRequired,
  columns: PropTypes.array.isRequired
};

function mapStateToProps (state) {
  return {
    apps: state.apps,
    sortBy: state.sortBy,
    columns: state.columns
  };
};

export default connect(mapStateToProps, { reverseAppsTableSort })(AppsTableContainer);
