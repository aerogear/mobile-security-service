import React, { useEffect } from 'react';
import { BrowserRouter as Router, Switch, Redirect, Route } from 'react-router-dom';
import LandingPage from '../components/LandingPage';
import PropTypes from 'prop-types';
import AppPageContainer from './appView/AppPageContainer';
import { authenticate } from '../actions/actions-ui';
import { connect } from 'react-redux';
import config from '../config/config';

/**
 * Main entry point for the entire UI view.
 *
 * @param {func} authenticate - Redux action to check if the user is logged in.
 * @param {bool} isLoggedIn - Whether the user is logged in or not.
 * @param {bool} isLoading - Whether the component has finished loading.
 */
export const App = ({ authenticate, isLoggedIn, isLoading }) => {
  useEffect(() => {
    authenticate();
  }, [authenticate]);

  /**
   * Check to see if the user is logged in.
   * If not, redirect them to the login screen.
   */
  const checkIfAuthenticated = () => {
    if (config.mode === 'production' && !isLoading && !isLoggedIn) {
      window.location.replace(config.auth.signInUrl);
    }
  };

  checkIfAuthenticated();

  return (
    <Router>
      <div className="App">
        <Switch>
          <Route exact path="/" component={LandingPage} />
          <Route path="/apps/:id" component={AppPageContainer} />
          {/* Default redirect */}
          <Redirect to="/" />
        </Switch>
      </div>
    </Router>
  );
};

App.propTypes = {
  authenticate: PropTypes.func.isRequired,
  isLoggedIn: PropTypes.bool.isRequired,
  isLoading: PropTypes.bool.isRequired
};

const mapStateToProps = (state) => ({
  isLoggedIn: state.user.isLoggedIn,
  isLoading: state.user.isLoading
});

const mapDispatchToProps = {
  authenticate
};

export default connect(mapStateToProps, mapDispatchToProps)(App);
