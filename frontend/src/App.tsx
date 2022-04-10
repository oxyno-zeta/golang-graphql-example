import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Box from '@mui/material/Box';
import Todos from './routes/Todos';

function App() {
  return (
    <Box sx={{ margin: '0 20px 20px 20px' }}>
      <Routes>
        <Route path="*" element={<Todos />} />
      </Routes>
    </Box>
  );
}

export default App;
