import React, { Component } from 'react';

class AppGridHeader extends Component {
  constructor() {
    super();
    this.state = {};
  }

  render() {
    return (
      <>
        <thead className="appGridHeader">
          <tr>
            <th>App</th>
            <th>No. Deployed Versions</th>
            <th>No. Clients</th>
            <th>No. App Startups</th>
          </tr>
        </thead>
      </>
    );
  }
}

export default AppGridHeader;
