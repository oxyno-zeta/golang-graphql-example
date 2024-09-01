import React from 'react';
import { useTranslation } from 'react-i18next';
import { DataGrid, GridColDef, GridRowParams } from '@mui/x-data-grid';
import IconButton from '@mui/material/IconButton';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiPencil, mdiDelete } from '@mdi/js';
import useMediaQuery from '@mui/material/useMediaQuery';
import { useTheme } from '@mui/material/styles';
import { TodoModel, TodoSortOrderModel } from '../../../../models/todos';
import { ConnectionModel } from '../../../../models/general';
import getDataGridCommonProps from '../../../../utils/dataGridCommon';
import { buildMUIXSort, setMUIXSortBuilder } from '../../../../components/sorts/utils';
import useTimezone from '../../../../components/timezone/useTimezone';
import { getDayjsTz } from '../../../../components/timezone/utils';

interface Props {
  readonly data: ConnectionModel<TodoModel> | undefined;
  readonly loading: boolean;
  readonly sorts: TodoSortOrderModel[];
  readonly setSorts: (data: TodoSortOrderModel[]) => void;
}

function TableView({ data, loading, sorts, setSorts }: Props) {
  // Setup translate
  const { t } = useTranslation();
  // Get if window have size matching request
  const theme = useTheme();
  const sizeMatching = useMediaQuery(theme.breakpoints.up('lg'));
  const timezone = useTimezone();

  let items: TodoModel[] = [];
  if (data && data.edges) {
    items = data.edges.map((it) => it.node);
  }

  const handleClick = (params: GridRowParams<TodoModel>) => {
    const { row } = params;

    console.log(row);
  };

  const columns: GridColDef<TodoModel>[] = [
    {
      field: 'text',
      headerName: t('todos.fields.text'),
      flex: 1,
      editable: false,
      filterable: false,
      sortable: true,
      disableColumnMenu: true,
    },
    {
      field: 'createdAt',
      headerName: t('common.fields.createdAt'),
      flex: 0.2,
      editable: false,
      filterable: false,
      sortable: true,
      disableColumnMenu: true,
      valueGetter: (value: string) => getDayjsTz(value, timezone).format('LLLL'),
    },
    {
      field: 'done',
      headerName: t('todos.fields.done'),
      flex: 0.1,
      editable: false,
      filterable: false,
      sortable: true,
      disableColumnMenu: true,
    },
    {
      field: 'actions',
      type: 'actions',
      headerName: t('common.actions'),
      width: 90,
      getActions: () => [
        <IconButton key="0" size="small">
          <SvgIcon style={{ fontSize: '1.25rem' }}>
            <path d={mdiPencil} />
          </SvgIcon>
        </IconButton>,
        <IconButton key="1" size="small">
          <SvgIcon style={{ fontSize: '1.25rem' }}>
            <path d={mdiDelete} />
          </SvgIcon>
        </IconButton>,
      ],
    },
  ];

  return (
    <DataGrid
      {...getDataGridCommonProps(t)}
      columns={columns}
      loading={loading}
      onRowClick={!sizeMatching ? handleClick : undefined}
      onRowDoubleClick={sizeMatching ? handleClick : undefined}
      onSortModelChange={setMUIXSortBuilder(setSorts)}
      rows={items}
      sortModel={buildMUIXSort(sorts, columns)}
      sortingMode="server"
    />
  );
}

export default TableView;
