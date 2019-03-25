import React from 'react';
import { mount, shallow } from 'enzyme';
import { AppVersionsTableContainer } from '../AppVersionsTableContainer';

describe('AppVersionsTableContainer', () => {
  const testProps = { appVersions: [], sortBy: {}, columns: [] };

  it('renders the expected component', () => {
    const appVersionsTableContainer = shallow(<AppVersionsTableContainer {...testProps} />);
    expect(appVersionsTableContainer).toHaveLength(1);
  });

  it('does not render an AppsTable when no apps present', () => {
    const wrapper = mount(<AppVersionsTableContainer {...testProps} />);
    expect(wrapper).toHaveLength(1);
    expect(wrapper.find('AppsTable').length).toBe(0);
  });

  it('contains the expected props', () => {
    const wrapper = mount(<AppVersionsTableContainer {...testProps} />);
    const props = wrapper.props();
    expect(props.sortBy).toBeTruthy();
    expect(props.columns).toBeTruthy();
    expect(props.appVersions).toBeTruthy();
  });
});
