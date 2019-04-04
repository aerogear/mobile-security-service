import React from 'react';
import { withRouter } from 'react-router-dom';
import { Toolbar, ToolbarGroup, ToolbarItem, Breadcrumb, BreadcrumbItem, Button } from '@patternfly/react-core';
import PropTypes from 'prop-types';
import './AppToolbar.css';

const AppToolbar = ({ app, onSaveAppClick, onDisableAppClick, history, isViewDirty }) => {
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
        <Button className='toolbar-button' onClick={onSaveAppClick} isDisabled={!isViewDirty} variant="primary">Save</Button>
      </ToolbarGroup>
    </Toolbar>
  );
};

AppToolbar.propTypes = {
  app: PropTypes.object.isRequired,
  onSaveAppClick: PropTypes.func.isRequired,
  onDisableAppClick: PropTypes.func.isRequired,
  isViewDirty: PropTypes.bool.isRequired,
  history: PropTypes.shape({
    push: PropTypes.func.isRequired
  }).isRequired
};

export default withRouter(AppToolbar);
