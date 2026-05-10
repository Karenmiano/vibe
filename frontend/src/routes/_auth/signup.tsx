import { createFileRoute, Link, useNavigate } from "@tanstack/react-router";
import { useForm, type FieldPath } from "react-hook-form";

import InputError from "@/ui/InputError";
import styles from "./Auth.module.css";
import { api } from "@/api";
import axios from "axios";
import toast from "react-hot-toast";

export const Route = createFileRoute("/_auth/signup")({
  component: SignUp,
});

type SignUpFormData = {
  fullName: string;
  username: string;
  email: string;
  password: string;
};

function SignUp() {
  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    setError,
  } = useForm<SignUpFormData>({
    mode: "onBlur",
  });

  const navigate = useNavigate();

  async function onSubmit(data: SignUpFormData) {
    if (isSubmitting) return;

    try {
      await api.post("/register", data);
      toast.success("Account created successfully");

      await api.post("/login", {
        identifier: data.username,
        password: data.password,
      });
      navigate({ to: "/app" });
    } catch (error) {
      if (axios.isAxiosError(error)) {
        if (error.response) {
          const errorData = error.response.data;
          switch (error.response.status) {
            case 422:
              Object.entries(errorData).forEach(([field, error]) =>
                setError(field as FieldPath<SignUpFormData>, {
                  message: error as string,
                }),
              );
              break;
            case 409: // email or username already taken
              Object.entries(errorData).forEach(([field, error]) =>
                setError(field as FieldPath<SignUpFormData>, {
                  type: "already_taken",
                  message: error as string,
                }),
              );
              break;
          }
        }
      }
    }
  }

  return (
    <form className={styles.form} onSubmit={handleSubmit(onSubmit)}>
      <h1 className={styles.formTitle}>Create an account</h1>

      <div className={styles.formField}>
        <label htmlFor="name">Full name</label>
        <input
          id="name"
          type="text"
          {...register("fullName", {
            required: "Full name is required",
            validate: (value) => {
              if (value.trim() === "") {
                return "Full name is required";
              }
              return true;
            },
          })}
        />

        {errors.fullName && <InputError>{errors.fullName.message}</InputError>}
      </div>

      <div className={styles.formField}>
        <label htmlFor="username">Username</label>
        <input
          id="username"
          type="text"
          {...register("username", {
            required: "Username is required",
            minLength: {
              value: 3,
              message: "Username must be 3-16 characters",
            },
            maxLength: {
              value: 16,
              message: "Username must be 3-16 characters",
            },
            pattern: {
              value: /^[a-zA-Z0-9._]+$/,
              message:
                "Username must contain only letters, numbers, dots (.) and underscores (_)",
            },
          })}
        />

        {errors.username && <InputError>{errors.username.message}</InputError>}
      </div>

      <div className={styles.formField}>
        <label htmlFor="email">Email</label>
        <input
          id="email"
          type="email"
          {...register("email", {
            required: "Email is required",
          })}
        />

        {errors.email && <InputError>{errors.email.message}</InputError>}
      </div>

      <div className={styles.formField}>
        <label htmlFor="password">Password</label>
        <input
          id="password"
          type="password"
          {...register("password", {
            required: "Password is required",
            minLength: {
              value: 6,
              message: "Password must be 6-72 characters",
            },
            maxLength: {
              value: 72,
              message: "Password must be 6-72 characters",
            },
          })}
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
          "Create account"
        )}
      </button>

      <p className={styles.formFooter}>
        Already have an account?{" "}
        <Link to="/signin" className={styles.linkAccent}>
          Log in
        </Link>
      </p>
    </form>
  );
}
