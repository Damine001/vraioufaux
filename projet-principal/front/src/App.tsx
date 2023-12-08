import { createBrowserRouter, RouterProvider } from "react-router-dom";
import RegisterPage from "./pages/register";
import Home from "./pages/home";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Toaster } from "sonner";
import DialogContextProvider from "./contexts/dialiog-context";
import UserContextProvider from "./contexts/user-context";

function App() {
  const router = createBrowserRouter([
    {
      path: "/",
      element: <Home />,
    },
    {
      path: "/inscrire",
      element: <RegisterPage />,
    },
  ]);

  const queryClient = new QueryClient();

  return (
    <QueryClientProvider client={queryClient}>
      <DialogContextProvider>
        <UserContextProvider>
          <Toaster />
          <RouterProvider router={router} />
        </UserContextProvider>
      </DialogContextProvider>
    </QueryClientProvider>
  );
}

export default App;
