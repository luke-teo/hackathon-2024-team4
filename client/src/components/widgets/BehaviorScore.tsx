import { TrendingDown, TrendingFlat, TrendingUp } from "@mui/icons-material";
import { Box, CircularProgress, Typography } from "@mui/material";
import { colors } from "../../utils/colors";
import { useContext } from "react";
import { SelectedUserContext } from "../context/SelectedUserContext";
import { DateTime } from "luxon";
import { useGetScoresByUserIdQuery } from "../../services/api/v1";

export const BehaviorScore = (): JSX.Element => {
	const { selectedUser } = useContext(SelectedUserContext);

	const now = DateTime.now();
	const start = now.startOf("month");
	const end = now.endOf("month");

	const { data: getUserScores } = useGetScoresByUserIdQuery({
		userId: selectedUser?.id.toString() ?? "",
		startDate: start.toISODate() ?? "",
		endDate: end.toISODate() ?? "",
	});

	const scores = getUserScores?.scores;
	const latestScore =
		scores && scores.length > 1 ? scores[scores.length - 1] : null;

	const isWarning = () => {
		const zScore = Math.abs(latestScore?.zScore ?? 0);
		return zScore > 2 && zScore < 3;
	};

	const isUrgent = () => {
		const zScore = Math.abs(latestScore?.zScore ?? 0);
		return zScore > 3;
	};

	console.log(latestScore?.zScore);
	console.log("iswarning", isWarning());
	console.log("isUrgent", isUrgent());
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
		);
	}

	if (!latestScore) {
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
				flex: 0,
				height: "fit-content",
				borderRadius: 2,
				background: colors.BackgroundBaseWhite,
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
					Behavior score
				</Typography>

				{isWarning() ? (
					<TrendingFlat
						sx={{
							color: colors.TextForegroundWarning,
						}}
					/>
				) : isUrgent() ? (
					<TrendingDown
						sx={{
							color: colors.TextForegroundDanger,
						}}
					/>
				) : (
					<TrendingUp
						sx={{
							color: colors.TextForegroundSuccess,
						}}
					/>
				)}
			</Box>

			<Box
				sx={{
					display: "grid",
					gridTemplateColumns: "1fr minmax(50%, 200px) 1fr",
					gridTemplateRows: "1fr minmax(50%, 200px) 1fr",
					p: 2,
				}}
			>
				<Box
					sx={{
						alignItems: "center",
						aspectRatio: "1/1",
						border: 1,
						borderColor: colors.BorderBase,
						borderRadius: "50%",
						display: "flex",
						gridColumn: "2 / 3",
						gridRow: "2 / 3",
						justifyContent: "center",
					}}
				>
					<Typography
						sx={{
							fontSize: 72,
							color: isWarning()
								? colors.TextForegroundWarning
								: isUrgent()
									? colors.TextForegroundDanger
									: colors.TextForegroundSuccess,
							fontWeight: "bold",
						}}
					>
						{isWarning() ? "B" : isUrgent() ? "C" : "A"}
					</Typography>
				</Box>
			</Box>

			<Box
				sx={{
					alignItems: "center",
					backgroundColor: colors.BackgroundHighlight,
					display: "flex",
					justifyContent: "center",
					p: 2,
				}}
			>
				<Typography
					sx={{
						fontSize: 16,
						color: colors.TextForegroundLow,
					}}
				>
					{isWarning()
						? "This person requires some attention."
						: isUrgent()
							? "This person's performance is worrisome."
							: "This person is doing well."}
				</Typography>
			</Box>
		</Box>
	);
};
