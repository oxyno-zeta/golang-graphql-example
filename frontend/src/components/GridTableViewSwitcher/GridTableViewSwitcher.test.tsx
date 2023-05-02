import React from 'react';
import { render, fireEvent } from '@testing-library/react';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';

import GridTableViewSwitcher from './GridTableViewSwitcher';

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('GridTableViewSwitcher', () => {
  it('should mark grid in primary color', async () => {
    const onChangeMock = jest.fn();

    const { container, findByRole } = render(<GridTableViewSwitcher onChange={onChangeMock} gridView />);

    expect(container).toMatchSnapshot();
    // Find group
    const groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');
    expect(onChangeMock).not.toHaveBeenCalled();
  });

  it('should mark table in primary color', async () => {
    const onChangeMock = jest.fn();

    const { container, findByRole } = render(<GridTableViewSwitcher onChange={onChangeMock} gridView={false} />);

    expect(container).toMatchSnapshot();
    // Find group
    const groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');
    expect(onChangeMock).not.toHaveBeenCalled();
  });

  it('should change grid marker when click on grid button and grid view is disabled', async () => {
    const onChangeMock = jest.fn();
    const gridViewValue = false;

    const { container, findByRole } = render(
      <GridTableViewSwitcher onChange={onChangeMock} gridView={gridViewValue} />,
    );

    expect(container).toMatchSnapshot();
    // Find group
    const groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');

    fireEvent.click(groupElement.children[1], 'click');

    expect(onChangeMock).toHaveBeenCalled();
    expect(onChangeMock).toHaveBeenCalledWith(true);
  });

  it('should change grid marker when click on table button and grid view is enabled', async () => {
    const onChangeMock = jest.fn();
    const gridViewValue = true;

    const { container, findByRole } = render(
      <GridTableViewSwitcher onChange={onChangeMock} gridView={gridViewValue} />,
    );

    expect(container).toMatchSnapshot();
    // Find group
    const groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');

    fireEvent.click(groupElement.children[0], 'click');

    expect(onChangeMock).toHaveBeenCalled();
    expect(onChangeMock).toHaveBeenCalledWith(false);
  });

  it("shouldn't change grid marker when click on grid button and grid view is already enabled", async () => {
    const onChangeMock = jest.fn();
    const gridViewValue = true;

    const { container, findByRole } = render(
      <GridTableViewSwitcher onChange={onChangeMock} gridView={gridViewValue} />,
    );

    expect(container).toMatchSnapshot();
    // Find group
    const groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');

    fireEvent.click(groupElement.children[1], 'click');

    expect(onChangeMock).not.toHaveBeenCalled();
  });

  it("shouldn't change grid marker when click on table button and grid view is already disabled", async () => {
    const onChangeMock = jest.fn();
    const gridViewValue = false;

    const { container, findByRole } = render(
      <GridTableViewSwitcher onChange={onChangeMock} gridView={gridViewValue} />,
    );

    expect(container).toMatchSnapshot();
    // Find group
    const groupElement = await findByRole('group');
    expect(groupElement).not.toBeNull();
    expect(groupElement.children).toHaveLength(2);
    expect(groupElement.children[0].firstChild).toHaveClass('MuiSvgIcon-colorPrimary');
    expect(groupElement.children[1].firstChild).not.toHaveClass('MuiSvgIcon-colorPrimary');

    fireEvent.click(groupElement.children[0], 'click');

    expect(onChangeMock).not.toHaveBeenCalled();
  });
});
