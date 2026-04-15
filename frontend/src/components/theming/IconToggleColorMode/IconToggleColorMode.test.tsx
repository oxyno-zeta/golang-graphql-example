import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/vitest';
import { mdiBrightness2, mdiBrightness7 } from '@mdi/js';
import ThemeProvider from '~components/theming/ThemeProvider';

import IconToggleColorMode from './IconToggleColorMode';

vi.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

let matchMedia: (q: string) => MediaQueryList;

describe('theming/IconToggleColorMode', () => {
  beforeAll(() => {
    matchMedia = globalThis.matchMedia;
  });

  afterAll(() => {
    globalThis.matchMedia = matchMedia;
  });

  it('should display a light theme icon with light theme and switch to dark', async () => {
    globalThis.matchMedia = vi.fn().mockImplementation((query) => {
      if (query === '(prefers-color-scheme: dark)') {
        return {
          matches: false,
          media: query,
          onchange: null,
          addListener: vi.fn(), // deprecated
          removeListener: vi.fn(), // deprecated
          addEventListener: vi.fn(),
          removeEventListener: vi.fn(),
          dispatchEvent: vi.fn(),
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
    globalThis.matchMedia = vi.fn().mockImplementation((query) => {
      if (query === '(prefers-color-scheme: dark)') {
        return {
          matches: true,
          media: query,
          onchange: null,
          addListener: vi.fn(), // deprecated
          removeListener: vi.fn(), // deprecated
          addEventListener: vi.fn(),
          removeEventListener: vi.fn(),
          dispatchEvent: vi.fn(),
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
