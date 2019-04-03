import React from 'react';
import { Button } from '@patternfly/react-core';
import { withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { toggleSaveAppModal } from '../actions/actions-ui';
import ConfirmationModal from '../components/common/ConfirmationModal';

const SaveAppModalContainer = ({ isSaveAppModalOpen, title, children, onConfirm, toggleSaveAppModal }) => {
  const handleModalClose = () => {
    toggleSaveAppModal(false);
  };

  return (
    <ConfirmationModal
      title={title}
      isOpen={isSaveAppModalOpen}
      onClose={handleModalClose}
      confirmAction={[<Button key="confirm" variant="primary" onClick={onConfirm}>
          Confirm
      </Button>]}>
      {children}
    </ConfirmationModal>
  );
};

SaveAppModalContainer.propTypes = {
  isSaveAppModalOpen: PropTypes.bool.isRequired,
  title: PropTypes.string.isRequired,
  children: PropTypes.node.isRequired,
  onConfirm: PropTypes.func.isRequired,
  toggleSaveAppModal: PropTypes.func.isRequired
};

function mapStateToProps (state) {
  return {
    isSaveAppModalOpen: state.isSaveAppModalOpen
  };
}

export default withRouter(connect(mapStateToProps, { toggleSaveAppModal })(SaveAppModalContainer));
