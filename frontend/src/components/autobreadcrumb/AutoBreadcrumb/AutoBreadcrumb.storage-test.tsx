import { gql } from '@apollo/client';
import type { MockedResponse } from '@apollo/client/testing';

export const SimpleQuery2 = gql`
  query SimpleQuery2 {
    name2
  }
`;

export const SimpleQuery1 = gql`
  query SimpleQuery1 {
    name
  }
`;

export const SimpleErrorQuery = gql`
  query SimpleError {
    error
  }
`;

export const SlowQuery = gql`
  query SlowQuery {
    slow
  }
`;

export const mockedResponses = [
  {
    request: {
      query: SimpleQuery1,
    },
    result: {
      data: {
        name: 'Query1',
      },
    },
  },
  {
    request: {
      query: SimpleQuery2,
    },
    result: {
      data: {
        name2: 'Query2',
      },
    },
  },
  {
    request: {
      query: SimpleErrorQuery,
    },
    error: new Error('Error !'),
  },
  {
    request: {
      query: SlowQuery,
    },
    delay: 5e3,
    result: {
      data: {
        slow: 'Slow',
      },
    },
  },
] as MockedResponse[];
