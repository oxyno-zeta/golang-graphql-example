import React, { useState } from 'react';
import { useTranslation } from 'react-i18next';
import Grid from '@mui/material/Grid';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import IconButton from '@mui/material/IconButton';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiDotsVertical } from '@mdi/js';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import { type TodoModel } from '../../../../models/todos';
import { getDayjsTz } from '../../../../components/timezone/utils';
import useTimezone from '../../../../components/timezone/useTimezone';

interface Props {
  readonly item: TodoModel;
}

function GridViewItem({ item }: Props) {
  // Setup translate
  const { t } = useTranslation();

  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const timezone = useTimezone();

  return (
    <Grid
      size={{
        xl: 3,
        lg: 3,
        md: 6,
        sm: 6,
        xs: 12,
      }}
    >
      <Card>
        <CardHeader
          action={
            <IconButton
              onClick={(event: React.MouseEvent<HTMLElement>) => {
                setAnchorEl(event.currentTarget);
              }}
            >
              <SvgIcon>
                <path d={mdiDotsVertical} />
              </SvgIcon>
            </IconButton>
          }
          subheader={getDayjsTz(item.createdAt, timezone).format('LLLL')}
          title={item.text}
        />
      </Card>
      <Menu
        anchorEl={anchorEl}
        open={Boolean(anchorEl)}
        onClose={() => {
          setAnchorEl(null);
        }}
        anchorOrigin={{
          vertical: 'bottom',
          horizontal: 'left',
        }}
        transformOrigin={{
          vertical: 'bottom',
          horizontal: 'left',
        }}
      >
        <MenuItem
          onClick={() => {
            setAnchorEl(null);
          }}
        >
          {t('common.editAction')}
        </MenuItem>
        <MenuItem
          onClick={() => {
            setAnchorEl(null);
          }}
        >
          {t('common.deleteAction')}
        </MenuItem>
      </Menu>
    </Grid>
  );
}

export default GridViewItem;
