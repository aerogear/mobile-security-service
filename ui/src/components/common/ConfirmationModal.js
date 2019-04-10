import React from 'react';
import PropTypes from 'prop-types';
import { Modal, Button } from '@patternfly/react-core';

/**
 * Stateless presentational component to display a confirmation modal
 *
 * @param {object} props Component props
 * @param {string} props.title String for the title header of the modal
 * @param {boolean} props.isOpen If the modal should be displayed on the screen or not
 * @param {func} props.onClose Function to execute on click of cancel or close of the modal
 * @param {array} props.confirmAction Additional UI components to display the confirmAction, normally a single button
 * @param {node} props.children All sub components to display in the modal body
 */
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
