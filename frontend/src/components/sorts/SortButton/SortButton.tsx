import React, { useState } from 'react';
import Button from '@mui/material/Button';
import Tooltip from '@mui/material/Tooltip';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiSortVariant } from '@mdi/js';
import { useTranslation } from 'react-i18next';
import { SortOrderModel, SortOrderFieldModel } from '../../../models/general';
import SortPopper from '../SortPopper';
import SortDialog from '../SortDialog';

type Props<T extends Record<string, SortOrderModel>> = {
  sort: null | undefined | T;
  setSort: React.Dispatch<T>;
  sortFields: SortOrderFieldModel[];
  isPopperEnabled?: boolean;
};

const defaultProps = { isPopperEnabled: false };

function SortButton<T extends Record<string, SortOrderModel>>({
  sort,
  setSort,
  sortFields,
  isPopperEnabled,
}: Props<T>) {
  // Setup translate
  const { t } = useTranslation();
  // States
  const [open, setOpen] = useState(false);
  const [anchorEl, setAnchorEl] = useState<HTMLButtonElement | null>(null);

  const handleClick = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  const sortDefaultProps = {
    initialSort: sort,
    open,
    sortFields,
    onSubmit: (args: T) => {
      setSort(args);
      handleClose();
    },
    onReset: () => {
      setSort({} as T);
      handleClose();
    },
    onClose: handleClose,
  };

  return (
    <>
      <Tooltip title={<>{t('common.sort.buttonTooltip')}</>}>
        <Button
          color="inherit"
          variant="outlined"
          sx={{ border: (theme) => `1px solid ${theme.palette.divider}`, padding: '5px 10px', minWidth: '46px' }}
          ref={(d: HTMLButtonElement) => {
            if (d && d !== anchorEl) {
              setAnchorEl(d);
            }
          }}
          onClick={handleClick}
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

SortButton.defaultProps = defaultProps;

export default SortButton;
