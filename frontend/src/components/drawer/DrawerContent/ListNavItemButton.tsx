import React, { ReactNode } from 'react';
import { Link as RouterLink, useMatch, useResolvedPath } from 'react-router-dom';
import ListItemButton, { ListItemButtonProps } from '@mui/material/ListItemButton';

export interface Props {
  to: string;
  children: ReactNode;
  listItemButtonProps?: Omit<ListItemButtonProps, 'component' | 'to' | 'selected'>;
}

const defaultProps = {
  listItemButtonProps: {},
};

function ListNavItemButton({ to, children, listItemButtonProps }: Props) {
  const match = useMatch({ path: useResolvedPath(to).pathname, end: true });

  /* eslint-disable @typescript-eslint/no-explicit-any */
  // Coming from official mui documentation
  // https://mui.com/material-ui/react-breadcrumbs/#integration-with-react-router
  return (
    <ListItemButton component={RouterLink as any} to={to} selected={!!match} {...listItemButtonProps}>
      {children}
    </ListItemButton>
  );
}

ListNavItemButton.defaultProps = defaultProps;

export default ListNavItemButton;
