import { createFileRoute, Outlet, useMatch } from "@tanstack/react-router";

import ChatList from "@/features/chats/components/ChatList";
import styles from "./AppLayout.module.css";

export const Route = createFileRoute("/app")({
  component: AppLayout,
});

function AppLayout() {
  const chatOpen = useMatch({ from: "/app/chat/$chatId", shouldThrow: false });

  return (
    <div className={`${styles.container} ${chatOpen ? styles.chatOpen : ""}`}>
      <ChatList />
      <Outlet />
    </div>
  );
}
