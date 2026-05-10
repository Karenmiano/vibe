import styles from "./InputError.module.css";

export default function InputError({
  children,
}: {
  children: React.ReactNode;
}) {
  return <p className={styles.error}>{children}</p>;
}
