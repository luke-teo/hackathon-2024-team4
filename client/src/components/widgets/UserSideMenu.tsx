import {
  Autocomplete,
  Avatar,
  Box,
  InputAdornment,
  TextField,
  Typography,
} from "@mui/material";
import { users } from "../../utils/users";
import { useContext } from "react";
import { SelectedUserContext } from "../context/SelectedUserContext";
import SearchIcon from "@mui/icons-material/Search";

export const UserSideMenu = () => {
  const { selectedUser, setSelectedUser } = useContext(SelectedUserContext);

  console.log(selectedUser);
  return (
    <Box
      sx={{
        height: "100%",
        flex: 0,
        width: "100%",
        borderRadius: 2,
        backgroundColor: "white",
      }}
    >
      <Box sx={{ p: 2 }}>
        <Autocomplete
          value={selectedUser}
          onChange={(_, v) => v && setSelectedUser(v)}
          size="small"
          disablePortal
          options={users}
          getOptionLabel={(u) => {
            return u.name;
          }}
          fullWidth
          sx={{ width: 300 }}
          renderOption={(props, option) => (
            <Box
              sx={{
                display: "flex",
                alignItems: "center",
                justifyContent: "start",
                gap: 2,
              }}
              component="li"
              {...props}
            >
              <Avatar src={option.avatarUrl}>
                {option.name
                  .split(" ")
                  .map((part) => part[0])
                  .join("")
                  .toUpperCase()}
              </Avatar>
              <Typography>{option.name}</Typography>
            </Box>
          )}
          renderInput={(params) => (
            // <TextField {...params} />
            <TextField
              {...params}
              label="Search for a user..."
              InputProps={{
                ...params.InputProps,
                startAdornment: (
                  <InputAdornment position="start">
                    <SearchIcon />
                  </InputAdornment>
                ),
                endAdornment: null,
              }}
            />
          )}
        />
      </Box>
    </Box>
  );
};
