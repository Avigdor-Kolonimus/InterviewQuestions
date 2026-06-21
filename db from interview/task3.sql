-- Task 1
-- Return all items with the maximum cost
CREATE TABLE example1(
    id    serial,
    title text,
    cost  int
);

INSERT INTO example1(title, cost)
VALUES ('item1', 10),
       ('item2', 12),
       ('item3', 15),
       ('item4', 10),
       ('item5', 15);

-- ===========================================
-- Answer
SELECT title, cost
FROM example1
WHERE cost = (SELECT MAX(cost) FROM example1);

-- ===========================================
-- Task 2
-- Return the latest comment from each author (based on the comment date)
CREATE TABLE example2(
    id      serial,
    comment text,
    author  text,
    date    timestamp with time zone
);

INSERT INTO example2(comment, author, date)
VALUES ('aaa', 'ivanov', '2024-03-18 09:45:53 +00:00'),
       ('bbb', 'petrov', '2024-03-16 07:24:53 +00:00'),
       ('ccc', 'ivanov', '2024-03-20 17:01:00 +00:00'),
       ('ddd', 'petrov', '2024-03-18 09:45:53 +00:00'),
       ('eee', 'sidorov', '2024-03-19 11:21:53 +00:00');

-- ===========================================
-- Answer
SELECT e.comment, e.author, e.date
FROM example2 as e
JOIN (
    SELECT author, MAX(date) AS max_date
    FROM example2
    GROUP BY author
) t
ON e.author = t.author AND e.date = t.max_date;

-- PostgreSQL
SELECT id, comment, author, date
FROM (
    SELECT id, comment, author, date,
           ROW_NUMBER() OVER (
               PARTITION BY author
               ORDER BY date DESC
           ) AS rn
    FROM example2
) t
WHERE rn = 1;