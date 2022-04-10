import React from 'react';
import { useTranslation } from 'react-i18next';
import Button from '@mui/material/Button';
import ButtonGroup from '@mui/material/ButtonGroup';
import Tooltip from '@mui/material/Tooltip';
import GridViewIcon from '@mui/icons-material/GridView';
import TableRowsIcon from '@mui/icons-material/TableRows';

interface Props {
  setGridView: (input: boolean) => void;
  gridView: boolean;
}

function GridTableViewSwitcher({ setGridView, gridView }: Props) {
  // Get translator
  const { t } = useTranslation();

  return (
    <ButtonGroup variant="outlined" sx={{ border: (theme) => `1px solid ${theme.palette.divider}` }}>
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
        >
          <GridViewIcon color={gridView ? 'primary' : 'inherit'} />
        </Button>
      </Tooltip>
    </ButtonGroup>
  );
}

export default GridTableViewSwitcher;
