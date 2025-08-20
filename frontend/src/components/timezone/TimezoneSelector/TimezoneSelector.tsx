import React, { type ReactNode, useContext } from 'react';
import TextField from '@mui/material/TextField';
import Autocomplete, { type AutocompleteProps, type AutocompleteRenderOptionState } from '@mui/material/Autocomplete';
import Typography from '@mui/material/Typography';
import parse from 'autosuggest-highlight/parse';
import match from 'autosuggest-highlight/match';
import { useTranslation } from 'react-i18next';
import { VariableSizeList, type ListChildComponentProps } from 'react-window';
import ListSubheader from '@mui/material/ListSubheader';
import { availableTimezones, getTimeZone } from '../utils';
import TimezoneContext from '../../../contexts/TimezoneContext';

// Copied and adapted from example: https://mui.com/material-ui/react-autocomplete/#virtualization

type GroupOrItemModel<G, I> = ItemModel<I> | GroupModel<G, I>;

interface ItemModel<T> {
  type: 'item';
  value: T;
  state: AutocompleteRenderOptionState;
  props: React.HTMLAttributes<HTMLLIElement>;
}

interface GroupModel<G, I> {
  type: 'group';
  value: G;
  children: ItemModel<I>[];
}

/* eslint-disable @typescript-eslint/no-explicit-any */
/* eslint-disable react/no-array-index-key */
function renderRow(props: ListChildComponentProps) {
  const { data, index, style } = props;
  const dataSet = data[index] as GroupOrItemModel<string, string>;
  const inlineStyle = {
    ...style,
    top: style.top as number,
  };

  if (dataSet.type === 'group') {
    const d = dataSet;

    return (
      <ListSubheader component="div" key={d.value} style={inlineStyle}>
        {d.value}
      </ListSubheader>
    );
  }

  const d = dataSet;

  const displayedOption = d.value;
  const matches = match(displayedOption, d.state.inputValue, { insideWords: true, findAllOccurrences: true });
  const parts = parse(displayedOption, matches);

  return (
    <li {...d.props} style={{ ...inlineStyle, display: 'block' }} key={d.value}>
      <Typography noWrap>
        {parts.map((part: any, i: number) => (
          <span
            key={i}
            style={{
              fontWeight: part.highlight ? 700 : 400,
            }}
          >
            {part.text}
          </span>
        ))}
      </Typography>
      <Typography color="text.secondary" style={{ fontSize: 12 }}>
        {getTimeZone(d.value)}
      </Typography>
    </li>
  );
}
/* eslint-enable */

const OuterElementContext = React.createContext({});

const OuterElementType = React.forwardRef<HTMLDivElement>((props, ref) => {
  const outerProps = React.useContext(OuterElementContext);
  return <div ref={ref} {...props} {...outerProps} />;
});

function useResetCache(length: number) {
  const ref = React.useRef<VariableSizeList>(null);
  React.useEffect(() => {
    if (ref.current != null) {
      ref.current.resetAfterIndex(0, true);
    }
  }, [length]);
  return ref;
}

// Adapter for react-window
const ListboxComponent = React.forwardRef<HTMLDivElement, React.HTMLAttributes<HTMLElement>>((props, ref) => {
  const { children, ...other } = props;
  const itemData: GroupOrItemModel<string, string>[] = [];

  (children as GroupModel<string, string>[]).forEach((item: GroupModel<string, string>) => {
    itemData.push(item);
    itemData.push(...(item.children || []));
  });

  const itemCount = itemData.length;

  const getChildSize = (child: GroupOrItemModel<string, string>) => {
    if (child.type === 'group') {
      return 48;
    }

    return 52;
  };

  const gridRef = useResetCache(itemCount);

  return (
    <div ref={ref}>
      <OuterElementContext.Provider value={other}>
        <VariableSizeList
          height={4 * 48}
          innerElementType="ul"
          itemCount={itemCount}
          itemData={itemData}
          itemSize={(index) => getChildSize(itemData[index])}
          outerElementType={OuterElementType}
          overscanCount={5}
          ref={gridRef}
          width="100%"
        >
          {renderRow}
        </VariableSizeList>
      </OuterElementContext.Provider>
    </div>
  );
});

export interface Props {
  readonly autocompleteProps?: Partial<AutocompleteProps<string, false, true, false>>;
}

function TimezoneSelector({
  autocompleteProps = {
    fullWidth: true,
    size: 'small',
    sx: { minWidth: { xs: '250px', md: '300px' } },
  },
}: Props) {
  // Get timezone context
  const timezoneCtx = useContext(TimezoneContext);
  // Setup translate
  const { t } = useTranslation();

  // Expand
  const { getTimezone, setTimezone } = timezoneCtx;

  return (
    <Autocomplete
      ListboxComponent={ListboxComponent}
      clearText={t('common.clearAction')}
      closeText={t('common.closeAction')}
      disableListWrap
      groupBy={(option) => option.split('/')[0]}
      noOptionsText={t('common.filter.noOptions')}
      onChange={(event, input) => {
        // Check if input exists
        if (input) {
          setTimezone(input);
        }
      }}
      openText={t('common.openAction')}
      options={availableTimezones}
      renderGroup={(params) => {
        const res: GroupModel<string, string> = {
          type: 'group',
          value: params.group,
          children: params.children as ItemModel<string>[],
        };

        // TODO: Post React 18 update - validate this conversion, look like a hidden bug
        return res as unknown as ReactNode;
      }}
      renderInput={(params) => <TextField {...params} label={t('common.timezone')} />}
      renderOption={(props, option, state) => {
        const res: ItemModel<string> = {
          type: 'item',
          value: option,
          state,
          props,
        };

        // TODO: Post React 18 update - validate this conversion, look like a hidden bug
        return res as unknown as ReactNode;
      }}
      value={getTimezone()}
      slotProps={{
        listbox: {
          component: ListboxComponent,
        },
      }}
      {...autocompleteProps}
    />
  );
}

export default TimezoneSelector;
