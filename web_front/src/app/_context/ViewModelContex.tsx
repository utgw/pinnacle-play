'use client'

import { createContext } from 'react'

export const ResourceContext = createContext({})
export default function ViewModelContext({ children }: { children: React.ReactNode }) {
  return <ResourceContext.Provider value={'test'}> {children}</ResourceContext.Provider>
}
