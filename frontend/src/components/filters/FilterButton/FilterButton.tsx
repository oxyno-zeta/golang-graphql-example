import React, { useState, memo } from 'react';
import Button from '@mui/material/Button';
import { mdiTune } from '@mdi/js';
import SvgIcon from '@mui/material/SvgIcon';
import Tooltip from '@mui/material/Tooltip';
import { useTranslation } from 'react-i18next';
import FilterPopper from '../FilterPopper';
import FilterDialog from '../FilterDialog';
import { FilterDefinitionFieldsModel } from '../../../models/general';
import { FilterValueObject, PredefinedFilter } from '../internal/types';

export type Props<T extends FilterValueObject> = {
  readonly filter: undefined | null | T;
  readonly onSubmit: (f: T) => void;
  readonly filterDefinitionModel: FilterDefinitionFieldsModel;
  readonly predefinedFilterObjects?: PredefinedFilter[];
  readonly isAdvancedFilterPopperEnabled?: boolean;
};

function FilterButton<T extends FilterValueObject>({
  filter,
  onSubmit,
  filterDefinitionModel,
  predefinedFilterObjects = undefined,
  isAdvancedFilterPopperEnabled = false,
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

  const defaultFilterProps = {
    filterDefinitionModel,
    initialFilter: filter,
    open,
    predefinedFilterObjects,
    onSubmit: (args: T) => {
      onSubmit(args);
      onClose();
    },
    onReset: () => {
      onSubmit({} as T);
      onClose();
    },
    onClose,
  };

  return (
    <>
      <Tooltip title={<>{t('common.filter.buttonTooltip')}</>}>
        <Button
          color={filter && Object.keys(filter).length !== 0 ? 'primary' : 'inherit'}
          onClick={onClick}
          ref={(d: HTMLButtonElement) => {
            if (d && d !== anchorEl) {
              setAnchorEl(d);
            }
          }}
          sx={{ border: (theme) => `1px solid ${theme.palette.divider}`, padding: '5px 10px', minWidth: '46px' }}
          variant="outlined"
        >
          <SvgIcon>
            <path d={mdiTune} />
          </SvgIcon>
        </Button>
      </Tooltip>
      {isAdvancedFilterPopperEnabled ? (
        <FilterPopper anchorElement={anchorEl} {...defaultFilterProps} />
      ) : (
        <FilterDialog {...defaultFilterProps} />
      )}
    </>
  );
}

export default memo(FilterButton);
