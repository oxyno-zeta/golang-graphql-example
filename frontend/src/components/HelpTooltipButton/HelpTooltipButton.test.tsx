import React from 'react';
import { fireEvent, render, waitFor } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';

import HelpTooltipButton from './HelpTooltipButton';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('HelpTooltipButton', () => {
  it('should display text tooltip on mouse click', async () => {
    const { container, getByText, findByRole } = render(<HelpTooltipButton tooltipTitle="fake-tooltip" />);

    expect(container).toMatchSnapshot();

    // Get child
    const child = container.querySelector('svg');

    expect(child).not.toBeNull();

    expect(fireEvent.click(child as SVGElement)).toBeTruthy();
    // Workaround to avoid "react component change without any act called"...
    await waitFor(() => 0);
    await waitFor(() => {
      expect(getByText('fake-tooltip')).toBeInTheDocument();
    });
    expect(await findByRole('tooltip')).not.toBeNull();
  });

  it('should display text tooltip on mouse over only', async () => {
    const { container, getByText, findByRole } = render(<HelpTooltipButton tooltipTitle="fake-tooltip" />);

    expect(container).toMatchSnapshot();

    // Get child
    const child = container.querySelector('span');

    expect(child).not.toBeNull();

    expect(fireEvent.mouseOver(child as HTMLSpanElement)).toBeTruthy();
    await waitFor(() => {
      expect(getByText('fake-tooltip')).toBeInTheDocument();
    });
    expect(await findByRole('tooltip')).not.toBeNull();

    expect(container).toMatchSnapshot();

    expect(fireEvent.mouseLeave(child as HTMLSpanElement)).toBeTruthy();
    // Workaround to avoid "react component change without any act called"...
    await waitFor(() => 0);
    await waitFor(() => {
      expect(container).not.toHaveTextContent('fake-tooltip');
    });

    expect(container).toMatchSnapshot();
  });

  it('should display element tooltip on mouse over only', async () => {
    const { container, getByText, findByRole } = render(<HelpTooltipButton tooltipTitle={<>fake-tooltip</>} />);

    expect(container).toMatchSnapshot();

    // Get child
    const child = container.querySelector('span');

    expect(child).not.toBeNull();

    expect(fireEvent.mouseOver(child as HTMLSpanElement)).toBeTruthy();
    await waitFor(() => {
      expect(getByText('fake-tooltip')).toBeInTheDocument();
    });
    expect(await findByRole('tooltip')).not.toBeNull();

    expect(container).toMatchSnapshot();

    expect(fireEvent.mouseLeave(child as HTMLSpanElement)).toBeTruthy();
    // Workaround to avoid "react component change without any act called"...
    await waitFor(() => 0);
    await waitFor(() => {
      expect(container).not.toHaveTextContent('fake-tooltip');
    });

    expect(container).toMatchSnapshot();
  });

  it('should display element tooltip on mouse click', async () => {
    const { container, getByText, findByRole } = render(<HelpTooltipButton tooltipTitle={<>fake-tooltip</>} />);

    expect(container).toMatchSnapshot();

    // Get child
    const child = container.querySelector('svg');

    expect(child).not.toBeNull();

    expect(fireEvent.click(child as SVGElement)).toBeTruthy();
    // Workaround to avoid "react component change without any act called"...
    await waitFor(() => 0);
    await waitFor(() => {
      expect(getByText('fake-tooltip')).toBeInTheDocument();
    });
    expect(await findByRole('tooltip')).not.toBeNull();

    expect(container).toMatchSnapshot();
    // Workaround to avoid "react component change without any act called"...
    await waitFor(() => 0);
  });
});
