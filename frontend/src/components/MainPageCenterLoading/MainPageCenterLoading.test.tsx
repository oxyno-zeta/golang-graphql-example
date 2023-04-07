import React from 'react';
import { render } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';

import MainPageCenterLoading from './MainPageCenterLoading';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('MainPageCenterLoading', () => {
  it('should display main page center loading', async () => {
    const { container, findByRole } = render(<MainPageCenterLoading />);

    // Find progressbar
    expect(await findByRole('progressbar', { hidden: true })).not.toBeNull();
    expect(container).toMatchSnapshot();
  });
});
