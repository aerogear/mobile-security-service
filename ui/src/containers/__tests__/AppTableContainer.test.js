import React from 'react';
import { shallow } from 'enzyme';
import { AppTableContainer } from '../AppTableContainer';

const props = { isAppsRequestFailed: false, app: [], sortBy: {}, columns: [] };

describe('AppTableContainer', () => {
  it('renders the expected component', () => {
    const appTableContainer = shallow(<AppTableContainer {...props} />);
    expect(appTableContainer).toBeDefined();
  });
});
