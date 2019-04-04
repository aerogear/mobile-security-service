import React from 'react';
import { shallow } from 'enzyme';
import DisableAppModalContainer from '../appView/DisableAppModalContainer';

const props = { isOpen: true };

describe('DisableAppModalContainer', () => {
  it('renders the expected component', () => {
    const wrapper = shallow(<DisableAppModalContainer {...props}/>);
    expect(wrapper).toHaveLength(1);
  });
});
