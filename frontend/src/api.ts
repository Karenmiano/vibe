import axios from "axios";
import toast from "react-hot-toast";

import { router } from "@/router";

const apiUrl = import.meta.env.VITE_API_URL;

export const api = axios.create({
  baseURL: apiUrl,
});

const GENERIC_ERROR = "An unexpected error occurred. Please try again.";

api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response) {
      switch (error.response.status) {
        case 400:
          console.error("Bad request:", error);
          toast.error(GENERIC_ERROR);
          break;
        case 401:
          if (router.state.location.pathname !== "/signin") {
            router.navigate({
              to: "/signin",
              search: { redir: router.state.location.href },
            });
          }
          break;
        case 403:
          toast.error("You don't have permission to do that.");
          break;
        case 500:
          toast.error("A server error occurred. Please try again later.");
      }
    } else if (error.request) {
      console.error("Network error:", error);
      toast.error(
        "A network error occurred. Please check your connection and try again.",
      );
    } else {
      console.error("Request set up error:", error);
      toast.error(GENERIC_ERROR);
    }

    return Promise.reject(error);
  },
);
