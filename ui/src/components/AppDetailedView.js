import React from 'react';
import Header from './common/Header';
import { getAppById } from '../actions/actions-ui';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';

class AppDetailedView extends React.Component {
  constructor (props) {
    super(props);
    this.state = {};
  }

  componentDidMount () {
    this.props.getAppById(this.props.match.params.id);
  }

  render () {
    return (
      <div className="app-detailed-view">
        <Header />
      </div>
    );
  };
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
