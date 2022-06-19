import React, { useState, memo } from 'react';
import Button from '@mui/material/Button';
import TuneIcon from '@mui/icons-material/Tune';
import Tooltip from '@mui/material/Tooltip';
import { useTranslation } from 'react-i18next';
import FilterPopper from '../FilterPopper';
import FilterDialog from '../FilterDialog';
import { FilterDefinitionFieldsModel } from '../../../models/general';
import { FilterValueObject, PredefinedFilter } from '../internal/types';

type Props<T extends FilterValueObject> = {
  filter: undefined | null | T;
  setFilter: React.Dispatch<T>;
  filterDefinitionModel: FilterDefinitionFieldsModel;
  predefinedFilterObjects?: PredefinedFilter[];
  isAdvancedFilterPopperEnabled?: boolean;
};

function FilterButton<T extends FilterValueObject>({
  filter,
  setFilter,
  filterDefinitionModel,
  predefinedFilterObjects,
  isAdvancedFilterPopperEnabled,
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

  const defaultFilterProps = {
    filterDefinitionModel,
    initialFilter: filter,
    open,
    predefinedFilterObjects,
    onSubmit: (args: T) => {
      setFilter(args);
      handleClose();
    },
    onReset: () => {
      setFilter({} as T);
      handleClose();
    },
    onClose: handleClose,
  };

  return (
    <>
      <Tooltip title={<>{t('common.filter.buttonTooltip')}</>}>
        <Button
          color="inherit"
          variant="outlined"
          sx={{ border: (theme) => `1px solid ${theme.palette.divider}`, padding: '5px 10px', minWidth: '46px' }}
          onClick={handleClick}
          ref={(d: HTMLButtonElement) => {
            if (d && d !== anchorEl) {
              setAnchorEl(d);
            }
          }}
        >
          <TuneIcon />
        </Button>
      </Tooltip>
      {isAdvancedFilterPopperEnabled ? (
        <FilterPopper<T> anchorElement={anchorEl} {...defaultFilterProps} />
      ) : (
        <FilterDialog<T> {...defaultFilterProps} />
      )}
    </>
  );
}

FilterButton.defaultProps = {
  predefinedFilterObjects: undefined,
  isAdvancedFilterPopperEnabled: false,
};

export default memo(FilterButton);
