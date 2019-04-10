import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import Content from './Content';
import { Grid, GridItem } from '@patternfly/react-core';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      className: 'content',
      children: <div />
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<Content {...props} />);
  const output = renderer.getRenderOutput();

  return {
    props: props,
    output: output
  };
};

describe('components', () => {
  describe('Content', () => {
    it('should render', () => {
      const { output } = setup();
      expect(output.type).toBe(Grid);
      expect(output.props.className).toBe('content');

      const [ leftGutter, mainContent, rightGutter ] = output.props.children;

      expect(leftGutter.type).toBe(GridItem);
      expect(leftGutter.props.span).toBe(1);

      expect(mainContent.type).toBe(GridItem);
      expect(mainContent.props.span).toBe(10);
      expect(mainContent.props.children.type).toBe('div');

      expect(rightGutter.type).toBe(GridItem);
      expect(rightGutter.props.span).toBe(1);
    });
  });
});
