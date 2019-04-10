import React from 'react';
import { Title } from '@patternfly/react-core';
import HeaderContainer from '../containers/HeaderContainer';
import AppsTableContainer from '../containers/AppsTableContainer';
import Content from './common/Content';
import './LandingPage.css';

/**
 * Initial landing page of the UI. First view the user sees
 */
const LandingPage = () => {
  return (
    <>
      <HeaderContainer />
      <Content className='container'>
        <Title className="table-title" size="3xl">Mobile Apps</Title>
        <AppsTableContainer className='table-scroll-x table-clickable-row' />
      </Content>
    </>
  );
};

export default LandingPage;
