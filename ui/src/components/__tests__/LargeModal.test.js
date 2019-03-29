import React from 'react';
import { LargeModal } from '../common/LargeModal';
import { shallow, mount } from 'enzyme';

describe('LargeModal', () => {
  const onClose = jest.fn();
  const testProps = { title: 'title', isOpen: true, onClose: onClose, actions: [], text: 'text' };

  it('renders the expected component', () => {
    const wrapper = shallow(<LargeModal {...testProps} />);
    expect(wrapper).toHaveLength(1);
  });

  it('passes the expected props', () => {
    const wrapper = mount(<LargeModal {...testProps} />);
    const props = wrapper.props();
    expect(props.title).toBeTruthy();
    expect(props.isOpen).toBeTruthy();
    expect(props.onClose).toBeTruthy();
    expect(props.actions).toBeTruthy();
    expect(props.text).toBeTruthy();
  });
});
