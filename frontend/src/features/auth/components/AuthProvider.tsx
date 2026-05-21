import { useEffect, useState } from "react";

import { api } from "@/api";
import { AuthContext } from "@/features/auth/contexts/authContext";

import type { User } from "@/features/auth/types";

export default function AuthProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [user, setUser] = useState<User | null>(null);
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    api
      .get("/whoami")
      .then((response) => {
        setUser(response.data);
        setIsAuthenticated(true);
      })
      .catch(() => {})
      .finally(() => setIsLoading(false));
  }, []);

  useEffect(() => {
    const unauthorizedInterceptor = api.interceptors.response.use(
      (response) => response,
      (error) => {
        if (error.response?.status === 401) {
          setUser(null);
          setIsAuthenticated(false);
        }

        return Promise.reject(error);
      },
    );

    return () => api.interceptors.response.eject(unauthorizedInterceptor);
  }, []);

  async function login(identifier: string, password: string) {
    const response = await api.post("/login", { identifier, password });
    setUser(response.data);
    setIsAuthenticated(true);
  }

  async function logout() {
    try {
      await api.post("/logout");
      setUser(null);
      setIsAuthenticated(false);
    } catch {
      // error has already been handled by interceptor so we ignore it
    }
  }

  return (
    <AuthContext.Provider value={{ isAuthenticated, user, login, logout }}>
      {isLoading ? <div>Loading...</div> : children}
    </AuthContext.Provider>
  );
}
