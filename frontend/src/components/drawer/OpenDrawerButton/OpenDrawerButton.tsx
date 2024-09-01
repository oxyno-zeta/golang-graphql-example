import React from 'react';
import IconButton, { IconButtonProps } from '@mui/material/IconButton';
import { mdiMenu } from '@mdi/js';
import SvgIcon from '@mui/material/SvgIcon';
import Tooltip from '@mui/material/Tooltip';
import { useTranslation } from 'react-i18next';

export interface Props {
  readonly onDrawerToggle: () => void;
  readonly iconButtonProps?: Partial<Omit<IconButtonProps, 'onClick'>>;
}

function OpenDrawerButton({ onDrawerToggle, iconButtonProps = {} }: Props) {
  // Setup translate
  const { t } = useTranslation();

  return (
    <Tooltip title={<>{t('common.openAction')}</>}>
      <span>
        <IconButton color="inherit" onClick={onDrawerToggle} sx={{ display: { lg: 'none' } }} {...iconButtonProps}>
          <SvgIcon>
            <path d={mdiMenu} />
          </SvgIcon>
        </IconButton>
      </span>
    </Tooltip>
  );
}

export default OpenDrawerButton;
