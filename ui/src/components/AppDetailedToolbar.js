import React from 'react';
import { withRouter } from 'react-router-dom';
import { Toolbar, ToolbarGroup, ToolbarItem, Breadcrumb, BreadcrumbItem, Button } from '@patternfly/react-core';
import './AppDetailedToolbar.css';

export class AppDetailedToolbar extends React.Component {
  onDisableApp = () => {
    console.log('Disable button clicked');
  }

  onHomeClick = () => {
    this.props.history.push('/');
  }

  onSaveApp = () => {
    console.log('Save button clicked');
  }

  render () {
    return (
      <Toolbar className='details-toolbar'>
        <ToolbarGroup>
          <ToolbarItem>
            <Breadcrumb>
              <BreadcrumbItem className='home-breadcrumb breadcrumb' onClick={this.onHomeClick}>Home</BreadcrumbItem>
              <BreadcrumbItem className='breadcrumb' isActive >{this.props.app.appName}</BreadcrumbItem>
            </Breadcrumb>
          </ToolbarItem>
        </ToolbarGroup>
        <ToolbarGroup className='toolbar-buttons'>
          <Button className='toolbar-button' onClick={this.onDisableApp} variant="primary">Disable App</Button>
          <Button className='toolbar-button' onClick={this.onSaveApp} variant="primary">Save</Button>
        </ToolbarGroup>
      </Toolbar>
    );
  };
}

export default withRouter(AppDetailedToolbar);
