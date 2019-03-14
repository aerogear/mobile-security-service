import React from 'react';
import { BrowserRouter as Router, Switch, Redirect, Route } from 'react-router-dom';

import LandingPage from './LandingPage';
import AppDetailedView from './AppDetailedView';

class App extends React.Component {
  render () {
    return (
      <Router>
        <div className="App">
          <Switch>
            <Route exact path="/" component={LandingPage} />
            <Route path="/apps/:id" component={AppDetailedView} />
            {/* Default redirect */}
            <Redirect to="/" />
          </Switch>
        </div>
      </Router>
    );
  }
}

export default App;
