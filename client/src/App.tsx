import { Box } from "@mui/material";
import "./App.css";
import { Layout } from "./components/layout/Layout";
import { UserSideMenu } from "./components/widgets/UserSideMenu";
import { SelectedUserProvider } from "./components/context/SelectedUserContext";
import { ScoreChart } from "./components/widgets/ScoreChart";
import { BehaviorScore } from "./components/widgets/BehaviorScore";
import { FollowUp } from "./components/widgets/FollowUp";

function App() {
  return (
    <SelectedUserProvider>
      <Layout>
        <UserSideMenu />
        <ScoreChart />
        <Box sx={{
          display: "flex",
          flex: 0,
          flexDirection: "column",
          height: "100%",
          gap: 2,
          minWidth: 360
        }}>
          <BehaviorScore />
          <FollowUp />
        </Box>
      </Layout>
    </SelectedUserProvider>
  );
}

export default App;
