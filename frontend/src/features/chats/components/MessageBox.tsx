import { PaperPlaneRightIcon } from "@phosphor-icons/react";

import styles from "./MessageBox.module.css";

export default function MessageBox() {
  return (
    <form className={styles.messageBox}>
      <input
        type="text"
        placeholder="Type a message"
        name="message"
        className={styles.messageInput}
        required
      />
      <button type="submit" className={styles.sendButton}>
        <PaperPlaneRightIcon size={24} />
      </button>
    </form>
  );
}
