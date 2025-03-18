import axios from 'axios';

export const clockIn = (userID) => {
  return axios.post('http://172.24.8.156:7777/api/clockin', { user_id: userID });
};

export const clockOut = (userID) => {
  return axios.post('http://172.24.8.156:7777/api/clockout', { user_id: userID });
};
