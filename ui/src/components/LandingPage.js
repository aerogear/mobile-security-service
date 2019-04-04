import React from 'react';
import { Title } from '@patternfly/react-core';
import HeaderContainer from '../containers/HeaderContainer';
import AppsTableContainer from '../containers/AppsTableContainer';
import './LandingPage.css';
import Content from './common/Content';

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
