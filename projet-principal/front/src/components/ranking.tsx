import { useEffect, useState } from "react";
import {
  BarChart,
  Bar,
  ResponsiveContainer,
  XAxis,
  YAxis,
  Legend,
  CartesianGrid,
} from "recharts";

const initialState = [
  {
    name: "Charlie",
    score: 40,
  },
  {
    name: "Youssef",
    score: 30,
  },
  {
    name: "Mohamed",
    score: 20,
  },
  {
    name: "Ali",
    score: 27,
  },
  {
    name: "Omar",
    score: 18,
  },
  {
    name: "Khalid",
    score: 23,
  },
  {
    name: "Hassan",
    score: 34,
  },
  {
    name: "Hassan",
    score: 34,
  },
];

export default function RankingCharts() {
  const [data, setData] = useState(initialState);
  useEffect(() => {
    const interval = setInterval(() => {
      setData((prev) => {
        const newData = [...prev];
        newData.forEach((item) => {
          const n = item.score - Math.floor(Math.random() * 10);

          if (n < 0) {
            item.score = 0;
            return;
          }

          item.score = item.score - Math.floor(Math.random() * 10);
        });
        return newData;
      });
    }, 10000);
    return () => clearInterval(interval);
  });

  return (
    <ResponsiveContainer width="100%" height="100%">
      <BarChart width={150} height={40} data={data}>
        <CartesianGrid />
        <Bar dataKey="score" fill="#8884d8" />
        <XAxis dataKey="name" />
        <YAxis />
        <Legend />
      </BarChart>
    </ResponsiveContainer>
  );
}
