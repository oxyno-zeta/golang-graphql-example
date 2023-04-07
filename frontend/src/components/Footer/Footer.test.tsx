import React from 'react';
import { render } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';

import Footer from './Footer';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('Footer', () => {
  it('should display footer', async () => {
    const { container } = render(<Footer />);
    // Now find errors
    expect(container).toHaveTextContent('Todo list application');
    expect(container).toMatchSnapshot();
  });
});
