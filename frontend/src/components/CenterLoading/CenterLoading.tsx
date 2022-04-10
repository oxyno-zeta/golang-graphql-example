import React from 'react';
import CircularProgress from '@mui/material/CircularProgress';

function CenterLoading() {
  return (
    <div style={{ display: 'flex', justifyContent: 'center', margin: '10px 0' }}>
      <CircularProgress />
    </div>
  );
}

export default CenterLoading;
