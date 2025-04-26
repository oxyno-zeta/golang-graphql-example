import React, { ReactNode, useState } from 'react';
import { useTranslation } from 'react-i18next';
import Tooltip from '@mui/material/Tooltip';
import IconButton from '@mui/material/IconButton';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiApps } from '@mdi/js';
import AppLinksPopper from './components/AppLinksPopper';
import type { SxProps } from '@mui/material';

export interface Props {
  readonly children: ReactNode;
  readonly appLinksPopperSx?: SxProps;
}

function AppLinksButton({ children, appLinksPopperSx = undefined }: Props) {
  // Setup translate
  const { t } = useTranslation();
  // States
  const [open, setOpen] = useState(false);
  const [anchorEl, setAnchorEl] = useState<HTMLButtonElement | null>(null);

  return (
    <>
      <Tooltip title={<>{t('common.appLinks')}</>}>
        <IconButton
          onClick={() => {
            setOpen(true);
          }}
          ref={(d: HTMLButtonElement) => {
            if (d && d !== anchorEl) {
              setAnchorEl(d);
            }
          }}
          sx={{ marginRight: '5px' }}
        >
          <SvgIcon>
            <path d={mdiApps} />
          </SvgIcon>
        </IconButton>
      </Tooltip>
      <AppLinksPopper
        anchorElement={anchorEl}
        onClose={() => {
          setOpen(false);
        }}
        open={open}
        popperSx={appLinksPopperSx}
      >
        {children}
      </AppLinksPopper>
    </>
  );
}

export default AppLinksButton;
