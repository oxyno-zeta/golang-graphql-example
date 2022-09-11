import React, { useState } from 'react';
import { gql, useQuery } from '@apollo/client';
import Paper from '@mui/material/Paper';
import Divider from '@mui/material/Divider';
import dayjs from 'dayjs';
import { useSearchParams } from 'react-router-dom';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useTheme } from '@mui/material/styles';
import {
  TodoSortOrderModel,
  TodoFilterModel,
  todoFilterDefinitionObject,
  todoSortFields,
  TodoModel,
} from '../../models/todos';
import Title from '../../components/Title';
import FilterSearchBar from '../../components/filters/FilterSearchBar';
import { ConnectionModel, FilterQueryParamName, SortQueryParamName, StringFilterModel } from '../../models/general';
import GraphqlErrors from '../../components/GraphqlErrors';
import SortButton from '../../components/sorts/SortButton';
import GridView from './components/GridView';
import TableView from './components/TableView';
import Pagination from '../../components/Pagination';
import { onMainSearchChange } from './utils';
import { getPaginationFromSearchParams, cleanAndSetCleanedPagination } from '../../utils/pagination';
import GridTableViewSwitcher from '../../components/GridTableViewSwitcher';
import { getJSONObjectFromSearchParam, setJSONObjectSearchParam } from '../../utils/urlSearchParams';

const GET_TODOS_QUERY = gql`
  query getTodos($first: Int, $last: Int, $before: String, $after: String, $filter: TodoFilter, $sort: TodoSortOrder) {
    todos(first: $first, last: $last, before: $before, after: $after, filter: $filter, sort: $sort) {
      edges {
        node {
          id
          createdAt
          updatedAt
          text
          done
        }
      }
      pageInfo {
        hasNextPage
        hasPreviousPage
        startCursor
        endCursor
      }
    }
  }
`;

interface QueryResult {
  todos: ConnectionModel<TodoModel>;
}

interface QueryVariables {
  first?: number;
  last?: number;
  before?: string;
  after?: string;
  sort?: TodoSortOrderModel;
  filter?: TodoFilterModel | null;
}

const maxPagination = 20;
const initialPagination = { first: maxPagination };

function Todos() {
  // Theming
  const theme = useTheme();
  const sizeMatching = useMediaQuery(theme.breakpoints.down('lg'));

  // Get search params
  const [searchParams, setSearchParams] = useSearchParams();
  // Filter, pagination and sort values
  const filter = getJSONObjectFromSearchParam<TodoFilterModel>(FilterQueryParamName, {}, searchParams);
  const sort = getJSONObjectFromSearchParam<TodoSortOrderModel>(
    SortQueryParamName,
    { createdAt: 'DESC' },
    searchParams,
  );
  const pagination = getPaginationFromSearchParams(initialPagination, maxPagination, searchParams, setSearchParams);

  // Setter
  const setSort = (data: TodoSortOrderModel) => {
    setJSONObjectSearchParam(SortQueryParamName, data, searchParams, setSearchParams);
  };
  const setFilter = (data: TodoFilterModel) => {
    setJSONObjectSearchParam(FilterQueryParamName, data, searchParams, setSearchParams);
  };

  // States
  const [gridView, setGridView] = useState(sizeMatching);
  // Call graphql
  const { data, loading, error } = useQuery<QueryResult, QueryVariables>(GET_TODOS_QUERY, {
    variables: { ...pagination, sort, filter },
    fetchPolicy: 'network-only',
  });

  let mainSearchInitialValue = '';
  if (filter && filter.text && filter.text.contains) {
    mainSearchInitialValue = filter.text.contains;
  } else if (filter && filter.AND) {
    const v = filter.AND.find((it) => it.text && it.text.contains);
    // Check if v exists
    if (v) {
      mainSearchInitialValue = (v.text as StringFilterModel).contains as string;
    }
  }

  return (
    <>
      <Title title="Todos" />
      {error && <GraphqlErrors error={error} />}
      {!error && (
        <Paper variant="outlined">
          <div style={{ display: 'flex', margin: '10px' }}>
            <FilterSearchBar
              filter={filter}
              setFilter={(f) => {
                // Flush pagination
                // When a pagination is set, like you are on second page and
                // a new filter is applied, you want to start from start again.
                cleanAndSetCleanedPagination(searchParams, setSearchParams);
                // Set filter
                setFilter(f);
              }}
              mainSearchInitialValue={mainSearchInitialValue}
              mainSearchDisplay="Todos contains"
              onMainSearchChange={(newValue: string, oldValue: string) => {
                // Flush pagination
                // When a pagination is set, like you are on second page and
                // a new filter is applied, you want to start from start again.
                cleanAndSetCleanedPagination(searchParams, setSearchParams);
                // Call on main search change
                onMainSearchChange(newValue, oldValue, (f: (input: TodoFilterModel) => TodoFilterModel) => {
                  // Execute filter generator
                  const nF = f(filter);

                  // Save new filter
                  setFilter(nF);
                });
              }}
              filterDefinitionModel={todoFilterDefinitionObject}
              predefinedFilterObjects={[
                {
                  display: 'code',
                  filter: { text: { eq: 'test1', notEq: 'ok' }, createdAt: { isNotNull: true } },
                },
                {
                  display: 'test1',
                  filter: {
                    OR: [{ text: { eq: 'test1' } }, { text: { notEq: 'test2' } }],
                    AND: [
                      { text: { eq: 'test1', notEq: 'ok' }, createdAt: { isNotNull: true } },
                      { text: { notEq: 'test2' } },
                      { createdAt: { isNull: true } },
                    ],
                  },
                },
                {
                  display: 'test2',
                  description: 'lonnnng description',
                  filter: {
                    OR: [{ text: { eq: 'test1' } }, { createdAt: { notEq: dayjs().format() } }],
                  },
                },
              ]}
            />
            <div style={{ width: '10px' }} />
            <SortButton
              sort={sort}
              setSort={(nSort) => {
                setSort({ ...nSort });
              }}
              sortFields={todoSortFields}
            />
            <div style={{ width: '10px' }} />
            <GridTableViewSwitcher setGridView={setGridView} gridView={gridView} />
          </div>
          <Divider />
          <div style={{ width: '100%' }}>
            {gridView && <GridView loading={loading} data={data?.todos} />}
            {!gridView && <TableView data={data?.todos} loading={loading} sort={sort} setSort={setSort} />}
            {data && data.todos && (
              <>
                <Divider />
                <Pagination pageInfo={data.todos.pageInfo} maxPaginationSize={maxPagination} />
              </>
            )}
          </div>
        </Paper>
      )}
    </>
  );
}

export default Todos;
