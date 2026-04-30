import React from 'react'
import ff from '../assets/ff.jpg'

const Header = () => {
  return (
    <div className='bg-[#0F172A]  mx-auto px-5 py-20 h-auto'>
      <div className='flex flex-col sm:flex-row items-start sm:items-center justify-between gap-2'>

        <h1 className='text-5xl sm:text-8xl font-medium text-center mt-10 text-white font-pixelify'>
          Header
        </h1>
          <div class="w-32 h-32  sm:w-64 sm:h-64 sm:ml-auto bg-gradient-to-r from-blue-500 to-cyan-400 rounded-full ">
                 <img src="" alt="" className='object-cover w-full h-full rounded-full' />
            </div>
        <div className='w-full sm:w-32 bg-[#14B8A6] py-2 px-4 sm:py-1.5 sm:px-2 rounded-md'>
          <p className='text-base font-bold font-pixelify text-white'>
            Lorem ipsum dolor sit amet.
          </p>
        </div>

      </div>

      <p className='font-medium font-pixelify text-white mt-auto text-sm sm:text-base'>
        Lorem ipsum dolor sit amet consectetur adipisicing elit. Porro.
      </p>
    </div>
  )
}

export default Header