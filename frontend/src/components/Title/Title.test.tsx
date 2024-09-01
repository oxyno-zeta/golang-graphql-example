import React from 'react';
import { render } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import { mdiMenu, mdiIdentifier } from '@mdi/js';
import IconButton from '@mui/material/IconButton';
import SvgIcon from '@mui/material/SvgIcon';

import Title from './Title';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('Title', () => {
  it('should display title only', async () => {
    const { container } = render(<Title title="Fake title" />);
    // Now find text
    expect(container).toHaveTextContent('Fake title');
    expect(container).toMatchSnapshot();
  });

  it('should display title with left element only', async () => {
    const { container, findByRole } = render(
      <Title
        leftElement={
          <IconButton>
            <SvgIcon>
              <path d={mdiMenu} />
            </SvgIcon>
          </IconButton>
        }
        title="Fake title"
      />,
    );
    // Now find text
    expect(container).toHaveTextContent('Fake title');
    // Find button
    const buttonElement = await findByRole('button');
    expect(buttonElement).not.toBeNull();
    expect(buttonElement).toEqual(container.firstChild?.firstChild);
    expect(container).toMatchSnapshot();
  });

  it('should display title with right element only', async () => {
    const { container, findByRole } = render(
      <Title
        rightElement={
          <IconButton>
            <SvgIcon>
              <path d={mdiMenu} />
            </SvgIcon>
          </IconButton>
        }
        title="Fake title"
      />,
    );
    // Now find text
    expect(container).toHaveTextContent('Fake title');
    // Find button
    const buttonElement = await findByRole('button');
    expect(buttonElement).not.toBeNull();
    expect(buttonElement).toEqual(container.firstChild?.lastChild);
    expect(container).toMatchSnapshot();
  });

  it('should display title with right and left element', async () => {
    const { container, findAllByRole } = render(
      <Title
        leftElement={
          <IconButton>
            <SvgIcon>
              <path d={mdiMenu} />
            </SvgIcon>
          </IconButton>
        }
        rightElement={
          <IconButton>
            <SvgIcon>
              <path d={mdiIdentifier} />
            </SvgIcon>
          </IconButton>
        }
        title="Fake title"
      />,
    );
    // Now find text
    expect(container).toHaveTextContent('Fake title');
    // Find button
    const buttonElements = await findAllByRole('button');
    expect(buttonElements).not.toBeNull();
    expect(buttonElements).toHaveLength(2);
    expect(buttonElements[0]).toEqual(container.firstChild?.firstChild);
    expect(buttonElements[0].firstChild?.firstChild).toHaveAttribute('d', mdiMenu);
    expect(buttonElements[1]).toEqual(container.firstChild?.lastChild);
    expect(buttonElements[1].firstChild?.firstChild).toHaveAttribute('d', mdiIdentifier);
    expect(container).toMatchSnapshot();
  });
});
