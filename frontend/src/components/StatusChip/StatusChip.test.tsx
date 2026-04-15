import React from 'react';
import { render } from '@testing-library/react';
import '@testing-library/jest-dom/vitest';

import StatusChip from './StatusChip';

vi.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('StatusChip', () => {
  it('should display label with default color (no value)', async () => {
    const { container } = render(<StatusChip label="fake-label" />);
    // Now find text
    expect(container).toHaveTextContent('fake-label');
    expect(container.firstChild).toHaveClass('MuiChip-colorDefault');
    expect(container).toMatchSnapshot();
  });

  it('should display label with default color', async () => {
    const { container } = render(<StatusChip color="default" label="fake-label" />);
    // Now find text
    expect(container).toHaveTextContent('fake-label');
    expect(container.firstChild).toHaveClass('MuiChip-colorDefault');
    expect(container).toMatchSnapshot();
  });

  it('should display label with error color', async () => {
    const { container } = render(<StatusChip color="error" label="fake-label" />);
    // Now find text
    expect(container).toHaveTextContent('fake-label');
    expect(container.firstChild).toHaveClass('MuiChip-colorError');
    expect(container).toMatchSnapshot();
  });
});
