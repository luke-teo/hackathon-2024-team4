import { Box, Button, Typography } from "@mui/material";
import { colors } from "../../utils/colors";
import { useContext, useState } from "react";
import { SelectedUserContext } from "../context/SelectedUserContext";
import { DateTime, type MonthNumbers } from "luxon";
import { mock } from "../../mock/mock";
import { ChevronLeft, ChevronRight } from "@mui/icons-material";
import { LineChartComponent } from "../LineChartComponent";

export const ScoreChart = () => {
	const { selectedUser } = useContext(SelectedUserContext);
	const now = DateTime.now();
	const [datetime, setDateTime] = useState<DateTime>(now);

	const incrementMonth = () => {
		setDateTime((prevDateTime) => {
			return prevDateTime.plus({ months: 1 });
		});
	};

	const decrementMonth = () => {
		setDateTime((prevDateTime) => {
			return prevDateTime.minus({ months: 1 });
		});
	};

	if (!selectedUser) {
		return (
			<Box
				sx={{
					height: "100%",
					flex: 1,
					borderRadius: 2,
					background: colors.BackgroundBaseWhite,
					display: "flex",
					justifyContent: "center",
					alignItems: "center",
				}}
			>
				<Typography
					sx={{
						color: colors.TextForegroundLow,
						fontWeight: 700,
						fontSize: 16,
					}}
				>
					Please select a user to view their chart
				</Typography>
			</Box>
		);
	}

	const userData = mock.find(
		(u) => Number.parseInt(u.userId) === selectedUser.id,
	);

	const scores = userData?.scores ?? [];

	return (
		<Box
			sx={{
				height: "100%",
				flex: 1,
				borderRadius: 2,
				background: colors.BackgroundBaseWhite,
				display: "flex",
				flexDirection: "column",
			}}
		>
			<Box
				sx={{
					minHeight: 60,
					display: "flex",
					alignItems: "center",
					justifyContent: "space-between",
					borderBottom: 1,
					borderColor: colors.BorderBase,
					px: 2,
					flexShrink: 0,
				}}
			>
				<Typography
					sx={{
						fontSize: 20,
						color: colors.TextForegroundLow,
						fontWeight: "bold",
					}}
				>
					User
				</Typography>
				<Box
					sx={{
						display: "flex",
						alignItems: "center",
						gap: 2,
					}}
				>
					<Typography
						sx={{
							fontWeight: "bold",
							color: colors.TextForegroundLow,
							fontSize: 20,
						}}
					>
						{datetime.monthLong}
					</Typography>
					<Box sx={{ display: "flex", alignItems: "center" }}>
						<Button
							onClick={decrementMonth}
							sx={{
								border: 1,
								borderColor: colors.TertiaryAlpha,
								borderRadius: "4px 0px 0px 4px",
							}}
						>
							<ChevronLeft htmlColor={colors.TextForegroundLow} />
						</Button>
						<Button
							onClick={incrementMonth}
							sx={{
								border: 1,
								borderColor: colors.TertiaryAlpha,
								borderRadius: "0px 4px 4px 0px",
							}}
						>
							<ChevronRight htmlColor={colors.TextForegroundLow} />
						</Button>
					</Box>
				</Box>
			</Box>
			<Box
				sx={{
					p: 2,
					flex: 1,
					display: "flex",
					flexDirection: "column",
					alignItems: "center",
					overflow: "hidden",
				}}
			>
				<Box sx={{ flexShrink: 0, width: "100%" }}>
					<Typography
						sx={{
							fontWeight: "bold",
							fontSize: 24,
							color: colors.TextForegroundLow,
						}}
					>
						{selectedUser?.name}
					</Typography>
				</Box>
				<Box
					sx={{
						flex: 1,
						width: "100%",
						p: 2,
					}}
				>
					<LineChartComponent datetime={datetime} scores={scores} />
				</Box>
			</Box>
		</Box>
	);
};
