import React, { Component } from 'react';
import Header from '../../containers/Header';
import AppGrid from '../../containers/AppGrid';

class LandingPage extends Component {
  constructor () {
    super();
    this.state = {};
  }

  render () {
    return (
      <div className="landingPage">
        <Header />
        <AppGrid />
      </div>
    );
  }
}

export default LandingPage;
