import React, { useContext } from 'react';
import { render, fireEvent } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import GridTableViewSwitcherContext from '~contexts/GridTableViewSwitcherContext';

import GridTableViewSwitcher from './GridTableViewSwitcher';
import GridTableViewSwitcherProvider from '../GridTableViewSwitcherProvider';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

function DisplayGridView() {
  const { isGridViewEnabled } = useContext(GridTableViewSwitcherContext);

  return <div data-testid="grid-view">{JSON.stringify(isGridViewEnabled())}</div>;
}

describe('GridTableViewSwitcher', () => {
  it('should display as table and allow to switch to grid and back to table', async () => {
    const { container, findByRole, findByTestId } = render(
      <GridTableViewSwitcherProvider>
        <GridTableViewSwitcher />
        <DisplayGridView />
      </GridTableViewSwitcherProvider>,
    );

    expect(container).toMatchSnapshot();
    // Find group
    let groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');
    let item = await findByTestId('grid-view');
    expect(item).toHaveTextContent('false');

    fireEvent.click(groupElement.children[1], 'click');

    expect(container).toMatchSnapshot();
    // Find group
    groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');
    item = await findByTestId('grid-view');
    expect(item).toHaveTextContent('true');

    fireEvent.click(groupElement.children[0], 'click');

    expect(container).toMatchSnapshot();
    // Find group
    groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');
    item = await findByTestId('grid-view');
    expect(item).toHaveTextContent('false');
  });

  it('should display as table and should not do anything to click on table again', async () => {
    const { container, findByRole, findByTestId } = render(
      <GridTableViewSwitcherProvider>
        <GridTableViewSwitcher />
        <DisplayGridView />
      </GridTableViewSwitcherProvider>,
    );

    expect(container).toMatchSnapshot();
    // Find group
    let groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');
    let item = await findByTestId('grid-view');
    expect(item).toHaveTextContent('false');

    fireEvent.click(groupElement.children[0], 'click');

    expect(container).toMatchSnapshot();
    // Find group
    groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');
    item = await findByTestId('grid-view');
    expect(item).toHaveTextContent('false');
  });

  it('should display as grid and should not do anything to click on grid again', async () => {
    const { container, findByRole, findByTestId } = render(
      <GridTableViewSwitcherProvider>
        <GridTableViewSwitcher />
        <DisplayGridView />
      </GridTableViewSwitcherProvider>,
    );

    // Find group
    let groupElement = await findByRole('group');

    fireEvent.click(groupElement.children[1], 'click');

    expect(container).toMatchSnapshot();
    // Find group
    groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');
    let item = await findByTestId('grid-view');
    expect(item).toHaveTextContent('true');

    fireEvent.click(groupElement.children[1], 'click');

    expect(container).toMatchSnapshot();
    // Find group
    groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');
    item = await findByTestId('grid-view');
    expect(item).toHaveTextContent('true');
  });
});
