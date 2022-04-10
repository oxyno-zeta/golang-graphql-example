import React from 'react';
import Typography from '@mui/material/Typography';
import FavoriteIcon from '@mui/icons-material/Favorite';
import Box from '@mui/material/Box';

function Footer() {
  return (
    <Box
      sx={{
        alignItems: 'center',
        display: 'flex',
        textAlign: 'center',
        flexDirection: 'column',
        margin: '10px 0',
      }}
    >
      <Typography sx={{ display: 'flex' }}>
        Todo list application / With <FavoriteIcon color="error" /> by Oxyno-zeta
      </Typography>
    </Box>
  );
}

export default Footer;
