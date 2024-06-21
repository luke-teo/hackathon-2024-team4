import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";

export const baseApiV1 = createApi({
    baseQuery: fetchBaseQuery({
        baseUrl: "",
        credentials: "include",
        prepareHeaders: (headers) => {
            headers.set("Accept", "application/json");
            headers.set("Content-Type", "application/json");
            return headers;
        },
    }),
    endpoints: () => ({}),
});
