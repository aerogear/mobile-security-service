import React from 'react';
import { Title, Grid, GridItem } from '@patternfly/react-core';
import Header from './common/Header';
import AppsTableContainer from '../containers/AppsTableContainer';
import './LandingPage.css';

class LandingPage extends React.Component {
  render () {
    return (
      <>
        <Header />
        <Grid className="container">
          <GridItem span={1} />
          <GridItem span={10}>
            <Title className="table-title" size="3xl">Mobile Apps</Title>
            <AppsTableContainer className='table-scroll-x table-clickable-row'/>
          </GridItem>
          <GridItem span={1} />
        </Grid>
      </>
    );
  }
}

export default LandingPage;
