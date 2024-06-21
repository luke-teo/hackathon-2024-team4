import { baseApiV1 as api } from "./client";
export const addTagTypes = [] as const;
const injectedRtkApi = api
  .enhanceEndpoints({
    addTagTypes,
  })
  .injectEndpoints({
    endpoints: (build) => ({
      getScoresByUserId: build.query<
        GetScoresByUserIdApiResponse,
        GetScoresByUserIdApiArg
      >({
        query: (queryArg) => ({
          url: `/scores/${queryArg.userId}`,
          body: queryArg.body,
          params: { startDate: queryArg.startDate, endDate: queryArg.endDate },
        }),
      }),
    }),
    overrideExisting: false,
  });
export { injectedRtkApi as apiV1 };
export type GetScoresByUserIdApiResponse = /** status 200 OK */ {
  userId: string;
  scores: Score[];
};
export type GetScoresByUserIdApiArg = {
  /** User ID */
  userId: string;
  /** Start date */
  startDate: string;
  /** End date */
  endDate: string;
  body: Blob;
};
export type Score = {
  date: string;
  currentScore: number;
  pastAverageScore: number;
};
export const { useGetScoresByUserIdQuery, useLazyGetScoresByUserIdQuery } =
  injectedRtkApi;
