import React, { Component } from "react";
import Header from './landingpage/Header'

class LandingPage extends Component {
  constructor() {
    super();
    this.state = {};
  }

  render() {
    return (
      <div className="appdetialedview">
        <Header />
      </div>
    );
  }
}

export default LandingPage;
