WITH past_stats AS (
    SELECT
        user_id,
        AVG(score) AS mean
    FROM
        user_behavior
    WHERE
        EXTRACT(DOW FROM date) = $1 -- 0: Sunday,....
    GROUP BY 
        user_id
)
SELECT
    ub.user_id,
    ub.date,
    ub.score,
    past_stats.mean
FROM
    user_behavior ub
    JOIN past_stats ON ub.user_id = past_stats.user_id
WHERE
    EXTRACT(DOW FROM ub.date) = $1 -- 0: Sunday,....
ORDER BY 
    ub.user_id, ub.date;