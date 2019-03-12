import React from 'react';
import { Title } from '@patternfly/react-core';
import Header from './common/Header';
import AppsTableContainer from '../containers/AppsTableContainer';
import './LandingPage.css';

class LandingPage extends React.Component {
  constructor () {
    super();
    this.state = {};
  }

  render () {
    return (
      <>
        <Header />
        <Title className="title" size="3xl">Mobile Apps</Title>
        <AppsTableContainer />
      </>
    );
  }
}

export default LandingPage;
