import React from 'react';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';

interface Props {
  handleDrawerToggle: () => void;
}

function OpenDrawerButton({ handleDrawerToggle }: Props) {
  return (
    <IconButton color="inherit" onClick={handleDrawerToggle} sx={{ display: { sm: 'none' } }}>
      <MenuIcon />
    </IconButton>
  );
}

export default OpenDrawerButton;
