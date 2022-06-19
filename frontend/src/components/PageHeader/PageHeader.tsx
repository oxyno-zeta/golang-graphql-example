import React, { ReactNode } from 'react';
import Typography from '@mui/material/Typography';
import Toolbar from '@mui/material/Toolbar';

interface Props {
  leftElement?: ReactNode;
  titleElement?: ReactNode;
  rightElement?: ReactNode;
  title?: string;
}

const defaultProps = {
  leftElement: null,
  title: '',
  titleElement: null,
  rightElement: null,
};

function PageHeader({ leftElement, title, titleElement, rightElement }: Props) {
  return (
    <Toolbar style={{ paddingLeft: '0px' }}>
      {leftElement}
      {title && (
        <Typography variant="h5" color="inherit" style={{ fontWeight: 'bold' }}>
          {title}
        </Typography>
      )}
      {titleElement}
      <div style={{ flexGrow: 1 }} />
      {rightElement}
    </Toolbar>
  );
}

PageHeader.defaultProps = defaultProps;

export default PageHeader;
