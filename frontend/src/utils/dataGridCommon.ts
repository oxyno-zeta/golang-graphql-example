import { Theme } from '@mui/material/styles';
import { TFunction } from 'react-i18next';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export default function getDataGridCommonProps(t: TFunction, sx: any = {}) {
  return {
    hideFooterPagination: true,
    hideFooter: true,
    autoHeight: true,
    sx: {
      border: '0px',
      borderRadius: '0px',
      '& .MuiDataGrid-cell:focus': { outline: 'none' },
      '& .MuiDataGrid-cell:focus-within': { outline: 'none' },
      '& .MuiDataGrid-cell:focus-visible': { outline: 'none' },
      '& .MuiDataGrid-columnHeader:focus': { outline: 'none' },
      '& .MuiDataGrid-columnHeader:focus-within': { outline: 'none' },
      '& .MuiDataGrid-columnHeader:focus-visible': { outline: 'none' },
      '& .MuiDataGrid-columnHeaders': {
        backgroundColor: (theme: Theme) => (theme.palette.mode === 'light' ? theme.palette.grey['200'] : 'inherit'),
      },
      ...sx,
    },
    localeText: {
      noRowsLabel: t('common.noData'),
    },
    isRowSelectable: () => false,
  };
}
