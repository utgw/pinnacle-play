import { Box, Button, Container, Typography } from "@mui/material";

export default function Home() {
  return (
    <Container
      sx={{ display: "flex", flexDirection: "column", alignItems: "center" }}
    >
      <Box
        sx={{
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          marginBottom: "20px",
        }}
      >
        <Typography variant="h2">Pinnacle Play</Typography>
      </Box>
      <Box
        sx={{
          display: "flex",
          justifyContent: "space-between",
          width: "600px",
          alignItems: "center",
        }}
      >
        <Box>ユーザー名</Box>
        <Box>グループ作成</Box>
        <Box>...</Box>
      </Box>
    </Container>
  );
}
