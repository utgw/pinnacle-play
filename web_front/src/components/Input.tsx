import { Box, Button, TextField } from '@mui/material'
import React from 'react'

export default function Input({
  label,
  onClick,
}: {
  label: string
  onClick: () => void
}) {
  return (
    <Box sx={{ display: 'flex', margin: '20px 0' }}>
      <TextField
        required
        label={label}
        sx={{ width: 'calc(100% - 80px)', borderRadius: '4px 0 0 4px' }}
      />
      <Button
        variant="contained"
        sx={{
          width: '80px',
          borderRadius: '0px 4px 4px 0',
          boxShadow: 'none',
          fontWeight: 'bold',
        }}
        onClick={onClick}
      >
        追加
      </Button>
    </Box>
  )
}
