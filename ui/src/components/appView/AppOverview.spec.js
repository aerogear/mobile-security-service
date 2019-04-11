import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import AppOverview from './AppOverview';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      app: {
        appId: 'my-app'
      },
      className: 'app-overview'
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<AppOverview {...props} />);
  const output = renderer.getRenderOutput();

  return {
    props: props,
    output: output
  };
};

describe('components', () => {
  describe('AppOverview', () => {
    it('should render', () => {
      const { output } = setup();
      expect(output.type).toBe('div');
      expect(output.props.className).toBe('app-overview');
    });
  });
});
