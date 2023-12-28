'use client'

import { Box, Button, Container, TextField, Typography } from '@mui/material'
import { useContext } from 'react'
import { ResourceContext } from '../_context/ViewModelContex'
import Header from '../../../components/Header'
import Input from '@/components/Input'
import NameTag from '@/components/NameTag'

export default function Groupe() {
  const txt = useContext(ResourceContext)

  const questions = [
    '怒ったら怖そうなのは？',
    '友達多そうなのは',
    '将来金持ちになりそうなの',
  ]

  const members = ['utgw', 'ikegaki', 'あばたろ', 'げれげれ']
  return (
    <Container sx={{ display: 'flex', flexDirection: 'column', alignItems: 'center' }}>
      <Header />
      <Box
        sx={{
          display: 'flex',
          flexDirection: 'column',
          width: '100%',
        }}
      >
        <TextField
          required
          label="グループ名"
          sx={{ marginBottom: '8px', width: '100%' }}
        />
        <Input label="メンバー" onClick={() => {}} />
        <Box>
          {members.map((m, index) => (
            <NameTag key={index} name={m} />
          ))}
        </Box>
        <Input label="お題" onClick={() => {}} />
        <Box sx={{ margin: '0 40px 40px', color: '#777' }}>
          {questions.map((q, index) => (
            <Typography key={index}>{`${index + 1}. ${q}`}</Typography>
          ))}
        </Box>
        <Button variant="contained" sx={{ height: '48px' }}>
          <Typography sx={{ fontWeight: 'bold', fontSize: '20px' }}>
            グループを作成
          </Typography>
        </Button>
      </Box>
    </Container>
  )
}
