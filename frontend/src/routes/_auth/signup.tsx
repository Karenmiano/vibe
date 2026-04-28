import { createFileRoute, Link } from "@tanstack/react-router";

import styles from "./SignUp.module.css";

export const Route = createFileRoute("/_auth/signup")({
  component: SignUp,
});

function SignUp() {
  return (
    <form className={styles.form}>
      <h1 className={styles.formTitle}>Create an account</h1>

      <div className={styles.formField}>
        <label htmlFor="name">Full name</label>
        <input id="name" name="name" type="text" />
      </div>

      <div className={styles.formField}>
        <label htmlFor="username">Username</label>
        <input id="username" name="username" type="text" />
      </div>

      <div className={styles.formField}>
        <label htmlFor="email">Email</label>
        <input id="email" name="email" type="email" />
      </div>

      <div className={styles.formField}>
        <label htmlFor="password">Password</label>
        <input id="password" name="password" type="password" />
      </div>

      <button type="submit" className={styles.submitButton}>
        Create account
      </button>

      <p className={styles.formFooter}>
        Already have an account? <Link to="/signin">Login</Link>
      </p>
    </form>
  );
}
