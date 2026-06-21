-- Write a SELECT query that produces the output shown below

CREATE TABLE test (
    id SERIAL,
    name VARCHAR
);

-- A regular SELECT would return:
-- | id | name  |
-- | 1  | name1 |
-- | 2  | name2 |
-- | 3  | name3 |
-- | 4  | name4 |
-- | 5  | name5 |

-- Expected output:
-- | id | name  |
-- | 1  | name1 |
-- | 3  | name3 |
-- | 2  | name2 |
-- | 4  | name4 |
-- | 5  | name5 |

-- ===========================================
-- Answer
SELECT id, name
FROM test
ORDER BY
    CASE
        WHEN id = 2 THEN 3
        WHEN id = 3 THEN 2
        ELSE id
    END;