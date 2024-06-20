import { Box, Typography } from "@mui/material";
import { colors } from "../../utils/colors";

export const Information = (): JSX.Element => {
  return (
    <Box
      sx={{
        borderRadius: 2,
        background: colors.BackgroundBaseWhite,
        flex: 1,
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
          Information
        </Typography>
      </Box>

      <Box
        sx={{
          p: 2,
        }}>
        <Typography
          sx={{
            color: colors.TextForegroundLow,
          }}
        >
          Placeholder
        </Typography></Box>
    </Box >
  )
}
