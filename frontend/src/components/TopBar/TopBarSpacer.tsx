import React from 'react';

function TopBarSpacer() {
  // First div is for upper div in flex
  return (
    <div>
      <div
        style={{
          height: '48px', // This size is the app bar dense mode size. This isn't present in the theme so...
        }}
      />
    </div>
  );
}

export default TopBarSpacer;
