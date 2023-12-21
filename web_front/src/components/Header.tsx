import Image from 'next/image'
import React from 'react'

export default function Header() {
  return (
    <div style={{ width: '100%', marginBottom: '40px' }}>
      <Image src={'/logo.png'} alt="logo" layout="responsive" width={1616} height={520} />
    </div>
  )
}
