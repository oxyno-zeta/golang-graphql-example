import React from 'react';
import { render } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';

import NoData from './NoData';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('NoData', () => {
  it('should display no data', async () => {
    const { container } = render(<NoData />);
    // Now find errors
    expect(container).toHaveTextContent('common.noData');
    expect(container).toMatchSnapshot();
  });

  it('should display no data with specific variant', async () => {
    const { container } = render(<NoData typographyProps={{ variant: 'body2' }} />);
    // Now find errors
    expect(container).toHaveTextContent('common.noData');
    expect(container.firstChild).toHaveClass('MuiTypography-body2');
    expect(container).toMatchSnapshot();
  });
});
