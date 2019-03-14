import React from 'react';
import { Title } from '@patternfly/react-core';
import Header from './common/Header';
import AppTableContainer from '../containers/AppTableContainer';
import './AppDetailedView.css';

class AppDetailedView extends React.Component {
  render () {
    return (
      <div className="app-detailed-view">
        <Header />
        <Title className="title" size="2xl">
          Deployed Versions
        </Title>
        <AppTableContainer />
      </div>
    );
  }
}

export default AppDetailedView;
