import { RouterProvider } from "@tanstack/react-router";

import AuthProvider from "@/features/auth/components/AuthProvider";
import { useAuth } from "./features/auth/contexts/authContext";
import { router } from "./router";
import { Toaster } from "react-hot-toast";
import { useEffect } from "react";

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

function InnerApp() {
  const auth = useAuth();

  useEffect(() => {
    // Forces routes to reload when the auth state changes
    router.invalidate();
  }, [auth]);

  return <RouterProvider router={router} context={{ auth }} />;
}

function App() {
  return (
    <>
      <AuthProvider>
        <InnerApp />
      </AuthProvider>
      <Toaster />
    </>
  );
}

export default App;
