import React, { useState } from 'react';

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
import { withRouter } from 'react-router-dom';
import { UserIcon } from '@patternfly/react-icons';
import accessibleStyles from '@patternfly/patternfly/utilities/Accessibility/accessibility.css';
import { css } from '@patternfly/react-styles';
import './Header.css';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import config from '../../config/config';

const Header = ({ currentUser, history }) => {
  const [isDropDownOpen, setIsDropDownOpen] = useState(false);

  const onTitleClick = () => {
    history.push('/');
  };

  const onUserDropdownToggle = () => {
    setIsDropDownOpen(!isDropDownOpen);
  };

  const onLogoutUser = () => {
    console.log('onLogoutUser()');
  };

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
            toggle={<DropdownToggle onToggle={onUserDropdownToggle}>{currentUser}</DropdownToggle>}
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

  const Header = <PageHeader logo={config.app.name.toUpperCase()} logoProps={{ onClick: onTitleClick }} toolbar={toolbar} />;

  return (
    <Page header={Header} className='mss-header'/>
  );
};

Header.propTypes = {
  currentUser: PropTypes.string.isRequired
};

function mapStateToProps (state) {
  return {
    currentUser: state.currentUser
  };
}

export default withRouter(connect(mapStateToProps)(Header));
