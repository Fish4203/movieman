import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { BrowserRouter, Routes, Route } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";


import Search from './pages/Search';
import Details from './pages/Details';
import NoPage from './pages/NoPage';
import Home from './pages/Home';
import Login from './pages/Login';
import Profile from './pages/Profile';
// import Torent from './pages/Torrent';
// import Watch from './pages/Watch';
import Nav from './components/Nav';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <BrowserRouter>
      <Routes>
        <Route path="/" element={<Nav />}>
          <Route index element={<Home />} />
          <Route path="search" element={<Search />} />
          <Route path="login" element={<Login />} />
          <Route path="Profile" element={<Profile />} />
          <Route path="details/:type/:id" element={<Details />} />
          <Route path="*" element={<NoPage />} />
        </Route>
      </Routes>
    </BrowserRouter>
);
