import {
  Autocomplete,
  Avatar,
  Box,
  Button,
  InputAdornment,
  TextField,
  Typography,
} from "@mui/material";
import { getUserInitials, users } from "../../utils/users";
import { type MouseEvent, useContext, useState } from "react";
import { SelectedUserContext } from "../context/SelectedUserContext";
import SearchIcon from "@mui/icons-material/Search";
import { colors } from "../../utils/colors";
import type { User } from "../../utils/types";
import { FileUploadDialog } from "./FileUploadDialog";
import { UserListItem } from "../UserListItem";

export const UserAvatar = ({ name }: { name: string }) => {
  return (
    <Avatar
      src={`/avatars/${name}.jpg`}
      sx={{
        backgroundColor: "#E3F7E8",
        color: "#00871D",
      }}
    >
      {getUserInitials(name)}
    </Avatar>
  );
};

export const UserSideMenu = () => {
  const { setSelectedUser } = useContext(SelectedUserContext);
  const [value, setValue] = useState<User | null>(null);
  const [isUploadDialogOpen, setIsUploadDialogOpen] = useState<boolean>(false);

  const handleClickUpload = (e: MouseEvent) => {
    e.preventDefault();
    setIsUploadDialogOpen(true);
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
              <UserAvatar name={option.name} />

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
            <UserListItem key={u.id} user={u} setValue={setValue} />
          ))}
        </Box>
      </Box>

      <Box
        sx={{
          display: "flex",
          justifyContent: "end",
          paddingY: "30px",
          paddingX: "18px",
        }}
      >
        <Button
          size="small"
          variant="contained"
          sx={{ backgroundColor: "#9C72ED" }}
          onClick={handleClickUpload}
        >
          Upload CSV
        </Button>
      </Box>

      {isUploadDialogOpen && (
        <FileUploadDialog closeDialog={() => setIsUploadDialogOpen(false)} />
      )}
    </Box>
  );
};
