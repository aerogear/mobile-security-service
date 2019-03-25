import React, { useState } from 'react';
import { Button, TextInput, Form, FormGroup, Stack, StackItem, Modal } from '@patternfly/react-core';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { toggleDisableAppModal } from '../actions/actions-ui';

const DisableAppModalContainer = ({ isOpen, toggleDisableAppModal }) => {
  function handleDisableAppSave () {
    // TODO: This needs to send a request to the backend and update the client state
    // this.props.toggleDisableAppModal();
  };

  const [disableMessage, setDisableMessage] = useState('');

  return (
    <Modal
      isLarge
      title="Disable All App Versions"
      isOpen={isOpen}
      onClose={toggleDisableAppModal}
      actions={[
        <Button key="cancel" variant="secondary" onClick={toggleDisableAppModal}>
          Cancel
        </Button>,
        <Button key="save" variant="primary" onClick={handleDisableAppSave}>
          Save
        </Button>
      ]}
    >
      <Stack gutter="md">
        <StackItem>
          You have requested App Disablement. This will disable all current versions of the App. Are you sure you want to proceed?
        </StackItem>
        <StackItem isMain>
          <Form isHorizontal>
            <FormGroup
              label="Disabled Message"
              fieldId="horizontal-form-disable-message"
              helperText="Optional Field. This will overwrite existing custom disabled messages for all versions"
            >
              <TextInput
                value={disableMessage}
                onChange={(value) => setDisableMessage(value)}
                placeholder="Enter disabled message"
                type="text"
                id="horizontal-form-disable-message"
                name="horizontal-form-disable-message"
              />
            </FormGroup>
          </Form>
        </StackItem>
      </Stack>
    </Modal>
  );
};

DisableAppModalContainer.propTypes = {
  isOpen: PropTypes.bool.isRequired
};

function mapStateToProps (state) {
  return {
    isOpen: state.modals.disableApp.isOpen
  };
}

export default connect(mapStateToProps, { toggleDisableAppModal })(DisableAppModalContainer);
