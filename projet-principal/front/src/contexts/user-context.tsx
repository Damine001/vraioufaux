import { createContext, useState } from "react";
import { useLocalStorage } from "@uidotdev/usehooks";
import useUserMe from "@/hooks/user-get-user-me";

export type User = {
  id: number;
  name: string;
  family_name: string;
  email: string;
  password: string;
  createdAt: string;
  updatedAt: string;
};

interface UserContextProps {
  user: User | null;
  setUser: (user: User | null) => void;
}

export const UserContext = createContext<UserContextProps>({
  user: null,
  setUser: () => {},
});

interface userContextProviderProps {
  children: React.ReactNode;
}

export default function UserContextProvider({
  children,
}: userContextProviderProps) {
  const [userMe, setUserMe] = useState<User | null>(null);
  const [accessToken, setAccessToken] = useLocalStorage<string | null>(
    "accessToken"
  );

  const { data } = useUserMe(accessToken);

  if (data && !userMe) {
    setUserMe(data);
  }

  return (
    <UserContext.Provider
      value={{
        user: userMe,
        setUser: setUserMe,
      }}
    >
      {children}
    </UserContext.Provider>
  );
}
