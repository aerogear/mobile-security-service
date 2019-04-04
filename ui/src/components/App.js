import React from 'react';
import { BrowserRouter as Router, Switch, Redirect, Route } from 'react-router-dom';
import LandingPage from './LandingPage';
import AppViewContainer from '../containers/appView/AppViewContainer';

const App = () => {
  return (
    <Router>
      <div className="App">
        <Switch>
          <Route exact path="/" component={LandingPage} />
          <Route path="/apps/:id" component={AppViewContainer} />
          {/* Default redirect */}
          <Redirect to="/" />
        </Switch>
      </div>
    </Router>
  );
};

export default App;
