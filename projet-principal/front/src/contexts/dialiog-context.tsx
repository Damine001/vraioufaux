import { createContext, useState } from "react";

interface DialogContextProps {
  loginDialogOpen: boolean;
  toggleLoginDialogOpen: () => void;
}

export const DialogContext = createContext<DialogContextProps>({
  loginDialogOpen: false,
  toggleLoginDialogOpen: () => {},
});

interface DialogProviderProps {
  children: React.ReactNode;
}

export default function DialogProvider({ children }: DialogProviderProps) {
  const [loginDialogOpen, setLoginDialogOpen] = useState(false);

  function toggleLoginDialogOpen() {
    setLoginDialogOpen((prev) => !prev);
  }

  return (
    <DialogContext.Provider
      value={{
        loginDialogOpen,
        toggleLoginDialogOpen,
      }}
    >
      {children}
    </DialogContext.Provider>
  );
}
