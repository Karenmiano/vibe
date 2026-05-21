export type User = {
  username: string;
  fullName: string;
};

export type AuthState = {
  isAuthenticated: boolean;
  user: User | null;
  login: (identifier: string, password: string) => Promise<void>;
  logout: () => void;
};
