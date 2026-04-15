import React from 'react';
import { render } from '@testing-library/react';
import '@testing-library/jest-dom/vitest';

import NoData from './NoData';

vi.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('NoData', () => {
  it('should display no data', async () => {
    const { container } = render(<NoData />);
    // Now find text
    expect(container).toHaveTextContent('common.noData');
    expect(container).toMatchSnapshot();
  });

  it('should display no data with specific variant', async () => {
    const { container } = render(<NoData typographyProps={{ variant: 'body2' }} />);
    // Now find text
    expect(container).toHaveTextContent('common.noData');
    expect(container.firstChild).toHaveClass('MuiTypography-body2');
    expect(container).toMatchSnapshot();
  });
});
