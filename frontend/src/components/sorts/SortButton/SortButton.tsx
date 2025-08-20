import React, { useState } from 'react';
import Button from '@mui/material/Button';
import Tooltip from '@mui/material/Tooltip';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiSortVariant } from '@mdi/js';
import { useTranslation } from 'react-i18next';
import { type SortOrderModel, type SortOrderFieldModel } from '../../../models/general';
import SortPopper from '../SortPopper';
import SortDialog from '../SortDialog';

export interface Props<T extends Record<string, SortOrderModel>> {
  readonly sorts: null | undefined | T[];
  readonly onSubmit: (s: T[]) => void;
  readonly sortFields: SortOrderFieldModel[];
  readonly isPopperEnabled?: boolean;
}

function SortButton<T extends Record<string, SortOrderModel>>({
  sorts,
  onSubmit,
  sortFields,
  isPopperEnabled = false,
}: Props<T>) {
  // Setup translate
  const { t } = useTranslation();
  // States
  const [open, setOpen] = useState(false);
  const [anchorEl, setAnchorEl] = useState<HTMLButtonElement | null>(null);

  const onClick = () => {
    setOpen(true);
  };

  const onClose = () => {
    setOpen(false);
  };

  const sortDefaultProps = {
    initialSorts: sorts,
    open,
    sortFields,
    onSubmit: (args: T[]) => {
      onSubmit(args);
      onClose();
    },
    onReset: () => {
      onSubmit([]);
      onClose();
    },
    onClose,
  };

  return (
    <>
      <Tooltip title={<>{t('common.sort.buttonTooltip')}</>}>
        <Button
          color="inherit"
          onClick={onClick}
          ref={(d: HTMLButtonElement) => {
            if (d && d !== anchorEl) {
              setAnchorEl(d);
            }
          }}
          sx={{ border: (theme) => `1px solid ${theme.palette.divider}`, padding: '5px 10px', minWidth: '46px' }}
          variant="outlined"
        >
          <SvgIcon color="inherit">
            <path d={mdiSortVariant} />
          </SvgIcon>
        </Button>
      </Tooltip>
      {isPopperEnabled ? (
        <SortPopper anchorElement={anchorEl} {...sortDefaultProps} />
      ) : (
        <SortDialog {...sortDefaultProps} />
      )}
    </>
  );
}

export default SortButton;
