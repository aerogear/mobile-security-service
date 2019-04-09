import React from 'react';
import { Button } from '@patternfly/react-core';
import { withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { toggleSaveAppModal } from '../../actions/actions-ui';
import ConfirmationModal from '../../components/common/ConfirmationModal';

/**
 * Redux container for the Save App Modal
 *
 * @param {boolean} props.isOpen - The open/closed state of the modal.
 * @param {string} props.title - The title to display on the modal.
 * @param {*} props.children - The child props of this component.
 * @param {*} props.onConfirm - The action to execute when `Confirm` is clicked.
 * @param {boolean} props.toggleSaveAppModal - The action to open or close the modal.
 */
export const SaveAppModalContainer = ({ isOpen, title, children, onConfirm, toggleSaveAppModal }) => {
  return (
    <ConfirmationModal
      title={title}
      isOpen={isOpen}
      onClose={toggleSaveAppModal}
      confirmAction={[<Button key="confirm" variant="primary" onClick={onConfirm}>
          Confirm
      </Button>]}>
      {children}
    </ConfirmationModal>
  );
};

SaveAppModalContainer.propTypes = {
  isOpen: PropTypes.bool.isRequired,
  title: PropTypes.string.isRequired,
  children: PropTypes.node.isRequired,
  onConfirm: PropTypes.func.isRequired,
  toggleSaveAppModal: PropTypes.func.isRequired
};

function mapStateToProps (state) {
  return {
    isOpen: state.modals.saveApp.isOpen
  };
}

export default withRouter(connect(mapStateToProps, { toggleSaveAppModal })(SaveAppModalContainer));
