import React from 'react';
import { Button } from '@patternfly/react-core';
import { withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { toggleNavigationModal, toggleAppDetailedIsDirty } from '../actions/actions-ui';
import { LargeModal } from '../components/common/LargeModal';

class NavigationModalContainer extends React.Component {
  handleModalClose = () => {
    this.props.toggleNavigationModal(false);
  };

  handleLeaveClick = () => {
    this.props.toggleAppDetailedIsDirty();
    this.props.history.goBack();
    this.handleModalClose();
  };

  render () {
    return (
      <LargeModal
        title={this.props.title}
        isOpen={this.props.isNavigationModalOpen}
        onClose={this.handleModalClose}
        actions={[
          <Button key="leave" variant="danger" onClick={this.handleLeaveClick}>
            Leave
          </Button>,
          <Button key="stay" variant="primary" onClick={this.handleModalClose}>
            Stay
          </Button>
        ]}
        text={this.props.text}
      />
    );
  }
}

NavigationModalContainer.propTypes = {
  isNavigationModalOpen: PropTypes.bool.isRequired,
  title: PropTypes.string.isRequired,
  text: PropTypes.string.isRequired
};

function mapStateToProps (state) {
  return {
    isNavigationModalOpen: state.isNavigationModalOpen
  };
}

export default withRouter(connect(mapStateToProps, { toggleNavigationModal, toggleAppDetailedIsDirty })(NavigationModalContainer));
