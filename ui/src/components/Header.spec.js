import React from 'react';
import { createRenderer } from 'react-test-renderer/shallow';
import Header from './Header';
import { Page, PageHeader, Toolbar, ToolbarItem, Dropdown, DropdownToggle } from '@patternfly/react-core';
import { UserIcon } from '@patternfly/react-icons';

const setup = (propOverrides) => {
  const props = Object.assign(
    {
      currentUser: 'username',
      appName: 'myApp',
      isDropDownOpen: false,
      onUserDropdownToggle: jest.fn(),
      onTitleClick: jest.fn(),
      logout: jest.fn()
    },
    propOverrides
  );

  const renderer = createRenderer();
  renderer.render(<Header {...props} />);
  const output = renderer.getRenderOutput();

  return {
    props: props,
    output: output
  };
};

describe('components', () => {
  describe('Header', () => {
    it('should render', () => {
      const { output } = setup();
      expect(output.type).toBe(Page);

      const pageHeaderView = output.props.header;
      expect(pageHeaderView.type).toBe(PageHeader);
      expect(pageHeaderView.props.logo).toBe('myApp');

      const toolbarView = pageHeaderView.props.toolbar;
      expect(toolbarView.type).toBe(Toolbar);

      const [ UserIconView, ToolbarItemView ] = toolbarView.props.children.props.children;
      expect(UserIconView.type).toBe(UserIcon);
      expect(ToolbarItemView.type).toBe(ToolbarItem);
    });

    it('logoClick should call onTitleClick', () => {
      const { output, props } = setup();

      output.props.header.props.logoProps.onClick();
      expect(props.onTitleClick).toBeCalled();
    });
  });

  describe('Dropdown', () => {
    it('should render', () => {
      const { output } = setup();

      const DropdownView = output.props.header.props.toolbar.props.children.props.children[1].props.children;
      expect(DropdownView.type).toBe(Dropdown);
      expect(DropdownView.props.isOpen).toBe(false);

      const DropdownToggleView = DropdownView.props.toggle;
      expect(DropdownToggleView.type).toBe(DropdownToggle);
      expect(DropdownToggleView.props.children).toBe('username');
    });

    it('Dropdown toggle should call onUserDropdownToggle', () => {
      const { output, props } = setup();

      const DropdownToggleView = output.props.header.props.toolbar.props.children.props.children[1].props.children.props.toggle;
      DropdownToggleView.props.onToggle();
      expect(props.onUserDropdownToggle).toBeCalled();
    });

    it('Dropdown logout click should call logout', () => {
      const { output, props } = setup();

      const DropdownItemsView = output.props.header.props.toolbar.props.children.props.children[1].props.children.props.dropdownItems[0];
      DropdownItemsView.props.onClick();
      expect(props.logout).toBeCalled();
    });
  });
});
