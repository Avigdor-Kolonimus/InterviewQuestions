-- You are given three tables: users, purchases, and banlist.
-- Write two queries based on these tables.
--
-- Query 1:
-- Select users and the IDs of purchases they made before being banned,
-- or all purchases if the user has never been banned.
--
-- Query 2:
-- Select users whose total purchase amount exceeds 5,000.

CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE purchases (
    purchase_id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    amount DECIMAL(10, 2) NOT NULL,
    purchase_date TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE banlist (
    user_id INT PRIMARY KEY,
    ban_date TIMESTAMP NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);
---------------------------------------------------
-- ANSWER
---------------------------------------------------
--Query 1
SELECT u.user_id, u.name, p.purchase_id
FROM users u
JOIN purchases p ON u.user_id = p.user_id
LEFT JOIN banlist b ON u.user_id = b.user_id
WHERE b.ban_date IS NULL OR p.purchase_date < b.ban_date;

--Query 2
SELECT u.user_id, u.name, SUM(p.amount) AS total_spent
FROM users u
JOIN purchases p ON u.user_id = p.user_id
GROUP BY u.user_id, u.name
HAVING SUM(p.amount) > 5000;