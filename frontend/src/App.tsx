import React from 'react';
import { Routes, Route } from 'react-router-dom';
import NotFoundRoute from '~components/NotFoundRoute';
import MainContentWrapper from '~components/MainContentWrapper';
import Todos from '~routes/Todos';

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
