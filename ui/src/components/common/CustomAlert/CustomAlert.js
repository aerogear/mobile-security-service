import React from 'react';
import './CustomAlert.css';
import PropTypes from 'prop-types';
import { Alert, AlertActionCloseButton } from '@patternfly/react-core';

/**
 * A custom alert component that wraps the Patternfly 4 Alert.
 *
 * @param {Object} visible - The visible state of the alert.
 * @param {func} onClose - The callback function to execute on close.
 * @param {*} children - Any child props - usually the description.
 * @param {Object} props - The component props
 */
const CustomAlert = ({ visible, onClose, children, ...props }) => {
  if (!visible) {
    return null;
  }

  return (
    <Alert
      className='Alert'
      {...props}
      action={<AlertActionCloseButton onClose={onClose} />}>
      {children}
    </Alert>
  );
};

CustomAlert.propTypes = {
  visible: PropTypes.bool.isRequired,
  onClose: PropTypes.func.isRequired,
  children: PropTypes.any
};

export default CustomAlert;
