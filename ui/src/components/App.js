import React from 'react';
import { BrowserRouter as Router, Switch, Redirect, Route } from 'react-router-dom';
import LandingPage from './LandingPage';
import AppPageContainer from '../containers/appView/AppPageContainer';

const App = () => {
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

export default App;
