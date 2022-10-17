import React from 'react';
import Typography from '@mui/material/Typography';

function UserInfo() {
  return (
    <>
      <Typography style={{ fontSize: 14 }} gutterBottom>
        Fake User
      </Typography>
      <Typography style={{ fontSize: 11 }} color="text.secondary" gutterBottom>
        fake@fake.com
      </Typography>
    </>
  );
}

export default UserInfo;
