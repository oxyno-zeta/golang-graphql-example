import React from 'react';
import { useTranslation } from 'react-i18next';
import { DataGrid, GridRowParams } from '@mui/x-data-grid';
import IconButton from '@mui/material/IconButton';
import DeleteIcon from '@mui/icons-material/Delete';
import EditIcon from '@mui/icons-material/Edit';
import dayjs from 'dayjs';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useTheme } from '@mui/material/styles';
import { TodoModel } from '../../../../models/todos';
import { ConnectionModel } from '../../../../models/general';

interface Props {
  data: ConnectionModel<TodoModel> | undefined;
  loading: boolean;
}

function TableView({ data, loading }: Props) {
  // Setup translate
  const { t } = useTranslation();
  // Get if window have size matching request
  const theme = useTheme();
  const sizeMatching = useMediaQuery(theme.breakpoints.up('lg'));

  let items: TodoModel[] = [];
  if (data && data.edges) {
    items = data.edges.map((it) => it.node);
  }

  const handleClick = (params: GridRowParams<TodoModel>) => {
    const { row } = params;

    console.log(row);
  };

  return (
    <DataGrid
      hideFooterPagination
      hideFooter
      autoHeight
      sx={{
        border: '0px',
        borderRadius: '0px',
        '& .MuiDataGrid-cell:focus': { outline: 'none' },
        '& .MuiDataGrid-cell:focus-within': { outline: 'none' },
        '& .MuiDataGrid-cell:focus-visible': { outline: 'none' },
        '& .MuiDataGrid-columnHeader:focus': { outline: 'none' },
        '& .MuiDataGrid-columnHeader:focus-within': { outline: 'none' },
        '& .MuiDataGrid-columnHeader:focus-visible': { outline: 'none' },
        '& .MuiDataGrid-columnHeaders': {
          backgroundColor: theme.palette.mode === 'light' ? theme.palette.grey['200'] : 'inherit',
        },
      }}
      localeText={{
        noRowsLabel: t('common.noData'),
      }}
      loading={loading}
      columns={[
        {
          field: 'text',
          headerName: t('todos.fields.text'),
          flex: 1,
          editable: false,
          filterable: false,
          sortable: false,
          disableColumnMenu: true,
        },
        {
          field: 'createdAt',
          headerName: t('common.fields.createdAt'),
          flex: 0.2,
          editable: false,
          filterable: false,
          sortable: false,
          disableColumnMenu: true,
          valueGetter: (params) => dayjs(params.value).format('LLLL'),
        },
        {
          field: 'done',
          headerName: t('todos.fields.done'),
          flex: 0.1,
          editable: false,
          filterable: false,
          sortable: false,
          disableColumnMenu: true,
        },
        {
          field: 'actions',
          headerName: t('common.actions'),
          flex: 0.1,
          editable: false,
          filterable: false,
          sortable: false,
          disableColumnMenu: true,
          renderCell: () => (
            <>
              <IconButton>
                <EditIcon />
              </IconButton>
              <IconButton>
                <DeleteIcon />
              </IconButton>
            </>
          ),
        },
      ]}
      isRowSelectable={() => false}
      onRowDoubleClick={sizeMatching ? handleClick : undefined}
      onRowClick={!sizeMatching ? handleClick : undefined}
      rows={items}
    />
  );
}

export default TableView;
