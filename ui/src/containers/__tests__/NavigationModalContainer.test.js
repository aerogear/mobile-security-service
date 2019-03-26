import React from 'react';
import { shallow } from 'enzyme';
import NavigationModalContainer from '../NavigationModalContainer';

const props = { isOpen: true };

describe('NavigationModalContainer', () => {
  it('renders the expected component', () => {
    const navigationModalContainer = shallow(<NavigationModalContainer {...props} />);
    expect(navigationModalContainer).toHaveLength(1);
  });
});
