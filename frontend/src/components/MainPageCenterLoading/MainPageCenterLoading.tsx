import React from 'react';
import { Theme } from '@mui/material/styles';
import Backdrop, { BackdropProps } from '@mui/material/Backdrop';
import CircularProgress, { CircularProgressProps } from '@mui/material/CircularProgress';

export interface Props {
  backdropProps?: BackdropProps;
  circularProgressProps?: CircularProgressProps;
}

const defaultProps = {
  backdropProps: {},
  circularProgressProps: {},
};

function MainPageCenterLoading({ backdropProps, circularProgressProps }: Props) {
  return (
    <Backdrop sx={{ color: '#fff', zIndex: (theme: Theme) => theme.zIndex.drawer + 1 }} open {...backdropProps}>
      <CircularProgress color="inherit" {...circularProgressProps} />
    </Backdrop>
  );
}

MainPageCenterLoading.defaultProps = defaultProps;

export default MainPageCenterLoading;
