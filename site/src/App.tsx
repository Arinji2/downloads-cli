import { Route, Routes } from "react-router-dom";
import Home from "./Home";

function App() {
  return (
    <div className="w-full h-fit flex flex-col items-center justify-start max-w-[1280px]">
      <Routes>
        <Route path="/" element={<Home />} />
      </Routes>
    </div>
  );
}

export default App;
