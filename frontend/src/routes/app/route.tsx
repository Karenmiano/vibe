import { createFileRoute, Outlet } from "@tanstack/react-router";

import ChatList from "@/features/chats/components/ChatList";
import styles from "./AppLayout.module.css";

export const Route = createFileRoute("/app")({
  component: AppLayout,
});

function AppLayout() {
  return (
    <div className={styles.container}>
      <ChatList />
      <Outlet />
    </div>
  );
}
