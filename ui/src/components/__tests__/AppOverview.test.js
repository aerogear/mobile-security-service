import React from 'react';
import { shallow } from 'enzyme';
import { AppOverview } from '../AppOverview';

describe('AppOverview', () => {
  const props = { app: {} };
  it('renders the expected component', () => {
    const Wrapper = shallow(<AppOverview {...props} />);
    expect(Wrapper).toHaveLength(1);
  });
});
