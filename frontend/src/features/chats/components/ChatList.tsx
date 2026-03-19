import { chats } from "@/mocks/chats";

import ChatListItem from "./ChatListItem";
import styles from "./ChatList.module.css";

export default function ChatList() {
  return (
    <nav className={styles.chats}>
      <header>
        <h2>Chats</h2>
      </header>

      <ul className={styles.chatFilters}>
        <li className="active">All</li>
        <li>Personal</li>
        <li>Groups</li>
      </ul>

      <ul className={styles.chatList}>
        {chats.map((chat) => (
          <ChatListItem chat={chat} />
        ))}
      </ul>
    </nav>
  );
}
