import { Box } from "@mui/material";
import type { ReactNode } from "react";

interface Props {
    children: ReactNode;
}
export const Layout = ({ children }: Props) => {
    return (
        <Box
            component="main"
            sx={{
                backgroundColor: "#ccc",
                display: "flex",
                flexDirection: "row",
                height: "100%",
                width: "100%",
            }}
        >
            {children}
        </Box>
    );
};
