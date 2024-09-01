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

export type Props<T extends FilterValueObject> = {
  readonly onSubmit: (filter: T) => void;
  readonly onReset: () => void;
  readonly onClose: () => void;
  readonly open: boolean;
  readonly filterDefinitionModel: FilterDefinitionFieldsModel;
  readonly predefinedFilterObjects?: PredefinedFilter[];
  readonly initialFilter?: undefined | null | T;
  readonly anchorElement: HTMLFormElement | HTMLButtonElement | null;
};

function FilterPopper<T extends FilterValueObject>({
  filterDefinitionModel,
  predefinedFilterObjects = undefined,
  onSubmit,
  onReset,
  onClose,
  initialFilter = undefined,
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
      anchorEl={anchorElement}
      disablePortal={false}
      modifiers={[]}
      open={open}
      placement="bottom-end"
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
          sx={{
            borderTopLeftRadius: 0,
            borderTopRightRadius: 0,
            borderTop: '0px',
          }}
          variant="outlined"
        >
          <div style={{ padding: '10px' }}>
            <FilterForm
              filterDefinitionModel={filterDefinitionModel}
              initialFilter={initialFilter}
              onChange={localOnChange}
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
              autoFocus
              disabled={error}
              onClick={() => {
                onSubmit(resultRef.current as T);
              }}
              variant="contained"
            >
              {t('common.applyAction')}
            </Button>
          </DialogActions>
        </Paper>
      </ClickAwayListener>
    </Popper>
  );
}

export default FilterPopper;
