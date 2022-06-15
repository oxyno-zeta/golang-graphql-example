import { Theme } from '@mui/material/styles';
import { TFunction } from 'react-i18next';

export default function getDataGridCommonProps(t: TFunction) {
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
    },
    localeText: {
      noRowsLabel: t('common.noData'),
    },
    isRowSelectable: () => false,
  };
}
