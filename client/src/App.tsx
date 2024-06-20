import { Box } from "@mui/material";
import "./App.css";
import { Layout } from "./components/layout/Layout";
import { UserSideMenu } from "./components/widgets/UserSideMenu";
import { SelectedUserProvider } from "./components/context/SelectedUserContext";

function App() {
  return (
    <SelectedUserProvider>
      <Layout>
        <UserSideMenu />
        <Box sx={{ border: 1, height: "100%", flex: 1 }}>2 here</Box>
        <Box sx={{ border: 1, height: "100%", flex: 0, minWidth: 360 }}>
          3 here
        </Box>
      </Layout>
    </SelectedUserProvider>
  );
}

export default App;
