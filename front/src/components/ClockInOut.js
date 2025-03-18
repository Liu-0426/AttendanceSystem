import React, { useState } from 'react';
import axios from 'axios';
import { jwtDecode } from 'jwt-decode'; // 使用命名導入
import Sidebar from './Sidebar';
import 'bootstrap/dist/css/bootstrap.min.css';
import './ClockInOut.css';

const ClockInOut = () => {
  const [message, setMessage] = useState('');

  const getUserID = () => {
    const token = localStorage.getItem('token');
    if (token) {
      const decoded = jwtDecode(token);
      return decoded.userID;
    }
    return null;
  };

  const handleClockIn = async () => {
    const userID = getUserID();
    if (!userID) {
      setMessage('User ID not found');
      return;
    }

    try {
      await axios.post('http://172.24.8.156:7777/api/clockin', { user_id: userID }, { 
        headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
      });
      setMessage('Clock-in successful');
    } catch (err) {
      setMessage('Error clocking in');
    }
  };

  const handleClockOut = async () => {
    const userID = getUserID();
    if (!userID) {
      setMessage('User ID not found');
      return;
    }

    try {
      await axios.post('http://172.24.8.156:7777/api/clockout', { user_id: userID }, {
        headers: { Authorization: `Bearer ${localStorage.getItem('token')}` }
      });
      setMessage('Clock-out successful');
    } catch (err) {
      setMessage('Error clocking out');
    }
  };

  return (
    <div className="dashboard-container">
      <Sidebar />
      <main className="main-content">
        <div className="container mt-5">
          <h2>Clock In / Clock Out</h2>
          <button className="btn btn-success mr-2" onClick={handleClockIn}>Clock In</button>
          <button className="btn btn-danger" onClick={handleClockOut}>Clock Out</button>
          {message && <p className="mt-3">{message}</p>}
        </div>
      </main>
    </div>
  );
};

export default ClockInOut;