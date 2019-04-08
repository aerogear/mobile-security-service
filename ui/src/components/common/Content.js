import React from 'react';
import { Grid, GridItem } from '@patternfly/react-core';
import PropTypes from 'prop-types';

/**
 * Stateless component that is a container for the main body of the application
 *
 * @param {*} children - Inner HTML and React elements
 * @param {string} className - Optionally provide a custom class for the component
 */
const Content = ({ className, children }) => (
  <Grid className={className}>
    <GridItem span={1} />
    <GridItem span={10}>
      {children}
    </GridItem>
    <GridItem span={1} />
  </Grid>
);

Content.propTypes = {
  className: PropTypes.string.isRequired,
  children: PropTypes.node.isRequired
};

export default Content;
