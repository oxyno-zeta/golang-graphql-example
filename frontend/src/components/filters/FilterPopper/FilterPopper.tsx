import React, { useCallback, useEffect, useRef, useState } from 'react';
import Button from '@mui/material/Button';
import DialogActions from '@mui/material/DialogActions';
import Popper from '@mui/material/Popper';
import Paper from '@mui/material/Paper';
import ClickAwayListener from '@mui/material/ClickAwayListener';
import { useTranslation } from 'react-i18next';
import FilterForm from '../internal/components/FilterForm';
import { FilterDefinitionFieldsModel } from '../../../models/general';
import { FilterValueObject, PredefinedFilter } from '../internal/types';

type Props<T extends FilterValueObject> = {
  onSubmit: (filter: T) => void;
  onReset: () => void;
  onClose: () => void;
  open: boolean;
  filterDefinitionModel: FilterDefinitionFieldsModel;
  predefinedFilterObjects?: PredefinedFilter[];
  initialFilter?: undefined | null | T;
  anchorElement: HTMLFormElement | HTMLButtonElement | null;
};

const defaultProps = {
  predefinedFilterObjects: undefined,
  initialFilter: undefined,
};

function FilterPopper<T extends FilterValueObject>({
  filterDefinitionModel,
  predefinedFilterObjects,
  onSubmit,
  onReset,
  onClose,
  initialFilter,
  open,
  anchorElement,
}: Props<T>) {
  // Setup translate
  const { t } = useTranslation();
  // State
  const resultRef = useRef<T | undefined | null>(null);
  const [error, setError] = useState(false);

  const localOnChange = useCallback((filter: null | FilterValueObject) => {
    resultRef.current = filter as T | null;

    if (filter === null || filter === undefined) {
      setError(true);
    } else {
      setError(false);
    }
  }, []);

  // Watch initial filter
  useEffect(() => {
    setError(false);
    resultRef.current = initialFilter;
  }, [initialFilter]);

  return (
    <Popper
      open={open}
      anchorEl={anchorElement}
      placement="bottom-end"
      disablePortal={false}
      modifiers={[]}
      sx={{
        minWidth: {
          xs: '300px',
          sm: '550px',
          md: '700px',
          lg: '950px',
          xl: '1100px',
        },
        maxHeight: '500px',
        overflowY: 'auto',
      }}
    >
      <ClickAwayListener onClickAway={onClose}>
        <Paper
          variant="outlined"
          sx={{
            borderTopLeftRadius: 0,
            borderTopRightRadius: 0,
            borderTop: '0px',
          }}
        >
          <div style={{ padding: '10px' }}>
            <FilterForm
              filterDefinitionModel={filterDefinitionModel}
              onChange={localOnChange}
              initialFilter={initialFilter}
              predefinedFilterObjects={predefinedFilterObjects}
            />
          </div>
          <DialogActions>
            <Button
              onClick={() => {
                onReset();
              }}
              sx={{ marginLeft: 'auto', marginRight: '5px' }}
            >
              {t('common.resetAction')}
            </Button>
            <Button
              variant="contained"
              disabled={error}
              onClick={() => {
                onSubmit(resultRef.current as T);
              }}
              autoFocus
            >
              {t('common.applyAction')}
            </Button>
          </DialogActions>
        </Paper>
      </ClickAwayListener>
    </Popper>
  );
}

FilterPopper.defaultProps = defaultProps;

export default FilterPopper;
