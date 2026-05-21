import { createContext, useContext } from "react";
import type { AuthState } from "@/features/auth/types";

export const AuthContext = createContext<AuthState | undefined>(undefined);

export function useAuth() {
  const context = useContext(AuthContext);

  if (context === undefined) {
    throw new Error("useAuth has to be used within <AuthContext.Provider>");
  }

  return context;
}
