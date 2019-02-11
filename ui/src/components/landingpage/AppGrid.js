import React, { Component } from 'react';
import AppGridHeader from './AppGridHeader';
import AppGridRows from './AppGridRows';
import { orderBy } from 'lodash';
import * as sort from 'sortabular';
import * as resolve from 'table-resolver';
import {
  actionHeaderCellFormatter,
  customHeaderFormattersDefinition,
  defaultSortingOrder,
  sortableHeaderCellFormatter,
  tableCellFormatter,
  Table,
  TABLE_SORT_DIRECTION
} from 'patternfly-react';
import { compose } from 'recompose';

const mockRows = [
  { id: 1000, app: 'App-A', versions: 3, clients: 1245, startups: 'male', birth_year: '1979', actions: null },
  { id: 1001, app: 'App-B', versions: 4, clients: 655, startups: 'male', birth_year: '1974', actions: null },
  { id: 1002, app: 'App-C', versions: 1, clients: 970, startups: 'female', birth_year: '1989', actions: null },
  { id: 1003, app: 'App-D', versions: 6, clients: 255, startups: 'male', birth_year: '1990', actions: null },
  { id: 1004, app: 'App-E', versions: 5, clients: 120, startups: 'female', birth_year: '1999', actions: null }
];

class AppGrid extends Component {
  constructor(props) {
    super(props);

    // Point the transform to your sortingColumns. React state can work for this purpose
    // but you can use a state manager as well.
    const getSortingColumns = () => this.state.sortingColumns || {};

    const sortableTransform = sort.sort({
      getSortingColumns,
      onSort: (selectedColumn) => {
        this.setState({
          sortingColumns: sort.byColumn({
            sortingColumns: this.state.sortingColumns,
            sortingOrder: defaultSortingOrder,
            selectedColumn
          })
        });
      },
      // Use property or index dependening on the sortingColumns structure specified
      strategy: sort.strategies.byProperty
    });

    const sortingFormatter = sort.header({
      sortableTransform,
      getSortingColumns,
      strategy: sort.strategies.byProperty
    });

    // enables our custom header formatters extensions to reactabular
    this.customHeaderFormatters = customHeaderFormattersDefinition;

    this.state = {
      // Sort the first column in an ascending way by default.
      sortingColumns: {
        name: {
          direction: TABLE_SORT_DIRECTION.ASC,
          position: 0
        }
      },
      columns: [
        {
          property: 'app',
          header: {
            label: 'Name',
            props: {
              index: 0,
              rowSpan: 1,
              colSpan: 1,
              sort: true
            },
            transforms: [ sortableTransform ],
            formatters: [ sortingFormatter ],
            customFormatters: [ sortableHeaderCellFormatter ]
          },
          cell: {
            props: {
              index: 0
            },
            formatters: [ tableCellFormatter ]
          }
        },
        {
          property: 'versions',
          header: {
            label: 'Versions',
            props: {
              index: 1,
              rowSpan: 1,
              colSpan: 1,
              sort: true
            },
            transforms: [ sortableTransform ],
            formatters: [ sortingFormatter ],
            customFormatters: [ sortableHeaderCellFormatter ]
          },
          cell: {
            props: {
              index: 1
            },
            formatters: [ tableCellFormatter ]
          }
        },
        {
          property: 'clients',
          header: {
            label: 'Clients',
            props: {
              index: 2,
              rowSpan: 1,
              colSpan: 1,
              sort: true
            },
            transforms: [ sortableTransform ],
            formatters: [ sortingFormatter ],
            customFormatters: [ sortableHeaderCellFormatter ]
          },
          cell: {
            props: {
              index: 2
            },
            formatters: [ tableCellFormatter ]
          }
        },
        {
          property: 'startups',
          header: {
            label: 'Startups',
            props: {
              index: 3,
              rowSpan: 1,
              colSpan: 1,
              sort: true
            },
            transforms: [ sortableTransform ],
            formatters: [ sortingFormatter ],
            customFormatters: [ sortableHeaderCellFormatter ]
          },
          cell: {
            props: {
              index: 3
            },
            formatters: [ tableCellFormatter ]
          }
        }
      ],
      rows: mockRows.slice(0, 6)
    };
  }
  render() {
    const { rows, sortingColumns, columns } = this.state;
    console.log('this.state', this.state);

    const sortedRows = compose(
      sort.sorter({
        columns: columns,
        sortingColumns,
        sort: orderBy,
        strategy: sort.strategies.byProperty
      })
    )(rows);

    return (
      <div>
        <Table.PfProvider
          striped
          bordered
          hover
          dataTable
          columns={columns}
          components={{
            header: {
              cell: (cellProps) => {
                return this.customHeaderFormatters({
                  cellProps,
                  columns,
                  sortingColumns
                });
              }
            }
          }}
        >
          <Table.Header headerRows={resolve.headerRows({ columns })} />
          <Table.Body
            rows={sortedRows}
            rowKey="id"
            onRow={() => {
              return {
                role: 'row'
              };
            }}
          />
        </Table.PfProvider>
      </div>
    );
  }
}

export default AppGrid;
