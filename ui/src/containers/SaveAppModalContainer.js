import React from 'react';
import { Button } from '@patternfly/react-core';
import { withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { toggleSaveAppModal } from '../actions/actions-ui';
import ConfirmationModal from '../components/common/ConfirmationModal';

class SaveAppModalContainer extends React.Component {
  handleModalClose = () => {
    this.props.toggleSaveAppModal(false);
  };

  handleOnConfirm = () => {
    this.handleModalClose();
  };

  render () {
    const { title, isSaveAppModalOpen, onConfirm, children } = this.props;

    return (
      <ConfirmationModal
        title={title}
        isOpen={isSaveAppModalOpen}
        onClose={this.handleModalClose}
        confirmAction={[<Button key="confirm" variant="primary" onClick={onConfirm}>
            Confirm
        </Button>]}>
        {children}
      </ConfirmationModal>
    );
  }
}

SaveAppModalContainer.propTypes = {
  isSaveAppModalOpen: PropTypes.bool.isRequired,
  title: PropTypes.string.isRequired,
  children: PropTypes.string.isRequired,
  onConfirm: PropTypes.func.isRequired
};

function mapStateToProps (state) {
  return {
    isSaveAppModalOpen: state.isSaveAppModalOpen
  };
}

export default withRouter(connect(mapStateToProps, { toggleSaveAppModal })(SaveAppModalContainer));
