import React from 'react';
import AppsTableContainer from '../containers/AppsTableContainer';

class LandingPage extends React.Component {
  constructor () {
    super();
    this.state = {};
  }

  render () {
    return (
      <AppsTableContainer />
    );
  }
}

export default LandingPage;
