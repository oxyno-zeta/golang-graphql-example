import React, { useContext } from 'react';
import { useTranslation } from 'react-i18next';
import Button, { ButtonProps } from '@mui/material/Button';
import ButtonGroup, { ButtonGroupProps } from '@mui/material/ButtonGroup';
import Tooltip from '@mui/material/Tooltip';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiViewGridOutline, mdiViewSequential } from '@mdi/js';
import GridTableViewSwitcherContext from '~contexts/GridTableViewSwitcherContext';

export interface Props {
  buttonGroupProps?: ButtonGroupProps;
  tableButtonProps?: Omit<ButtonProps, 'onClick'>;
  gridButtonProps?: Omit<ButtonProps, 'onClick'>;
}

const defaultProps = {
  buttonGroupProps: {},
  tableButtonProps: {},
  gridButtonProps: {},
};

function GridTableViewSwitcher({ buttonGroupProps, tableButtonProps, gridButtonProps }: Props) {
  // Get translator
  const { t } = useTranslation();
  // Get context
  const { isGridViewEnabled, toggleGridTableView } = useContext(GridTableViewSwitcherContext);
  // Get grid view
  const gridView = isGridViewEnabled();

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
            // Optimization
            if (gridView) {
              toggleGridTableView();
            }
          }}
          {...tableButtonProps}
        >
          <SvgIcon color={gridView ? 'inherit' : 'primary'}>
            <path d={mdiViewSequential} />
          </SvgIcon>
        </Button>
      </Tooltip>
      <Tooltip title={<>{t('common.gridViewTooltip')}</>}>
        <Button
          color="inherit"
          sx={{ border: 'none', padding: '5px 10px', minWidth: '46px' }}
          onClick={() => {
            // Optimization
            if (!gridView) {
              toggleGridTableView();
            }
          }}
          {...gridButtonProps}
        >
          <SvgIcon color={gridView ? 'primary' : 'inherit'}>
            <path d={mdiViewGridOutline} />
          </SvgIcon>
        </Button>
      </Tooltip>
    </ButtonGroup>
  );
}

GridTableViewSwitcher.defaultProps = defaultProps;

export default GridTableViewSwitcher;
