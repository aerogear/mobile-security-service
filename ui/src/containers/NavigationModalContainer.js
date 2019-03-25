import React from 'react';
import { Button, Modal } from '@patternfly/react-core';
import { withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { toggleNavigationModal, toggleAppDetailedIsDirty } from '../actions/actions-ui';

class NavigationModalContainer extends React.Component {
  handleModalClose = () => {
    this.props.toggleNavigationModal(false);
  };

  handleLeaveClick = () => {
    this.props.toggleAppDetailedIsDirty();
    this.props.unblockHistory();
    this.props.history.push(this.props.targetLocation);
    this.handleModalClose();
  };

  render () {
    return (
      <Modal
        isLarge
        title={this.props.title}
        isOpen={this.props.isOpen}
        onClose={this.handleModalClose}
        actions={[
          <Button key="leave" variant="danger" onClick={this.handleLeaveClick}>
            Leave
          </Button>,
          <Button key="stay" variant="primary" onClick={this.handleModalClose}>
            Stay
          </Button>
        ]}>
        {this.props.children}
      </Modal>
    );
  }
}

NavigationModalContainer.propTypes = {
  isOpen: PropTypes.bool.isRequired,
  targetLocation: PropTypes.string,
  title: PropTypes.string.isRequired,
  unblockHistory: PropTypes.func.isRequired,
  children: PropTypes.string.isRequired
};

function mapStateToProps (state) {
  return {
    isOpen: state.navigationModal.isOpen,
    targetLocation: state.navigationModal.targetLocation
  };
}

export default withRouter(connect(mapStateToProps, { toggleNavigationModal, toggleAppDetailedIsDirty })(NavigationModalContainer));
