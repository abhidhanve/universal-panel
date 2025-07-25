import axios from 'axios';

// Use environment variable or fallback to production for testing
const API_BASE_URL = process.env.REACT_APP_API_URL || 'https://universal-panel-production-51f6.up.railway.app/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor to add auth token
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('auth_token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Response interceptor to handle auth errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('auth_token');
      localStorage.removeItem('user');
      // Don't force redirect here - let the auth context handle it
    }
    return Promise.reject(error);
  }
);

export default api;
