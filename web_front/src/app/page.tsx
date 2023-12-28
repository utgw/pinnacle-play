'use client'

import { Button, Container, Typography } from '@mui/material'
import { useContext } from 'react'
import { ResourceContext } from './(user)/_context/ViewModelContex'
import Header from '../components/Header'
import { useRouter } from 'next/navigation'

export default function Home() {
  const txt = useContext(ResourceContext)
  const router = useRouter()

  return (
    <Container sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
      <Header />
      <Typography fontSize={'24px'} fontWeight={'bold'} margin={'40px 0'}>
        この中で1番〇〇な人は誰？
      </Typography>
      <Typography color="#777" margin={'20px 20px 40px'} lineHeight={2}>
        「1番早く結婚しそうなのは誰?」
        <br />
        「好きな人にデレデレしそうなのは？」
        <br />
        グループ内での自分の立ち位置を知りたくなった時はありませんか？
        <br />
        各質問に対して、グループ内で最も適したメンバーを選んで投票。集計されるランキングで、自分のポジションや友達からの印象が一目でわかるサービスです。
      </Typography>
      <Button
        variant="contained"
        sx={{
          width: 'calc(100% - 40px)',
          height: '60px',
          fontSize: '20px',
          fontWeight: 'bold',
          marginBottom: '40px',
        }}
        onClick={() => router.push('/group')}
      >
        はじめる
      </Button>
    </Container>
  )
}
