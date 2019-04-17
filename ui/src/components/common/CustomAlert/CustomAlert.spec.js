import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import CustomAlert from './CustomAlert';
import { Alert } from '@patternfly/react-core';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      title: '',
      visible: false,
      onClose: jest.fn(),
      children: ''
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<CustomAlert {...props} />);
  const output = renderer.getRenderOutput();

  return {
    props,
    output
  };
};

describe('components', () => {
  describe('CustomAlert', () => {
    it('should render', () => {
      const { output, props } = setup({ visible: true, title: 'Alert Title', variant: 'info' });

      expect(output.type).toBe(Alert);
      expect(output.props.className).toBe('Alert');
      expect(output.props.title).toBe(props.title);
      expect(output.props.variant).toBe(props.variant);
      expect(output.props.children).toBe('');
    });

    it('should not render', () => {
      const { output } = setup({ title: 'Alert Title', variant: 'info' });
      expect(output).toBeNull();
    });
  });
});
