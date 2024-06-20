import { Box, Typography } from "@mui/material";
import { colors } from "../../utils/colors";
import { useContext } from "react";
import { SelectedUserContext } from "../context/SelectedUserContext";
import {
	Chart as ChartJS,
	CategoryScale,
	LinearScale,
	PointElement,
	LineElement,
	Title,
	Tooltip,
	Legend,
} from "chart.js";
import { Line } from "react-chartjs-2";
import { faker } from "@faker-js/faker";

export const ScoreChart = () => {
	const { selectedUser } = useContext(SelectedUserContext);
	ChartJS.register(
		CategoryScale,
		LinearScale,
		PointElement,
		LineElement,
		Title,
		Tooltip,
		Legend,
	);

	const options = {
		responsive: true,
		plugins: {
			legend: {
				position: "top" as const,
			},
			title: {
				display: true,
				text: "Chart.js Line Chart",
			},
		},
	};

	const labels = [
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
	];

	const data = {
		labels,
		datasets: [
			{
				label: "Dataset 1",
				data: labels.map(() =>
					faker.datatype.number({ min: -1000, max: 1000 }),
				),
				borderColor: "rgb(255, 99, 132)",
				backgroundColor: "rgba(255, 99, 132, 0.5)",
			},
			{
				label: "Dataset 2",
				data: labels.map(() =>
					faker.datatype.number({ min: -1000, max: 1000 }),
				),
				borderColor: "rgb(53, 162, 235)",
				backgroundColor: "rgba(53, 162, 235, 0.5)",
			},
		],
	};
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
					<Typography>{selectedUser?.name}</Typography>
				</Box>
				<Line options={options} data={data} />
			</Box>
		</Box>
	);
};
