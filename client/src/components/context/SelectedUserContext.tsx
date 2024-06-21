import { type ReactNode, createContext, useState } from "react";
import type { User } from "../../utils/types";

type SelectedUserContextType = {
    selectedUser: User | null;
    setSelectedUser: (u: User | null) => void;
};
export const SelectedUserContext = createContext<SelectedUserContextType>({
    selectedUser: null,
    setSelectedUser: () => { },
});

interface ProviderProps {
    children: ReactNode;
}
export const SelectedUserProvider = ({ children }: ProviderProps) => {
    const [selectedUser, setSelectedUser] = useState<User | null>(null);

    return (
        <SelectedUserContext.Provider value={{ selectedUser, setSelectedUser }}>
            {children}
        </SelectedUserContext.Provider>
    );
};
