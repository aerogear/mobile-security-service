import React from 'react';
import { Modal, Button } from '@patternfly/react-core';

class ConfirmationModal extends React.Component {
  render () {
    const { title, isOpen, onClose, confirmAction, children } = this.props;

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
  }
}

export default ConfirmationModal;
