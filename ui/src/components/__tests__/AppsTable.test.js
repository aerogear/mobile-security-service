import React from 'react';
import { AppsTable } from '../AppsTable';
import { shallow, mount } from 'enzyme';

describe('AppsTable', () => {
  const onSort = jest.fn();
  const onRowClick = jest.fn();
  const testProps = { columns: [], rows: [], sortBy: {}, onSort: onSort, onRowClick: onRowClick };

  it('renders the expected component', () => {
    const wrapper = shallow(<AppsTable {...testProps} />);
    expect(wrapper).toHaveLength(1);
  });

  it('passes the expected props', () => {
    const wrapper = mount(<AppsTable {...testProps} />);
    const props = wrapper.props();
    expect(props.columns).toBeTruthy();
    expect(props.rows).toBeTruthy();
    expect(props.sortBy).toBeTruthy();
    expect(props.onSort).toBeTruthy();
    expect(props.onRowClick).toBeTruthy();
  });
});
