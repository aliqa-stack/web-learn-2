import React from 'react'
import { Routes, Route, BrowserRouter } from "react-router-dom";
import Navbar from './component/Navbar'
import Header from './component/Header'
import Sections from './component/Section/Sections'
import Heading from './component/Heading/Heading'
import Information from './component/Informations/Information'
import Login from './component/Login.jsx/LoginSection'

// Buat komponen khusus untuk halaman utama
const HomePage = () => (
  <>
    <Navbar />
    <Header />
    <Heading />
    <Sections />
    <Information />
  </>
)

const App = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />}/>
        <Route path="/login" element={<Login />}/>  {/* Tambah slash di depan */}
      </Routes>
    </BrowserRouter>
  )
}

export default App