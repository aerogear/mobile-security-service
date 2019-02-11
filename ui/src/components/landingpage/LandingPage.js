import React, { Component } from "react";
import Header from './Header'
import AppGrid from "./AppGrid";

class LandingPage extends Component {
  constructor() {
    super();
    this.state = {};
  }

  render() {
    return (
      <div className="landingPage">
        <Header />
        <AppGrid />
      </div>
    );
  }
}

export default LandingPage;
