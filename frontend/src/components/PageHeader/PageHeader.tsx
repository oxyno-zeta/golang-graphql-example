import React, { ReactNode } from 'react';
import Typography from '@mui/material/Typography';
import Toolbar from '@mui/material/Toolbar';
import Box from '@mui/material/Box';

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
    <Toolbar>
      <Box>{leftElement}</Box>
      <Box>
        {title && (
          <Typography variant="h5" color="inherit" style={{ fontWeight: 'bold' }}>
            {title}
          </Typography>
        )}
        {titleElement}
      </Box>
      <Box sx={{ marginLeft: 'auto' }}>{rightElement}</Box>
    </Toolbar>
  );
}

PageHeader.defaultProps = defaultProps;

export default PageHeader;
