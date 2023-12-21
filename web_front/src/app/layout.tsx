import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'
import ViewModelContext from './_context/ViewModelContex'
import ThemeContext from './_context/ThemeContext'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: 'Pinnacle Play',
  description: '',
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <ViewModelContext>
          <ThemeContext>{children}</ThemeContext>
        </ViewModelContext>
      </body>
    </html>
  )
}
