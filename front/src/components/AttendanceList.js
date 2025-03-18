import React, { useEffect, useState } from 'react';
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.min.css';

const AttendanceList = () => {
  const [attendanceList, setAttendanceList] = useState([]);

  useEffect(() => {
    const fetchAttendanceList = async () => {
      try {
        const response = await axios.get('http://172.24.8.156:7777/api/clockins/1', { // Replace 1 with the actual user ID
          headers: { Authorization: `Bearer ${localStorage.getItem('token')}` },
        });
        setAttendanceList(response.data);
      } catch (err) {
        console.error('Error fetching attendance list', err);
      }
    };

    fetchAttendanceList();
    
  }, []);

  return (
    <div className="container mt-5">
      <h2>Your Attendance</h2>
      <table className="table table-striped">
        <thead>
          <tr>
            <th>Clock In</th>
            <th>Clock Out</th>
            <th>Date</th>
          </tr>
        </thead>
        <tbody>
          {attendanceList.map((attendance) => (
            <tr key={attendance.id}>
              <td>{attendance.clock_in ? attendance.clock_in : 'N/A'}</td>
              <td>{attendance.clock_out ? attendance.clock_out : 'N/A'}</td>
              <td>{attendance.date.split('T')[0]}</td>

            </tr>
          ))}
        </tbody>
      </table>

    </div>
  );
};

export default AttendanceList;