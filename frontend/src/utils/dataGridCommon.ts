import type { Theme } from '@mui/material/styles';
import type { SxProps } from '@mui/material';

export default function getDataGridCommonProps(t: (d: string) => string | undefined, sx: SxProps = {}) {
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
