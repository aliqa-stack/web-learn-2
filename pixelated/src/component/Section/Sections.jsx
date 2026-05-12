import React from 'react'

const Sections = () => {
  return (
    <div className='  bg-[#F0F9FF] py-16'>

        <div className='max-w-7xl mx-auto px-6'>
            <div className='flex items-center gap-4 mb-10 '>
                <div className='h-[5px] flex-grow bg-gradient-to-r from-transparent to-blue-600/20'/>
                    <h1 className='font-pixelify font-medium text-2xl text-blue-600 tracking-widest uppercase'>test</h1>
            
                <div className='h-[2px] flex-grow bg-gradient-to-l from-transparent to-blue-500/20' />
            </div>

        <div className='grid grid-cols-1  md:grid-rows-2 lg:grid-cols-6 auto-rows-[180px] gap-6 '>

       <div className='lg:col-span-3 lg:row-span-2 group cursor-pointer'>
            <div className='relative h-full border border-blue-200 bg-white hover:bg-blue-600 transition-all duration-500 rounded-3xl p-8 overflow-hidden shadow-sm hover:shadow-2xl hover:shadow-blue-200'>
              <div className='absolute top-0 right-0 p-4 opacity-10 group-hover:opacity-100 transition-opacity'>
                 <div className='w-20 h-20 bg-blue-400 rounded-full blur-3xl' />
              </div>
              <h2 className='font-pixelify text-3xl text-blue-600 group-hover:text-white transition-colors duration-300'>01. Missions</h2>
              <p className='mt-4 text-slate-500 group-hover:text-blue-100 max-w-[200px] transition-colors'>Explore the latest objectives and tactical rewards.</p>
            </div>
          </div>

          <div className='lg:col-span-2 lg:row-span-1 group cursor-pointer'>
            <div className='h-full border border-cyan-200 bg-cyan-500 rounded-3xl p-6 flex flex-col justify-end hover:brightness-110 transition-all'>
               <h2 className='font-pixelify text-2xl text-white'>03. Database</h2>
            </div>
          </div>

            <div className='lg:col-span-2 lg:rowspan-2 group cursor-pointer'>
                <div className='h-full border border-cyan-200 bg-cyan-800 rounded-3xl p-6 flex flex-col justify-end  '>
                    <h2 className='font-pixelify text-2xl text-white'>Sections</h2>
                </div>
            </div>

         <div className='lg:col-span-1 lg:row-span-1 group cursor-pointer'>
            <div className='h-full border border-blue-100 bg-white rounded-3xl p-6 flex items-center justify-center hover:bg-slate-50 transition-colors'>
               <span className='font-pixelify text-blue-600 text-4xl'>+</span>
            </div>
          </div>
          
         <div className='lg:col-span-1 group'>
  <div className='h-full w-full border-purple-300 bg-cyan-300 rounded-3xl flex items-center justify-center'>
    <h2 className='font-pixelify'>Photo</h2>
  </div>
</div>

<div className='lg:col-span-4 group'>
  <div className='h-full w-full border border-purple-300 bg-white hover:bg-blue-500 transition-all duration-300 rounded-3xl flex  justify-center'>
    <h2 className='font-pixelify text-blue-700 mt-5 group-hover:text-white transition-all duration-200'>Lorem ipsum dolor sit amet.</h2>
  </div>
</div>
     

        </div>

        </div>
    </div>
  )
}

export default Sections
