import React from 'react'

const Header = () => {
  return (
    <div className='relative overflow-hidden bg-radial from-[#0F172A] via-[#15203b] to-[#0a0f1d] mx-auto px-6 py-24 min-h-[500px] flex flex-col justify-center'>
      
      {/* Dekorasi Latar Belakang (Ambient Light) */}
      <div className="absolute top-0 right-0 -translate-y-1/4 translate-x-1/4 w-96 h-96 bg-emerald-500/10 blur-[120px] rounded-full" />
      <div className="absolute bottom-0 left-0 translate-y-1/4 -translate-x-1/4 w-80 h-80 bg-purple-500/10 blur-[100px] rounded-full" />

      <div className='relative z-10 flex flex-col md:flex-row items-center justify-between gap-12 max-w-7xl mx-auto w-full'>
        
        {/* Konten Teks */}
        <div className='flex flex-col items-center md:items-start text-center md:text-left space-y-2'>
          <h1 className='text-6xl sm:text-9xl font-bold tracking-tighter text-transparent bg-clip-text bg-gradient-to-b from-emerald-100 to-emerald-400 drop-shadow-[0_0_15px_rgba(110,231,183,0.3)]'>
            ALIK AL MALIKI
          </h1>
          
          <h2 className='text-5xl sm:text-8xl font-medium text-transparent bg-clip-text bg-gradient-to-r from-purple-400 via-sky-400 to-emerald-400 font-pixelify leading-none'>
            React
          </h2>

          {/* Badge Tag */}
          <div className='inline-flex items-center gap-2 mt-6 bg-white/5 backdrop-blur-xl border border-white/10 rounded-full py-2 px-6 shadow-2xl shadow-purple-500/20'>
            <div className='w-2 h-2 rounded-full bg-emerald-400 animate-pulse' />
            <p className='text-sm sm:text-lg font-bold font-pixelify text-emerald-100 tracking-widest uppercase'>
              LOCATION : SUMEDANG
            </p>
          </div>
        </div>
        
        {/* Avatar/Visual Container */}
        <div className="relative group">
          {/* Ring Animasi */}
          <div className="absolute -inset-1 bg-gradient-to-r from-blue-500 to-emerald-400 rounded-full blur opacity-40 group-hover:opacity-75 transition duration-1000 group-hover:duration-200 animate-tilt"></div>
          
          <div className="relative w-48 h-48 sm:w-72 sm:h-72 bg-slate-900 rounded-full border-2 border-white/10 overflow-hidden flex items-center justify-center shadow-2xl">
            {/* Placeholder untuk Image ff.jpg */}
            <div className="w-full h-full bg-gradient-to-tr from-blue-600/20 to-emerald-400/20 flex items-center justify-center">
              <span className="text-white/20 font-pixelify">IMG</span>
            </div>
          </div>
        </div>

      </div>

      {/* Deskripsi Bawah */}
      <div className='max-w-7xl mx-auto w-full mt-16 relative z-10'>
        <div className='h-[1px] w-full bg-gradient-to-r from-emerald-500/50 via-transparent to-transparent mb-4' />
        <p className='font-medium font-pixelify text-emerald-100/70 max-w-lg leading-relaxed text-sm sm:text-lg'>
          // STATUS_REPORT: Lorem ipsum dolor sit amet consectetur adipisicing elit. 
          Expedita porro laboriosam hic veniam.
        </p>
      </div>
    </div>
  )
}

export default Header