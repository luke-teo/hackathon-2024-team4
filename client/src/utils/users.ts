import type { User } from "./types";

export const users: User[] = [
    {
        id: 1,
        name: "Ichiki Kouta",
        avatarUrl:
            "https://stg-dashboard.colorkrew-id.com/avatar/a-61f52c12-aab1-4a08-b4d2-259fc9c239a7-d5a20184-1d68-482f-8e04-cbf08640c4eb?sv=2017-11-09&sr=b&st=2024-06-20T07:55:14Z&se=2024-06-20T08:20:14Z&sp=r&spr=https&sig=0eVD%2FOdOKtoTwNeIQjP%2BwMOjcvKMpsrXFa5JhF1gbE4%3D",
    },
    {
        id: 2,
        name: "James Aparis",
        avatarUrl:
            "https://stg-dashboard.colorkrew-id.com/avatar/a-61f52c12-aab1-4a08-b4d2-259fc9c239a7-50c94534-bdb7-40fd-8698-f6da833499a5?sv=2017-11-09&sr=b&st=2024-06-20T07:56:14Z&se=2024-06-20T08:21:14Z&sp=r&spr=https&sig=dtdZXoSplJ0qPtTy53wPMOLLgbGg5%2Br5Wj%2FDNTct2ao%3D",
    },
    {
        id: 3,
        name: "Evan de Graaff",
        avatarUrl:
            "https://stg-dashboard.colorkrew-id.com/avatar/a-61f52c12-aab1-4a08-b4d2-259fc9c239a7-e430c9de-119f-471a-8602-89adea73dcee?sv=2017-11-09&sr=b&st=2024-06-20T07:55:12Z&se=2024-06-20T08:20:12Z&sp=r&spr=https&sig=ZKy%2BbmBQ0DMBYAZpZG1ZwnFAJEN8uBJesWg4szmlOsE%3D",
    },
    {
        id: 4,
        name: "Luke Teo",
        avatarUrl:
            "https://stg-dashboard.colorkrew-id.com/avatar/a-61f52c12-aab1-4a08-b4d2-259fc9c239a7-225d3213-7501-4bc3-a1ab-360abf74cba5?sv=2017-11-09&sr=b&st=2024-06-20T07:54:29Z&se=2024-06-20T08:19:29Z&sp=r&spr=https&sig=5r854zAT2RFsndy2MOBwez4thsTxN5%2FGpg9bS77mtfc%3D",
    },
    {
        id: 5,
        name: "Nomura Jinya",
        avatarUrl:
            "https://stg-dashboard.colorkrew-id.com/avatar/a-61f52c12-aab1-4a08-b4d2-259fc9c239a7-17dbb396-1d55-4b6f-98cf-8bcfa093610c?sv=2017-11-09&sr=b&st=2024-06-20T07:55:08Z&se=2024-06-20T08:20:08Z&sp=r&spr=https&sig=TbQHqTCCO%2BrzWA9buxdvGVOf2taAVko0ZCJVAUWWWoM%3D",
    },
    {
        id: 6,
        name: "Nisato Tatsuya",
        avatarUrl:
            "https://stg-dashboard.colorkrew-id.com/avatar/a-61f52c12-aab1-4a08-b4d2-259fc9c239a7-48703220-d30b-4f78-8df4-e5bf0065fea5?sv=2017-11-09&sr=b&st=2024-06-20T07:56:52Z&se=2024-06-20T08:21:52Z&sp=r&spr=https&sig=GIqquxTHVWbI%2BJLtrSwUu4enhq8lTZV%2FOYEsReoVcBE%3D",
    },
];

export const getUserInitials = (name: string): string => {
    return name.split(" ")
        .reduce((acc: string, cur, i, src) => {
            if (i === 0) {
                acc += cur[0];
            }
            if (i === src.length - 1) {
                acc += cur[0];
            }
            return acc;
        }, '')
        .toUpperCase()
}
