import React from 'react';
import { fireEvent, render, waitFor, screen } from '@testing-library/react';
import { mdiPlus, mdiDelete, mdiPlusBoxMultiple } from '@mdi/js';
// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';
import * as dayjs from 'dayjs';
import localizedFormat from 'dayjs/plugin/localizedFormat';
import utc from 'dayjs/plugin/utc';
import timezone from 'dayjs/plugin/timezone';

import {
  type BooleanFilterModel,
  type DateFilterModel,
  type FilterDefinitionFieldsModel,
  type StringFilterModel,
} from '~models/general';
import { booleanOperations, dateOperations, stringOperations } from '~models/general-operations';
import FilterForm from './FilterForm';

// Extend dayjs
dayjs.extend(localizedFormat);
dayjs.extend(utc);
dayjs.extend(timezone);

interface TestFilterModel {
  AND?: TestFilterModel[];
  OR?: TestFilterModel[];
  createdAt?: DateFilterModel;
  text?: StringFilterModel;
  done?: BooleanFilterModel;
}

const testFilterDefinitionObject: FilterDefinitionFieldsModel = {
  createdAt: {
    display: 'common.fields.createdAt',
    description: 'longgggggggggggggggggggg description',
    operations: dateOperations,
  },
  text: {
    display: 'todos.fields.text',
    operations: stringOperations,
  },
  done: {
    display: 'todos.fields.done',
    operations: booleanOperations,
  },
};

jest.mock('react-i18next', () => ({
  useTranslation: () => ({ t: (key: string) => key }),
}));

function generatePseudoRandomSuffix(random: number) {
  return (random + 1).toString(36).substring(2);
}

