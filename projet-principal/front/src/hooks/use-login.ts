import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { toast } from "sonner";
import { useLocalStorage } from "@uidotdev/usehooks";
import { useContext } from "react";
import { DialogContext } from "@/contexts/dialiog-context";

async function login(params: { email: string; password: string }) {
  const apiUrl = import.meta.env.VITE_API_URL as string;

  const response = await axios.post(`${apiUrl}/api/auth/signin`, params);

  return response.data;
}

export function useLogin() {
  const navigate = useNavigate();
  const [, setAccessToken] = useLocalStorage<string | null>(
    "accessToken",
    null
  );
  const { toggleLoginDialogOpen } = useContext(DialogContext);

  return useMutation({
    mutationKey: ["login"],
    mutationFn: login,
    onSuccess({ token }) {
      toast.success("Connexion avec succès");
      setAccessToken(token);
      toggleLoginDialogOpen();
      navigate("/");
    },
    onError() {
      toast.error("Erreur lors de la connexion");
    },
  });
}
