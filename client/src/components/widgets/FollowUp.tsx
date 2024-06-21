import { Box, Button, CircularProgress, Input, Typography } from "@mui/material";
import { colors } from "../../utils/colors";
import React, { useContext, useEffect } from "react";
import { SelectedUserContext } from "../context/SelectedUserContext";
import { HrNote } from "../../utils/types";
import { hrNotesMock } from "../../mock/hrNotesMock";

export const FollowUp = (): JSX.Element => {
  const { selectedUser } = useContext(SelectedUserContext);
  const [hrNotes, setHrNotes] = React.useState<HrNote[] | undefined>(undefined);
  const [currentHrNote, setCurrentHrNote] = React.useState<string>('');

  useEffect(() => {
    if (selectedUser === undefined) {
      setHrNotes(undefined);
    }

    for (let i = 0; i < hrNotesMock.length; i++) {
      if (hrNotesMock[i].userId === selectedUser?.id) {
        setHrNotes(hrNotesMock[i].notes);
      }
    }
  }, [selectedUser, setHrNotes]);

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

  if (hrNotes === undefined) {
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
        backgroundColor: colors.BackgroundBaseWhite,
        borderRadius: 2,
        display: "flex",
        flexDirection: "column",
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
          Follow Up
        </Typography>
      </Box>


      <Box sx={{
        overflowX: "hidden",
        overflowY: "scroll",
        flex: 1,
        p: 2
      }}>
        {hrNotes.map((note, index) => (
          <Box key={index} sx={{
            borderBottom: 1,
            borderBottomColor: colors.BorderBase,
            display: "flex",
            gap: 2,
            justifyContent: "space-between",
            p: 2,
          }}>
            <Typography>{note.body}</Typography>
            <Typography>{note.timestamp}</Typography>
            </Box>))}
      </Box>

      <Box sx={{
        display: "flex",
        borderTop: 1,
        borderTopColor: colors.BorderBase,
        gap: 2,
        justifyContent: "center",
        p: 2,
      }}>
        <Input minRows="3" maxRows="5" multiline sx={{
          flex: 1,
        }} value={currentHrNote} onChange={(e) => {
          setCurrentHrNote(e.target.value);
        }} />

        <Button variant="outlined" onClick={() => {


          setHrNotes([
            ...hrNotes,
            {
              body: currentHrNote,
              timestamp: 42344232,
            }
          ])
        }}>Send</Button>
      </Box>
    </Box>
  )
};
