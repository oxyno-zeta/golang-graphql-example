import React from 'react';
import Typography, { type TypographyProps } from '@mui/material/Typography';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiHeart } from '@mdi/js';
import Box from '@mui/material/Box';
import type { SxProps } from '@mui/material';

export interface Props {
  readonly containerBoxSx?: SxProps;
  readonly typographyProps?: TypographyProps;
}

function Footer({ containerBoxSx = {}, typographyProps = {} }: Props) {
  return (
    <Box
      sx={{
        alignItems: 'center',
        display: 'flex',
        textAlign: 'center',
        flexDirection: 'column',
        margin: '10px 0',
        ...containerBoxSx,
      }}
    >
      <Typography sx={{ display: 'flex' }} {...typographyProps}>
        Todo list application / With{' '}
        <SvgIcon color="error">
          <path d={mdiHeart} />
        </SvgIcon>{' '}
        by Oxyno-zeta
      </Typography>
    </Box>
  );
}

export default Footer;
