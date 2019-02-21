import React from 'react';
import { shallow } from 'enzyme';
import Header from '../common/Header';

it('renders the expected components without crashing', () => {
  const headerComponent = shallow(<Header />);
  expect(headerComponent).toBeDefined();
});
