import { useState } from "react";

export default function OrderForm() {
  const [options, setOptions] = useState([]);
  const [orderNumber, setOrderNumber] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    const res = await fetch("http://localhost:8080/orders", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ options }),
    });
    const data = await res.json();
    setOrderNumber(data.orderNumber);
  };

  const toggleOption = (option) => {
    setOptions((prev) =>
      prev.includes(option) ? prev.filter((o) => o !== option) : [...prev, option]
    );
  };

  return (
    <form onSubmit={handleSubmit}>
      <label>
        <input type="checkbox" onChange={() => toggleOption("Option 1")} />
        Option 1
      </label>
      <label>
        <input type="checkbox" onChange={() => toggleOption("Option 2")} />
        Option 2
      </label>
      <button type="submit">Złóż zamówienie</button>
      {orderNumber && <p>Numer zamówienia: {orderNumber}</p>}
    </form>
  );
}
