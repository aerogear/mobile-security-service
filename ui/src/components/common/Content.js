import React from 'react';
import { Grid, GridItem } from '@patternfly/react-core';

const Content = ({ children, className }) => (
  <Grid className={className}>
    <GridItem span={1} />
    <GridItem span={10}>
      {children}
    </GridItem>
    <GridItem span={1} />
  </Grid>
);

export default Content;
