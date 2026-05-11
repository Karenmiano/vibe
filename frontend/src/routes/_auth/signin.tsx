import axios from "axios";
import { z } from "zod";
import { createFileRoute, Link, useNavigate } from "@tanstack/react-router";
import { useForm, type FieldPath } from "react-hook-form";

import InputError from "@/ui/InputError";
import { api } from "@/api";
import styles from "./Auth.module.css";

export const Route = createFileRoute("/_auth/signin")({
  component: SignIn,
  validateSearch: z.object({
    redir: z.string().default("/app").catch("/app"),
  }),
});

type SignInFormData = {
  identifier: string;
  password: string;
};

function SignIn() {
  const {
    register,
    handleSubmit,
    formState: { isSubmitting, errors },
    setError,
  } = useForm<SignInFormData>();

  const navigate = useNavigate();
  const { redir } = Route.useSearch();
  
  async function onSubmit(data: SignInFormData) {
    if (isSubmitting) return;

    try {
      await api.post("/login", data);

      // on success, redirect to the page the user originally wanted to visit, or to the app home
      navigate({ to: redir });
    } catch (error) {
      if (axios.isAxiosError(error)) {
        if (error.response) {
          const errorData = error.response.data;
          switch (error.response.status) {
            case 422:
              Object.entries(errorData).forEach(([field, error]) =>
                setError(field as FieldPath<SignInFormData>, {
                  message: error as string,
                }),
              );
              break;
            case 401:
              setError("root.authError", { message: errorData.message });
          }
        }
      }
    }
  }

  return (
    <form className={styles.form} onSubmit={handleSubmit(onSubmit)}>
      <h1 className={styles.formTitle}>Welcome back!</h1>

      {errors.root?.authError && (
        <div style={{ textAlign: "center", marginBottom: "1rem" }}>
          <InputError>{errors.root.authError.message}</InputError>
        </div>
      )}

      <div className={styles.formField}>
        <label htmlFor="identifier">Username or Email</label>
        <input
          id="identifier"
          type="text"
          {...register("identifier", {
            required: "Please enter your username or email",
          })}
        />

        {errors.identifier && (
          <InputError>{errors.identifier.message}</InputError>
        )}
      </div>

      <div className={styles.formField}>
        <label htmlFor="password">Password</label>
        <input
          id="password"
          type="password"
          {...register("password", { required: "Password is required" })}
        />

        {errors.password && <InputError>{errors.password.message}</InputError>}
      </div>

      <button
        type="submit"
        className={styles.submitButton}
        disabled={isSubmitting}
      >
        {isSubmitting ? (
          <div className={`loader ${styles.centerLoader}`}></div>
        ) : (
          "Sign in"
        )}
      </button>

      <p className={styles.formFooter}>
        Don't have an account?{" "}
        <Link to="/signup" className={styles.linkAccent}>
          Sign up here
        </Link>
      </p>
    </form>
  );
}
