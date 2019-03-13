import React from 'react';
import { shallow } from 'enzyme';
import AppDetailedToolbar from '../AppDetailedToolbar';

describe('AppDetailedView', () => {
  const props = { app: {} };

  it('renders the expected components without crashing', () => {
    const wrapper = shallow(<AppDetailedToolbar {...props}/>);
    expect(wrapper).toHaveLength(1);
  });
});
