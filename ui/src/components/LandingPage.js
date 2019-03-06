import React from 'react';
import Header from './common/Header';
import AppsTableContainer from '../containers/AppsTableContainer';

class LandingPage extends React.Component {
  constructor () {
    super();
    this.state = {};
  }

  render () {
    return (
      <>
        <Header />
        <AppsTableContainer />
      </>
    );
  }
}

export default LandingPage;
