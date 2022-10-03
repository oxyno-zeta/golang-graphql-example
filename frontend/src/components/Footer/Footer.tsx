import React from 'react';
import Typography, { TypographyProps } from '@mui/material/Typography';
import FavoriteIcon from '@mui/icons-material/Favorite';
import Box from '@mui/material/Box';
import type { SxProps } from '@mui/material';

interface Props {
  containerBoxSx?: SxProps;
  typographyProps?: TypographyProps;
}

const defaultProps = {
  containerBoxSx: {},
  typographyProps: {},
};

function Footer({ containerBoxSx, typographyProps }: Props) {
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
        Todo list application / With <FavoriteIcon color="error" /> by Oxyno-zeta
      </Typography>
    </Box>
  );
}

Footer.defaultProps = defaultProps;

export default Footer;
