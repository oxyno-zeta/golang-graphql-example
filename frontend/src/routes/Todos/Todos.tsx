import React, { useContext } from 'react';
import { gql, useQuery } from '@apollo/client';
import Paper from '@mui/material/Paper';
import Divider from '@mui/material/Divider';
import dayjs from 'dayjs';
import { useSearchParams } from 'react-router-dom';
import Title from '~components/Title';
import FilterSearchBar from '~components/filters/FilterSearchBar';
import ErrorsDisplay from '~components/ErrorsDisplay';
import SortButton from '~components/sorts/SortButton';
import Pagination from '~components/Pagination';
import GridTableViewSwitcher from '~components/gridTableViewSwitch/GridTableViewSwitcher';
import { onMainSearchChangeContains } from '~components/filters/utils/mainSearch';
import TopListContainer from '~components/TopListContainer';
import {
  TodoSortOrderModel,
  TodoFilterModel,
  todoFilterDefinitionObject,
  todoSortFields,
  TodoModel,
} from '~models/todos';
import { ConnectionModel, FilterQueryParamName, SortQueryParamName, StringFilterModel } from '~models/general';
import { getPaginationFromSearchParams, cleanAndSetCleanedPagination } from '~utils/pagination';
import { getJSONObjectFromSearchParam, setJSONObjectSearchParam } from '~utils/urlSearchParams';
import GridTableViewSwitcherContext from '~contexts/GridTableViewSwitcherContext';
import GridView from './components/GridView';
import TableView from './components/TableView';

const GET_TODOS_QUERY = gql`
  query getTodos(
    $first: Int
    $last: Int
    $before: String
    $after: String
    $filter: TodoFilter
    $sorts: [TodoSortOrder]
  ) {
    todos(first: $first, last: $last, before: $before, after: $after, filter: $filter, sorts: $sorts) {
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
  sorts?: TodoSortOrderModel[];
  filter?: TodoFilterModel | null;
}

const maxPagination = 20;
const initialPagination = { first: maxPagination };

function Todos() {
  // Get search params
  const [searchParams, setSearchParams] = useSearchParams();
  // Filter, pagination and sort values
  const filter = getJSONObjectFromSearchParam<TodoFilterModel>(FilterQueryParamName, {}, searchParams);
  const sorts = getJSONObjectFromSearchParam<TodoSortOrderModel[]>(
    SortQueryParamName,
    [{ createdAt: 'DESC' }],
    searchParams,
  );
  const pagination = getPaginationFromSearchParams(initialPagination, maxPagination, searchParams, setSearchParams);

  // Setter
  const setSorts = (data: TodoSortOrderModel[]) => {
    setJSONObjectSearchParam(SortQueryParamName, data, searchParams, setSearchParams);
  };
  const setFilter = (data: TodoFilterModel) => {
    setJSONObjectSearchParam(FilterQueryParamName, data, searchParams, setSearchParams);
  };

  // Get data from context
  const gridView = useContext(GridTableViewSwitcherContext).isGridViewEnabled();

  // Call graphql
  const { data, loading, error } = useQuery<QueryResult, QueryVariables>(GET_TODOS_QUERY, {
    variables: { ...pagination, sorts, filter },
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
      {error && <ErrorsDisplay error={error} />}
      {!error && (
        <Paper variant="outlined">
          <TopListContainer>
            <FilterSearchBar
              filter={filter}
              onSubmit={(f) => {
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
                onMainSearchChangeContains<TodoFilterModel>(
                  'text',
                  newValue,
                  oldValue,
                  (f: (input: TodoFilterModel) => TodoFilterModel) => {
                    // Execute filter generator
                    const nF = f(filter);

                    // Save new filter
                    setFilter(nF);
                  },
                );
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
            <SortButton
              sorts={sorts}
              onSubmit={(nSort) => {
                setSorts(nSort);
              }}
              sortFields={todoSortFields}
            />
            <GridTableViewSwitcher />
          </TopListContainer>
          <Divider />
          <div style={{ width: '100%' }}>
            {gridView && <GridView loading={loading} data={data?.todos} />}
            {!gridView && <TableView data={data?.todos} loading={loading} sorts={sorts} setSorts={setSorts} />}
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
