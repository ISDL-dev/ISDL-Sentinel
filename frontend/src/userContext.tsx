import React, { createContext, useContext, useState, useEffect } from "react";

export interface AuthUser {
  user_id: number;
  user_name: string;
  status: string;
  role_list: string[];
  avatar_id: number;
  avatar_img_path: string;
}

interface UserContextType {
  authUser: AuthUser | undefined;
  setAuthUser: React.Dispatch<React.SetStateAction<AuthUser | undefined>>;
}

const UserContext = createContext<UserContextType | undefined>(undefined);

export const UserProvider: React.FC<{ children: React.ReactNode }> = ({
  children,
}) => {
  const [authUser, setAuthUser] = useState<AuthUser | undefined>(() => {
    const savedUser = localStorage.getItem("auth_user");
    return savedUser ? JSON.parse(savedUser) : undefined;
  });

  useEffect(() => {
    if (authUser) {
      localStorage.setItem("auth_user", JSON.stringify(authUser));
    } else {
      localStorage.removeItem("auth_user");
    }
  }, [authUser]);

  return (
    <UserContext.Provider value={{ authUser, setAuthUser }}>
      {children}
    </UserContext.Provider>
  );
};

export const useUser = () => {
  const context = useContext(UserContext);
  if (context === undefined) {
    throw new Error("useUser must be used within a UserProvider");
  }
  return context;
};
