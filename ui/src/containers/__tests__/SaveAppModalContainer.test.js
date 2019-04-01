import React from 'react';
import { shallow } from 'enzyme';
import SaveAppModalContainer from '../SaveAppModalContainer';

const props = { isSaveAppModalOpen: true, title: 'Confirmation Modal', children: 'This is some text.' };

describe('SaveAppModalContainer', () => {
  it('renders the expected component', () => {
    const wrapper = shallow(<SaveAppModalContainer {...props}/>);
    expect(wrapper).toHaveLength(1);
  });
});
