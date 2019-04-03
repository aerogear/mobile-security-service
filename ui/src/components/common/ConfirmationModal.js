import React from 'react';
import { Modal, Button } from '@patternfly/react-core';
import PropTypes from 'prop-types';

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

ConfirmationModal.propTypes = {
  title: PropTypes.string.isRequired,
  isOpen: PropTypes.bool.isRequired,
  onClose: PropTypes.func.isRequired,
  confirmAction: PropTypes.array.isRequired,
  children: PropTypes.node.isRequired
};

export default ConfirmationModal;
