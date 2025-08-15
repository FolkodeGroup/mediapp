import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import './App.css';
import MainLayout from './components/layout/MainLayout';
import Login from './components/auth/Login';
import Dashboard from './pages/Dashboard';
import Patients from './pages/Patients';
import { AuthProvider, useAuth } from './auth/AuthContext';

import { ReactNode } from 'react';
function PrivateRoute({ children }: { children: ReactNode }) {
  const { isAuthenticated } = useAuth();
  return isAuthenticated ? <>{children}</> : <Navigate to="/login" />;
}

function AppRoutes() {
  return (
    <Routes>
      <Route path="/login" element={<Login />} />
      <Route
        path="/dashboard"
        element={
          <PrivateRoute>
            <Dashboard />
          </PrivateRoute>
        }
      />
      <Route
        path="/patients"
        element={
          <PrivateRoute>
            <Patients />
          </PrivateRoute>
        }
      />
      <Route path="*" element={<Navigate to="/login" />} />
    </Routes>
  );
}


function App() {
  return (
    <AuthProvider>
      <BrowserRouter>
        <MainLayout>
          <AppRoutes />
        </MainLayout>
      </BrowserRouter>
    </AuthProvider>
  );
}

export default App;
