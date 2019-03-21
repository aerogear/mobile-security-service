import React from 'react';
import { shallow } from 'enzyme';
import { AppOverviewContainer } from '../AppOverviewContainer';

describe('AppsTableContainer', () => {
  const props = { app: {} };
  it('renders the expected component', () => {
    const Wrapper = shallow(<AppOverviewContainer {...props} />);
    expect(Wrapper).toHaveLength(1);
  });
});
