import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import './App.css';
import MainLayout from './components/layout/MainLayout';
import LoginView from './components/LoginSection/LoginView';
import Dashboard from './pages/Dashboard';
import Patients from './pages/Patients';
import { AuthProvider, useAuth } from './auth/AuthContext';
import ErrorBoundary from './components/ErrorBoundary';

import { ReactNode } from 'react';

function PrivateRoute({ children }: { children: ReactNode }) {

  const { isAuthenticated, loading } = useAuth();

  if (loading) {
    return <div className="w-full flex justify-center items-center min-h-[200px]">Cargando...</div>;
  }
  return isAuthenticated ? <>{children}</> : <Navigate to="/login" />;
}
function AppRoutes() {
  return (
    <Routes>
      <Route path="/login" element={<LoginView />} />
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
        <ErrorBoundary>
          <MainLayout>
            <AppRoutes />
          </MainLayout>
        </ErrorBoundary>
      </BrowserRouter>
    </AuthProvider>
  );
}

export default App;
