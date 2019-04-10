import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import ConfirmationModal from './ConfirmationModal';
import { Modal } from '@patternfly/react-core';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      isLarge: true,
      title: 'title',
      isOpen: false,
      onClose: jest.fn(),
      confirmAction: [],
      children: <div />
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<ConfirmationModal {...props} />);
  const output = renderer.getRenderOutput();

  return {
    props: props,
    output: output
  };
};

describe('components', () => {
  describe('Modal', () => {
    it('should render', () => {
      const { output } = setup();
      expect(output.type).toBe(Modal);
    });
  });
});
