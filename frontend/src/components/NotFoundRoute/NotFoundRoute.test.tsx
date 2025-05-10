import React from 'react';
import { render } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';

import NotFoundRoute from './NotFoundRoute';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('NotFoundRoute', () => {
  it('should display no data', async () => {
    const { container } = render(<NotFoundRoute />);
    // Now find errors
    expect(container).toHaveTextContent('common.routeNotFound');
    expect(container).toMatchSnapshot();

    expect(document.title).toEqual('common.routeNotFound');
  });

  it('should display no data with specific variant', async () => {
    const { container } = render(<NotFoundRoute typographyProps={{ variant: 'body2' }} />);
    // Now find errors
    expect(container).toHaveTextContent('common.routeNotFound');
    expect(container.firstChild).toHaveClass('MuiTypography-body2');
    expect(container).toMatchSnapshot();

    expect(document.title).toEqual('common.routeNotFound');
  });
});
