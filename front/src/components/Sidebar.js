import React from 'react';
import { Link } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';
import './Sidebar.css';

const Sidebar = () => {
  return (
    <aside className="sidebar">
      <h2>公司名稱</h2>
      <ul>
        <li>
          <Link to="/dashboard" className="sidebar-link" style={{ display: 'flex', alignItems: 'center' }}>
            <i className="fas fa-tachometer-alt" style={{ marginRight: '10px' }}></i>
            <span>Dashboard</span>
          </Link>
        </li>
        <li>
          <Link to="/clockinout" className="sidebar-link" style={{ display: 'flex', alignItems: 'center' }}>
            <i className="fas fa-clock" style={{ marginRight: '10px' }}></i>
            <span>Clock In / Out</span>
          </Link>
        </li>
        <li>
          <Link to="/settings" className="sidebar-link" style={{ display: 'flex', alignItems: 'center' }}>
            <i className="fas fa-cog" style={{ marginRight: '10px' }}></i>
            <span>Settings</span>
          </Link>
        </li>
      </ul>
    </aside>
  );
};

export default Sidebar;
