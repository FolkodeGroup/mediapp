import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import './App.css';
import MainLayout from './components/layout/MainLayout';
import Login from './components/auth/Login';
import Dashboard from './pages/Dashboard'
import Patients from './pages/Patients'
function App() {
const isAuthenticated = false 

  return (
    <MainLayout>
      <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route
          path="/dashboard"
          element={isAuthenticated ? <Dashboard /> : <Navigate to="/login" />}
        />
        <Route
          path="/patients"
          element={isAuthenticated ? <Patients /> : <Navigate to="/login" />}
        />
        <Route path="*" element={<Navigate to="/login" />} />
      </Routes>
    </BrowserRouter>
    </MainLayout>
  );
}

export default App;
