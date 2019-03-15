import React from 'react';
import { shallow } from 'enzyme';
import { AppVersionsTableContainer } from '../AppVersionsTableContainer';

const props = { appVersions: [], sortBy: {}, columns: [] };

describe('AppVersionsTableContainer', () => {
  it('renders the expected component', () => {
    const appVersionsTableContainer = shallow(<AppVersionsTableContainer {...props} />);
    expect(appVersionsTableContainer).toBeDefined();
  });
});
