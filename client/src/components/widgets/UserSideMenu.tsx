import {
  Autocomplete,
  Avatar,
  Box,
  IconButton,
  InputAdornment,
  ListItem,
  TextField,
  Typography,
} from "@mui/material";
import { getUserInitials, users } from "../../utils/users";
import { type MouseEvent, useContext, useState } from "react";
import { SelectedUserContext } from "../context/SelectedUserContext";
import SearchIcon from "@mui/icons-material/Search";
import CloseIcon from "@mui/icons-material/Close";
import PersonIcon from "@mui/icons-material/Person";
import { colors } from "../../utils/colors";
import type { User } from "../../utils/types";

export const UserSideMenu = () => {
  const { selectedUser, setSelectedUser } = useContext(SelectedUserContext);
  const [value, setValue] = useState<User | null>(null);

  const handleSelectUser = (u: User) => {
    setSelectedUser(u);
    setValue(u);
  };

  const handleDeleteUser = (e: MouseEvent) => {
    e.preventDefault();
    setSelectedUser(null);
    setValue(null);
  };

  return (
    <Box
      sx={{
        height: "100%",
        flex: 0,
        width: "100%",
        borderRadius: 2,
        background: colors.BackgroundBaseWhite,
      }}
    >
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
        }}
      >
        <Autocomplete
          value={value}
          onChange={(_, v) => {
            if (v) {
              setSelectedUser(v);
              setValue(v);
            }
          }}
          size="small"
          disablePortal
          options={users}
          getOptionLabel={(u) => {
            return u.name;
          }}
          fullWidth
          sx={{
            width: 300,
            p: 2,
          }}
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
                {getUserInitials(option.name)}
              </Avatar>
              <Typography>{option.name}</Typography>
            </Box>
          )}
          renderInput={(params) => (
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
        <Box
          sx={{
            display: "flex",
            flexDirection: "column",
          }}
        >
          {users.map((u) => (
            <ListItem
              key={u.id}
              secondaryAction={
                <>
                  {selectedUser?.id === u.id ? (
                    <IconButton onClick={(e) => handleDeleteUser(e)}>
                      <CloseIcon />
                    </IconButton>
                  ) : (
                    <IconButton onClick={() => handleSelectUser(u)}>
                      <PersonIcon
                        sx={{
                          color:
                            selectedUser?.id === u.id
                              ? colors.IconSelected
                              : undefined,
                        }}
                      />
                    </IconButton>
                  )}
                </>
              }
              sx={{
                backgroundColor:
                  selectedUser?.id === u.id
                    ? colors.BackgroundSelected
                    : undefined,
              }}
            >
              <Box
                sx={{
                  display: "flex",
                  alignItems: "center",
                  justifyContent: "space-between",
                }}
              >
                <Box
                  sx={{
                    display: "flex",
                    alignItems: "center",
                    gap: 1,
                  }}
                >
                  <Avatar src={u.avatarUrl}>
                    {getUserInitials(u.name)}
                  </Avatar>
                  <Typography
                    sx={{
                      fontSize: 14,
                      fontWeight: "bold",
                    }}
                  >
                    {u.name}
                  </Typography>
                </Box>
              </Box>
            </ListItem>
          ))}
        </Box>
      </Box>
    </Box>
  );
};
