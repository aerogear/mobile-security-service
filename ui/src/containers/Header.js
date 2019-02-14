import React, { Component } from 'react';
import NavHeader from '../components/common/NavHeader';

class Header extends Component {
  render () {
    console.log('navheader called');
    const { user } = 'Username';
    const userDropdowns = [
      {
        text: 'Some Shit',
        // the sign in endpoint will perform the sign out action, and return the login page
        href: '/oauth/sign_in'
      }
    ];
    return (
      <NavHeader
        title="Mobile Developer Console"
        titleHref="/overview"
        user={user}
        userDropdownItems={userDropdowns}
      />
    );
  }
}

export default Header;
