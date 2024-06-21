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
          params: { startDate: queryArg.startDate, endDate: queryArg.endDate },
        }),
      }),
      postUploadCsv: build.mutation<
        PostUploadCsvApiResponse,
        PostUploadCsvApiArg
      >({
        query: (queryArg) => ({
          url: `/upload_csv`,
          method: "POST",
          body: queryArg.body,
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
};
export type PostUploadCsvApiResponse = /** status 200 OK */ void;
export type PostUploadCsvApiArg = {
  body: {
    filename?: string;
    file?: Blob;
  };
};
export type Score = {
  date: string;
  currentScore: number;
  mean: number;
  standardDeviation: number;
  zScore: number;
};
export const {
  useGetScoresByUserIdQuery,
  useLazyGetScoresByUserIdQuery,
  usePostUploadCsvMutation,
} = injectedRtkApi;
