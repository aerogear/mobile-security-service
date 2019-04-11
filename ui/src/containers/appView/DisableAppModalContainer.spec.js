import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import { DisableAppModalContainer } from './DisableAppModalContainer';
import { Modal } from '@patternfly/react-core';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      id: '1234',
      isOpen: false,
      toggleDisableAppModal: jest.fn(),
      disableAppVersions: jest.fn()
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<DisableAppModalContainer {...props} />);
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

describe('events', () => {
  describe('handles Modal close', () => {
    it('should call toggleSaveAppModal', () => {
      const { output, props } = setup();
      output.props.onClose();
      expect(props.toggleDisableAppModal).toBeCalled();
    });
  });

  describe('handles cancel click', () => {
    it('should call onConfirm', () => {
      const { output, props } = setup();
      const [ cancelButton ] = output.props.actions;
      cancelButton.props.onClick();
      expect(props.toggleDisableAppModal).toBeCalled();
    });
  });

  describe('handles save click', () => {
    it('should call onConfirm', () => {
      const { output, props } = setup();
      const [ , saveButton ] = output.props.actions;
      saveButton.props.onClick();
      expect(props.disableAppVersions).toBeCalled();
    });
  });
});
