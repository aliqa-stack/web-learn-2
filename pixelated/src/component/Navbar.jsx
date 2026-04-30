import React from 'react'
import { IoMdNotifications } from "react-icons/io";
import { ImPacman } from "react-icons/im";

const Navbar = () => {


  const notif = ()=> {
    window.alert("Soon");
  }

  return (
     <nav className='relative bg-gradient-to-br from-[#2563EB] to-[#0EA5E9]   after:pointer-events-none after:absolute after:inset-x-0 after:bottom-0 after:h-px after:bg-white/10'>
      <div className='mx-auto max-w-7xl px-2 sm:px-6 lg:px-8'>
   
          <div className='relative h-16 flex items-center justify-between'>
            <div className=' flex items-center '>
            <ImPacman className='h-12 w-12 text-[#38BDF8]'/>
              <div className='flex flex-1 justify-center items-center sm:items-stretch sm:justify-start'>
                  <div className='hidden sm:ml-6 sm:block'>
                        <div className='flex space-x-7'>

                              <a href="#" className='font-bold text-l  px-2 duration-200 rounded-md '>test</a>
                              <a href="#" className='font-medium text-l hover:bg-[#14B8A6] duration-200 rounded-md text-white'>test</a>
                              <a href="#" className='font-medium text-l hover:bg-[#14B8A6] duration-200 rounded-md text-white'>test</a>
                             
                        </div>
                    </div>
              </div>
            </div>
              <div className='flesx items-center space-x-4'>
                    <IoMdNotifications className='w-9  h-16 text-emerald-500' onClick={notif}/>
              </div>
          </div>
      </div>
     </nav>
  )
}

export default Navbar;