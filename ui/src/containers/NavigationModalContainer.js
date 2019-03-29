import React from 'react';
import { Button, Modal } from '@patternfly/react-core';
import PropTypes from 'prop-types';

const NavigationModalContainer = ({ title, text, isOpen, handleLeaveClick, handleModalClose }) => {
  return (
    <Modal
      title={title}
      isOpen={isOpen}
      onClose={handleModalClose}
      actions={[
        <Button key="leave" variant="danger" onClick={handleLeaveClick}>
          Leave
        </Button>,
        <Button key="stay" variant="primary" onClick={handleModalClose}>
          Stay
        </Button>
      ]}
    >
      {text}
    </Modal>
  );
};

NavigationModalContainer.propTypes = {
  isOpen: PropTypes.bool.isRequired,
  title: PropTypes.string.isRequired,
  text: PropTypes.string.isRequired,
  handleLeaveClick: PropTypes.func.isRequired,
  handleModalClose: PropTypes.func.isRequired
};

export default NavigationModalContainer;
