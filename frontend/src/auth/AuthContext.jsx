import { createContext, useState, useContext, useEffect } from "react";

// Crear el contexto de autenticación
const AuthContext = createContext();

// Proveedor de contexto de autenticación
export const AuthProvider = ({ children }) => {
    const [isAuthenticated, setIsAuthenticated] = useState(false);
    const [user, setUser] = useState(null);

    // Login: guarda usuario en localStorage
    const login = (userData) => {
        setIsAuthenticated(true);
        setUser(userData);
        localStorage.setItem("user", JSON.stringify(userData));
    };

    // Logout: limpia usuario
    const logout = () => {
        setIsAuthenticated(false);
        setUser(null);
        localStorage.removeItem("user");
    };

    // Mantener sesión si hay usuario guardado
    useEffect(() => {
        const stored = localStorage.getItem("user");
        if (stored) {
            setUser(JSON.parse(stored));
            setIsAuthenticated(true);
        }
    }, []);

    return (
        <AuthContext.Provider value={{ isAuthenticated, user, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
};

// Hook para usar el contexto de autenticación
export const useAuth = () => {
  const context = useContext(AuthContext);
  
  if (!context) {
    throw new Error("useAuth debe usarse dentro de un AuthProvider");
  }
  return context;
};