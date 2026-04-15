import React from 'react';
import { render } from '@testing-library/react';
import '@testing-library/jest-dom/vitest';

import Footer from './Footer';

vi.mock('react-i18next', () => ({
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
