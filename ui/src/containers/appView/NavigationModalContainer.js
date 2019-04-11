import React from 'react';
import { Button, Modal } from '@patternfly/react-core';
import { withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { toggleNavigationModal, setAppDetailedDirtyState } from '../../actions/actions-ui';

/**
 * Stateful container component for the Navigation App Modal.
 *
 * @param {string} props.id - ID of the app to disable versions for
 * @param {boolean} props.isOpen - The opened state of the modal
 * @param {string} props.targetLocation - Target location of page for routing history
 * @param {string} props.title - Title of modal
 * @param {func} props.unblockHistory - Logic to unblock routing history
 * @param {*} props.children - Children props of modal
 * @param {object} props.history Contains functions to modify the react-router-dom
 * @param {func} props.toggleNavigationModal - Logic to toggle visibility of modal
 * @param {func} props.setAppDetailedDirtyState - Logic to toggle dirty state of modal
 */
export const NavigationModalContainer = ({
  isOpen,
  targetLocation,
  title,
  unblockHistory,
  children,
  history,
  toggleNavigationModal,
  setAppDetailedDirtyState
}) => {
  const handleModalClose = () => {
    toggleNavigationModal(false);
  };

  const handleLeaveClick = () => {
    setAppDetailedDirtyState();
    unblockHistory();
    history.push(targetLocation);
    handleModalClose();
  };

  return (
    <Modal
      isLarge
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
      {children}
    </Modal>
  );
};

NavigationModalContainer.propTypes = {
  isOpen: PropTypes.bool.isRequired,
  targetLocation: PropTypes.string,
  title: PropTypes.string.isRequired,
  unblockHistory: PropTypes.func.isRequired,
  children: PropTypes.node.isRequired
};

function mapStateToProps (state) {
  return {
    isOpen: state.modals.navigationModal.isOpen,
    targetLocation: state.modals.navigationModal.targetLocation
  };
}

export default withRouter(
  connect(mapStateToProps, { toggleNavigationModal, setAppDetailedDirtyState })(NavigationModalContainer)
);
