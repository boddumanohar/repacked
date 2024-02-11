import React, { useState } from 'react';
import './App.css';
import axios from 'axios';

function App() {
  const [packSizes, setPackSizes] = useState('');
  const [orderSize, setOrderSize] = useState('');
  const [packs, setPacks] = useState([]);
  const [error, setError] = useState('');

  const submitPackSizes = async () => {
    try {
      const response = await axios.post('http://localhost:8080/pack', { packSizes: packSizes.split(',').map(Number) });
      alert('Pack sizes updated successfully!');
    } catch (err) {
      setError(err.message);
    }
  };

  const calculatePacks = async () => {
    try {
      const response = await axios.get(`http://localhost:8080/pack?orderSize=${orderSize}`);
      setPacks(response.data.packs.map((pack, index) => `${pack} X ${response.data.packSizes[index]}`));
    } catch (err) {
      setError(err.message);
    }
  };

  return (
    <div className="App">
      <h1>Order Packs Calculator</h1>
      {error && <p className="error">{error}</p>}
      <div>
        <input
          type="text"
          value={packSizes}
          onChange={(e) => setPackSizes(e.target.value)}
          placeholder="Enter pack sizes (comma-separated)"
        />
        <button onClick={submitPackSizes}>Submit Pack Size Changes</button>
      </div>
      <div>
        <input
          type="number"
          value={orderSize}
          onChange={(e) => setOrderSize(e.target.value)}
          placeholder="Enter order size"
        />
        <button onClick={calculatePacks}>Calculate</button>
      </div>
      <div>
        <h2>Packs:</h2>
        {packs.map((pack, index) => (
          <p key={index}>{pack}</p>
        ))}
      </div>
    </div>
  );
}

export default App;
