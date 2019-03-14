import React from 'react';
import { shallow } from 'enzyme';
import { AppVersionTableContainer } from '../AppVersionTableContainer';

const props = { isAppsRequestFailed: false, app: [], sortBy: {}, columns: [] };

describe('AppVersionTableContainer', () => {
  it('renders the expected component', () => {
    const AppVersionTableContainer = shallow(<AppVersionTableContainer {...props} />);
    expect(AppVersionTableContainer).toBeDefined();
  });
});
