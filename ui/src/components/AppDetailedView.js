import React from 'react';
import { Title } from '@patternfly/react-core';
import Header from './common/Header';
import AppVersionsTableContainer from '../containers/AppVersionsTableContainer';
import './AppDetailedView.css';
import { getAppById } from '../actions/actions-ui';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';

class AppDetailedView extends React.Component {
  componentWillMount () {
    this.props.getAppById(this.props.match.params.id);
  }

  render () {
    return (
      <div className="app-detailed-view">
        <Header />
        <Title className="title" size="2xl">
          Deployed Versions
        </Title>
        <AppVersionsTableContainer />
      </div>
    );
  }
}

AppDetailedView.propTypes = {
  app: PropTypes.object.isRequired,
  getAppById: PropTypes.func.isRequired
};

function mapStateToProps (state) {
  return {
    app: state.app.data,
    getAppById: PropTypes.func.isRequired
  };
};

const mapDispatchToProps = {
  getAppById
};

export default connect(mapStateToProps, mapDispatchToProps)(AppDetailedView);
