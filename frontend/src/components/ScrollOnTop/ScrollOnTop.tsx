import React from 'react';
import useScrollTrigger from '@mui/material/useScrollTrigger';
import Box from '@mui/material/Box';
import Fab from '@mui/material/Fab';
import Fade from '@mui/material/Fade';
import SvgIcon from '@mui/material/SvgIcon';
import { mdiChevronUp } from '@mdi/js';

function ScrollOnTop() {
  const trigger = useScrollTrigger({
    disableHysteresis: true,
    threshold: 100,
  });

  const handleClick = (event: React.MouseEvent<HTMLDivElement>) => {
    const anchor = ((event.target as HTMLDivElement).ownerDocument || document).querySelector('#topbar');

    if (anchor) {
      anchor.scrollIntoView({
        block: 'center',
      });
    }
  };

  return (
    <Fade in={trigger}>
      <Box onClick={handleClick} role="presentation" sx={{ position: 'fixed', bottom: 16, right: 16 }}>
        <Fab size="small">
          <SvgIcon>
            <path d={mdiChevronUp} />
          </SvgIcon>
        </Fab>
      </Box>
    </Fade>
  );
}

export default ScrollOnTop;
