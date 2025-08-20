import React, { useState, memo, useEffect } from 'react';
import Paper from '@mui/material/Paper';
import Divider from '@mui/material/Divider';
import Button from '@mui/material/Button';
import { mdiTune } from '@mdi/js';
import SvgIcon from '@mui/material/SvgIcon';
import InputBase from '@mui/material/InputBase';
import Tooltip from '@mui/material/Tooltip';
import { useTranslation } from 'react-i18next';
import FilterPopper from '../FilterPopper';
import FilterDialog from '../FilterDialog';
import { type FilterDefinitionFieldsModel } from '../../../models/general';
import { type FilterValueObject, type PredefinedFilter } from '../internal/types';

export interface Props<T extends FilterValueObject> {
  readonly filter: undefined | null | T;
  readonly onSubmit: (f: T) => void;
  readonly filterDefinitionModel: FilterDefinitionFieldsModel;
  readonly predefinedFilterObjects?: PredefinedFilter[];
  readonly isAdvancedFilterPopperEnabled?: boolean;
  readonly onMainSearchChange: (newValue: string, oldValue: string) => void;
  readonly mainSearchInitialValue: string;
  readonly mainSearchDisplay: string;
}

function FilterSearchBar<T extends FilterValueObject>({
  filter,
  onSubmit,
  filterDefinitionModel,
  predefinedFilterObjects = undefined,
  isAdvancedFilterPopperEnabled = undefined,
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

  const onClick = () => {
    setOpen(true);
  };

  const onClose = () => {
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
      ref={(d: HTMLFormElement) => {
        if (d && d !== anchorEl) {
          setAnchorEl(d);
        }
      }}
      style={open ? { borderBottomRightRadius: 0 } : {}}
      sx={{ display: 'flex', flex: 1, alignItems: 'center' }}
      variant="outlined"
    >
      <InputBase
        disabled={open}
        fullWidth
        onChange={(event) => {
          const newValue = event.target.value;
          // Check if values are different
          if (newValue !== value) {
            setValue(newValue);
          }
        }}
        placeholder={mainSearchDisplay}
        sx={{ ml: 1, flex: 1, minWidth: '200px' }}
        value={value}
      />
      <Divider orientation="vertical" sx={{ height: 28 }} />
      <Tooltip title={<>{t('common.filter.buttonTooltip')}</>}>
        <Button
          color={filter && Object.keys(filter).length !== 0 ? 'primary' : 'inherit'}
          onClick={onClick}
          sx={{
            padding: '5px 10px',
            minWidth: '46px',
            height: 38,
            borderRadius: '0px 2px 2px 0px',
          }}
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
    </Paper>
  );
}

export default memo(FilterSearchBar);
