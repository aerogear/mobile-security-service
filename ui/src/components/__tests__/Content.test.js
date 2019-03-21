import React from 'react';
import { shallow } from 'enzyme';
import Content from '../common/Content';

it('renders the expected components without crashing', () => {
  const wrapper = shallow(<Content><h1 className="heading">Heading</h1></Content>);
  expect(wrapper).toHaveLength(1);
  expect(wrapper.find('h1.heading')).toHaveLength(1);
});
