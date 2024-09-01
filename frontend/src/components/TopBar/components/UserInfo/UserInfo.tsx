import React from 'react';
import Typography from '@mui/material/Typography';

function UserInfo() {
  return (
    <>
      <Typography gutterBottom style={{ fontSize: 14 }}>
        Fake User
      </Typography>
      <Typography color="text.secondary" gutterBottom style={{ fontSize: 11 }}>
        fake@fake.com
      </Typography>
    </>
  );
}

export default UserInfo;
