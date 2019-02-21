import React from 'react';
import { shallow } from 'enzyme';
import AppDetailedView from '../AppDetailedView';

it('renders the expected components without crashing', () => {
  const wrapper = shallow(<AppDetailedView />);
  expect(wrapper.find('div.app-detailed-view')).toHaveLength(1);
});
