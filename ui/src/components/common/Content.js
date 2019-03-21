import React from 'react';
import { Grid, GridItem } from '@patternfly/react-core';

/**
 * Stateless component that is a container for the main body of the application
 *
 * @param {*} children - Inner HTML and React elements
 * @param {String} className - Optionally provide a custom class for the component
 */
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
