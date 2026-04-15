import React from 'react';
import { render } from '@testing-library/react';
import '@testing-library/jest-dom/vitest';

import MainPageCenterLoading from './MainPageCenterLoading';

vi.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('MainPageCenterLoading', () => {
  it('should display main page center loading', async () => {
    const { container, findByRole } = render(<MainPageCenterLoading />);

    // Find progressbar
    expect(await findByRole('progressbar', { hidden: true })).not.toBeNull();
    expect(container).toMatchSnapshot();
  });
});