describe('filters/internal/FilterForm', () => {
  let randomCount = 0;

  beforeEach(() => {
    jest.spyOn(global.Math, 'random').mockImplementation(() => {
      randomCount += 0.0001;

      return randomCount;
    });
  });

  afterEach(() => {
    jest.spyOn(global.Math, 'random').mockRestore();
    randomCount = 0;
  });

  it('should be ok to just display no initial filters', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm filterDefinitionModel={testFilterDefinitionObject} initialFilter={{}} onChange={onChange} />,
    );

    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenCalledWith(null);

    const group = await screen.findByRole('group');
    expect(group.children[0]).toHaveAttribute('type', 'button');
    expect(group.children[1]).toHaveAttribute('type', 'button');
    expect(group.children[0]).toHaveClass('MuiButton-contained MuiButton-containedPrimary');
    expect(group.children[1]).toHaveClass('MuiButton-outlined MuiButton-outlinedPrimary');
    expect(group.children[0]).toHaveTextContent('common.operations.and');
    expect(group.children[1]).toHaveTextContent('common.operations.or');

    // Get parent
    const parentElem = group.parentElement as Element;
    expect(parentElem.children[1]).toHaveAttribute('type', 'button');
    expect(parentElem.children[2]).toHaveAttribute('type', 'button');
    expect(parentElem.children[1].firstChild).toHaveClass('MuiSvgIcon-root');
    expect(parentElem.children[1].firstChild?.firstChild).toHaveAttribute('d', mdiPlus);
    expect(parentElem.children[2].firstChild).toHaveClass('MuiSvgIcon-root');
    expect(parentElem.children[2].firstChild?.firstChild).toHaveAttribute('d', mdiPlusBoxMultiple);

    const firstLine = await screen.findByTestId('undefinedroot');

    expect(firstLine.firstChild?.firstChild).toHaveAttribute('type', 'button');
    expect(firstLine.firstChild?.firstChild).toHaveClass('MuiButtonBase-root');
    expect(firstLine.firstChild?.firstChild?.firstChild?.firstChild).toHaveAttribute('d', mdiDelete);

    expect(firstLine).toHaveTextContent('common.filter.field');

    const inputElement = await screen.findByRole('combobox');
    expect(inputElement).toHaveAttribute('placeholder', 'common.filter.field');
    expect(inputElement).toHaveAttribute('value', '');

    expect(container).toHaveTextContent('common.fieldValidationError.required');
  });

  it('should be ok to interact with no initial filters and perform a full clean', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm filterDefinitionModel={testFilterDefinitionObject} initialFilter={{}} onChange={onChange} />,
    );

    expect(container).toMatchSnapshot();

    const group = await screen.findByRole('group');
    expect(group.children[0]).toHaveClass('MuiButton-contained MuiButton-containedPrimary');
    expect(group.children[1]).toHaveClass('MuiButton-outlined MuiButton-outlinedPrimary');
    expect(group.children[0]).toHaveTextContent('common.operations.and');
    expect(group.children[1]).toHaveTextContent('common.operations.or');

    // Click on and
    expect(fireEvent.click(group.children[0])).toBeTruthy();
    expect(group.children[0]).toHaveClass('MuiButton-contained MuiButton-containedPrimary');
    expect(group.children[1]).toHaveClass('MuiButton-outlined MuiButton-outlinedPrimary');
    expect(onChange).toHaveBeenLastCalledWith(null);

    expect(container).toMatchSnapshot();

    // Click on or
    expect(fireEvent.click(group.children[1])).toBeTruthy();
    expect(group.children[0]).toHaveClass('MuiButton-outlined MuiButton-outlinedPrimary');
    expect(group.children[1]).toHaveClass('MuiButton-contained MuiButton-containedPrimary');
    expect(onChange).toHaveBeenLastCalledWith(null);

    const firstLine = await screen.findByTestId('undefinedroot');

    // Delete line
    expect(fireEvent.click(firstLine.firstChild?.firstChild as Element)).toBeTruthy();
    expect(container).toMatchSnapshot();
    const firstLine2 = screen.queryByTestId('undefinedroot');
    expect(firstLine2).toBeNull();

    // Get parent
    const parentElem = group.parentElement as Element;

    // Add line
    expect(fireEvent.click(parentElem.children[1])).toBeTruthy();
    await waitFor(() => 0);
    expect(onChange).toHaveBeenLastCalledWith(null);
    expect(container).toMatchSnapshot();
    let line3 = screen.queryByTestId(`line-${generatePseudoRandomSuffix(randomCount)}`);
    expect(line3).not.toBeNull();
    // Delete line
    expect(fireEvent.click((line3 as Element).firstChild?.firstChild as Element)).toBeTruthy();
    expect(container).toMatchSnapshot();
    line3 = screen.queryByTestId(`line-${generatePseudoRandomSuffix(randomCount)}`);
    expect(line3).toBeNull();

    // Add group
    expect(fireEvent.click(parentElem.children[2])).toBeTruthy();
    await waitFor(() => 0);
    expect(onChange).toHaveBeenLastCalledWith(null);
    expect(container).toMatchSnapshot();
    const groups = await screen.findAllByRole('group');
    expect(groups).toHaveLength(2);
    const line4 = screen.queryByTestId('undefinedroot');
    expect(line4).not.toBeNull();
    expect(groups[1].children[0]).toHaveClass('MuiButton-contained MuiButton-containedPrimary');
    expect(groups[1].children[1]).toHaveClass('MuiButton-outlined MuiButton-outlinedPrimary');
    expect(groups[1].children[0]).toHaveTextContent('common.operations.and');
    expect(groups[1].children[1]).toHaveTextContent('common.operations.or');
    // Get parent
    const parentElem2 = groups[1].parentElement as Element;
    expect(parentElem2.children[1]).toHaveAttribute('type', 'button');
    expect(parentElem2.children[2]).toHaveAttribute('type', 'button');
    expect(parentElem2.children[3]).toHaveAttribute('type', 'button');
    expect(parentElem2.children[1].firstChild).toHaveClass('MuiSvgIcon-root');
    expect(parentElem2.children[1].firstChild?.firstChild).toHaveAttribute('d', mdiPlus);
    expect(parentElem2.children[2].firstChild).toHaveClass('MuiSvgIcon-root');
    expect(parentElem2.children[2].firstChild?.firstChild).toHaveAttribute('d', mdiPlusBoxMultiple);
    expect(parentElem2.children[3].firstChild).toHaveClass('MuiSvgIcon-root');
    expect(parentElem2.children[3].firstChild?.firstChild).toHaveAttribute('d', mdiDelete);

    const inputElement = await screen.findByRole('combobox');
    expect(inputElement).toHaveAttribute('placeholder', 'common.filter.field');
    expect(inputElement).toHaveAttribute('value', '');
    expect(container).toHaveTextContent('common.fieldValidationError.required');

    // Delete line
    expect(fireEvent.click(parentElem2.children[3].firstChild?.firstChild as Element)).toBeTruthy();
    expect(container).toMatchSnapshot();
    line3 = screen.queryByTestId('undefinedroot');
    expect(line3).toBeNull();
    const groups2 = await screen.findAllByRole('group');
    expect(groups2).toHaveLength(1);
  });

  it('should be ok to just display a simple field without value and should be able to change it', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm
        filterDefinitionModel={testFilterDefinitionObject}
        initialFilter={{ text: { isNull: true } } as TestFilterModel}
        onChange={onChange}
      />,
    );

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenCalledWith({ text: { isNull: true } } as TestFilterModel);

    const inputElements = container.querySelectorAll('input');
    expect(inputElements).toHaveLength(2);

    expect(container).toHaveTextContent('common.filter.field');
    expect(container).toHaveTextContent('common.filter.operation');

    expect(inputElements[0]).toHaveAttribute('placeholder', 'common.filter.field');
    expect(inputElements[0]).toHaveAttribute('value', 'todos.fields.text');
    expect(inputElements[1]).toHaveAttribute('placeholder', 'common.filter.operation');
    expect(inputElements[1]).toHaveAttribute('value', 'common.operations.isNull');

    expect(container).not.toHaveTextContent('common.fieldValidationError.required');

    // Get second action
    const buttons = await screen.findAllByTitle('common.openAction');
    expect(buttons).toHaveLength(2);

    // Open second autocomplete
    expect(fireEvent.click(buttons[1])).toBeTruthy();
    const role2 = await screen.findByRole('presentation');
    expect(role2).toMatchSnapshot();
    Object.values(stringOperations).forEach((v) => {
      expect(role2).toHaveTextContent(v.display);
      if (v.description) {
        expect(role2).toHaveTextContent(v.description);
      }
    });
    // Select value
    fireEvent.change(inputElements[1], { target: { value: 'isNotNull' } });
    fireEvent.keyDown(inputElements[1], { key: 'ArrowDown' });
    fireEvent.keyDown(inputElements[1], { key: 'Enter' });
    expect(onChange).toHaveBeenLastCalledWith({ text: { isNotNull: true } } as TestFilterModel);

    // Open first autocomplete
    expect(fireEvent.click(buttons[0])).toBeTruthy();
    const role1 = await screen.findByRole('presentation');
    expect(role1).toMatchSnapshot();
    Object.values(testFilterDefinitionObject).forEach((v) => {
      expect(role1).toHaveTextContent(v.display);
      if (v.description) {
        expect(role1).toHaveTextContent(v.description);
      }
    });
  });

  it('should be ok to just display a simple field with simple text value and should be able to change it', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm
        filterDefinitionModel={testFilterDefinitionObject}
        initialFilter={{ text: { eq: 'fake' } } as TestFilterModel}
        onChange={onChange}
      />,
    );

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenCalledWith({ text: { eq: 'fake', caseInsensitive: true } } as TestFilterModel);

    const inputElements = container.querySelectorAll('input');
    expect(inputElements).toHaveLength(3);

    expect(container).toHaveTextContent('common.filter.field');
    expect(container).toHaveTextContent('common.filter.operation');
    expect(container).toHaveTextContent('common.filter.value');

    expect(inputElements[0]).toHaveAttribute('placeholder', 'common.filter.field');
    expect(inputElements[0]).toHaveAttribute('value', 'todos.fields.text');
    expect(inputElements[1]).toHaveAttribute('placeholder', 'common.filter.operation');
    expect(inputElements[1]).toHaveAttribute('value', 'common.operations.eq');
    expect(inputElements[2]).toHaveAttribute('placeholder', 'common.filter.value');
    expect(inputElements[2]).toHaveAttribute('value', 'fake');

    expect(container).not.toHaveTextContent('common.fieldValidationError.required');

    // Get second action
    const buttons = await screen.findAllByTitle('common.openAction');
    expect(buttons).toHaveLength(2);

    fireEvent.change(inputElements[2], { target: { value: 'foo' } });
    expect(onChange).toHaveBeenLastCalledWith({ text: { eq: 'foo', caseInsensitive: true } } as TestFilterModel);
  });

  it('should be ok to just display a simple field with multiples text values and should be able to change it', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm
        filterDefinitionModel={testFilterDefinitionObject}
        initialFilter={{ text: { in: ['fake'] } } as TestFilterModel}
        onChange={onChange}
      />,
    );

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenCalledWith({ text: { in: ['fake'], caseInsensitive: true } } as TestFilterModel);

    const inputElements = container.querySelectorAll('input');
    expect(inputElements).toHaveLength(3);

    expect(container).toHaveTextContent('common.filter.field');
    expect(container).toHaveTextContent('common.filter.operation');
    expect(container).toHaveTextContent('common.filter.value');

    expect(inputElements[0]).toHaveAttribute('placeholder', 'common.filter.field');
    expect(inputElements[0]).toHaveAttribute('value', 'todos.fields.text');
    expect(inputElements[1]).toHaveAttribute('placeholder', 'common.filter.operation');
    expect(inputElements[1]).toHaveAttribute('value', 'common.operations.in');
    expect(inputElements[2]).toHaveAttribute('placeholder', 'common.filter.value');
    expect(inputElements[2]).toHaveAttribute('value', '');
    expect(container).toHaveTextContent('fake');
    expect(container).not.toHaveTextContent('common.fieldValidationError.required');

    // Get chip
    const chipElement1 = inputElements[2].parentElement?.firstChild;
    expect(chipElement1).not.toBeNull();
    // Click on it to remove it
    fireEvent.click(chipElement1?.lastChild as Element);
    expect(container).toMatchSnapshot();
    expect(container).not.toHaveTextContent('fake');
    expect(container).toHaveTextContent('common.fieldValidationError.required');
    expect(onChange).toHaveBeenLastCalledWith(null);
  });

  it('should be ok to just display a simple field with simple enum value and should be able to change it', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm
        filterDefinitionModel={testFilterDefinitionObject}
        initialFilter={{ done: { eq: true } } as TestFilterModel}
        onChange={onChange}
      />,
    );

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenCalledWith({ done: { eq: true } } as TestFilterModel);

    const inputElements = container.querySelectorAll('input');
    expect(inputElements).toHaveLength(3);

    expect(container).toHaveTextContent('common.filter.field');
    expect(container).toHaveTextContent('common.filter.operation');
    expect(container).toHaveTextContent('common.filter.value');

    expect(inputElements[0]).toHaveAttribute('placeholder', 'common.filter.field');
    expect(inputElements[0]).toHaveAttribute('value', 'todos.fields.done');
    expect(inputElements[1]).toHaveAttribute('placeholder', 'common.filter.operation');
    expect(inputElements[1]).toHaveAttribute('value', 'common.operations.eq');
    expect(inputElements[2]).toHaveAttribute('placeholder', 'common.filter.value');
    expect(inputElements[2]).toHaveAttribute('value', 'common.boolean.true');

    expect(container).not.toHaveTextContent('common.fieldValidationError.required');

    // Get third action
    const buttons = await screen.findAllByTitle('common.openAction');
    expect(buttons).toHaveLength(3);

    // Open third autocomplete
    expect(fireEvent.click(buttons[1])).toBeTruthy();
    const role3 = await screen.findByRole('presentation');
    expect(role3).toMatchSnapshot();
    Object.values(booleanOperations).forEach((v) => {
      expect(role3).toHaveTextContent(v.display);
      if (v.description) {
        expect(role3).toHaveTextContent(v.description);
      }
    });
    // Select value
    fireEvent.change(inputElements[2], { target: { value: 'false' } });
    fireEvent.keyDown(inputElements[2], { key: 'ArrowDown' });
    fireEvent.keyDown(inputElements[2], { key: 'Enter' });
    expect(onChange).toHaveBeenLastCalledWith({ done: { eq: false } } as TestFilterModel);
  });

  it('should be ok to just display two simple fields at root level and should be able to change it', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm
        filterDefinitionModel={testFilterDefinitionObject}
        initialFilter={{ done: { eq: true }, text: { eq: 'foo' } } as TestFilterModel}
        onChange={onChange}
      />,
    );

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenCalledWith({
      AND: [{ done: { eq: true } }, { text: { eq: 'foo', caseInsensitive: true } }],
    } as TestFilterModel);

    const inputElements = container.querySelectorAll('input');
    expect(inputElements).toHaveLength(6);

    expect(container).toHaveTextContent('common.filter.field');
    expect(container).toHaveTextContent('common.filter.operation');
    expect(container).toHaveTextContent('common.filter.value');

    expect(inputElements[0]).toHaveAttribute('placeholder', 'common.filter.field');
    expect(inputElements[0]).toHaveAttribute('value', 'todos.fields.done');
    expect(inputElements[1]).toHaveAttribute('placeholder', 'common.filter.operation');
    expect(inputElements[1]).toHaveAttribute('value', 'common.operations.eq');
    expect(inputElements[2]).toHaveAttribute('placeholder', 'common.filter.value');
    expect(inputElements[2]).toHaveAttribute('value', 'common.boolean.true');
    expect(inputElements[3]).toHaveAttribute('placeholder', 'common.filter.field');
    expect(inputElements[3]).toHaveAttribute('value', 'todos.fields.text');
    expect(inputElements[4]).toHaveAttribute('placeholder', 'common.filter.operation');
    expect(inputElements[4]).toHaveAttribute('value', 'common.operations.eq');
    expect(inputElements[5]).toHaveAttribute('placeholder', 'common.filter.value');
    expect(inputElements[5]).toHaveAttribute('value', 'foo');

    expect(container).not.toHaveTextContent('common.fieldValidationError.required');

    // Select value
    fireEvent.change(inputElements[2], { target: { value: 'false' } });
    fireEvent.keyDown(inputElements[2], { key: 'ArrowDown' });
    fireEvent.keyDown(inputElements[2], { key: 'Enter' });
    expect(onChange).toHaveBeenLastCalledWith({
      AND: [{ done: { eq: false } }, { text: { eq: 'foo', caseInsensitive: true } }],
    } as TestFilterModel);

    const buttons = container.querySelectorAll('button');

    // Click on OR
    expect(buttons[1]).not.toBeNull();
    fireEvent.click(buttons[1]);
    expect(onChange).toHaveBeenLastCalledWith({
      OR: [{ done: { eq: false } }, { text: { eq: 'foo', caseInsensitive: true } }],
    } as TestFilterModel);
    expect(container).toMatchSnapshot();
  });

  it('should be ok to just display two simple fields at root level and should be able to add line', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm
        filterDefinitionModel={testFilterDefinitionObject}
        initialFilter={{ done: { eq: true }, text: { eq: 'foo' } } as TestFilterModel}
        onChange={onChange}
      />,
    );

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenCalledWith({
      AND: [{ done: { eq: true } }, { text: { eq: 'foo', caseInsensitive: true } }],
    } as TestFilterModel);

    const inputElements = container.querySelectorAll('input');
    expect(inputElements).toHaveLength(6);

    expect(container).not.toHaveTextContent('common.fieldValidationError.required');

    const buttons = container.querySelectorAll('button');

    // Add new line
    fireEvent.click(buttons[2]);
    expect(onChange).toHaveBeenLastCalledWith(null);
    expect(container).toMatchSnapshot();

    expect(container).toHaveTextContent('common.fieldValidationError.required');

    // Set value
    let inputElements2 = container.querySelectorAll('input');
    expect(inputElements2).toHaveLength(7);
    let lastInput = inputElements2[inputElements2.length - 1];
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'Enter' });
    expect(container).toMatchSnapshot();
    expect(container).toHaveTextContent('common.fieldValidationError.required');

    inputElements2 = container.querySelectorAll('input');
    expect(inputElements2).toHaveLength(8);
    lastInput = inputElements2[inputElements2.length - 1];
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'Enter' });
    screen.debug(lastInput.parentElement?.parentElement as Element);
    expect(container).toMatchSnapshot();

    inputElements2 = container.querySelectorAll('input');
    expect(inputElements2).toHaveLength(9);
    lastInput = inputElements2[inputElements2.length - 1];
    fireEvent.change(lastInput, { target: { value: 'bar' } });
    expect(container).toMatchSnapshot();

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenLastCalledWith({
      AND: [
        { done: { eq: true } },
        { text: { eq: 'foo', caseInsensitive: true } },
        { text: { notEq: 'bar', caseInsensitive: true } },
      ],
    } as TestFilterModel);
  });

  it('should be ok to just display two simple fields at root level and should be remove one line', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm
        filterDefinitionModel={testFilterDefinitionObject}
        initialFilter={
          { AND: [{ done: { eq: false } }, { text: { eq: 'foo', caseInsensitive: true } }] } as TestFilterModel
        }
        onChange={onChange}
      />,
    );

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenCalledWith({
      AND: [{ done: { eq: false } }, { text: { eq: 'foo', caseInsensitive: true } }],
    } as TestFilterModel);

    const inputElements = container.querySelectorAll('input');
    expect(inputElements).toHaveLength(6);

    const buttons = container.querySelectorAll('button');
    // Click on Delete
    expect(buttons[4]).not.toBeNull();
    fireEvent.click(buttons[4]);
    expect(onChange).toHaveBeenLastCalledWith({
      text: { eq: 'foo', caseInsensitive: true },
    } as TestFilterModel);
    expect(container).toMatchSnapshot();
  });

  it('should be ok to just display two fields on first level and should be able to add group', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm
        filterDefinitionModel={testFilterDefinitionObject}
        initialFilter={
          { AND: [{ done: { eq: true } }, { text: { eq: 'foo', caseInsensitive: true } }] } as TestFilterModel
        }
        onChange={onChange}
      />,
    );

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenCalledWith({
      AND: [{ done: { eq: true } }, { text: { eq: 'foo', caseInsensitive: true } }],
    } as TestFilterModel);

    const inputElements = container.querySelectorAll('input');
    expect(inputElements).toHaveLength(6);

    expect(container).not.toHaveTextContent('common.fieldValidationError.required');

    let buttons = container.querySelectorAll('button');

    // Add new group
    fireEvent.click(buttons[3]);
    expect(onChange).toHaveBeenLastCalledWith(null);
    expect(container).toMatchSnapshot();

    expect(container).toHaveTextContent('common.fieldValidationError.required');
    buttons = container.querySelectorAll('button');
    expect(buttons).toHaveLength(23);

    // Set value
    let inputElements2 = container.querySelectorAll('input');
    expect(inputElements2).toHaveLength(7);
    let lastInput = inputElements2[inputElements2.length - 1];
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'Enter' });
    expect(container).toMatchSnapshot();
    expect(container).toHaveTextContent('common.fieldValidationError.required');

    inputElements2 = container.querySelectorAll('input');
    expect(inputElements2).toHaveLength(8);
    lastInput = inputElements2[inputElements2.length - 1];
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'Enter' });
    screen.debug(lastInput.parentElement?.parentElement as Element);
    expect(container).toMatchSnapshot();

    inputElements2 = container.querySelectorAll('input');
    expect(inputElements2).toHaveLength(9);
    lastInput = inputElements2[inputElements2.length - 1];
    fireEvent.change(lastInput, { target: { value: 'bar' } });
    expect(container).toMatchSnapshot();

    // Add second line
    fireEvent.click(buttons[19]);
    expect(container).toMatchSnapshot();
    // Set value
    inputElements2 = container.querySelectorAll('input');
    expect(inputElements2).toHaveLength(10);
    lastInput = inputElements2[inputElements2.length - 1];
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'Enter' });
    expect(container).toMatchSnapshot();
    expect(container).toHaveTextContent('common.fieldValidationError.required');

    inputElements2 = container.querySelectorAll('input');
    expect(inputElements2).toHaveLength(11);
    lastInput = inputElements2[inputElements2.length - 1];
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'ArrowDown' });
    fireEvent.keyDown(lastInput, { key: 'Enter' });
    screen.debug(lastInput.parentElement?.parentElement as Element);
    expect(container).toMatchSnapshot();

    inputElements2 = container.querySelectorAll('input');
    expect(inputElements2).toHaveLength(12);
    lastInput = inputElements2[inputElements2.length - 1];
    fireEvent.change(lastInput, { target: { value: 'bar' } });
    expect(container).toMatchSnapshot();

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenLastCalledWith({
      AND: [
        { done: { eq: true } },
        { text: { eq: 'foo', caseInsensitive: true } },
        { AND: [{ text: { notEq: 'bar', caseInsensitive: true } }, { text: { notEq: 'bar', caseInsensitive: true } }] },
      ],
    } as TestFilterModel);
  });

  it('should be ok to just display two fields on first level and should be able to edit second group', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm
        filterDefinitionModel={testFilterDefinitionObject}
        initialFilter={
          {
            AND: [
              { done: { eq: true } },
              { text: { eq: 'foo' } },
              { AND: [{ text: { notEq: 'bar' } }, { text: { notEq: 'bar' } }] },
            ],
          } as TestFilterModel
        }
        onChange={onChange}
      />,
    );

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenCalledWith({
      AND: [
        { done: { eq: true } },
        { text: { eq: 'foo', caseInsensitive: true } },
        { AND: [{ text: { notEq: 'bar', caseInsensitive: true } }, { text: { notEq: 'bar', caseInsensitive: true } }] },
      ],
    } as TestFilterModel);

    const inputElements = container.querySelectorAll('input');
    expect(inputElements).toHaveLength(12);

    expect(container).not.toHaveTextContent('common.fieldValidationError.required');

    const buttons = container.querySelectorAll('button');

    // Change group for OR
    fireEvent.click(buttons[17]);

    await waitFor(() => 0);
    expect(onChange).toHaveBeenLastCalledWith({
      AND: [
        { done: { eq: true } },
        { text: { eq: 'foo', caseInsensitive: true } },
        { OR: [{ text: { notEq: 'bar', caseInsensitive: true } }, { text: { notEq: 'bar', caseInsensitive: true } }] },
      ],
    } as TestFilterModel);
    expect(container).toMatchSnapshot();
  });

  it('should be ok to just display two fields on first level and should be able to delete second group', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm
        filterDefinitionModel={testFilterDefinitionObject}
        initialFilter={
          {
            AND: [
              { done: { eq: true } },
              { text: { eq: 'foo' } },
              { AND: [{ text: { notEq: 'bar' } }, { text: { notEq: 'bar' } }] },
            ],
          } as TestFilterModel
        }
        onChange={onChange}
      />,
    );

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(onChange).toHaveBeenCalledWith({
      AND: [
        { done: { eq: true } },
        { text: { eq: 'foo', caseInsensitive: true } },
        { AND: [{ text: { notEq: 'bar', caseInsensitive: true } }, { text: { notEq: 'bar', caseInsensitive: true } }] },
      ],
    } as TestFilterModel);

    const inputElements = container.querySelectorAll('input');
    expect(inputElements).toHaveLength(12);

    expect(container).not.toHaveTextContent('common.fieldValidationError.required');

    const buttons = container.querySelectorAll('button');

    // Delete group for OR
    fireEvent.click(buttons[20]);

    await waitFor(() => 0);
    expect(onChange).toHaveBeenLastCalledWith({
      AND: [{ done: { eq: true } }, { text: { eq: 'foo', caseInsensitive: true } }],
    } as TestFilterModel);
    expect(container).toMatchSnapshot();
  });

  it('should be ok to just display and select predefined filter no initial filters', async () => {
    const onChange = jest.fn().mockImplementation((i) => i);

    const { container } = render(
      <FilterForm
        filterDefinitionModel={testFilterDefinitionObject}
        initialFilter={{}}
        onChange={onChange}
        predefinedFilterObjects={[{ display: 'fake', filter: { done: { eq: true } } }]}
      />,
    );

    await waitFor(() => 0);
    expect(container).toMatchSnapshot();
    expect(container).toHaveTextContent('common.fieldValidationError.required');

    let inputElements = container.querySelectorAll('input');
    expect(inputElements).toHaveLength(2);

    const firstInput = inputElements[0];
    fireEvent.keyDown(firstInput, { key: 'ArrowDown' });
    fireEvent.keyDown(firstInput, { key: 'ArrowDown' });
    fireEvent.keyDown(firstInput, { key: 'ArrowDown' });

    // Find content
    const presentationElement = await screen.findByRole('presentation');
    expect(presentationElement).toHaveTextContent('fake');
    expect(presentationElement).toMatchSnapshot();

    // Enter
    fireEvent.keyDown(firstInput, { key: 'Enter' });
    const buttons = container.querySelectorAll('button');
    // Load
    fireEvent.click(buttons[1]);
    expect(container).toMatchSnapshot();

    expect(container).toHaveTextContent('common.filter.field');
    expect(container).toHaveTextContent('common.filter.operation');
    expect(container).toHaveTextContent('common.filter.value');

    inputElements = container.querySelectorAll('input');
    expect(inputElements).toHaveLength(4);
    expect(inputElements[1]).toHaveAttribute('placeholder', 'common.filter.field');
    expect(inputElements[1]).toHaveAttribute('value', 'todos.fields.done');
    expect(inputElements[2]).toHaveAttribute('placeholder', 'common.filter.operation');
    expect(inputElements[2]).toHaveAttribute('value', 'common.operations.eq');
    expect(inputElements[3]).toHaveAttribute('placeholder', 'common.filter.value');
    expect(inputElements[3]).toHaveAttribute('value', 'common.boolean.true');

    expect(container).not.toHaveTextContent('common.fieldValidationError.required');
    expect(onChange).toHaveBeenLastCalledWith({ done: { eq: true } } as TestFilterModel);
  });
});
