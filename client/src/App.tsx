import { Box, CircularProgress } from "@mui/material";
import "./App.css";
import { Layout } from "./components/layout/Layout";
import { UserSideMenu } from "./components/widgets/UserSideMenu";
import { SelectedUserProvider } from "./components/context/SelectedUserContext";
import { ScoreChart } from "./components/widgets/ScoreChart";
import { BehaviorScore } from "./components/widgets/BehaviorScore";
import { Information } from "./components/widgets/Information";
import { useEffect, useState } from "react";

type UserStatusStatus = 'normal' | 'review' | 'required';
type UserStatus = {
  user: number;
  status: UserStatusStatus;
}

function App() {
  const [scoreData, setScoreData] = useState<any>(undefined);
  const [usersStatus, setUsersStatus] = useState<UserStatus[]>([]);

  useEffect(() => {
    fetch('/mock.json')
      .then((response) => response.json())
      .then((data: any) => {
        setScoreData(data)
        const userStatusData: UserStatus[] = [];
        
        for (let i = 0; i < data.length; i++) {
          const userId = data[i].userId;
          const latestScore = data[i].scores[data[i].scores.length - 1].zScore;
          
          let status: UserStatusStatus = 'normal';
          if (latestScore > 2) {
            status = 'review';
          }
          if (latestScore > 3) {
            status = 'required';
          }

          userStatusData.push({
            user: userId,
            status: status
          });
        }
        
        setUsersStatus(userStatusData)
      });
  }, [setScoreData])

  if (!scoreData && !usersStatus) {
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
