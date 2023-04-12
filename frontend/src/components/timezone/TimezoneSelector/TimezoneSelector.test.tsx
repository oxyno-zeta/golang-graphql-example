import React from 'react';
import { fireEvent, render, screen, waitFor } from '@testing-library/react';
import * as dayjs from 'dayjs';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';
import TimezoneProvider from '~components/timezone/TimezoneProvider';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';

import TimezoneSelector from './TimezoneSelector';
import useTimezone from '../useTimezone';

// Extend dayjs
dayjs.extend(localizedFormat);
dayjs.extend(utc);
dayjs.extend(timezone);

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

function Result() {
  return <div id="result">{useTimezone()}</div>;
}

describe('timezone/TimezoneSelector', () => {
  it('should display an autocomplete with default value', async () => {
    const { container, findByRole, findByTitle } = render(
      <TimezoneProvider>
        <TimezoneSelector />
      </TimezoneProvider>,
    );

    expect(container).toMatchSnapshot();

    expect(container.firstChild).toHaveClass('MuiAutocomplete-root');

    expect(container).toHaveTextContent('common.timezone');

    const combobox = await findByRole('combobox');
    expect(combobox).toHaveAttribute('type', 'text');
    expect(combobox).toHaveAttribute('value', 'UTC');

    expect(await findByTitle('common.clearAction')).not.toBeNull();
    expect(await findByTitle('common.openAction')).not.toBeNull();
  });

  it('should display a list of timezone when a click is done and to select it', async () => {
    const { container, findByTitle, rerender } = render(
      <TimezoneProvider>
        <TimezoneSelector />
      </TimezoneProvider>,
    );

    expect(container).toMatchSnapshot();

    expect(container.firstChild).toHaveClass('MuiAutocomplete-root');

    rerender(
      <TimezoneProvider>
        <Result />
      </TimezoneProvider>,
    );
    expect(container).toMatchSnapshot();
    expect(container).toHaveTextContent('UTC');

    rerender(
      <TimezoneProvider>
        <TimezoneSelector />
      </TimezoneProvider>,
    );

    const openButton = await findByTitle('common.openAction');
    expect(openButton).not.toBeNull();

    expect(fireEvent.click(openButton)).toBeTruthy();

    await waitFor(async () => {
      expect(await screen.findByRole('presentation')).toBeInTheDocument();
    });

    expect(container).toMatchSnapshot();
    expect(container.firstChild).toHaveClass('Mui-expanded');
    expect(container.firstChild).toHaveClass('Mui-focused');

    const valuesContainer = await screen.findByRole('presentation');
    expect(valuesContainer).toMatchSnapshot();

    // Select
    const options = await screen.findAllByRole('option');
    fireEvent.click(options[0] as Element);

    expect(valuesContainer).toMatchSnapshot();

    rerender(
      <TimezoneProvider>
        <Result />
      </TimezoneProvider>,
    );
    expect(container).toMatchSnapshot();
    expect(container).toHaveTextContent('Africa/Abidjan');
  });
});
