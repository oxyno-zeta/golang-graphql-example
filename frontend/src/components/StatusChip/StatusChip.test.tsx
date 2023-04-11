import React from 'react';
import { render } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';

import StatusChip from './StatusChip';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('StatusChip', () => {
  it('should display label with default color (no value)', async () => {
    const { container } = render(<StatusChip label="fake-label" />);
    // Now find text
    expect(container).toHaveTextContent('fake-label');
    expect(container.firstChild).toHaveClass('MuiChip-colorDefault');
    expect(container.firstChild).toHaveClass('MuiChip-outlinedDefault');
    expect(container).toMatchSnapshot();
  });

  it('should display label with default color', async () => {
    const { container } = render(<StatusChip label="fake-label" color="default" />);
    // Now find text
    expect(container).toHaveTextContent('fake-label');
    expect(container.firstChild).toHaveClass('MuiChip-colorDefault');
    expect(container.firstChild).toHaveClass('MuiChip-outlinedDefault');
    expect(container).toMatchSnapshot();
  });

  it('should display label with error color', async () => {
    const { container } = render(<StatusChip label="fake-label" color="error" />);
    // Now find text
    expect(container).toHaveTextContent('fake-label');
    expect(container.firstChild).toHaveClass('MuiChip-colorError');
    expect(container.firstChild).toHaveClass('MuiChip-outlinedError');
    expect(container).toMatchSnapshot();
  });
});
