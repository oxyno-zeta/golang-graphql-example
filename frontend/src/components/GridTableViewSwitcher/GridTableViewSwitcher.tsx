import React from 'react';
import { useTranslation } from 'react-i18next';
import Button, { ButtonProps } from '@mui/material/Button';
import ButtonGroup, { ButtonGroupProps } from '@mui/material/ButtonGroup';
import Tooltip from '@mui/material/Tooltip';
import GridViewIcon from '@mui/icons-material/GridView';
import TableRowsIcon from '@mui/icons-material/TableRows';

interface Props {
  setGridView: (input: boolean) => void;
  gridView: boolean;
  buttonGroupProps?: ButtonGroupProps;
  tableButtonProps?: Omit<ButtonProps, 'onClick'>;
  gridButtonProps?: Omit<ButtonProps, 'onClick'>;
}

const defaultProps = {
  buttonGroupProps: {},
  tableButtonProps: {},
  gridButtonProps: {},
};

function GridTableViewSwitcher({ setGridView, gridView, buttonGroupProps, tableButtonProps, gridButtonProps }: Props) {
  // Get translator
  const { t } = useTranslation();

  return (
    <ButtonGroup
      variant="outlined"
      sx={{ border: (theme) => `1px solid ${theme.palette.divider}` }}
      {...buttonGroupProps}
    >
      <Tooltip title={<>{t('common.tableViewTooltip')}</>}>
        <Button
          color="inherit"
          sx={{
            border: 'none',
            padding: '5px 10px',
            minWidth: '46px',
          }}
          onClick={() => {
            setGridView(false);
          }}
          {...tableButtonProps}
        >
          <TableRowsIcon color={gridView ? 'inherit' : 'primary'} />
        </Button>
      </Tooltip>
      <Tooltip title={<>{t('common.gridViewTooltip')}</>}>
        <Button
          color="inherit"
          sx={{ border: 'none', padding: '5px 10px', minWidth: '46px' }}
          onClick={() => {
            setGridView(true);
          }}
          {...gridButtonProps}
        >
          <GridViewIcon color={gridView ? 'primary' : 'inherit'} />
        </Button>
      </Tooltip>
    </ButtonGroup>
  );
}

GridTableViewSwitcher.defaultProps = defaultProps;

export default GridTableViewSwitcher;
