import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import { NavigationModalContainer } from './NavigationModalContainer';
import { Modal } from '@patternfly/react-core';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      isOpen: false,
      targetLocation: '/',
      title: 'title',
      unblockHistory: jest.fn(),
      children: [],
      history: [],
      toggleNavigationModal: jest.fn(),
      setAppDetailedDirtyState: jest.fn()
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<NavigationModalContainer {...props} />);
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
    it('should call toggleNavigationModal', () => {
      const { output, props } = setup();
      output.props.onClose();
      expect(props.toggleNavigationModal).toBeCalledWith(false);
    });
  });

  describe('handles leave click', () => {
    it('should call setAppDetailedDirtyState', () => {
      const { output, props } = setup();
      const [ leaveButton ] = output.props.actions;
      leaveButton.props.onClick();
      expect(props.setAppDetailedDirtyState).toBeCalled();
    });

    it('should call unblockHistory', () => {
      const { output, props } = setup();
      const [ leaveButton ] = output.props.actions;
      leaveButton.props.onClick();
      expect(props.unblockHistory).toBeCalled();
    });

    it('should push location to history', () => {
      // TODO
    });

    it('should call toggleNavigationModal', () => {
      const { output, props } = setup();
      const [ leaveButton ] = output.props.actions;
      leaveButton.props.onClick();
      expect(props.toggleNavigationModal).toBeCalledWith(false);
    });
  });
});
