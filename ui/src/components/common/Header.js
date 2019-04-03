import React from 'react';

import {
  Dropdown,
  DropdownToggle,
  DropdownItem,
  Page,
  PageHeader,
  Toolbar,
  ToolbarGroup,
  ToolbarItem
} from '@patternfly/react-core';
import { UserIcon } from '@patternfly/react-icons';
import accessibleStyles from '@patternfly/patternfly/utilities/Accessibility/accessibility.css';
import { css } from '@patternfly/react-styles';
import './Header.css';
import PropTypes from 'prop-types';

const Header = ({ currentUser, appName, isDropDownOpen, onUserDropdownToggle, onTitleClick, onLogoutUser }) => {
  const toolbar = (
    <Toolbar>
      <ToolbarGroup className={css(accessibleStyles.screenReader, accessibleStyles.visibleOnLg)}>
        <UserIcon />
        <ToolbarItem className={css(accessibleStyles.screenReader, accessibleStyles.visibleOnMd)}>
          <Dropdown
            isPlain
            position="right"
            onSelect={onUserDropdownToggle}
            isOpen={isDropDownOpen}
            toggle={
              <DropdownToggle onToggle={onUserDropdownToggle}>
                {currentUser}
              </DropdownToggle>
            }
            dropdownItems={[
              <DropdownItem key="logout" component="button" href="#logout" onClick={onLogoutUser}>
                Log out
              </DropdownItem>
            ]}
          />
        </ToolbarItem>
      </ToolbarGroup>
    </Toolbar>
  );

  const Header = <PageHeader logo={appName} logoProps={{ onClick: onTitleClick }} toolbar={toolbar} />;

  return (
    <Page header={Header} className='mss-header'/>
  );
};

Header.propTypes = {
  currentUser: PropTypes.string.isRequired,
  appName: PropTypes.string.isRequired,
  isDropDownOpen: PropTypes.bool.isRequired,
  onUserDropdownToggle: PropTypes.func.isRequired,
  onTitleClick: PropTypes.func.isRequired,
  onLogoutUser: PropTypes.func.isRequired
};

export default Header;
