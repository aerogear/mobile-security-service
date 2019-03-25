import React from 'react';
import { shallow } from 'enzyme';
import LandingPage from '../LandingPage';

describe('AppVersionsTableContainer', () => {
  it('renders the expected components', () => {
    const landingPageComponent = shallow(<LandingPage />);
    expect(landingPageComponent).toHaveLength(1);
  });
});
