import { AppBar, Box, Toolbar, Typography } from "@mui/material";
import type { ReactNode } from "react";
import { colors } from "../../utils/colors";

interface Props {
    children: ReactNode;
}
export const Layout = ({ children }: Props) => {
    return (
        <Box
            sx={{
                height: "100%",
                width: "100%",
                display: "flex",
                flexDirection: "column",
            }}
        >
            <AppBar
                position="sticky"
                sx={{
                    backgroundColor: "white",
                    color: colors.TextBase,
                    boxShadow: "none",
                    borderBottom: 1,
                    borderColor: colors.BorderBase,
                    px: 2,
                    flex: 0,
                }}
            >
                <Toolbar disableGutters>
                    <Typography>First Move</Typography>
                </Toolbar>
            </AppBar>
            <Box
                component="main"
                sx={{
                    background:
                        "linear-gradient(283.89deg, #E3E8FE 32.07%, #C6A9FF 80.17%)",
                    flex: 1,
                    height: "100%",
                    display: "flex",
                    gap: 4,
                    alignItems: "center",
                    justifyContent: "space-evenly",
                    p: 2,
                }}
            >
                {children}
            </Box>
        </Box>
    );
};
