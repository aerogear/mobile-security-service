import React, { useState, useEffect } from 'react';
import { withRouter } from 'react-router-dom';
import { connect } from 'react-redux';
import PropTypes from 'prop-types';
import config from '../config/config';
import Header from '../components/Header';
import { getUser } from '../actions/actions-ui';

/**
 * Stateful container component to manage the Header state
 *
 * @param {object} props Component props
 * @param {string} props.currentUser The current logged in user
 * @param {object} props.history Contains functions to modify the react-router-dom
 * @param {object} props.getUser Contains functions to modify the react-router-dom
 */
export const HeaderContainer = ({ currentUser, history, getUser }) => {
  const [ isDropDownOpen, setIsDropDownOpen ] = useState(false);

  useEffect(() => {
    getUser();
  }, [getUser]);

  const onUserDropdownToggle = () => {
    setIsDropDownOpen(!isDropDownOpen);
  };

  const onTitleClick = () => {
    history.push('/');
  };

  const logout = () => {
    window.location.replace(config.auth.signInUrl);
  };

  return (
    <Header
      username={currentUser.username}
      appName={config.app.name.toUpperCase()}
      isDropDownOpen={isDropDownOpen}
      onUserDropdownToggle={onUserDropdownToggle}
      onTitleClick={onTitleClick}
      logout={logout}
    />
  );
};

HeaderContainer.propTypes = {
  currentUser: PropTypes.object.isRequired,
  isUserRequestFailed: PropTypes.bool.isRequired,
  getUser: PropTypes.func.isRequired,
  history: PropTypes.shape({
    push: PropTypes.func.isRequired
  }).isRequired
};

function mapStateToProps (state) {
  return {
    currentUser: state.user.data,
    isUserRequestFailed: state.user.isUserRequestFailed
  };
}

export default withRouter(connect(mapStateToProps, { getUser })(HeaderContainer));
