import React from 'react';
import { shallow } from 'enzyme';
import AppsTableContainer from '../AppsTableContainer';

it('renders the without crashing', () => {
  shallow(<AppsTableContainer />);
});
