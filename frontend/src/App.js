import React from 'react';
import Header from './components/Header';
import AlbumList from './pages/AlbumList';
import AlbumDetail from './pages/AlbumDetail';

import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';

function App() {
  return (
    <Router>
      <div>
        <Header />
        {/* <Navbar /> */}
        <Routes>
          <Route path="/" element={<AlbumList />} />
          <Route path="/album/:id" element={<AlbumDetail />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
