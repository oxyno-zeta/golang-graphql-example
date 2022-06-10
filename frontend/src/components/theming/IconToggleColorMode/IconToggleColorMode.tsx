import React from 'react';
import { useTranslation } from 'react-i18next';
import IconButton from '@mui/material/IconButton';
import Tooltip from '@mui/material/Tooltip';
import { useTheme } from '@mui/material/styles';
import Brightness4Icon from '@mui/icons-material/Brightness4';
import Brightness7Icon from '@mui/icons-material/Brightness7';
import ColorModeContext from '../../../contexts/ColorModeContext';

function IconToggleColorMode() {
  // Setup translate
  const { t } = useTranslation();
  // Get theme
  const theme = useTheme();
  // Get color mode context
  const colorMode = React.useContext(ColorModeContext);

  return (
    <Tooltip title={<>{t(theme.palette.mode === 'dark' ? 'common.lightThemeTooltip' : 'common.darkThemeTooltip')}</>}>
      <span>
        <IconButton onClick={colorMode.toggleColorMode} color="inherit">
          {theme.palette.mode === 'dark' ? <Brightness7Icon /> : <Brightness4Icon />}
        </IconButton>
      </span>
    </Tooltip>
  );
}

export default IconToggleColorMode;
