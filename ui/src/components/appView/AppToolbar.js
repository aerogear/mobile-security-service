import React from 'react';
import { withRouter } from 'react-router-dom';
import PropTypes from 'prop-types';
import { Toolbar, ToolbarGroup, ToolbarItem, Breadcrumb, BreadcrumbItem, Button } from '@patternfly/react-core';
import './AppToolbar.css';

/**
 * Stateless presentational component to display breadcrumbs and app level action buttons
 * @param {object} props Component props
 * @param {object} props.app Contains all app properties
 * @param {func} props.onSaveAppClick Function to execute on click of Save button
 * @param {func} props.onDisableAppClick Function to execute on click of Disable app button
 * @param {boolean} props.isViewDirty Boolean on if the app view is dirty due to unsaved changes
 * @param {object} props.history Contains functions to modify the react-router-dom
 */
export const AppToolbar = ({ app, onSaveAppClick, onDisableAppClick, isViewDirty, history }) => {
  const onHomeClick = () => {
    history.push('/');
  };

  return (
    <Toolbar className="details-toolbar">
      <ToolbarGroup>
        <ToolbarItem>
          <Breadcrumb>
            <BreadcrumbItem className="home-breadcrumb breadcrumb" onClick={onHomeClick}>
              Home
            </BreadcrumbItem>
            <BreadcrumbItem className="breadcrumb" isActive>
              {app.appName}
            </BreadcrumbItem>
          </Breadcrumb>
        </ToolbarItem>
      </ToolbarGroup>
      <ToolbarGroup className="toolbar-buttons">
        <Button className="toolbar-button" onClick={onDisableAppClick} variant="primary">
          Disable App
        </Button>
        <Button className="toolbar-button" onClick={onSaveAppClick} isDisabled={!isViewDirty} variant="primary">
          Save
        </Button>
      </ToolbarGroup>
    </Toolbar>
  );
};

AppToolbar.propTypes = {
  app: PropTypes.shape({
    appName: PropTypes.string.isRequired
  }).isRequired,
  onSaveAppClick: PropTypes.func.isRequired,
  onDisableAppClick: PropTypes.func.isRequired,
  isViewDirty: PropTypes.bool.isRequired,
  history: PropTypes.shape({
    push: PropTypes.func.isRequired
  }).isRequired
};

export default withRouter(AppToolbar);
