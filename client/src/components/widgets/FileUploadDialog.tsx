import { AttachFile } from "@mui/icons-material";
import { Box, Button, Dialog, Typography } from "@mui/material";
import { LoadingButton } from "@mui/lab";
import { MuiFileInput } from "mui-file-input";
import { useState } from "react";

type Props = {
  closeDialog: () => void;
};

export const FileUploadDialog = ({ closeDialog }: Props): JSX.Element => {
  const [file, setFile] = useState<File | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const handleFileChange = (newFile: File | null): void => {
    setFile(newFile);
  };

  const handleUploadFile = (): void => {
    if (!file) {
      return;
    }

    setIsLoading(true);

    const data = new FormData();
    data.append("file", file);

    fetch("/api/upload_csv", {
      body: data,
      method: "POST",
      headers: {
        contentType: "multipart/form-data",
      },
    }).then((resp) => {
      console.log(resp);
      setIsLoading(false);
      closeDialog();
    });
  };

  return (
    <Dialog open={true} onClose={() => closeDialog()}>
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
          <Button
            variant="outlined"
            color="inherit"
            onClick={() => closeDialog()}
          >
            Cancel
          </Button>
          <LoadingButton
            loading={isLoading}
            variant="contained"
            onClick={handleUploadFile}
            sx={{ backgroundColor: "#9C72ED" }}
          >
            Upload
          </LoadingButton>
        </Box>
      </Box>
    </Dialog>
  );
};
