import { TrendingUp } from "@mui/icons-material";
import { Box, CircularProgress, Typography } from "@mui/material";
import { colors } from "../../utils/colors";
import React, { useContext, useEffect } from "react";
import { mock } from "../../mock/mock";
import { SelectedUserContext } from "../context/SelectedUserContext";

export const BehaviorScore = (): JSX.Element => {
  const { selectedUser } = useContext(SelectedUserContext);
  const [userZScore, setUserZScore] = React.useState<number | undefined>(undefined);

  useEffect(() => {
    if (selectedUser === undefined) {
      setUserZScore(undefined);
    }

    const userScoreMock = mock;

    for (let i = 0; i < userScoreMock.length; i++) {
      if (Number(userScoreMock[i].userId) === selectedUser?.id) {

        setUserZScore(userScoreMock[i].scores[userScoreMock[i].scores.length - 1].zScore);
      }
    }
  }, [selectedUser, setUserZScore]);

  if (selectedUser === null) {
    return (
      <Box
        sx={{
          alignItems: "center",
          background: colors.BackgroundBaseWhite,
          borderRadius: 2,
          display: "flex",
          justifyContent: "center",
          p: 2,
        }}
      >
        <Typography sx={{ color: colors.TextForegroundLow }}>
          No user selected
        </Typography>
      </Box>
    )
  }

  if (userZScore === undefined) {
    return (
      <Box
        sx={{
          alignItems: "center",
          background: colors.BackgroundBaseWhite,
          borderRadius: 2,
          display: "flex",
          justifyContent: "center",
          p: 2,
        }}
      >
        <CircularProgress />
      </Box>
    );
  }

  return (
    <Box
      sx={{
        borderRadius: 2,
        background: colors.BackgroundBaseWhite,
      }}
    >
      <Box
        sx={{
          alignItems: "center",
          borderBottom: 1,
          borderColor: colors.BorderBase,
          display: "flex",
          height: 60,
          justifyContent: "space-between",
          px: 2,
        }}
      >
        <Typography
          sx={{
            fontSize: 20,
            color: colors.TextForegroundLow,
            fontWeight: "bold",
          }}
        >
          Behavior score
        </Typography>

        <TrendingUp
          sx={{
            color: colors.TextForegroundHighlight,
          }}
        />
      </Box>

      <Box
        sx={{
          display: "grid",
          gridTemplateColumns: "1fr minmax(50%, 200px) 1fr",
          gridTemplateRows: "1fr minmax(50%, 200px) 1fr",
          p: 2,
        }}
      >
        <Box
          sx={{
            alignItems: "center",
            aspectRatio: "1/1",
            border: 1,
            borderColor: colors.BorderBase,
            borderRadius: "50%",
            display: "flex",
            gridColumn: "2 / 3",
            gridRow: "2 / 3",
            justifyContent: "center",
          }}
        >
          <Typography
            sx={{
              fontSize: 72,
              color: colors.TextForegroundHighlight,
              fontWeight: "bold",
            }}
          >
            {userZScore <= 2 ? "A" : userZScore > 3 ? "C" : "B"}
          </Typography>
        </Box>
      </Box>

      <Box
        sx={{
          alignItems: "center",
          backgroundColor: colors.BackgroundHighlight,
          display: "flex",
          height: 60,
          justifyContent: "center",
          px: 2,
        }}
      >
        <Typography
          sx={{
            fontSize: 20,
            color: colors.TextForegroundLow,
            fontWeight: "bold",
          }}
        >
          Rating: Average
        </Typography>
      </Box>
    </Box>
  )
}
