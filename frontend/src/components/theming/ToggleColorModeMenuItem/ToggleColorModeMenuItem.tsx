import React, { useContext } from 'react';
import ToggleButton from '@mui/material/ToggleButton';
import ToggleButtonGroup from '@mui/material/ToggleButtonGroup';
import { useTranslation } from 'react-i18next';
import Typography from '@mui/material/Typography';
import { useTheme } from '@mui/material/styles';
import DarkModeIcon from '@mui/icons-material/DarkMode';
import LightModeIcon from '@mui/icons-material/LightMode';
import { PaletteMode } from '@mui/material';
import ColorModeContext from '../../../contexts/ColorModeContext';

interface Props {
  titleStyle?: React.CSSProperties;
}

const defaultProps = {
  titleStyle: { fontSize: 13, marginBottom: '2px' },
};

function ToggleColorModeMenuItem({ titleStyle }: Props) {
  // Setup translate
  const { t } = useTranslation();
  // Get theme
  const theme = useTheme();
  // Get color mode context
  const colorMode = useContext(ColorModeContext);

  return (
    <>
      <Typography style={titleStyle}>{t('common.themeTitle')}</Typography>
      <ToggleButtonGroup
        fullWidth
        size="small"
        value={theme.palette.mode}
        exclusive
        onChange={(event, value) => {
          colorMode.setColorMode(value as PaletteMode);
        }}
      >
        <ToggleButton value="dark">
          <DarkModeIcon /> {t('common.darkThemeSelector')}
        </ToggleButton>
        <ToggleButton value="light">
          <LightModeIcon /> {t('common.lightThemeSelector')}
        </ToggleButton>
      </ToggleButtonGroup>
    </>
  );
}

ToggleColorModeMenuItem.defaultProps = defaultProps;

export default ToggleColorModeMenuItem;