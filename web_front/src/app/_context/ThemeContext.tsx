'use client'
import { ThemeProvider, createTheme, CssBaseline } from '@mui/material'
import React from 'react'

export default function ThemeContext({ children }: { children: React.ReactNode }) {
  const theme = createTheme({
    palette: {
      primary: {
        main: '#644AF6',
      },
      info: {
        main: '#6D6D6D',
      },
    },
  })

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      {children}
    </ThemeProvider>
  )
}
