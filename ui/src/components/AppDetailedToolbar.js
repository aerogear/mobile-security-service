import React from 'react';
import { withRouter } from 'react-router-dom';
import { Toolbar, ToolbarGroup, ToolbarItem, Breadcrumb, BreadcrumbItem, Button } from '@patternfly/react-core';
import PropTypes from 'prop-types';
import './AppDetailedToolbar.css';

const AppDetailedToolbar = ({ app, onSaveAppClick, onDisableAppClick, history }) => {
  const onHomeClick = () => {
    history.push('/');
  };

  return (
    <Toolbar className='details-toolbar'>
      <ToolbarGroup>
        <ToolbarItem>
          <Breadcrumb>
            <BreadcrumbItem className='home-breadcrumb breadcrumb' onClick={onHomeClick}>Home</BreadcrumbItem>
            <BreadcrumbItem className='breadcrumb' isActive>{app.appName}</BreadcrumbItem>
          </Breadcrumb>
        </ToolbarItem>
      </ToolbarGroup>
      <ToolbarGroup className='toolbar-buttons'>
        <Button className='toolbar-button' onClick={onDisableAppClick} variant="primary">Disable App</Button>
        <Button className='toolbar-button' onClick={onSaveAppClick} variant="primary">Save</Button>
      </ToolbarGroup>
    </Toolbar>
  );
};

AppDetailedToolbar.propTypes = {
  app: PropTypes.object.isRequired,
  onSaveAppClick: PropTypes.func.isRequired,
  onDisableAppClick: PropTypes.func.isRequired,
  history: PropTypes.shape({
    push: PropTypes.func.isRequired
  }).isRequired
};

export default withRouter(AppDetailedToolbar);
