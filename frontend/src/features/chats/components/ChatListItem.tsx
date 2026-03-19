import { useNavigate } from "@tanstack/react-router";

import styles from "./ChatListItem.module.css";

type ChatListItemProps = {
  chat: {
    avatar: string;
    recentMessage: string;
    name: string;
    timestamp: string;
    unreadCount: number;
  };
};

export default function ChatListItem({ chat }: ChatListItemProps) {
  const navigate = useNavigate();

  return (
    <li
      onClick={() =>
        navigate({ to: "/app/chat/$chatId", params: { chatId: "1" } })
      }
    >
      <img
        src={chat.avatar}
        alt={`${chat.name}'s avatar`}
        className={styles.avatar}
      />
      <div className={styles.chatContent}>
        <div>
          <p>{chat.name}</p>
          <p className={styles.timestamp}>{chat.timestamp}</p>
        </div>
        <div>
          <p className={styles.recentMessage}>{chat.recentMessage}</p>
          <div className={styles.unreadCount}>{chat.unreadCount}</div>
        </div>
      </div>
    </li>
  );
}
