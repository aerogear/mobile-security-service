import React from 'react';
import { shallow } from 'enzyme';
import ConfirmationModal from '../common/ConfirmationModal';

it('renders the expected components without crashing', () => {
  const props = { isSaveAppModalOpen: true, title: 'Confirmation Modal', children: 'This is some text.' };

  const wrapper = shallow(<ConfirmationModal {...props} />);
  expect(wrapper).toHaveLength(1);
});
