import React, { useState } from 'react';
import Grid from '@mui/material/Grid';
import Card from '@mui/material/Card';
import CardHeader from '@mui/material/CardHeader';
import IconButton from '@mui/material/IconButton';
import MoreVertIcon from '@mui/icons-material/MoreVert';
import Menu from '@mui/material/Menu';
import MenuItem from '@mui/material/MenuItem';
import dayjs from 'dayjs';
import { TodoModel } from '../../../../models/todos';

interface Props {
  item: TodoModel;
}

function GridViewItem({ item }: Props) {
  const [anchorEl, setAnchorEl] = useState<null | HTMLElement>(null);

  return (
    <Grid item xl={3} lg={3} md={6} sm={6} xs={12}>
      <Card>
        <CardHeader
          action={
            <IconButton
              onClick={(event: React.MouseEvent<HTMLElement>) => {
                setAnchorEl(event.currentTarget);
              }}
            >
              <MoreVertIcon />
            </IconButton>
          }
          title={item.text}
          subheader={dayjs(item.createdAt).format('LLLL')}
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
