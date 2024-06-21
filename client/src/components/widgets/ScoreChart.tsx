import { Box, Typography } from "@mui/material";
import { colors } from "../../utils/colors";
import { useContext, useState } from "react";
import { SelectedUserContext } from "../context/SelectedUserContext";
import { DateTime, type MonthNumbers } from "luxon";
import { mock } from "../../mock/mock";
import { LineChart } from "../LineChart";

export const ScoreChart = () => {
	const { selectedUser } = useContext(SelectedUserContext);
	const now = DateTime.now();
	const currentMonth = now.month;
	const [month, setMonth] = useState<MonthNumbers>(currentMonth);

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
			}}
		>
			<Box
				sx={{
					height: 60,
					display: "flex",
					alignItems: "center",
					borderBottom: 1,
					borderColor: colors.BorderBase,
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
					User
				</Typography>
			</Box>
			<Box
				sx={{
					p: 2,
				}}
			>
				<Box>
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
				<LineChart month={month} scores={scores} />
			</Box>
		</Box>
	);
};
