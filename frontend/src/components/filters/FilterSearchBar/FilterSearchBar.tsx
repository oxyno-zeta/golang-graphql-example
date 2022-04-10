import React, { useState, memo, useEffect } from 'react';
import Paper from '@mui/material/Paper';
import Divider from '@mui/material/Divider';
import Button from '@mui/material/Button';
import TuneIcon from '@mui/icons-material/Tune';
import InputBase from '@mui/material/InputBase';
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
  onMainSearchChange: (newValue: string, oldValue: string) => void;
  mainSearchInitialValue: string;
  mainSearchDisplay: string;
};

function FilterSearchBar<T extends FilterValueObject>({
  filter,
  setFilter,
  filterDefinitionModel,
  predefinedFilterObjects,
  isAdvancedFilterPopperEnabled,
  onMainSearchChange,
  mainSearchInitialValue,
  mainSearchDisplay,
}: Props<T>) {
  // Setup translate
  const { t } = useTranslation();
  // States
  const [open, setOpen] = useState(false);
  const [anchorEl, setAnchorEl] = useState<HTMLFormElement | null>(null);
  const [value, setValue] = useState(mainSearchInitialValue);

  const handleClick = () => {
    setOpen(true);
  };

  const handleClose = () => {
    setOpen(false);
  };

  // Watch main search initial value
  useEffect(() => {
    setValue(mainSearchInitialValue);
  }, [mainSearchInitialValue]);

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
    <Paper
      component="form"
      onSubmit={(event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();

        // Check if initial value and value now are different
        if (value !== mainSearchInitialValue) {
          // Call the hook
          onMainSearchChange(value, mainSearchInitialValue);
        }
      }}
      sx={{ display: 'flex', flex: 1, alignItems: 'center' }}
      style={open ? { borderBottomRightRadius: 0 } : {}}
      variant="outlined"
      ref={(d: HTMLFormElement) => {
        if (d && d !== anchorEl) {
          setAnchorEl(d);
        }
      }}
    >
      <InputBase
        sx={{ ml: 1, flex: 1 }}
        placeholder={mainSearchDisplay}
        fullWidth
        disabled={open}
        value={value}
        onChange={(event) => {
          const newValue = event.target.value;
          // Check if values are different
          if (newValue !== value) {
            setValue(newValue);
          }
        }}
      />
      <Divider sx={{ height: 28 }} orientation="vertical" />
      <Tooltip title={<>{t('common.filter.buttonTooltip')}</>}>
        <Button
          color={filter && Object.keys(filter).length !== 0 ? 'primary' : 'inherit'}
          sx={{ padding: '5px 10px', minWidth: '46px', height: 38 }}
          onClick={handleClick}
        >
          <TuneIcon />
        </Button>
      </Tooltip>
      {isAdvancedFilterPopperEnabled ? (
        <FilterPopper<T> anchorElement={anchorEl} {...defaultFilterProps} />
      ) : (
        <FilterDialog<T> {...defaultFilterProps} />
      )}
    </Paper>
  );
}

FilterSearchBar.defaultProps = {
  predefinedFilterObjects: undefined,
  isAdvancedFilterPopperEnabled: false,
};

export default memo(FilterSearchBar);
