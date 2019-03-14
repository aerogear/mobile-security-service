import React from 'react';
import { Title } from '@patternfly/react-core';
import Header from './common/Header';
import AppTableContainer from '../containers/AppTableContainer';
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
        <AppTableContainer />
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
    app: state.app,
    getAppById: PropTypes.func.isRequired
  };
};

const mapDispatchToProps = {
  getAppById
};

export default connect(mapStateToProps, mapDispatchToProps)(AppDetailedView);
