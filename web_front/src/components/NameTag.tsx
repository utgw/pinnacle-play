import { Button } from '@mui/material'
import React from 'react'
import ClearIcon from '@mui/icons-material/Clear'

export default function NameTag({ name }: { name: string }) {
  return (
    <Button
      variant="outlined"
      color="info"
      sx={{ margin: '4px', borderRadius: '20px' }}
      startIcon={<ClearIcon />}
    >
      {name}
    </Button>
  )
}
