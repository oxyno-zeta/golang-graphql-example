import React from 'react';
import { Theme } from '@mui/material/styles';
import Backdrop, { BackdropProps } from '@mui/material/Backdrop';
import CircularProgress, { CircularProgressProps } from '@mui/material/CircularProgress';

export interface Props {
  readonly backdropProps?: Omit<BackdropProps, 'open'>;
  readonly circularProgressProps?: CircularProgressProps;
}

function MainPageCenterLoading({ backdropProps = {}, circularProgressProps = {} }: Props) {
  return (
    <Backdrop sx={{ color: '#fff', zIndex: (theme: Theme) => theme.zIndex.drawer + 1 }} {...backdropProps} open>
      <CircularProgress color="inherit" {...circularProgressProps} />
    </Backdrop>
  );
}

export default MainPageCenterLoading;
