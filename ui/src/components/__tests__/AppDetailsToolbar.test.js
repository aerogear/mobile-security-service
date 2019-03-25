import React from 'react';
import { shallow, mount } from 'enzyme';
import { AppDetailedToolbar } from '../AppDetailedToolbar';

describe('AppDetailedToolbar', () => {
  const onSaveApp = jest.fn();
  const onDisableApp = jest.fn();
  const props = { app: {}, onSaveApp, onDisableApp };

  it('renders the expected components without crashing', () => {
    const wrapper = shallow(<AppDetailedToolbar {...props}/>);
    expect(wrapper).toHaveLength(1);
  });

  it('renders the breadcrumb and buttons', () => {
    const wrapper = mount(<AppDetailedToolbar {...props}/>);
    expect(wrapper.find('Breadcrumb')).toHaveLength(1);
    expect(wrapper.find('Button')).toHaveLength(2);
  });
});
