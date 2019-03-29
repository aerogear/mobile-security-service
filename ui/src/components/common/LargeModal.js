import React from 'react';
import { Modal } from '@patternfly/react-core';
import PropTypes from 'prop-types';
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

LargeModal.propTypes = {
  title: PropTypes.string.isRequired,
  isOpen: PropTypes.bool.isRequired,
  onClose: PropTypes.func.isRequired,
  actions: PropTypes.array.isRequired,
  text: PropTypes.string.isRequired
};

export default withRouter(LargeModal);
