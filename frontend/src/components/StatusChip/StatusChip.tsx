import React from 'react';
import Chip, { ChipProps } from '@mui/material/Chip';

export type Props = ChipProps;

function StatusChip({ label, color, sx = {}, ...rest }: Props) {
  return (
    <Chip
      clickable={false}
      color={color}
      label={label}
      size="small"
      sx={{
        ...sx,
        borderRadius: '5px',
        backgroundColor: (theme) =>
          color === 'default' || !color ? theme.palette.action.selected : `${theme.palette[color].main}60`,
      }}
      variant="outlined"
      {...rest}
    />
  );
}

export default StatusChip;
