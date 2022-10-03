import React, { ReactNode } from 'react';
import { Link as RouterLink, useMatch, useResolvedPath } from 'react-router-dom';
import ListItemButton from '@mui/material/ListItemButton';

interface Props {
  to: string;
  children: ReactNode;
}

function ListNavItemButton({ to, children }: Props) {
  const match = useMatch({ path: useResolvedPath(to).pathname, end: true });

  return (
    <ListItemButton component={RouterLink} to={to} selected={!!match}>
      {children}
    </ListItemButton>
  );
}

export default ListNavItemButton;
