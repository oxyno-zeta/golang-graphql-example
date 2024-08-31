import React, { useState } from 'react';
import Grid from '@mui/material/Grid2';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import IconButton from '@mui/material/IconButton';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiDotsVertical } from '@mdi/js';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import { TodoModel } from '../../../../models/todos';
import { getDayjsTz } from '../../../../components/timezone/utils';
import useTimezone from '../../../../components/timezone/useTimezone';

interface Props {
  item: TodoModel;
}

function GridViewItem({ item }: Props) {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);
  const timezone = useTimezone();

  return (
    <Grid size={{ xl: 3, lg: 3, md: 6, sm: 6, xs: 12 }}>
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
          title={item.text}
          subheader={getDayjsTz(item.createdAt, timezone).format('LLLL')}
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
          Edit
        </MenuItem>
        <MenuItem
          onClick={() => {
            setAnchorEl(null);
          }}
        >
          Delete
        </MenuItem>
      </Menu>
    </Grid>
  );
}

export default GridViewItem;
