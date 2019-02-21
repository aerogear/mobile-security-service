import React from 'react';
import Header from './common/Header';

class AppDetailedView extends React.Component {
  constructor () {
    super();
    this.state = {};
  }

  render () {
    return (
      <div className="app-detailed-view">
        <Header />
      </div>
    );
  };
}

export default AppDetailedView;
