import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import LoginForm from "../forms/login-form";
import { Link } from "react-router-dom";
import { useContext } from "react";
import { DialogContext } from "@/contexts/dialiog-context";

export default function LoginDialog() {
  const { loginDialogOpen, toggleLoginDialogOpen } = useContext(DialogContext);

  return (
    <Dialog open={loginDialogOpen} onOpenChange={toggleLoginDialogOpen}>
      <DialogTrigger>Connecter</DialogTrigger>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Se connecter</DialogTitle>
        </DialogHeader>
        <LoginForm />
        <p>
          Pas encore de compte ?{" "}
          <Link to="/inscrire" className="underline">
            S'inscrire
          </Link>
        </p>
      </DialogContent>
    </Dialog>
  );
}
