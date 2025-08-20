import React, { useContext } from 'react';
import { useTranslation } from 'react-i18next';
import IconButton, { type IconButtonProps } from '@mui/material/IconButton';
import Tooltip from '@mui/material/Tooltip';
import { useTheme } from '@mui/material/styles';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiBrightness2, mdiBrightness7 } from '@mdi/js';
import ColorModeContext from '../../../contexts/ColorModeContext';

export interface Props {
  readonly iconButtonProps?: IconButtonProps;
}

function IconToggleColorMode({ iconButtonProps = {} }: Props) {
  // Setup translate
  const { t } = useTranslation();
  // Get theme
  const theme = useTheme();
  // Get color mode context
  const colorMode = useContext(ColorModeContext);

  // Expand
  const { toggleColorMode } = colorMode;

  return (
    <Tooltip title={<>{t(theme.palette.mode === 'dark' ? 'common.lightThemeTooltip' : 'common.darkThemeTooltip')}</>}>
      <span>
        <IconButton color="inherit" onClick={toggleColorMode} {...iconButtonProps}>
          {theme.palette.mode === 'dark' ? (
            <SvgIcon>
              <path d={mdiBrightness2} />
            </SvgIcon>
          ) : (
            <SvgIcon>
              <path d={mdiBrightness7} />
            </SvgIcon>
          )}
        </IconButton>
      </span>
    </Tooltip>
  );
}

export default IconToggleColorMode;
