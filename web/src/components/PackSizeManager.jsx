import React, { useState } from 'react';
import { 
  Box, 
  TextField, 
  Button, 
  Typography, 
  Alert 
} from '@mui/material';
import axios from 'axios';

const PackSizeManager = () => {
  const [sizes, setSizes] = useState('');
  const [message, setMessage] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async () => {
    try {
      const sizeArray = sizes.split(',').map(size => parseInt(size.trim()));
      await axios.post('http://localhost:3001/packs', sizeArray);
      setMessage('Pack sizes stored successfully');
      setError('');
      setSizes('');
    } catch (err) {
      setError(err.response?.data?.error || 'An error occurred');
      setMessage('');
    }
  };

  return (
    <Box sx={{ maxWidth: 600, mx: 'auto', p: 3 }}>
      <Typography variant="h4" gutterBottom>
        Manage Pack Sizes
      </Typography>
      
      <Box sx={{ mb: 3 }}>
        <TextField
          fullWidth
          label="Pack Sizes (comma-separated)"
          placeholder="e.g., 250,500,1000,2000,5000"
          value={sizes}
          onChange={(e) => setSizes(e.target.value)}
          sx={{ mb: 2 }}
        />
        <Button 
          variant="contained" 
          onClick={handleSubmit}
          fullWidth
        >
          Store Pack Sizes
        </Button>
      </Box>

      {message && (
        <Alert severity="success" sx={{ mb: 2 }}>
          {message}
        </Alert>
      )}

      {error && (
        <Alert severity="error" sx={{ mb: 2 }}>
          {error}
        </Alert>
      )}
    </Box>
  );
};

export default PackSizeManager;