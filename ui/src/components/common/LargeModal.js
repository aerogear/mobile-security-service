import React from 'react';
import { Modal } from '@patternfly/react-core';
import { withRouter } from 'react-router-dom';

export const LargeModal = ({ title, isOpen, onClose, actions, text }) => (
  <Modal
    isLarge
    title={title}
    isOpen={isOpen}
    onClose={onClose}
    actions={actions}
  >
    {text}
  </Modal>
);

export default withRouter(LargeModal);
