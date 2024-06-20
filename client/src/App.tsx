import { Box } from "@mui/material";
import "./App.css";
import { Layout } from "./components/layout/Layout";
import { UserSideMenu } from "./components/widgets/UserSideMenu";
import { SelectedUserProvider } from "./components/context/SelectedUserContext";
import { ScoreChart } from "./components/widgets/ScoreChart";

function App() {
	return (
		<SelectedUserProvider>
			<Layout>
				<UserSideMenu />
				<ScoreChart />
				<Box sx={{ border: 1, height: "100%", flex: 0, minWidth: 360 }}>
					3 here
				</Box>
			</Layout>
		</SelectedUserProvider>
	);
}

export default App;
