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
import type { Score } from "../services/api/v1";
import { Box, Typography } from "@mui/material";
import { colors } from "../utils/colors";

type Props = {
	datetime: DateTime;
	scores: Score[];
};

export const LineChartComponent = ({ datetime, scores }: Props) => {
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
			s.currentScore - s.standardDeviation,
			s.currentScore + s.standardDeviation,
		],
	}));

	if (data.length < 1) {
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
				<Typography
					sx={{
						color: colors.TextForegroundLow,
						fontSize: 16,
						fontWeight: "bold",
					}}
				>
					No Data is Available for this Period
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
