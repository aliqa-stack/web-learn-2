import React from 'react'
import Login from '../Login.jsx/LoginSection'
import { Link } from 'react-router-dom'
const Information = () => {
  return (
    <div className='bg-[#0F172A] py-6 mt-4'>
        <div className='text-center bg-[#14B8A6] mx-auto w-11/12 sm:w-1/2 py-4 rounded-md'>
            <h1 className='font-pixelify text-white'>Lorem ipsum dolor sit amet.</h1>
        </div>

        <div>
            <p className='text-white text-center px-4 mt-10 '>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam auctor, nisl eget ultricies tincidunt, nunc nisl aliquam nisl, eget ultricies nunc nisl eget nunc.
            </p>
        </div>

        <div className='flex mt-4 px-5 flex-col gap-3 sm:ml-12'>
            <p className='text-white font-pixelify font-light '>I Carry, We Win</p>
            <p className='text-white font-pixelify font-light '>God Mode: ON</p>
            <p className='text-white font-pixelify font-light '>This? Absolute Meta!</p>
            <p className='text-white font-pixelify font-light '>
Honkai-DMG Demo</p>
        </div>

        <div className='flex justify-center mt-6'>
            <Link to="/login" className='bg-[#14B8A6] hover:bg-[#14B8A6]/80 text-white font-pixelify py-2 px-4 sm:px-8 sm:py-4 rounded-md'>
                Go to Login
            </Link>
        </div>
    </div>
  )
}

export default Information;