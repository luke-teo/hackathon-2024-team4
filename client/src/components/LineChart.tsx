import { Line } from "react-chartjs-2";
import { DateTime } from "luxon";
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

ChartJS.register(
	CategoryScale,
	LinearScale,
	PointElement,
	LineElement,
	Title,
	Tooltip,
	Legend,
);

type Props = {
	month: number;
	scores: Array<{
		date: string;
		currentScore: number;
		mean: number;
		standardDeviation: number;
	}>;
};

export const LineChart = ({ month, scores }: Props) => {
	const options = {
		responsive: true,
		plugins: {
			legend: {
				position: "top" as const,
				labels: {
					generateLabels: (chart) => {
						const original =
							ChartJS.defaults.plugins.legend.labels.generateLabels;
						const labelsOriginal = original.call(this, chart);

						// Filter out the labels you don't want to show in the legend
						return labelsOriginal.filter(
							(label) => label.text !== "std1-top" && label.text !== "std1-bot",
						);
					},
				},
			},
		},
	};

	const relevantScores = scores.filter(
		(s) => DateTime.fromISO(s.date).month === month,
	);

	const labels = relevantScores.map((s) => DateTime.fromISO(s.date).day);

	const scoreData = relevantScores.map((s) => s.currentScore);

	const meanData = relevantScores.map((s) => s.mean);

	const stdTopData = relevantScores.map((s) => {
		const std = s.standardDeviation;
		return s.currentScore + std;
	});
	const stdBotData = relevantScores.map((s) => {
		const std = s.standardDeviation;
		return s.currentScore - std;
	});

	const data = {
		labels,
		datasets: [
			{
				label: "Score",
				data: scoreData,
				borderColor: "rgb(255, 99, 132)",
				backgroundColor: "rgba(255, 99, 132, 0.5)",
				borderWidth: 1,
				pointRadius: 0,
				pointHoverRadius: 0,
				tension: 0.4,
				pointHitRadius: 0,
				fill: false,
			},
			{
				label: "Avg.",
				data: meanData,
				borderColor: "rgb(53, 162, 235)",
				backgroundColor: "rgba(53, 162, 235, 0.5)",
				borderWidth: 1,
				pointRadius: 0,
				pointHoverRadius: 0,
				tension: 0.4,
				pointHitRadius: 0,
				fill: false,
			},
			{
				label: "std1-top",
				data: stdTopData,
				borderColor: "rgba(75, 192, 192, 0)",
				backgroundColor: "rgba(75, 192, 192, 0.2)",
				borderWidth: 1,
				pointRadius: 0,
				pointHoverRadius: 0,
				tension: 0.4,
				pointHitRadius: 0,
				fill: true, // Fill to the dataset below (std1-bot)
			},
			{
				label: "std1-bot",
				data: stdBotData,
				borderColor: "rgba(75, 192, 192, 1)",
				backgroundColor: "rgba(75, 192, 192, 0.2)",
				borderWidth: 1,
				pointRadius: 0,
				pointHoverRadius: 0,
				tension: 0.4,
				pointHitRadius: 0,
				fill: false, // Fill to the dataset above (std1-top)
			},
		],
	};

	return <Line options={options} data={data} />;
};
