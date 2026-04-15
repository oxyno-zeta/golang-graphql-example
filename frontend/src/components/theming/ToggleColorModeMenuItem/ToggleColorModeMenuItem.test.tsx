import React from 'react';
import { render, fireEvent } from '@testing-library/react';
import '@testing-library/jest-dom/vitest';
import ThemeProvider from '~components/theming/ThemeProvider';

import ToggleColorModeMenuItem from './ToggleColorModeMenuItem';

vi.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

let matchMedia: (q: string) => MediaQueryList;

describe('theming/ToggleColorModeMenuItem', () => {
  beforeAll(() => {
    matchMedia = globalThis.matchMedia;
  });

  afterAll(() => {
    globalThis.matchMedia = matchMedia;
  });

  it('should display the light theme as selected and switch to dark', async () => {
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
        <ToggleColorModeMenuItem />
      </ThemeProvider>,
    );

    expect(container).toMatchSnapshot();

    expect(container).toHaveTextContent('common.themeTitle');
    expect(container).toHaveTextContent('common.darkThemeSelector');
    expect(container).toHaveTextContent('common.lightThemeSelector');

    const allButtons = container.querySelectorAll('button');

    expect(allButtons[0]).toHaveAttribute('value', 'dark');
    expect(allButtons[1]).toHaveAttribute('value', 'light');

    expect(allButtons[1]).toHaveClass('Mui-selected');

    expect(fireEvent.click(allButtons[0])).toBeTruthy();

    expect(allButtons[0]).toHaveClass('Mui-selected');

    expect(container).toMatchSnapshot();
  });

  it('should display the dark theme as selected and switch to light', async () => {
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
        <ToggleColorModeMenuItem />
      </ThemeProvider>,
    );

    expect(container).toMatchSnapshot();

    expect(container).toHaveTextContent('common.themeTitle');
    expect(container).toHaveTextContent('common.darkThemeSelector');
    expect(container).toHaveTextContent('common.lightThemeSelector');

    const allButtons = container.querySelectorAll('button');

    expect(allButtons[0]).toHaveAttribute('value', 'dark');
    expect(allButtons[1]).toHaveAttribute('value', 'light');

    expect(allButtons[0]).toHaveClass('Mui-selected');

    expect(fireEvent.click(allButtons[1])).toBeTruthy();

    expect(allButtons[1]).toHaveClass('Mui-selected');

    expect(container).toMatchSnapshot();
  });
});
