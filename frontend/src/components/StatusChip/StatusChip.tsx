import React from 'react';
import Chip, { ChipProps } from '@mui/material/Chip';

type Props = ChipProps;

function StatusChip({ label, color, sx = {}, ...rest }: Props) {
  return (
    <Chip
      label={label}
      color={color}
      variant="outlined"
      size="small"
      sx={{
        ...sx,
        borderRadius: '5px',
        backgroundColor: (theme) =>
          color === 'default' || !color ? theme.palette.action.selected : `${theme.palette[color].main}60`,
      }}
      {...rest}
    />
  );
}

export default StatusChip;
