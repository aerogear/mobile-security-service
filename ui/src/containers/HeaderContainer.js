import React, { useState } from 'react';
import { withRouter } from 'react-router-dom';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import config from '../config/config';
import Header from '../components/Header';

/**
 * Container component to manage the Header state
 *
 * @param {object} props Component props
 * @param {string} props.currentUser The current logged in user
 * @param {object} props.history Contains functions to modify the react-router-dom
 */
export const HeaderContainer = ({ currentUser, history }) => {
  const [isDropDownOpen, setIsDropDownOpen] = useState(false);

  const onUserDropdownToggle = () => {
    setIsDropDownOpen(!isDropDownOpen);
  };

  const onTitleClick = () => {
    history.push('/');
  };

  const logout = () => {
    window.location.replace('/oauth/sign_in');
  };

  return (
    <Header
      currentUser={currentUser}
      appName={config.app.name.toUpperCase()}
      isDropDownOpen={isDropDownOpen}
      onUserDropdownToggle={onUserDropdownToggle}
      onTitleClick={onTitleClick}
      logout={logout}
    />
  );
};

HeaderContainer.propTypes = {
  currentUser: PropTypes.string.isRequired,
  history: PropTypes.shape({
    push: PropTypes.func.isRequired
  }).isRequired
};

function mapStateToProps (state) {
  return {
    currentUser: state.header.currentUser
  };
}

export default withRouter(connect(mapStateToProps)(HeaderContainer));
