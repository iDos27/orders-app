import { useEffect, useState } from "react";

export default function WarehouseView() {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/orders")
      .then((res) => res.json())
      .then(setOrders);
  }, []);

  const updateStatus = (id, status) => {
    fetch(`http://localhost:8080/orders/${id}/status`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ status }),
    }).then(() => {
      setOrders((prev) =>
        prev.map((o) => (o.id === id ? { ...o, status } : o))
      );
    });
  };

  return (
    <div>
      <h1>Aktualne zamówienia</h1>
      {orders.map((order) => (
        <div key={order.id}>
          <p>{order.orderNumber} - Status: {order.status}</p>
          <button onClick={() => updateStatus(order.id, "in_progress")}>
            W trakcie
          </button>
          <button onClick={() => updateStatus(order.id, "completed")}>
            Zakończone
          </button>
        </div>
      ))}
    </div>
  );
}
