import React, { UIEventHandler, useEffect, useState } from 'react';
import MuiAutocomplete, { AutocompleteProps } from '@mui/material/Autocomplete';

// Imported and edited from issue: https://github.com/mui/material-ui/issues/30249

interface Props<
  T,
  Multiple extends boolean | undefined,
  DisableClearable extends boolean | undefined,
  FreeSolo extends boolean | undefined,
> extends AutocompleteProps<T, Multiple, DisableClearable, FreeSolo> {
  loadMore: () => void;
}

function AutocompleteInfiniteScroll<
  T,
  Multiple extends boolean | undefined,
  DisableClearable extends boolean | undefined,
  FreeSolo extends boolean | undefined,
>({ loadMore, ...rest }: Props<T, Multiple, DisableClearable, FreeSolo>) {
  // take the scroll position at the moment
  const [position, setPosition] = useState(0);

  // take the ListboxNode's DOM Node
  // this value is unlikely to change,
  // but it is necessary to specify it in useEffect
  const [listboxNode, setListboxNode] = useState<HTMLUListElement | null>(null);

  // change scroll position,
  // if the position's state has been changed
  useEffect(() => {
    // this condition checks if there is a necessary DOM Node
    if (listboxNode !== null) {
      listboxNode.scrollTop = position;
    }
    // it will only work when the position or node changes
  }, [position, listboxNode, rest.options]);

  const onScroll: UIEventHandler<HTMLUListElement> = (event) => {
    // const ListboxNode = event.currentTarget;
    // replaced by this
    setListboxNode(event.currentTarget);

    // Check if list box node is set
    if (listboxNode) {
      // this is necessary constant,
      // because if we change the position state
      // in this place then, with any scrolling,
      // we will be immediately returned to a position
      // equal to this value,
      // since our useEffect is triggered immediately
      const x = listboxNode.scrollTop + listboxNode.clientHeight;

      // only when checking this condition we change the position
      if (listboxNode.scrollHeight - x <= 1) {
        setPosition(x);
        loadMore();
      }
    }
  };

  return (
    <MuiAutocomplete<T, Multiple, DisableClearable, FreeSolo>
      {...rest}
      ListboxProps={{
        onScroll,
      }}
    />
  );
}

export default AutocompleteInfiniteScroll;
