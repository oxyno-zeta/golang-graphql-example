import React from 'react';
import { fireEvent, render, waitFor } from '@testing-library/react';
import { mdiPlus, mdiDelete, mdiChevronDown, mdiChevronUp } from '@mdi/js';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';

import { SortOrderFieldModel, SortOrderModel } from '~models/general';
import SortForm from './SortForm';

const testSortFields: SortOrderFieldModel[] = [
  { field: 'createdAt', display: 'common.fields.createdAt' },
  { field: 'updatedAt', display: 'common.fields.updatedAt' },
  { field: 'text', display: 'test.fields.text' },
];
type TestSortOrderModel = {
  createdAt?: SortOrderModel;
  updatedAt?: SortOrderModel;
  text?: SortOrderModel;
};

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

describe('sorts/internal/SortForm', () => {
  it('should be ok with no initial sorts', () => {
    const onSubmit = jest.fn().mockImplementation((i) => i);
    const onReset = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <SortForm<TestSortOrderModel>
        initialSorts={[]}
        onReset={onReset}
        onSubmit={onSubmit}
        sortFields={testSortFields}
      />,
    );

    expect(container).toMatchSnapshot();

    expect(container).toHaveTextContent('common.resetAction');
    expect(container).toHaveTextContent('common.applyAction');

    const buttons = container.querySelectorAll('button');
    expect(buttons).toHaveLength(3);

    const pathElement = container.querySelector('path');
    expect(pathElement).toHaveAttribute('d', mdiPlus);
  });

  it('should be ok with 1 initial sort', async () => {
    const onSubmit = jest.fn().mockImplementation((i) => i);
    const onReset = jest.fn().mockImplementation((i) => i);

    const { container, findAllByRole, findByTestId, queryAllByRole } = render(
      <SortForm<TestSortOrderModel>
        initialSorts={[{ text: 'ASC' }]}
        onReset={onReset}
        onSubmit={onSubmit}
        sortFields={testSortFields}
      />,
    );

    expect(container).toMatchSnapshot();

    expect(container).toHaveTextContent('common.resetAction');
    expect(container).toHaveTextContent('common.applyAction');

    const buttons = container.querySelectorAll('button');
    expect(buttons).toHaveLength(6);

    const autocomplInputList = await findAllByRole('combobox');
    expect(autocomplInputList).toHaveLength(2);

    expect(autocomplInputList[0]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[0]).toHaveAttribute('value', 'test.fields.text');
    expect(autocomplInputList[1]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[1]).toHaveAttribute('value', 'common.sort.asc');

    const pathElements = container.querySelectorAll('path');
    // Last button must be the add
    expect(pathElements[pathElements.length - 1]).toHaveAttribute('d', mdiPlus);

    // Find one line
    const line = await findByTestId('sort-0-text');

    expect(line.firstElementChild?.firstElementChild).toEqual(buttons[0]);
    expect(pathElements[0]).toHaveAttribute('d', mdiDelete);

    // No button group in this case
    expect(queryAllByRole('group')).toHaveLength(0);
  });

  it('should be ok with 2 initial sorts', async () => {
    const onSubmit = jest.fn().mockImplementation((i) => i);
    const onReset = jest.fn().mockImplementation((i) => i);

    const { container, findAllByRole, findByTestId } = render(
      <SortForm<TestSortOrderModel>
        initialSorts={[{ text: 'ASC' }, { createdAt: 'DESC' }]}
        onReset={onReset}
        onSubmit={onSubmit}
        sortFields={testSortFields}
      />,
    );

    expect(container).toMatchSnapshot();

    expect(container).toHaveTextContent('common.resetAction');
    expect(container).toHaveTextContent('common.applyAction');

    const buttons = container.querySelectorAll('button');
    expect(buttons).toHaveLength(11);

    const autocomplInputList = await findAllByRole('combobox');
    expect(autocomplInputList).toHaveLength(4);

    expect(autocomplInputList[0]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[0]).toHaveAttribute('value', 'test.fields.text');
    expect(autocomplInputList[1]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[1]).toHaveAttribute('value', 'common.sort.asc');
    expect(autocomplInputList[2]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[2]).toHaveAttribute('value', 'common.fields.createdAt');
    expect(autocomplInputList[3]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[3]).toHaveAttribute('value', 'common.sort.desc');

    const pathElements = container.querySelectorAll('path');
    // Last button must be the add
    expect(pathElements[pathElements.length - 1]).toHaveAttribute('d', mdiPlus);

    // Find lines
    const line1 = await findByTestId('sort-0-text');
    const line2 = await findByTestId('sort-1-createdAt');

    const buttonGroups = await findAllByRole('group');
    expect(buttonGroups).toHaveLength(2);

    expect(line1.firstElementChild?.children[0]).toEqual(buttonGroups[0]);
    expect(line1.firstElementChild?.children[1]).toEqual(buttons[1]);
    expect(pathElements[1]).toHaveAttribute('d', mdiDelete);
    expect(line2.firstElementChild?.children[0]).toEqual(buttonGroups[1]);
    expect(line2.firstElementChild?.children[1]).toEqual(buttons[5]);
    expect(pathElements[5]).toHaveAttribute('d', mdiDelete);

    expect(buttonGroups[0].children).toHaveLength(1);
    expect(buttonGroups[0].children[0].firstElementChild?.firstElementChild).toHaveAttribute('d', mdiChevronDown);
    expect(buttonGroups[1].children).toHaveLength(1);
    expect(buttonGroups[1].children[0].firstElementChild?.firstElementChild).toHaveAttribute('d', mdiChevronUp);
  });

  it('should be ok with all initial sorts', async () => {
    const onSubmit = jest.fn().mockImplementation((i) => i);
    const onReset = jest.fn().mockImplementation((i) => i);

    const { container, findAllByRole, findByTestId } = render(
      <SortForm<TestSortOrderModel>
        initialSorts={[{ text: 'ASC' }, { createdAt: 'DESC' }, { updatedAt: 'ASC' }]}
        onReset={onReset}
        onSubmit={onSubmit}
        sortFields={testSortFields}
      />,
    );

    expect(container).toMatchSnapshot();

    expect(container).toHaveTextContent('common.resetAction');
    expect(container).toHaveTextContent('common.applyAction');

    const buttons = container.querySelectorAll('button');
    expect(buttons).toHaveLength(15);

    const autocomplInputList = await findAllByRole('combobox');
    expect(autocomplInputList).toHaveLength(6);

    expect(autocomplInputList[0]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[0]).toHaveAttribute('value', 'test.fields.text');
    expect(autocomplInputList[1]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[1]).toHaveAttribute('value', 'common.sort.asc');
    expect(autocomplInputList[2]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[2]).toHaveAttribute('value', 'common.fields.createdAt');
    expect(autocomplInputList[3]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[3]).toHaveAttribute('value', 'common.sort.desc');
    expect(autocomplInputList[4]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[4]).toHaveAttribute('value', 'common.fields.updatedAt');
    expect(autocomplInputList[5]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[5]).toHaveAttribute('value', 'common.sort.asc');

    const pathElements = container.querySelectorAll('path');
    // Last button mustn't be the add (as it should be removed)
    expect(pathElements[pathElements.length - 1]).not.toHaveAttribute('d', mdiPlus);

    // Find lines
    const line1 = await findByTestId('sort-0-text');
    const line2 = await findByTestId('sort-1-createdAt');
    const line3 = await findByTestId('sort-2-updatedAt');

    const buttonGroups = await findAllByRole('group');
    expect(buttonGroups).toHaveLength(3);

    expect(line1.firstElementChild?.children[0]).toEqual(buttonGroups[0]);
    expect(line1.firstElementChild?.children[1]).toEqual(buttons[1]);
    expect(pathElements[1]).toHaveAttribute('d', mdiDelete);
    expect(line2.firstElementChild?.children[0]).toEqual(buttonGroups[1]);
    expect(line2.firstElementChild?.children[1]).toEqual(buttons[6]);
    expect(pathElements[6]).toHaveAttribute('d', mdiDelete);
    expect(line3.firstElementChild?.children[0]).toEqual(buttonGroups[2]);
    expect(line3.firstElementChild?.children[1]).toEqual(buttons[10]);
    expect(pathElements[10]).toHaveAttribute('d', mdiDelete);

    expect(buttonGroups[0].children).toHaveLength(1);
    expect(buttonGroups[0].children[0].firstElementChild?.firstElementChild).toHaveAttribute('d', mdiChevronDown);
    expect(buttonGroups[1].children).toHaveLength(2);
    expect(buttonGroups[1].children[0].firstElementChild?.firstElementChild).toHaveAttribute('d', mdiChevronUp);
    expect(buttonGroups[1].children[1].firstElementChild?.firstElementChild).toHaveAttribute('d', mdiChevronDown);
    expect(buttonGroups[2].children).toHaveLength(1);
    expect(buttonGroups[2].children[0].firstElementChild?.firstElementChild).toHaveAttribute('d', mdiChevronUp);
  });

  it('should ignore a second key in object', async () => {
    const onSubmit = jest.fn().mockImplementation((i) => i);
    const onReset = jest.fn().mockImplementation((i) => i);

    const { container, findAllByRole, findByTestId, queryAllByRole } = render(
      <SortForm<TestSortOrderModel>
        initialSorts={[{ text: 'ASC', createdAt: 'ASC' }]}
        onReset={onReset}
        onSubmit={onSubmit}
        sortFields={testSortFields}
      />,
    );

    expect(container).toMatchSnapshot();

    expect(container).toHaveTextContent('common.resetAction');
    expect(container).toHaveTextContent('common.applyAction');

    const buttons = container.querySelectorAll('button');
    expect(buttons).toHaveLength(6);

    const autocomplInputList = await findAllByRole('combobox');
    expect(autocomplInputList).toHaveLength(2);

    expect(autocomplInputList[0]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[0]).toHaveAttribute('value', 'test.fields.text');
    expect(autocomplInputList[1]).toHaveAttribute('type', 'text');
    expect(autocomplInputList[1]).toHaveAttribute('value', 'common.sort.asc');

    const pathElements = container.querySelectorAll('path');
    // Last button must be the add
    expect(pathElements[pathElements.length - 1]).toHaveAttribute('d', mdiPlus);

    // Find one line
    const line = await findByTestId('sort-0-text');

    expect(line.firstElementChild?.firstElementChild).toEqual(buttons[0]);
    expect(pathElements[0]).toHaveAttribute('d', mdiDelete);

    // No button group in this case
    expect(queryAllByRole('group')).toHaveLength(0);
  });

  it('should be ok to reset value', async () => {
    const onSubmit = jest.fn().mockImplementation((i) => i);
    const onReset = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <SortForm<TestSortOrderModel>
        initialSorts={[{ text: 'ASC' }]}
        onReset={onReset}
        onSubmit={onSubmit}
        sortFields={testSortFields}
      />,
    );

    const buttons = container.querySelectorAll('button');
    expect(buttons).toHaveLength(6);

    expect(fireEvent.click(buttons[4])).toBeTruthy();

    expect(onReset).toHaveBeenCalled();
  });

  it('should be ok to click on submit without changing anything', async () => {
    const onSubmit = jest.fn().mockImplementation((i) => i);
    const onReset = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <SortForm<TestSortOrderModel>
        initialSorts={[{ text: 'ASC' }]}
        onReset={onReset}
        onSubmit={onSubmit}
        sortFields={testSortFields}
      />,
    );

    const buttons = container.querySelectorAll('button');
    expect(buttons).toHaveLength(6);

    expect(fireEvent.click(buttons[5])).toBeTruthy();

    expect(onSubmit).toHaveBeenCalled();
    expect(onSubmit).toHaveBeenCalledWith([{ text: 'ASC' }]);
  });

  it('should be ok to click on submit with a change on value', async () => {
    const onSubmit = jest.fn().mockImplementation((i) => i);
    const onReset = jest.fn().mockImplementation((i) => i);

    const { container, findAllByRole } = render(
      <SortForm<TestSortOrderModel>
        initialSorts={[{ text: 'ASC' }]}
        onReset={onReset}
        onSubmit={onSubmit}
        sortFields={testSortFields}
      />,
    );

    const buttons = container.querySelectorAll('button');
    expect(buttons).toHaveLength(6);

    const autocomplInputList = await findAllByRole('combobox');

    expect(fireEvent.change(autocomplInputList[1], { target: { value: 'DESC' } })).toBeTruthy();
    fireEvent.keyDown(autocomplInputList[1], { key: 'ArrowDown' });
    fireEvent.keyDown(autocomplInputList[1], { key: 'Enter' });

    await waitFor(() => 0);

    expect(fireEvent.click(buttons[5])).toBeTruthy();

    expect(onSubmit).toHaveBeenCalled();
    expect(onSubmit).toHaveBeenCalledWith([{ text: 'DESC' }]);
  });

  it('should be ok to click on submit with a change on field', async () => {
    const onSubmit = jest.fn().mockImplementation((i) => i);
    const onReset = jest.fn().mockImplementation((i) => i);

    const { container, findAllByRole } = render(
      <SortForm<TestSortOrderModel>
        initialSorts={[{ text: 'ASC' }]}
        onReset={onReset}
        onSubmit={onSubmit}
        sortFields={testSortFields}
      />,
    );

    const buttons = container.querySelectorAll('button');
    expect(buttons).toHaveLength(6);

    const autocomplInputList = await findAllByRole('combobox');

    expect(fireEvent.change(autocomplInputList[0], { target: { value: 'common.fields.createdAt' } })).toBeTruthy();
    fireEvent.keyDown(autocomplInputList[0], { key: 'ArrowDown' });
    fireEvent.keyDown(autocomplInputList[0], { key: 'Enter' });

    await waitFor(() => 0);

    expect(fireEvent.click(buttons[5])).toBeTruthy();

    expect(onSubmit).toHaveBeenCalled();
    expect(onSubmit).toHaveBeenCalledWith([{ createdAt: 'ASC' }]);
  });

  it('should be ok to remove last line', async () => {
    const onSubmit = jest.fn().mockImplementation((i) => i);
    const onReset = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <SortForm<TestSortOrderModel>
        initialSorts={[{ text: 'ASC' }]}
        onReset={onReset}
        onSubmit={onSubmit}
        sortFields={testSortFields}
      />,
    );

    const buttons = container.querySelectorAll('button');
    expect(buttons).toHaveLength(6);

    expect(fireEvent.click(buttons[0])).toBeTruthy();

    expect(fireEvent.click(buttons[5])).toBeTruthy();

    expect(onSubmit).toHaveBeenCalled();
    expect(onSubmit).toHaveBeenCalledWith([]);
  });

  it('should be ok to reorder', async () => {
    const onSubmit = jest.fn().mockImplementation((i) => i);
    const onReset = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <SortForm<TestSortOrderModel>
        initialSorts={[{ text: 'ASC' }, { createdAt: 'DESC' }, { updatedAt: 'ASC' }]}
        onReset={onReset}
        onSubmit={onSubmit}
        sortFields={testSortFields}
      />,
    );

    expect(container).toMatchSnapshot();

    let buttons = container.querySelectorAll('button');
    expect(buttons).toHaveLength(15);

    // Go down
    expect(fireEvent.click(buttons[0])).toBeTruthy();
    expect(fireEvent.click(buttons[14])).toBeTruthy();

    expect(onSubmit).toHaveBeenCalled();
    expect(onSubmit).toHaveBeenNthCalledWith(1, [{ createdAt: 'DESC' }, { text: 'ASC' }, { updatedAt: 'ASC' }]);
    expect(container).toMatchSnapshot();

    buttons = container.querySelectorAll('button');
    // Go down
    expect(fireEvent.click(buttons[5])).toBeTruthy();
    expect(fireEvent.click(buttons[14])).toBeTruthy();

    expect(onSubmit).toHaveBeenCalled();
    expect(onSubmit).toHaveBeenNthCalledWith(2, [{ createdAt: 'DESC' }, { updatedAt: 'ASC' }, { text: 'ASC' }]);
    expect(container).toMatchSnapshot();

    buttons = container.querySelectorAll('button');
    // Go up
    expect(fireEvent.click(buttons[4])).toBeTruthy();
    expect(fireEvent.click(buttons[14])).toBeTruthy();

    expect(onSubmit).toHaveBeenCalled();
    expect(onSubmit).toHaveBeenNthCalledWith(3, [{ updatedAt: 'ASC' }, { createdAt: 'DESC' }, { text: 'ASC' }]);
    expect(container).toMatchSnapshot();
  });

  it('should be ok to add last available field', async () => {
    const onSubmit = jest.fn().mockImplementation((i) => i);
    const onReset = jest.fn().mockImplementation((i) => i);

    const { container, findAllByRole } = render(
      <SortForm<TestSortOrderModel>
        initialSorts={[{ text: 'ASC' }, { createdAt: 'DESC' }]}
        onReset={onReset}
        onSubmit={onSubmit}
        sortFields={testSortFields}
      />,
    );

    expect(container).toMatchSnapshot();

    const buttons = container.querySelectorAll('button');

    let autocomplInputList = await findAllByRole('combobox');

    expect(autocomplInputList[0]).toHaveAttribute('value', 'test.fields.text');
    expect(autocomplInputList[1]).toHaveAttribute('value', 'common.sort.asc');
    expect(autocomplInputList[2]).toHaveAttribute('value', 'common.fields.createdAt');
    expect(autocomplInputList[3]).toHaveAttribute('value', 'common.sort.desc');

    // Click on Add
    expect(fireEvent.click(buttons[buttons.length - 3])).toBeTruthy();

    autocomplInputList = await findAllByRole('combobox');

    expect(autocomplInputList[4]).toHaveAttribute('value', 'common.fields.updatedAt');
    expect(autocomplInputList[5]).toHaveAttribute('value', 'common.sort.asc');
    expect(container).toMatchSnapshot();
  });
});
