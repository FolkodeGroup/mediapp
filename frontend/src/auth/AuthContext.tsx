import { createContext, useState, useContext, useEffect, ReactNode } from "react";
import axios from "axios";
import { useMemo } from "react";
// Tipos para usuario y contexto
export interface User {
  id?: string;
  username?: string;
  email?: string;
  [key: string]: any;
}

interface AuthContextType {
  // isAuthenticated: boolean;
  user: User | null;
  token: string | null;
  isAuthenticated: boolean;
  login: (userData: User) => void;
  logout: () => void;
  loading: boolean;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {


  const [user, setUser] = useState<User | null>(null);
  const [token, setToken] = useState<string | null>(null);
  const [loading, setLoading] = useState(true);

    useEffect(() => {
      const storedUser = localStorage.getItem("user");
      const storedToken = localStorage.getItem("token");

      if (storedUser && storedToken) {
        setUser(JSON.parse(storedUser));
        setToken(JSON.parse(storedToken));
        setLoading(false);
      }

      setLoading(false);
    }, []);

  const login = async (userData: User) => {
    try {
      const response = await (axios.post('http://localhost:8080/login', userData));
      const {token: receivedToken, user: receivedUser} = response.data;

      setToken(receivedToken);
      // setIsAuthenticated(true);
      setUser(receivedUser);
      
      localStorage.setItem("user", JSON.stringify(receivedUser));
      localStorage.setItem("token", JSON.stringify(receivedToken));
      
    } catch (error) {
      console.error(error);
      throw error;
    } finally {
      setLoading(false);
    }
  };

  const logout = () => {
    // setIsAuthenticated(false);
    setToken(null);
    setUser(null);
    localStorage.removeItem("user");
    localStorage.removeItem("token");
  };

  const value = useMemo(() => ({
    token,
    user,
    isAuthenticated: !!token,
    loading,
    login,
    logout,
  }), [token, user, loading]);

  return (
    <AuthContext.Provider value={value}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth debe usarse dentro de un AuthProvider");
  }
  return context;
};
