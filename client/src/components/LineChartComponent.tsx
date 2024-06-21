import {
	Line,
	XAxis,
	YAxis,
	Tooltip,
	Legend,
	Area,
	ResponsiveContainer,
	ComposedChart,
} from "recharts";
import { DateTime } from "luxon";
import { useGetScoresByUserIdQuery } from "../services/api/v1";
import { Box, CircularProgress, Typography } from "@mui/material";
import { colors } from "../utils/colors";
import { useContext } from "react";
import { SelectedUserContext } from "./context/SelectedUserContext";

type Props = {
	datetime: DateTime;
};

export const LineChartComponent = ({ datetime }: Props) => {
	const { selectedUser } = useContext(SelectedUserContext);
	const start = datetime.startOf("month");
	const end = datetime.endOf("month");
	const userId = selectedUser?.id.toString() ?? "";

	const { data: getUserScores } = useGetScoresByUserIdQuery({
		userId: userId,
		startDate: start.toISODate() ?? "",
		endDate: end.toISODate() ?? "",
	});

	const scores = getUserScores?.scores ?? [];

	const relevantScores = scores.filter(
		(s) =>
			DateTime.fromISO(s.date).month === datetime.month &&
			DateTime.fromISO(s.date).year === datetime.year,
	);

	const data = relevantScores.map((s) => ({
		day: DateTime.fromISO(s.date).day,
		score: s.currentScore,
		mean: s.mean,
		stdRange: [
			s.mean - s.standardDeviation * 2,
			s.mean + s.standardDeviation * 2,
		],
	}));

	const currentMonth = DateTime.now().month;
	// Check if data is empty and the month is between May and the current month
	if (
		data.length < 1 &&
		datetime.month >= 4 &&
		datetime.month <= currentMonth
	) {
		console.log("Data is empty and month is between May and the current month");
		return (
			<Box
				sx={{
					height: "100%",
					width: "100%",
					display: "flex",
					alignItems: "center",
					justifyContent: "center",
				}}
			>
				<CircularProgress />
			</Box>
		);
	}

	// Check if data is empty and the month is less than April
	if (data.length < 1 && datetime.month < 4) {
		console.log("Data is empty and month is less than April");
		return (
			<Box
				sx={{
					height: "100%",
					width: "100%",
					display: "flex",
					alignItems: "center",
					justifyContent: "center",
				}}
			>
				{/* Handle this case if needed */}
				<Typography
					sx={{
						color: colors.TextForegroundLow,
						fontSize: 16,
						fontWeight: "bold",
					}}
				>
					No Data Available Before April
				</Typography>
			</Box>
		);
	}

	// Check if data is empty and the month is greater than the current month
	if (data.length < 1 && datetime.month > currentMonth) {
		console.log("Data is empty and month is in the future");
		return (
			<Box
				sx={{
					height: "100%",
					width: "100%",
					display: "flex",
					alignItems: "center",
					justifyContent: "center",
				}}
			>
				{/* Handle this case if needed */}
				<Typography
					sx={{
						color: colors.TextForegroundLow,
						fontSize: 16,
						fontWeight: "bold",
					}}
				>
					No Data Available for Future Dates
				</Typography>
			</Box>
		);
	}

	return (
		<ResponsiveContainer width="100%" height="100%">
			<ComposedChart data={data}>
				<XAxis dataKey="day" />
				<YAxis />
				<Tooltip />
				<Legend />
				<Area
					type="monotone"
					dataKey="stdRange"
					stroke="none"
					fill="rgba(185, 192, 204, 0.3)"
					activeDot={{ r: 0 }}
					legendType="none"
					tooltipType="none"
				/>
				<Line
					type="monotone"
					dataKey="score"
					stroke="#5E1CDE"
					dot={false}
					label="Score"
				/>
				<Line
					type="monotone"
					dataKey="mean"
					stroke="#FFB200"
					dot={false}
					label="Avg."
				/>
			</ComposedChart>
		</ResponsiveContainer>
	);
};
