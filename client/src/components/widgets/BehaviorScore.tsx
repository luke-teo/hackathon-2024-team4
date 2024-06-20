import { ArrowUpward, TrendingUp } from "@mui/icons-material";
import { Box, Card, Divider, Icon, Stack, Typography } from "@mui/material";
import { colors } from "../../utils/colors";

export const BehaviorScore = (): JSX.Element => {
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
            70
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
