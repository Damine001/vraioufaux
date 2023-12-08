import { HTMLProps, useContext } from "react";
import { ChevronDown } from "lucide-react";
import { cn } from "@/lib/utils";
import { Link } from "react-router-dom";
import LoginDialog from "./dialogs/login-dialog";
import { UserContext } from "@/contexts/user-context";

interface Props extends HTMLProps<HTMLDivElement> {}

export default function Navbar({ className }: Props) {
  function handleScrollToNextSection() {
    window.scrollTo({
      top: window.innerHeight,
      behavior: "smooth",
    });
  }

  const { user } = useContext(UserContext);

  return (
    <div
      className={cn(
        className,
        "container flex items-center backdrop-blur-xl border-b border-gray-300 z-50"
      )}
    >
      <nav className="flex-[3] flex space-x-4 justify-start">
        <Link to="/about-us">About us</Link>
        <Link to="/deadlines">Deadlines</Link>
        <Link to="/posts">Posts</Link>
      </nav>
      <ChevronDown
        size={35}
        className="flex-1 mt-1 animate-bounce cursor-pointer"
        onClick={handleScrollToNextSection}
      />
      <nav className="flex-[3] flex space-x-4  justify-end">
        <Link to="/leader-board">Leader board</Link>
        <Link to="/boite-a-idees">Boite à idées</Link>
        {user ? <span>{user.name}</span> : <LoginDialog />}
      </nav>
    </div>
  );
}
