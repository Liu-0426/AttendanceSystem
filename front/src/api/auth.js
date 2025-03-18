import axios from 'axios';

export const login = (username, password) => {
  return axios.post('http://172.24.8.156:7777/login', { username, password });
};
