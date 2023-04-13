import React from 'react';
import { render, fireEvent } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import { mdiBrightness2, mdiBrightness7 } from '@mdi/js';
import ThemeProvider from '~components/theming/ThemeProvider';

import IconToggleColorMode from './IconToggleColorMode';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

let matchMedia: (q: string) => MediaQueryList;

describe('theming/IconToggleColorMode', () => {
  beforeAll(() => {
    matchMedia = window.matchMedia;
  });

  afterAll(() => {
    window.matchMedia = matchMedia;
  });

  it('should display a light theme icon with light theme and switch to dark', async () => {
    window.matchMedia = jest.fn().mockImplementation((query) => {
      if (query === '(prefers-color-scheme: dark)') {
        return {
          matches: false,
          media: query,
          onchange: null,
          addListener: jest.fn(), // deprecated
          removeListener: jest.fn(), // deprecated
          addEventListener: jest.fn(),
          removeEventListener: jest.fn(),
          dispatchEvent: jest.fn(),
        };
      }

      return matchMedia(query);
    });

    const { container } = render(
      <ThemeProvider themeOptions={{}}>
        <IconToggleColorMode />
      </ThemeProvider>,
    );

    expect(container).toMatchSnapshot();

    // Find path
    const pathElement = container.querySelector('path');
    expect(pathElement).not.toBeNull();
    expect(pathElement).toHaveAttribute('d', mdiBrightness7);

    // Find button
    const buttonElement = container.querySelector('button');
    expect(buttonElement).not.toBeNull();
    expect(fireEvent.click(buttonElement as Element)).toBeTruthy();

    expect(pathElement).toHaveAttribute('d', mdiBrightness2);

    expect(container).toMatchSnapshot();
  });

  it('should display a dark theme icon with dark theme and switch to light', async () => {
    window.matchMedia = jest.fn().mockImplementation((query) => {
      if (query === '(prefers-color-scheme: dark)') {
        return {
          matches: true,
          media: query,
          onchange: null,
          addListener: jest.fn(), // deprecated
          removeListener: jest.fn(), // deprecated
          addEventListener: jest.fn(),
          removeEventListener: jest.fn(),
          dispatchEvent: jest.fn(),
        };
      }

      return matchMedia(query);
    });

    const { container } = render(
      <ThemeProvider themeOptions={{}}>
        <IconToggleColorMode />
      </ThemeProvider>,
    );

    expect(container).toMatchSnapshot();

    // Find path
    const pathElement = container.querySelector('path');
    expect(pathElement).not.toBeNull();
    expect(pathElement).toHaveAttribute('d', mdiBrightness2);

    // Find button
    const buttonElement = container.querySelector('button');
    expect(buttonElement).not.toBeNull();
    expect(fireEvent.click(buttonElement as Element)).toBeTruthy();

    expect(pathElement).toHaveAttribute('d', mdiBrightness7);

    expect(container).toMatchSnapshot();
  });
});
