import React from 'react'
import Navbar from './component/Navbar'
import Header from './component/Header'
import Sections from './component/Section/Sections'
import Heading from './component/Heading/Heading'
import Information from './component/Informations/Information'

const App = () => {
  return (
    <div>

      <Navbar />
      <Header />
      <Heading />
      <Sections />
      <Information />
    </div>

  )
}

export default App