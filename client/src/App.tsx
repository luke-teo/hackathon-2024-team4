import { Box, CircularProgress } from "@mui/material";
import "./App.css";
import { Layout } from "./components/layout/Layout";
import { UserSideMenu } from "./components/widgets/UserSideMenu";
import { SelectedUserProvider } from "./components/context/SelectedUserContext";
import { ScoreChart } from "./components/widgets/ScoreChart";
import { BehaviorScore } from "./components/widgets/BehaviorScore";
import { Information } from "./components/widgets/Information";
import { useEffect, useState } from "react";

function App() {
  const [scoreData, setScoreData] = useState<any>(undefined);

  useEffect(() => {
    fetch('/mock.json')
      .then((response) => response.json())
      .then((data) => {
        setScoreData(data)
      });
  }, [setScoreData])

  if (!scoreData) {
    return (<Box
      sx={{
        alignItems: "center",
        display: "flex",
        height: "100%",
        justifyContent: "center",
        width: "100%"
      }}
    >
      <CircularProgress />
    </Box>);
  }

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
          <Information />
        </Box>
      </Layout>
    </SelectedUserProvider>
  );
}

export default App;
