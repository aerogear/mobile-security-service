import React, { Component } from 'react';

class AppGridRows extends Component {
  constructor() {
    super();
    this.state = {};
  }

  render() {
    return (
      <>
        <tbody className="appGridRows">
          <tr>
            <td data-label="Repository Name">Repository 1</td>
            <td data-label="Branches">10</td>
            <td data-label="Pull Requests">25</td>
            <td data-label="Workspaces">5</td>
          </tr>
          <tr>
            <td data-label="Repository Name">Repository 2</td>
            <td data-label="Branches">10</td>
            <td data-label="Pull Requests">25</td>
            <td data-label="Workspaces">5</td>
          </tr>
        </tbody>
      </>
    );
  }
}

export default AppGridRows;
