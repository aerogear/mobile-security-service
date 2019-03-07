import React from 'react';
import { shallow } from 'enzyme';
import { AppDetailsContainer } from '../AppDetailsContainer';

const props = { isAppsRequestFailed: false, apps: [], sortBy: {}, columns: [] };

describe('AppDetailsContainer', () => {
  it('renders the expected component', () => {
    const appDetailsContainer = shallow(<AppDetailsContainer {...props} />);
    expect(appDetailsContainer).toBeDefined();
  });
});
