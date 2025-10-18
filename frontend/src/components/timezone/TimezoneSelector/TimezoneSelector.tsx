import React, { type ReactNode, useContext } from 'react';
import TextField from '@mui/material/TextField';
import Autocomplete, { type AutocompleteProps, type AutocompleteRenderOptionState } from '@mui/material/Autocomplete';
import Typography from '@mui/material/Typography';
import parse from 'autosuggest-highlight/parse';
import match from 'autosuggest-highlight/match';
import { useTranslation } from 'react-i18next';
import { List, type RowComponentProps, type ListImperativeAPI } from 'react-window';
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
function RowComponent({
  index,
  itemData,
  style,
}: RowComponentProps & {
  itemData: GroupOrItemModel<string, string>[];
}) {
  const dataSet = itemData[index];
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

// Adapter for react-window v2
const ListboxComponent = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLElement> & {
    internalListRef: React.Ref<ListImperativeAPI>;
    onItemsBuilt: (optionIndexMap: Map<string, number>) => void;
  }
>((props, ref) => {
  const { children, internalListRef, onItemsBuilt, ...other } = props;
  const itemData: GroupOrItemModel<string, string>[] = [];
  const optionIndexMap = React.useMemo(() => new Map<string, number>(), []);

  (children as GroupModel<string, string>[]).forEach((item) => {
    itemData.push(item);
    itemData.push(...(item.children || []));
  });

  // Map option values to their indices in the flattened array
  itemData.forEach((item, index) => {
    if (Array.isArray(item) && item[1]) {
      // eslint-disable-next-line react-hooks/immutability
      optionIndexMap.set(item[1], index);
    }
  });

  React.useEffect(() => {
    if (onItemsBuilt) {
      onItemsBuilt(optionIndexMap);
    }
  }, [onItemsBuilt, optionIndexMap]);

  const itemCount = itemData.length;

  const getChildSize = (child: GroupOrItemModel<string, string>) => {
    if (child.type === 'group') {
      return 48;
    }

    return 52;
  };

  // Separate className for List, other props for wrapper div (ARIA, handlers)
  // eslint-disable-next-line @typescript-eslint/naming-convention
  const { className, style: _style, ...otherProps } = other;

  return (
    <div ref={ref} {...otherProps}>
      <List
        className={className}
        listRef={internalListRef}
        key={itemCount}
        rowCount={itemCount}
        rowHeight={(index) => getChildSize(itemData[index])}
        rowComponent={RowComponent}
        rowProps={{ itemData }}
        style={{
          height: 4 * 48,
          width: '100%',
        }}
        overscanCount={5}
        tagName="ul"
      />
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
