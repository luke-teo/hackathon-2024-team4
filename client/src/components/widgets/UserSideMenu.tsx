import {
  Autocomplete,
  Avatar,
  Box,
  Button,
  Chip,
  CircularProgress,
  IconButton,
  InputAdornment,
  ListItem,
  TextField,
  Typography,
} from "@mui/material";
import { getUserInitials, users } from "../../utils/users";
import React, { type MouseEvent, useContext, useState } from "react";
import { SelectedUserContext } from "../context/SelectedUserContext";
import SearchIcon from "@mui/icons-material/Search";
import CloseIcon from "@mui/icons-material/Close";
import PersonIcon from "@mui/icons-material/Person";
import { colors } from "../../utils/colors";
import type { User } from "../../utils/types";
import { mock } from "../../mock/mock";
import { FileUploadDialog } from "./FileUploadDialog";

const UserAvatar = ({
  avatarUrl,
  name,
}: {
  avatarUrl: string;
  name: string;
}) => {
  return (
    <Avatar
      src={avatarUrl}
      sx={{
        backgroundColor: "#E3F7E8",
        color: "#00871D",
      }}
    >
      {getUserInitials(name)}
    </Avatar>
  );
};

type UserStatus = { userId: number; zScore: number };

const getZScore = (usersStatus: UserStatus[], userId: number): number => {
  return usersStatus.find((e) => e.userId === userId)?.zScore ?? 0;
};

export const UserSideMenu = () => {
  const { selectedUser, setSelectedUser } = useContext(SelectedUserContext);
  const [value, setValue] = useState<User | null>(null);
  const [usersStatus, setUsersStatus] = useState<UserStatus[] | undefined>(
    undefined,
  );
  const [isUploadDialogOpen, setIsUploadDialogOpen] = useState<boolean>(false);

  React.useEffect(() => {
    const usersStatusCalc: UserStatus[] = [];

    for (let j = 0; j < mock.length; j++) {
      usersStatusCalc.push({
        userId: Number(mock[j].userId),
        zScore: mock[j].scores[mock[j].scores.length - 1].zScore,
      });
    }

    setUsersStatus(usersStatusCalc);
  });

  const handleSelectUser = (u: User) => {
    setSelectedUser(u);
    setValue(u);
  };

  const handleDeleteUser = (e: MouseEvent) => {
    e.preventDefault();
    setSelectedUser(null);
    setValue(null);
  };

  const handleClickUpload = (e: MouseEvent) => {
    e.preventDefault();
    setIsUploadDialogOpen(true);
  };

  if (usersStatus === undefined) {
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
              <UserAvatar avatarUrl={option.avatarUrl} name={option.name} />

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
                  <UserAvatar avatarUrl={u.avatarUrl} name={u.name} />

                  <Typography
                    sx={{
                      fontSize: 14,
                      fontWeight: "bold",
                    }}
                  >
                    {u.name}
                  </Typography>

                  {getZScore(usersStatus, u.id) <= 2 ? (
                    <></>
                  ) : getZScore(usersStatus, u.id) > 3 ? (
                    <Chip color="error" label="Required" size="small" />
                  ) : (
                    <Chip color="warning" label="Review" size="small" />
                  )}
                </Box>
              </Box>
            </ListItem>
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
        <FileUploadDialog onClose={() => setIsUploadDialogOpen(false)} />
      )}
    </Box>
  );
};
