import React from 'react';
import PropTypes from 'prop-types';
import { Grid, GridItem } from '@patternfly/react-core';

/**
 * Stateless component that is a container for the main body of the application
 *
 * @param {object} props Component props
 * @param {node} props.children Inner HTML and React elements
 * @param {String} props.className Optionally provide a custom class for the component
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
