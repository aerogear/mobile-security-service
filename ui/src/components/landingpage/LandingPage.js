import React, { Component } from 'react';
import Header from '../../containers/Header';

class LandingPage extends Component {
  constructor () {
    super();
    this.state = {};
  }
  render () {
    return (
      <div className="landingPage">
        <Header />
      </div>
    );
  }
}

export default LandingPage;
