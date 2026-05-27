import styles from "./Logo.module.css";

export default function Logo() {
  return (
    <div className={styles.logo}>
      <span className={styles.square}>V</span>
      <span>Vibe</span>
    </div>
  );
}
