import styles from "./ChatThreadHeader.module.css";
import { recipient } from "@/mocks/chatThread";

export default function ChatThreadHeader() {
  return (
    <header>
      <div className={styles.recipient}>
        <img
          src={recipient.avatar}
          alt={recipient.name}
          className={styles.avatar}
        />
        <h2 className={styles.recipientName}>{recipient.name}</h2>
      </div>
    </header>
  );
}
