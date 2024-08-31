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
      '& .MuiDataGrid-cell': { display: 'flex', alignItems: 'center' },
      '& .MuiDataGrid-cell:focus': { outline: 'none' },
      '& .MuiDataGrid-cell:focus-within': { outline: 'none' },
      '& .MuiDataGrid-cell:focus-visible': { outline: 'none' },
      '& .MuiDataGrid-columnHeader:focus': { outline: 'none' },
      '& .MuiDataGrid-columnHeader:focus-within': { outline: 'none' },
      '& .MuiDataGrid-columnHeader:focus-visible': { outline: 'none' },
      '& .MuiDataGrid-columnHeaders': {
        borderTopLeftRadius: '0px',
        borderTopRightRadius: '0px',
      },
      '& .MuiDataGrid-columnHeader': {
        borderTopLeftRadius: '0px !important',
        borderTopRightRadius: '0px !important',
        backgroundColor: (theme: Theme) =>
          theme.palette.mode === 'light' ? theme.palette.grey['200'] : theme.palette.grey['800'],
      },
      ...sx,
    },
    localeText: {
      noRowsLabel: t('common.noData'),
    },
    isRowSelectable: () => false,
  };
}
