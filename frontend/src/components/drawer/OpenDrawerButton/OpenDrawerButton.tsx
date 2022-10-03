import React from 'react';
import IconButton, { IconButtonProps } from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';
import Tooltip from '@mui/material/Tooltip';
import { useTranslation } from 'react-i18next';

interface Props {
  handleDrawerToggle: () => void;
  iconButtonProps?: Partial<Omit<IconButtonProps, 'onClick'>>;
}

const defaultProps = {
  iconButtonProps: {},
};

function OpenDrawerButton({ handleDrawerToggle, iconButtonProps }: Props) {
  // Setup translate
  const { t } = useTranslation();

  return (
    <Tooltip title={<>{t('common.openAction')}</>}>
      <span>
        <IconButton color="inherit" onClick={handleDrawerToggle} sx={{ display: { lg: 'none' } }} {...iconButtonProps}>
          <MenuIcon />
        </IconButton>
      </span>
    </Tooltip>
  );
}

OpenDrawerButton.defaultProps = defaultProps;

export default OpenDrawerButton;
