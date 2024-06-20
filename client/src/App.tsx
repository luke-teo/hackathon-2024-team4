import "./App.css";
import { Layout } from "./components/layout/Layout";
import { UserSideMenu } from "./components/UserSideMenu";
import { ScoreChart } from "./components/ScoreChart";

function App() {
  return (
    <Layout>
      <UserSideMenu />
      <ScoreChart />
    </Layout>
  );
}

export default App;
