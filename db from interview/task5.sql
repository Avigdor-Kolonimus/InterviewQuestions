-- Build the optimal index for the following query:
SELECT *
FROM employee
WHERE sex = 'm'
  AND salary > 300000
  AND age = 20
ORDER BY created_at;

-- (P.S. A single answer is not enough here — explain your reasoning and justify your choice.)

-- ===========================================
-- Answer
CREATE INDEX idx_employee ON employee (age, salary, created_at);

-- The most likely candidate is a composite B-tree index on (age, sex, salary, created_at) or (age, salary, created_at), 
-- depending on the selectivity of the sex column.