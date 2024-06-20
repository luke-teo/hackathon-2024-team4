import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import { Box } from '@mui/material'

function App() {

  return (
    <Box
      sx={{
        backgroundColor: "#ccc",
        display: "flex",
        flexDirection: "row",
        height: "100%",
        width: "100%"
      }}
    >
      <Box
        sx={{
          borderRight: "1px solid black",
          height: "100%",
          width: "500px"
        }}
      >
        Left
      </Box>
      <Box
        sx={{
          flex: 1,
        }}
      >
        Right
      </Box>
    </Box>
  )
}

export default App
