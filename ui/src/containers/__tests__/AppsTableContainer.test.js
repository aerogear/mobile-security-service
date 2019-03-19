import React from 'react';
import { shallow } from 'enzyme';
import { AppsTableContainer } from '../AppsTableContainer';
import AppsTable from '../../components/AppsTable';

describe('AppsTableContainer', () => {
  const getApps = jest.fn();
  const props = { isAppsRequestFailed: false, getApps: getApps, apps: {}, sortBy: {}, columns: [] };

  it('renders without crashing', () => {
    const Wrapper = shallow(<AppsTableContainer {...props} />);
    expect(Wrapper.find(AppsTable)).toHaveLength(1);
  });

  it('renders the expected view on app request', () => {
    props.isAppsRequestFailed = true;
    const Wrapper = shallow(<AppsTableContainer {...props} />);
    expect(Wrapper.find(AppsTable)).toHaveLength(0);
    expect(Wrapper.find('div.no-apps')).toHaveLength(1);
  });
});
