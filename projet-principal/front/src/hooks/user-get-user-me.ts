import { useQuery } from "@tanstack/react-query";
import axios from "axios";

async function getUserMe(access_token: string) {
  const apiUrl = import.meta.env.VITE_API_URL as string;

  const { data } = await axios.get(`${apiUrl}/api/app/user/me`, {
    headers: {
      Authorization: `Bearer ${access_token}`,
    },
  });
  const { id } = data;

  const response = await axios.get(`${apiUrl}/api/app/user/${id}`, {
    headers: {
      Authorization: `Bearer ${access_token}`,
    },
  });

  return response.data;
}

export default function useUserMe(accessToken: string | null) {
  return useQuery({
    queryKey: ["user-me"],
    queryFn: () => getUserMe(accessToken as string),
    enabled: !!accessToken,
  });
}
