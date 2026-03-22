import styles from "./Message.module.css";
type MessageProps = {
  message: {
    content: string;
    senderId: string;
    createdAt: string;
    senderName: string;
    senderAvatar: string;
  };
};

export default function Message({ message }: MessageProps) {
  return message.senderId === "pixel-panda" ? (
    <li className={`${styles.msg} ${styles.outgoingMessage}`}>
      <p>{message.content}</p>
      <p className={styles.timestamp}>{message.createdAt}</p>
    </li>
  ) : (
    <li className={`${styles.msg} ${styles.incomingMessage}`}>
      {/* {message.senderAvatar && (
        <img src={message.senderAvatar} alt={message.senderName} />
      )} */}
      {message.senderName && (
        <p className={styles.senderName}>{message.senderName}</p>
      )}
      <p>{message.content}</p>
      <p className={styles.timestamp}>{message.createdAt}</p>
    </li>
  );
}
