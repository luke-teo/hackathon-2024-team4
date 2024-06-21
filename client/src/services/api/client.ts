import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const baseApiV1 = createApi({
  baseQuery: fetchBaseQuery({
    baseUrl: "/api",
    credentials: "include",
    prepareHeaders: (headers) => {
      headers.set("Accept", "application/json");
      headers.set("Content-Type", "application/json");
      headers.set("ngrok-skip-browser-warning", "21312");
      return headers;
    },
  }),
  endpoints: () => ({}),
});
