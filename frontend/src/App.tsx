import { RouterProvider } from "@tanstack/react-router";

import { router } from "./router";
import { Toaster } from "react-hot-toast";

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

function App() {
  return (
    <>
      <RouterProvider router={router} />
      <Toaster />
    </>
  );
}

export default App;
