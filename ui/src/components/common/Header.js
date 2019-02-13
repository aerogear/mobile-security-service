import React, { Component } from 'react';
import { Masthead , MenuItem , Icon } from 'patternfly-react';
import { withRouter } from 'react-router-dom';
import './Header.css';

class Header extends Component {
  render() {

    const userDropdowns = [
        {
          text: 'Logout',
          // the sign in endpoint will perform the sign out action, and return the login page
          href: '/oauth/sign_in'
        }
      ];
    const user = "Username";
    const title = "Mobile Security Service";
    const titleHref ="/overview";

    return (
     
        <Masthead title={title} navToggle={false} href={titleHref}>
        <Masthead.Collapse>
            <Masthead.Dropdown
                    id="app-user-dropdown"
                    title={[
                    <Icon type="pf" name="user" key="user-icon" />,
                    <span className="dropdown-title" key="dropdown-title">
                        {user}
                    </span>
                    ]}
                >
                    {userDropdowns.map((item, index) => (
                    <MenuItem key={index} eventKey={index} href={item.href}>
                        {' '}
                        {item.text}{' '}
                    </MenuItem>
                    ))}
                </Masthead.Dropdown>
            </Masthead.Collapse>
        </Masthead>
);

  }
}

export default withRouter(Header);
