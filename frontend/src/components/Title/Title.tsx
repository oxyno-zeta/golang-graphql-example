import React, { ReactNode } from 'react';
import Typography, { TypographyProps } from '@mui/material/Typography';
import Toolbar, { ToolbarProps } from '@mui/material/Toolbar';

export interface Props {
  leftElement?: ReactNode;
  titleElement?: ReactNode;
  rightElement?: ReactNode;
  title?: string;
  titleTypographyProps?: Partial<TypographyProps>;
  toolbarProps?: Partial<ToolbarProps>;
}

function Title({
  leftElement = null,
  title = '',
  titleElement = null,
  rightElement = null,
  titleTypographyProps = {
    variant: 'h5',
    color: 'inherit',
    style: { fontWeight: 'bold' },
  },
  toolbarProps = {
    style: { paddingLeft: '0px', paddingRight: '0px' },
  },
}: Props) {
  return (
    <Toolbar {...toolbarProps}>
      {leftElement}
      {title && <Typography {...titleTypographyProps}>{title}</Typography>}
      {titleElement}
      <div style={{ flexGrow: 1 }} />
      {rightElement}
    </Toolbar>
  );
}

export default Title;
