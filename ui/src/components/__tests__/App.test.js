import React from 'react';
import { BrowserRouter as Router, Route } from 'react-router-dom';
import { shallow } from 'enzyme';
import App from '../App';

it('renders expected components without crashing', () => {
  const wrapper = shallow(<App />);
  expect(wrapper.find(Router)).toHaveLength(1);
  expect(wrapper.find(Route)).toHaveLength(2);
  expect(wrapper.find('div.App')).toHaveLength(1);
});
