import React from 'react';
import { Masthead, MenuItem, Icon } from 'patternfly-react';
import { withRouter } from 'react-router-dom';
import './Header.css';

export const _NavHeader = ({ title, titleHref, user, userDropdownItems }) => (
  <Masthead title={title} navToggle={false} href={titleHref}>
    <Masthead.Collapse>
      <Masthead.Dropdown
        id="app-user-dropdown"
        title={[
          <Icon type="pf" name="user" key="user-icon" />,
          <span className="dropdown-title" key="dropdown-title">
            {user && user.name}
          </span>
        ]}
      >
        {userDropdownItems.map((item, index) => (
          <MenuItem key={index} eventKey={index} href={item.href}>
            {' '}
            {item.text}{' '}
          </MenuItem>
        ))}
      </Masthead.Dropdown>
    </Masthead.Collapse>
  </Masthead>
);

export default withRouter(_NavHeader);
