import React from 'react';
import { Routes, Route } from 'react-router-dom';
import Todos from './routes/Todos';
import NotFoundRoute from './components/NotFoundRoute';
import MainContentWrapper from './components/MainContentWrapper';

function App() {
  return (
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
  );
}

export default App;
