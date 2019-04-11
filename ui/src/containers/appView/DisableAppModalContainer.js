import React from 'react';
import { Button, TextInput, Form, FormGroup, Stack, StackItem, Modal } from '@patternfly/react-core';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import { toggleDisableAppModal, disableAppVersions, setModalDisableMessage } from '../../actions/actions-ui';

/**
 * Redux container component for the Disable App Modal.
 *
 * @param {string} props.id - ID of the app to disable versions for
 * @param {boolean} props.isOpen - The opened state of the modal
 * @param {*} props.toggleDisableAppModal - Action to toggle opened/closed state of modal
 * @param {*} props.disableAppVersions - Action to disable all app versions
 */
export const DisableAppModalContainer = ({ id, isOpen, toggleDisableAppModal, disableAppVersions, disableMessage, setModalDisableMessage }) => {
  const handleDisableAppSave = () => {
    disableAppVersions(id, disableMessage);
  };

  const onChangeDisableTextInput = (value) => {
    setModalDisableMessage(value);
  };

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
                onChange={onChangeDisableTextInput}
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
    id: state.app.data.id,
    isOpen: state.modals.disableApp.isOpen,
    disableMessage: state.modals.disableApp.disableMessage
  };
}

const mapDispatchToProps = {
  disableAppVersions,
  toggleDisableAppModal,
  setModalDisableMessage
};

export default connect(mapStateToProps, mapDispatchToProps)(DisableAppModalContainer);
