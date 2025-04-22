import { BrowserRouter, Routes, Route } from "react-router-dom";
import ClientView from "./pages/ClientView.jsx"
import WarehouseView from "./pages/WarehouseView.jsx"

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<ClientView />} />
        <Route path="/warehouse" element={<WarehouseView />} />
      </Routes>
    </BrowserRouter>
  );
}