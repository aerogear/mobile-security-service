import React from 'react';
import { BrowserRouter as Router, Switch, Redirect, Route } from 'react-router-dom';
import LandingPage from './components/LandingPage'
import AppView from './components/AppDetailedView'
import Header from './containers/Header';
import '../node_modules/patternfly-react/dist/css/patternfly-react.css'

class App extends React.Component {
  render () {
    return (
      <Router>
        <div className="App">
          <Header />
          <Switch>
            <Route exact path="/overview" component={LandingPage} />
            <Route exact path="/app/:id" component={AppView} />
            {/* Default redirect */}
            <Redirect to="/overview" />
          </Switch>
        </div>
      </Router>
    );
  }
}
export default App;
