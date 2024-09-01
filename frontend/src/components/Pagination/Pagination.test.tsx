import React from 'react';
import { MemoryRouter, Route, Routes } from 'react-router-dom';
import { fireEvent, render } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';

import Pagination, { Props as PaginationProps } from './Pagination';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

interface Props {
  readonly paginationProps: PaginationProps;
}

function TestComponent({ paginationProps }: Props) {
  return (
    <MemoryRouter initialEntries={['/route']}>
      <Routes>
        <Route element={<Pagination {...paginationProps} />} path="/route" />
      </Routes>
    </MemoryRouter>
  );
}

describe('Pagination', () => {
  it('should display 3 disable links when no next or previous page are present', async () => {
    const { container } = render(
      <TestComponent
        paginationProps={{
          maxPaginationSize: 10,
          pageInfo: { hasNextPage: false, hasPreviousPage: false },
        }}
      />,
    );

    const allA = container.querySelectorAll('a');
    expect(allA).toHaveLength(3);
    // Loop over
    allA.forEach((item) => {
      expect(item).toHaveClass('Mui-disabled');
    });
    expect(container).toMatchSnapshot();
  });

  it('should display only next page link as enabled with correct url', async () => {
    const { container } = render(
      <TestComponent
        paginationProps={{
          maxPaginationSize: 10,
          pageInfo: { hasNextPage: true, endCursor: 'fake-end', hasPreviousPage: false },
        }}
      />,
    );

    const allA = container.querySelectorAll('a');
    expect(allA).toHaveLength(3);
    expect(allA[0]).toHaveClass('Mui-disabled');
    expect(allA[1]).toHaveClass('Mui-disabled');
    expect(allA[2]).not.toHaveClass('Mui-disabled');
    expect(allA[2]).toHaveAttribute('href', '/route?after=fake-end&first=10');
    expect(container).toMatchSnapshot();
  });

  it('should display only previous page links as enabled with correct url', async () => {
    const { container } = render(
      <TestComponent
        paginationProps={{
          maxPaginationSize: 10,
          pageInfo: { hasNextPage: false, startCursor: 'fake-start', hasPreviousPage: true },
        }}
      />,
    );

    const allA = container.querySelectorAll('a');
    expect(allA).toHaveLength(3);
    expect(allA[0]).not.toHaveClass('Mui-disabled');
    expect(allA[0]).toHaveAttribute('href', '/route');
    expect(allA[1]).not.toHaveClass('Mui-disabled');
    expect(allA[1]).toHaveAttribute('href', '/route?before=fake-start&last=10');
    expect(allA[2]).toHaveClass('Mui-disabled');
    expect(container).toMatchSnapshot();
  });

  it('should display previous and next page links as enabled with correct url', async () => {
    const { container } = render(
      <TestComponent
        paginationProps={{
          maxPaginationSize: 10,
          pageInfo: {
            hasNextPage: true,
            startCursor: 'fake-start',
            hasPreviousPage: true,
            endCursor: 'fake-end',
          },
        }}
      />,
    );

    const allA = container.querySelectorAll('a');
    expect(allA).toHaveLength(3);
    expect(allA[0]).not.toHaveClass('Mui-disabled');
    expect(allA[0]).toHaveAttribute('href', '/route');
    expect(allA[1]).not.toHaveClass('Mui-disabled');
    expect(allA[1]).toHaveAttribute('href', '/route?before=fake-start&last=10');
    expect(allA[2]).not.toHaveClass('Mui-disabled');
    expect(allA[2]).toHaveAttribute('href', '/route?after=fake-end&first=10');
    expect(container).toMatchSnapshot();
  });

  it('should display only next page button as enabled and clickable', async () => {
    let clicked = false;
    const { container } = render(
      <TestComponent
        paginationProps={{
          maxPaginationSize: 10,
          pageInfo: { hasNextPage: true, endCursor: 'fake-end', hasPreviousPage: false },
          onNextPage: () => {
            clicked = true;
          },
        }}
      />,
    );

    const allA = container.querySelectorAll('a');
    expect(allA).toHaveLength(2);
    expect(allA[0]).toHaveClass('Mui-disabled');
    expect(allA[1]).toHaveClass('Mui-disabled');
    const allButtons = container.querySelectorAll('button');
    expect(allButtons).toHaveLength(1);
    expect(allButtons[0]).not.toHaveClass('Mui-disabled');
    expect(allButtons[0]).not.toHaveAttribute('href', '/route?after=fake-end&first=10');
    expect(container).toMatchSnapshot();

    expect(fireEvent.click(allButtons[0])).toBeTruthy();
    expect(clicked).toBeTruthy();
  });

  it('should display only previous page buttons as enabled and clickable', async () => {
    let previousPageClicked = false;
    let firstPageClicked = false;
    const { container } = render(
      <TestComponent
        paginationProps={{
          maxPaginationSize: 10,
          pageInfo: { hasNextPage: false, startCursor: 'fake-start', hasPreviousPage: true },
          onPreviousPage: () => {
            previousPageClicked = true;
          },
          onFirstPage: () => {
            firstPageClicked = true;
          },
        }}
      />,
    );

    const allA = container.querySelectorAll('a');
    expect(allA).toHaveLength(1);
    expect(allA[0]).toHaveClass('Mui-disabled');
    const allButtons = container.querySelectorAll('button');
    expect(allButtons).toHaveLength(2);
    expect(allButtons[0]).not.toHaveClass('Mui-disabled');
    expect(allButtons[0]).not.toHaveAttribute('href', '/route');
    expect(allButtons[1]).not.toHaveClass('Mui-disabled');
    expect(allButtons[1]).not.toHaveAttribute('href', '/route?before=fake-start&last=10');
    expect(container).toMatchSnapshot();

    expect(fireEvent.click(allButtons[0])).toBeTruthy();
    expect(firstPageClicked).toBeTruthy();
    expect(fireEvent.click(allButtons[1])).toBeTruthy();
    expect(previousPageClicked).toBeTruthy();
  });

  it('should display next and previous page buttons as enabled and clickable', async () => {
    let nextPageClicked = false;
    let previousPageClicked = false;
    let firstPageClicked = false;
    const { container } = render(
      <TestComponent
        paginationProps={{
          maxPaginationSize: 10,
          pageInfo: {
            hasNextPage: true,
            endCursor: 'fake-end',
            startCursor: 'fake-start',
            hasPreviousPage: true,
          },
          onNextPage: () => {
            nextPageClicked = true;
          },
          onPreviousPage: () => {
            previousPageClicked = true;
          },
          onFirstPage: () => {
            firstPageClicked = true;
          },
        }}
      />,
    );

    const allA = container.querySelectorAll('a');
    expect(allA).toHaveLength(0);
    const allButtons = container.querySelectorAll('button');
    expect(allButtons).toHaveLength(3);
    expect(allButtons[0]).not.toHaveClass('Mui-disabled');
    expect(allButtons[0]).not.toHaveAttribute('href', '/route');
    expect(allButtons[1]).not.toHaveClass('Mui-disabled');
    expect(allButtons[1]).not.toHaveAttribute('href', '/route?before=fake-start&last=10');
    expect(allButtons[2]).not.toHaveClass('Mui-disabled');
    expect(allButtons[2]).not.toHaveAttribute('href', '/route?after=fake-end&first=10');
    expect(container).toMatchSnapshot();

    expect(fireEvent.click(allButtons[0])).toBeTruthy();
    expect(firstPageClicked).toBeTruthy();
    expect(fireEvent.click(allButtons[1])).toBeTruthy();
    expect(previousPageClicked).toBeTruthy();
    expect(fireEvent.click(allButtons[2])).toBeTruthy();
    expect(nextPageClicked).toBeTruthy();
  });
});
