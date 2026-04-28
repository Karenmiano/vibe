import { createFileRoute, Link } from "@tanstack/react-router";

import styles from "./Auth.module.css";

export const Route = createFileRoute("/_auth/signin")({
  component: SignIn,
});

function SignIn() {
  return (
    <form className={styles.form}>
      <h1 className={styles.formTitle}>Welcome back!</h1>

      <div className={styles.formField}>
        <label htmlFor="identifier">Username or Email</label>
        <input id="identifier" name="identifier" type="text" />
      </div>

      <div className={styles.formField}>
        <label htmlFor="password">Password</label>
        <input id="password" name="password" type="password" />
      </div>

      <button type="submit" className={styles.submitButton}>
        Sign in
      </button>

      <p className={styles.formFooter}>
        Don't have an account? <Link to="/signup">Sign up here</Link>
      </p>
    </form>
  );
}
