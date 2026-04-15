import React from 'react';
import { render } from '@testing-library/react';
import '@testing-library/jest-dom/vitest';

import CenterLoading from './CenterLoading';

vi.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('CenterLoading', () => {
  it('should display main page center loading', async () => {
    const { container, findByRole } = render(<CenterLoading />);

    // Find progressbar
    expect(await findByRole('progressbar')).not.toBeNull();
    expect(container).toMatchSnapshot();
  });
});
