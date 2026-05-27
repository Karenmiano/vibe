import {
  createFileRoute,
  Outlet,
  redirect,
  useMatch,
} from "@tanstack/react-router";

import AppBar from "@/ui/AppBar";
import ChatList from "@/features/chats/components/ChatList";
import styles from "./AppLayout.module.css";

export const Route = createFileRoute("/app")({
  beforeLoad: ({ context, location }) => {
    if (!context.auth?.isAuthenticated) {
      throw redirect({ to: "/signin", search: { redir: location.href } });
    }
  },
  component: AppLayout,
});

function AppLayout() {
  const chatOpen = useMatch({ from: "/app/chat/$chatId", shouldThrow: false });

  return (
    <div className={`${styles.container} ${chatOpen ? styles.chatOpen : ""}`}>
      <div className={styles.leadPanel}>
        <AppBar />
        <ChatList />
      </div>
      <Outlet />
    </div>
  );
}
