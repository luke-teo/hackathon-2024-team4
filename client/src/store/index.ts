import { configureStore } from "@reduxjs/toolkit";
import { apiV1 } from "../services/api/v1";

export const store = configureStore({
  reducer: {
    [apiV1.reducerPath]: apiV1.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(apiV1.middleware),
});
