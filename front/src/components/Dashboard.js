import React, { useEffect, useState } from 'react';
import axios from 'axios';
import Sidebar from './Sidebar';
import 'bootstrap/dist/css/bootstrap.min.css';
import './Dashboard.css';

const Dashboard = () => {
  const [attendanceList, setAttendanceList] = useState([]);
  const [totalEmployees, setTotalEmployees] = useState(0);
  const [totalClockCount, setTotalClockCount] = useState(0);

  useEffect(() => {
    const fetchAttendanceList = async () => {
      try {
        const response = await axios.get('http://172.24.8.156:7777/api/clocklist', {
          headers: { Authorization: `Bearer ${localStorage.getItem('token')}` },
        });
        setAttendanceList(response.data);
      } catch (err) {
        console.error('Error fetching attendance list', err);
      }
    };

    const fetchTotalEmployees = async () => {
      try {
        const response = await axios.get('http://172.24.8.156:7777/api/users', {
          headers: { Authorization: `Bearer ${localStorage.getItem('token')}` },
        });
        setTotalEmployees(response.data.length);
      } catch (err) {
        console.error('Error fetching total employees', err);
      }
    };
    const fetchTotalClockCount = async () => {
      try {
        const response = await axios.get('http://172.24.8.156:7777/api/todayClockin', {
          headers: { Authorization: `Bearer ${localStorage.getItem('token')}` },
        });
        setTotalClockCount(response.data.count);
      } catch (err) {
        console.error('Error fetching total clock count', err);
      }
    };

    fetchAttendanceList();
    fetchTotalEmployees();
    fetchTotalClockCount();
  }, []);

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleString();
  };

  return (
    <div className="dashboard-container">
      <Sidebar />
      <main className="main-content">
        <header>
          <h1>Dashboard</h1>
          <div className="user-profile">
            <span>👤 User</span>
          </div>
        </header>
        <div className="summary-cards">
          <div className="card">
            <h3>Total Employees</h3>
            <p>{totalEmployees}</p>
          </div>
          <div className="card">
            <h3>Today’s Check-ins</h3>
            <p>{totalClockCount}</p>
          </div>
        </div>
        <div className="table-container">
          <h2>Employee Attendance</h2>
          <table className="styled-table">
            <thead>
              <tr>
                <th>User ID</th>
                <th>Clock In</th>
                <th>Clock Out</th>
                <th>Date</th>
              </tr>
            </thead>
            <tbody>
              {attendanceList.map((attendance) => (
                <tr key={attendance.ID}>
                  <td>{attendance.UserID}</td>
                  <td>{attendance.ClockIn.Valid ? formatDate(attendance.ClockIn.Time) : 'N/A'}</td>
                  <td>{attendance.ClockOut.Valid ? formatDate(attendance.ClockOut.Time) : 'N/A'}</td>
                  <td>{formatDate(attendance.Date)}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </main>
    </div>
  );
};

export default Dashboard;