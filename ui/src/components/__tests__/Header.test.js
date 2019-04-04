import React from 'react';
import { mount, shallow } from 'enzyme';
import { Header } from '../Header';

describe('Header', () => {
  const testProps = { currentUser: 'jdoe', history: [] };

  it('renders the expected components without crashing', () => {
    const headerComponent = shallow(<Header {...testProps} />);
    expect(headerComponent).toHaveLength(1);
  });

  it('has expected props', () => {
    const wrapper = mount(<Header {...testProps} />);
    const props = wrapper.props();
    expect(props.currentUser).toBeTruthy();
  });

  it('onTitleClick() works and pushes to page history prop', () => {
    const wrapper = mount(<Header {...testProps} />);
    wrapper.find('.pf-c-page__header-brand-link').simulate('click');
    expect(wrapper.props().history).toContain('/');
  });

  it('toolbar and components render', () => {
    const wrapper = mount(<Header {...testProps} />);
    expect(wrapper.find('Toolbar').length).toBe(1);
    expect(wrapper.find('UserIcon').length).toBe(1);
    expect(wrapper.find('Dropdown').length).toBe(1);
  });
});
