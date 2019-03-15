import React from 'react';
import { shallow } from 'enzyme';
import AppDetailedView from '../AppDetailedView';

describe('AppDetailedView', () => {
  const getAppById = jest.fn();
  const props = { getAppById: getAppById, app: {} };

  it('renders the expected components without crashing', () => {
    const wrapper = shallow(<AppDetailedView {...props}/>);
    expect(wrapper).toHaveLength(1);
  });
});
