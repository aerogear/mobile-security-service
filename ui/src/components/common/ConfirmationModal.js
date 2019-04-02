import React from 'react';
import { Modal, Button } from '@patternfly/react-core';

const ConfirmationModal = ({ title, isOpen, onClose, confirmAction, children }) => {
  return (
    <Modal
      isLarge
      title={title}
      isOpen={isOpen}
      onClose={onClose}
      actions={[
        <Button key="cancel" variant="secondary" onClick={onClose}>
          Cancel
        </Button>,
        confirmAction
      ]}>
      {children}
    </Modal>
  );
};

export default ConfirmationModal;
