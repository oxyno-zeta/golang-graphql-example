import React from 'react';
import { Theme } from '@mui/material/styles';
import Backdrop from '@mui/material/Backdrop';
import CircularProgress from '@mui/material/CircularProgress';

function MainPageCenterLoading() {
  return (
    <Backdrop sx={{ color: '#fff', zIndex: (theme: Theme) => theme.zIndex.drawer + 1 }} open>
      <CircularProgress color="inherit" />
    </Backdrop>
  );
}

export default MainPageCenterLoading;
