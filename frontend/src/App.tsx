import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Box from '@mui/material/Box';
import Todos from './routes/Todos';
import NotFoundRoute from './components/NotFoundRoute';
import MainContentWrapper from './components/MainContentWrapper';

function App() {
  return (
    <Box>
      <Routes>
        <Route path="/">
          <Route
            index
            element={
              <MainContentWrapper>
                <Todos />
              </MainContentWrapper>
            }
          />
          <Route
            path="*"
            element={
              <MainContentWrapper>
                <NotFoundRoute />
              </MainContentWrapper>
            }
          />
        </Route>
      </Routes>
    </Box>
  );
}

export default App;
