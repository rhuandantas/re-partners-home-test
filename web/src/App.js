import React from 'react';
import { Container, Box, Typography, Divider } from '@mui/material';
import PackCalculator from './components/PackCalculator';
import PackSizeManager from './components/PackSizeManager';

function App() {
  return (
    <Container>
      <Box sx={{ my: 4 }}>
        <Typography variant="h3" component="h1" gutterBottom align="center">
          Pack Management System
        </Typography>
        <PackSizeManager />
        <Divider sx={{ my: 4 }} />
        <PackCalculator />
      </Box>
    </Container>
  );
}

export default App;