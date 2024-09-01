import React from 'react';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Accordion from '@mui/material/Accordion';
import AccordionSummary from '@mui/material/AccordionSummary';
import AccordionDetails from '@mui/material/AccordionDetails';

interface Props {
  readonly error: Error;
}

function FallbackErrorBoundary({ error }: Props) {
  return (
    <Box sx={{ p: 5 }}>
      <Typography component="h3" sx={{ color: 'red', marginBottom: '10px' }} variant="h5">
        Something went wrong: {error.message}
      </Typography>

      <Accordion>
        <AccordionSummary>
          <Typography>Open for technical details</Typography>
        </AccordionSummary>
        <AccordionDetails>
          <Typography>
            <pre>{error.stack}</pre>
          </Typography>
        </AccordionDetails>
      </Accordion>
    </Box>
  );
}

export default FallbackErrorBoundary;
