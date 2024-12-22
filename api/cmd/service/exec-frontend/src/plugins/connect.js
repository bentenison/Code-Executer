import axios from 'axios';


// Axios instance for Server 2
const AdminAPI = axios.create({
  baseURL: '/adminapi', // Base URL of Server 2
  headers: {
    'Content-Type': 'application/json',
  },
});

// // Axios instance for Server 3
// const server3API = axios.create({
//   baseURL: 'https://server3.com/api', // Base URL of Server 3
//   headers: {
//     'Content-Type': 'application/json',
//   },
// });

export { AdminAPI };