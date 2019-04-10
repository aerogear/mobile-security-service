import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import { BrowserRouter as Router } from 'react-router-dom';
import App from './App';

const setup = (propOverrides) => {
  const props = Object.assign(
    {},
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<App {...props} />);
  const output = renderer.getRenderOutput();

  return {
    props: props,
    output: output
  };
};

describe('components', () => {
  describe('App', () => {
    it('should render', () => {
      const { output } = setup();
      expect(output.type).toBe(Router);
    });
  });
});
