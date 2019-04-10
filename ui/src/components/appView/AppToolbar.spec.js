import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import { AppToolbar } from './AppToolbar';
import { Toolbar, ToolbarGroup, ToolbarItem, Breadcrumb, BreadcrumbItem, Button } from '@patternfly/react-core';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      app: {
        appName: 'my-app'
      },
      onSaveAppClick: jest.fn(),
      onDisableAppClick: jest.fn(),
      onHomeClick: jest.fn(),
      isViewDirty: false,
      history: {
        push: jest.fn()
      }
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<AppToolbar {...props} />);
  const output = renderer.getRenderOutput();

  return {
    props: props,
    output: output
  };
};

describe('components', () => {
  describe('Toolbar', () => {
    it('should render', () => {
      const { output } = setup();
      expect(output.type).toBe(Toolbar);
    });
  });

  describe('Home breadcrumb', () => {
    it('should render', () => {
      const { output, props } = setup();
      expect(output.type).toBe(Toolbar);

      const [ toolbarGroup ] = output.props.children;
      expect(toolbarGroup.type).toBe(ToolbarGroup);

      const toolbarItem = toolbarGroup.props.children;
      expect(toolbarItem.type).toBe(ToolbarItem);

      const breadcrumb = toolbarItem.props.children;
      expect(breadcrumb.type).toBe(Breadcrumb);

      const [ homeBreadcrumb ] = breadcrumb.props.children;
      expect(homeBreadcrumb.props.className).toBe('home-breadcrumb breadcrumb');
      expect(homeBreadcrumb.props.children).toBe('Home');

      homeBreadcrumb.props.onClick();
      expect(props.history.push).toBeCalled();
    });

    it('onClick should call history push', () => {
      const { output, props } = setup();
      const [ toolbarGroup ] = output.props.children;
      const toolbarItem = toolbarGroup.props.children;
      const breadcrumb = toolbarItem.props.children;

      const [ homeBreadcrumb ] = breadcrumb.props.children;
      homeBreadcrumb.props.onClick();
      expect(props.history.push).toBeCalled();
    });
  });

  describe('App Breadcrumb', () => {
    it('should render', () => {
      const { output } = setup();
      const [ toolbarGroup ] = output.props.children;
      const toolbarItem = toolbarGroup.props.children;
      const breadcrumb = toolbarItem.props.children;

      const [ , appBreadcrumb ] = breadcrumb.props.children;
      expect(appBreadcrumb.type).toBe(BreadcrumbItem);
      expect(appBreadcrumb.props.className).toBe('breadcrumb');
      expect(appBreadcrumb.props.children).toBe('my-app');
    });
  });

  describe('Disable App Button', () => {
    it('should render', () => {
      const { output } = setup();
      const [ , toolbarGroup ] = output.props.children;

      const [ disableAppButton ] = toolbarGroup.props.children;
      expect(disableAppButton.type).toBe(Button);
    });

    it('onClick should call onDisableAppClick', () => {
      const { output, props } = setup();
      const [ , toolbarGroup ] = output.props.children;

      const [ disableAppButton ] = toolbarGroup.props.children;
      disableAppButton.props.onClick();
      expect(props.onDisableAppClick).toBeCalled();
    });
  });

  describe('Save App Button', () => {
    it('should render', () => {
      const { output } = setup();
      const [ , toolbarGroup ] = output.props.children;

      const [ , saveAppButton ] = toolbarGroup.props.children;
      expect(saveAppButton.type).toBe(Button);
    });

    it('onClick should call onSaveAppClick', () => {
      const { output, props } = setup();
      const [ , toolbarGroup ] = output.props.children;

      const [ , saveAppButton ] = toolbarGroup.props.children;
      saveAppButton.props.onClick();
      expect(props.onSaveAppClick).toBeCalled();
    });
  });
});
