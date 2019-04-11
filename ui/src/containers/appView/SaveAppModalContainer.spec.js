import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import { SaveAppModalContainer } from './SaveAppModalContainer';
import ConfirmationModal from '../../components/common/ConfirmationModal';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      isOpen: false,
      title: 'title',
      children: [],
      onConfirm: jest.fn(),
      toggleSaveAppModal: jest.fn()
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<SaveAppModalContainer {...props} />);
  const output = renderer.getRenderOutput();

  return {
    props: props,
    output: output
  };
};

describe('components', () => {
  describe('ConfirmationModal', () => {
    it('should render', () => {
      const { output } = setup();
      expect(output.type).toBe(ConfirmationModal);
    });
  });
});

describe('events', () => {
  describe('handles Modal close', () => {
    it('should call toggleSaveAppModal', () => {
      const { output, props } = setup();
      output.props.onClose();
      expect(props.toggleSaveAppModal).toBeCalled();
    });
  });

  describe('handles confirm click', () => {
    it('should call onConfirm', () => {
      const { output, props } = setup();
      const [ confirmButton ] = output.props.confirmAction;
      confirmButton.props.onClick();
      expect(props.onConfirm).toBeCalled();
    });
  });
});
