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
import { withRouter } from 'react-router-dom';
import { UserIcon } from '@patternfly/react-icons';
import accessibleStyles from '@patternfly/patternfly/utilities/Accessibility/accessibility.css';
import { css } from '@patternfly/react-styles';
import './Header.css';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';
import config from '../../config/config';
import { toggleHeaderDropdown } from '../../actions/actions-ui';

export class Header extends React.Component {
  onTitleClick = () => {
    this.props.history.push('/');
  };

  onUserDropdownToggle = () => {
    this.props.toggleHeaderDropdown();
  };

  onUserDropdownSelect = () => {
    this.onUserDropdownToggle();
  };

  onLogoutUser = () => {
    console.log('onLogoutUser()');
  };

  render () {
    const toolbar = (
      <Toolbar>
        <ToolbarGroup className={css(accessibleStyles.screenReader, accessibleStyles.visibleOnLg)}>
          <UserIcon />
          <ToolbarItem className={css(accessibleStyles.screenReader, accessibleStyles.visibleOnMd)}>
            <Dropdown
              isPlain
              position="right"
              onSelect={this.onUserDropdownSelect}
              isOpen={this.props.isUserDropdownOpen}
              toggle={<DropdownToggle onToggle={this.onUserDropdownToggle}>{this.props.currentUser}</DropdownToggle>}
              dropdownItems={[
                <DropdownItem key="logout" component="button" href="#logout" onClick={this.onLogoutUser}>
                  Log out
                </DropdownItem>
              ]}
            />
          </ToolbarItem>
        </ToolbarGroup>
      </Toolbar>
    );

    const Header = <PageHeader logo={config.app.name.toUpperCase()} logoProps={{ onClick: this.onTitleClick }} toolbar={toolbar} />;

    return (
      <div className="mssHeader">
        <Page header={Header} />
      </div>
    );
  }
}

Header.propTypes = {
  currentUser: PropTypes.string.isRequired,
  isUserDropdownOpen: PropTypes.bool.isRequired
};

function mapStateToProps (state) {
  return {
    currentUser: state.currentUser,
    isUserDropdownOpen: state.isUserDropdownOpen
  };
}

export default withRouter(connect(mapStateToProps, { toggleHeaderDropdown })(Header));
