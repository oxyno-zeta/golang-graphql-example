import React, { type ReactNode } from 'react';
import { Link as RouterLink, useMatch, useResolvedPath } from 'react-router';
import ListItemButton, { type ListItemButtonProps } from '@mui/material/ListItemButton';

export interface Props {
  readonly to: string;
  readonly children: ReactNode;
  readonly listItemButtonProps?: Omit<ListItemButtonProps, 'component' | 'to' | 'selected'>;
}

function ListNavItemButton({ to, children, listItemButtonProps = {} }: Props) {
  const match = useMatch({ path: useResolvedPath(to).pathname, end: true });

  /* eslint-disable @typescript-eslint/no-explicit-any */
  // Coming from official mui documentation
  // https://mui.com/material-ui/react-breadcrumbs/#integration-with-react-router
  return (
    <ListItemButton component={RouterLink as any} selected={!!match} to={to} {...listItemButtonProps}>
      {children}
    </ListItemButton>
  );
}

export default ListNavItemButton;
