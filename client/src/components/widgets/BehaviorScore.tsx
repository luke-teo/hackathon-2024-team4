import { TrendingDown, TrendingFlat, TrendingUp } from "@mui/icons-material";
import { Box, CircularProgress, Typography } from "@mui/material";
import { colors } from "../../utils/colors";
import { useContext } from "react";
import { SelectedUserContext } from "../context/SelectedUserContext";
import { DateTime } from "luxon";
import { useGetScoresByUserIdQuery } from "../../services/api/v1";

export const BehaviorScore = (): JSX.Element => {
  const { selectedUser } = useContext(SelectedUserContext);

  const now = DateTime.now();
  const start = now.startOf("month");
  const end = now.endOf("month");

  const { data: getUserScores } = useGetScoresByUserIdQuery({
    userId: selectedUser?.id.toString() ?? "",
    startDate: start.toISODate() ?? "",
    endDate: end.toISODate() ?? "",
  });

  const scores = getUserScores?.scores;
  const latestScore =
    scores && scores.length > 1 ? scores[scores.length - 1] : null;

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
    );
  }

  if (!latestScore) {
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
        flex: 0,
        height: "fit-content",
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

        {latestScore.zScore <= 2 ? (
          <TrendingUp
            sx={{
              color: colors.TextForegroundSuccess,
            }}
          />
        ) : latestScore.zScore > 3 ? (
          <TrendingDown
            sx={{
              color: colors.TextForegroundDanger,
            }}
          />
        ) : (
          <TrendingFlat
            sx={{
              color: colors.TextForegroundWarning,
            }}
          />
        )}
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
              color:
                latestScore.zScore <= 2
                  ? colors.TextForegroundSuccess
                  : latestScore.zScore > 3
                    ? colors.TextForegroundDanger
                    : colors.TextForegroundWarning,
              fontWeight: "bold",
            }}
          >
            {latestScore.zScore <= 2 ? "A" : latestScore.zScore > 3 ? "C" : "B"}
          </Typography>
        </Box>
      </Box>

      <Box
        sx={{
          alignItems: "center",
          backgroundColor: colors.BackgroundHighlight,
          display: "flex",
          justifyContent: "center",
          p: 2,
        }}
      >
        <Typography
          sx={{
            fontSize: 16,
            color: colors.TextForegroundLow,
          }}
        >
          {latestScore.zScore <= 2
            ? "This person is doing well."
            : latestScore.zScore > 3
              ? "This person's performance is worrisome."
              : "This person requires some attention."}
        </Typography>
      </Box>
    </Box>
  );
};
