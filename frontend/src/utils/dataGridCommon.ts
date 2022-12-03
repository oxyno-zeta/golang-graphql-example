import type { Theme } from '@mui/material/styles';

// eslint-disable-next-line @typescript-eslint/no-explicit-any
export default function getDataGridCommonProps(t: (d: string) => string | undefined, sx: any = {}) {
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
        backgroundColor: (theme: Theme) =>
          theme.palette.mode === 'light' ? theme.palette.grey['200'] : theme.palette.grey['800'],
        borderTopLeftRadius: '0px',
        borderTopRightRadius: '0px',
      },
      ...sx,
    },
    localeText: {
      noRowsLabel: t('common.noData'),
    },
    isRowSelectable: () => false,
  };
}
