import { useMutation } from "@tanstack/react-query";
import axios from "axios";
import { useNavigate } from "react-router-dom";
import { toast } from "sonner";

async function register(params: {
  name: string;
  family_name: string;
  email: string;
  password: string;
}) {
  const apiUrl = import.meta.env.VITE_API_URL as string;

  const response = await axios.post(`${apiUrl}/api/auth/signup`, params);

  return response.data;
}

export function useRegister() {
  const navigate = useNavigate();

  return useMutation({
    mutationKey: ["register"],
    mutationFn: register,
    onSuccess() {
      toast.success("Creation de compte avec succès");
      navigate("/");
    },
    onError() {
      toast.error("Erreur lors de la création du compte");
    },
  });
}
