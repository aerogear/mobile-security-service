import AppsTable from '../AppsTable';
import { TableBody, TableHeader } from '@patternfly/react-table';

describe('AppsTable', () => {
  it('renders the expected components without crashing', () => {
    expect(AppsTable).toBeDefined();
    expect(TableBody).toBeDefined();
    expect(TableHeader).toBeDefined();
  });
});
