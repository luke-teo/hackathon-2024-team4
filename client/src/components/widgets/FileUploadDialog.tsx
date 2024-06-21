import { AttachFile } from "@mui/icons-material";
import { Box, Button, Dialog, Typography } from "@mui/material";
import { MuiFileInput } from "mui-file-input";
import { useState } from "react";
import { usePostUploadCsvMutation } from "../../services/api/v1";

type Props = {
  onClose: (e: React.MouseEvent) => void;
};

export const FileUploadDialog = ({ onClose }: Props): JSX.Element => {
  const [file, setFile] = useState<File | null>(null);

  const [postUploadFile, postUploadFileRes] = usePostUploadCsvMutation();

  const handleFileChange = (newFile: File | null): void => {
    setFile(newFile);
  };

  const handleUploadFile = (): void => {
    postUploadFile({
      body: {
        filename: file?.name,
        file: file ?? undefined,
      },
    })
      .unwrap()
      .then((resp) => console.log(resp));
  };

  return (
    <Dialog open={true} onClose={onClose}>
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          width: "428px",
          height: "188px",
          padding: "24px",
          gap: "16px",
        }}
      >
        <Typography variant="h5">Upload CSV</Typography>
        <MuiFileInput
          value={file}
          variant="outlined"
          size="small"
          placeholder="Upload text chat file"
          onChange={handleFileChange}
          InputProps={{
            inputProps: {
              accept: ".csv",
            },
            startAdornment: <AttachFile />,
          }}
          hideSizeText
        />
        <Box
          sx={{
            display: "flex",
            justifyContent: "end",
            paddingY: "24px",
            gap: "8px",
          }}
        >
          <Button variant="outlined" color="inherit" onClick={onClose}>
            Cancel
          </Button>
          <Button
            variant="contained"
            onClick={handleUploadFile}
            sx={{ backgroundColor: "#9C72ED" }}
          >
            Upload
          </Button>
        </Box>
      </Box>
    </Dialog>
  );
};
