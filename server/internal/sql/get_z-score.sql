WITH main_stats AS (
    SELECT
	user_id,
        AVG(score) AS mean,
        STDDEV(score) AS stddev
    FROM
        user_behavior
    WHERE
        desinated_date > CURRENT_DATE - interval '1 days'*$1 -- Desinated By Backend
    GROUP BY 
	    user_id
)
SELECT
    ub.user_id,
    ub.date,
    ub.score,
    main_stats.mean,
    main_stats.stddev,
    (ub.score - main_stats.mean) / main_stats.stddev AS z_score
FROM
   user_behavior ub
   JOIN main_stats ON ub.user_id = main_stats.user_id
WHERE
        desinated_date > CURRENT_DATE - interval '1 days'*$1 -- Desinated By Backend
ORDER BY
    ub.user_id, ub.date;