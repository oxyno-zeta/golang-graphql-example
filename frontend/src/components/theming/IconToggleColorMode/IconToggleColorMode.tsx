import React, { useContext } from 'react';
import { useTranslation } from 'react-i18next';
import IconButton, { IconButtonProps } from '@mui/material/IconButton';
import Tooltip from '@mui/material/Tooltip';
import { useTheme } from '@mui/material/styles';
import DarkModeIcon from '@mui/icons-material/DarkMode';
import LightModeIcon from '@mui/icons-material/LightMode';
import ColorModeContext from '../../../contexts/ColorModeContext';

interface Props {
  iconButtonProps?: IconButtonProps;
}

const defaultProps = {
  iconButtonProps: {},
};

function IconToggleColorMode({ iconButtonProps }: Props) {
  // Setup translate
  const { t } = useTranslation();
  // Get theme
  const theme = useTheme();
  // Get color mode context
  const colorMode = useContext(ColorModeContext);

  return (
    <Tooltip title={<>{t(theme.palette.mode === 'dark' ? 'common.lightThemeTooltip' : 'common.darkThemeTooltip')}</>}>
      <span>
        <IconButton onClick={colorMode.toggleColorMode} color="inherit" {...iconButtonProps}>
          {theme.palette.mode === 'dark' ? <DarkModeIcon /> : <LightModeIcon />}
        </IconButton>
      </span>
    </Tooltip>
  );
}

IconToggleColorMode.defaultProps = defaultProps;

export default IconToggleColorMode;
