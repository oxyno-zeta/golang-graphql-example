import React from 'react';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import Tooltip from '@mui/material/Tooltip';
import { useTranslation } from 'react-i18next';

interface Props {
  handleDrawerToggle: () => void;
}

function OpenDrawerButton({ handleDrawerToggle }: Props) {
  // Setup translate
  const { t } = useTranslation();

  return (
    <Tooltip title={<>{t('common.openAction')}</>}>
      <span>
        <IconButton color="inherit" onClick={handleDrawerToggle} sx={{ display: { lg: 'none' } }}>
          <MenuIcon />
        </IconButton>
      </span>
    </Tooltip>
  );
}

export default OpenDrawerButton;
