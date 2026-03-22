import { createFileRoute } from "@tanstack/react-router";

import ChatThreadHeader from "@/features/chats/components/ChatThreadHeader";
import styles from "./ChatThread.module.css";
import { conversation } from "@/mocks/chatThread";
import Message from "@/features/chats/components/Message";

export const Route = createFileRoute("/app/chat/$chatId")({
  component: ChatThread,
});

function ChatThread() {
  return (
    <main className={styles.chatThread}>
      <ChatThreadHeader />
      <ul role="log" aria-live="polite" className={styles.feed}>
        {conversation.map((message) => (
          <Message message={message} />
        ))}
      </ul>
      <form className={styles.messageForm}>
        <input type="text" placeholder="Type a message" name="message" />
      </form>
    </main>
  );
}
