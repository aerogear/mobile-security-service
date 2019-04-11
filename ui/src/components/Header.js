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
import PropTypes from 'prop-types';
import accessibleStyles from '@patternfly/patternfly/utilities/Accessibility/accessibility.css';
import { css } from '@patternfly/react-styles';
import './Header.css';

/**
 * Stateless presentational component to display the UI header
 *
 * @param {object} props Component props
 * @param {string} props.currentUser Name of the currently logged in user
 * @param {string} props.appName Name of the UI
 * @param {boolean} props.isDropDownOpen IF the dropdown should be open or not
 * @param {func} props.onUserDropdownToggle logic to determine what should happen on dropdown toggle
 * @param {func} props.onTitleClick logic to execute on title click
 * @param {func} props.logout logic to execute on logout click
 */
const Header = ({ username, appName, isDropDownOpen, onUserDropdownToggle, onTitleClick, logout }) => {
  const toolbar = (
    <Toolbar>
      <ToolbarGroup className={css(accessibleStyles.screenReader, accessibleStyles.visibleOnLg)}>
        <UserIcon />
        <ToolbarItem className={css(accessibleStyles.screenReader, accessibleStyles.visibleOnMd)}>
          <Dropdown
            isPlain
            position="right"
            isOpen={isDropDownOpen}
            toggle={
              <DropdownToggle onToggle={onUserDropdownToggle}>
                {username}
              </DropdownToggle>
            }
            dropdownItems={[
              <DropdownItem key="logout" component="button" onClick={logout}>
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
  username: PropTypes.string.isRequired,
  appName: PropTypes.string.isRequired,
  isDropDownOpen: PropTypes.bool.isRequired,
  onUserDropdownToggle: PropTypes.func.isRequired,
  onTitleClick: PropTypes.func.isRequired,
  logout: PropTypes.func.isRequired
};

export default Header;
