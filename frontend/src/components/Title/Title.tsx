import React, { ReactNode } from 'react';
import Typography, { TypographyProps } from '@mui/material/Typography';
import Toolbar, { ToolbarProps } from '@mui/material/Toolbar';

export interface Props {
  readonly leftElement?: ReactNode;
  readonly titleElement?: ReactNode;
  readonly rightElement?: ReactNode;
  readonly title?: string;
  readonly titleTypographyProps?: Partial<TypographyProps>;
  readonly toolbarProps?: Partial<ToolbarProps>;
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
      {title ? <Typography {...titleTypographyProps}>{title}</Typography> : null}
      {titleElement}
      <div style={{ flexGrow: 1 }} />
      {rightElement}
    </Toolbar>
  );
}

export default Title;
