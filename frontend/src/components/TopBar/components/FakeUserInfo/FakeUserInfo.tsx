import React from 'react';
import Typography from '@mui/material/Typography';

function FakeUserInfo() {
  return (
    <>
      <Typography gutterBottom style={{ fontSize: 14 }}>
        Fake User
      </Typography>
      <Typography
        gutterBottom
        style={{ fontSize: 11 }}
        sx={{
          color: 'text.secondary',
        }}
      >
        fake@fake.com
      </Typography>
    </>
  );
}

export default FakeUserInfo;
