import React, { useState } from 'react';
import { 
  Box, 
  TextField, 
  Button, 
  Typography, 
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow
} from '@mui/material';
import axios from 'axios';

const PackCalculator = () => {
  const [items, setItems] = useState('');
  const [result, setResult] = useState(null);
  const [error, setError] = useState('');

  const calculatePacks = async () => {
    try {
      const response = await axios.get(`http://localhost:3001/packs/${items}`);
      setResult(response.data);
      setError('');
    } catch (err) {
      setError(err.response?.data?.error || 'An error occurred');
      setResult(null);
    }
  };

  return (
    <Box sx={{ maxWidth: 600, mx: 'auto', p: 3 }}>
      <Typography variant="h4" gutterBottom>
        Pack Calculator
      </Typography>
      
      <Box sx={{ mb: 3 }}>
        <TextField
          fullWidth
          label="Number of Items"
          type="number"
          value={items}
          onChange={(e) => setItems(e.target.value)}
          sx={{ mb: 2 }}
        />
        <Button 
          variant="contained" 
          onClick={calculatePacks}
          fullWidth
        >
          Calculate Packs
        </Button>
      </Box>

      {error && (
        <Typography color="error" sx={{ mb: 2 }}>
          {error}
        </Typography>
      )}

      {result && (
        <TableContainer component={Paper}>
          <Table>
            <TableHead>
              <TableRow>
                <TableCell>Pack Size</TableCell>
                <TableCell>Number of Packs</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {Object.entries(result).map(([size, count]) => (
                <TableRow key={size}>
                  <TableCell>{size}</TableCell>
                  <TableCell>{count}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      )}
    </Box>
  );
};

export default PackCalculator;