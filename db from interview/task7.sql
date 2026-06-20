CREATE TABLE customer(
    id      INTEGER PRIMARY KEY,
    email   VARCHAR(100) NOT NULL,
    country CHAR(2) NOT NULL
);

CREATE TABLE cart_item(
    id          INTEGER PRIMARY KEY,
    customer_id INTEGER NOT NULL,
    title       VARCHAR(20) NOT NULL,
    amount      INTEGER NOT NULL,
    price       INTEGER NOT NULL
);

-- Return all customers row by row (id, email)
-- together with all items in their shopping cart (title, amount)
-- ===========================================
-- Answer
SELECT c.id, c.email, ci.title, ci.amount
FROM customer as c
LEFT JOIN cart_item as ci ON c.id = ci.customer_id

-- Return the top 10 customers (id, email)
-- by the total value of items in their shopping cart
-- ===========================================
-- Answer
SELECT c.id, c.email, SUM(ci.amount * ci.price) AS total
FROM customer as c
JOIN cart_item as ci ON c.id=ci.customer_id
GROUP BY c.id, c.email
ORDER BY total DESC
LIMIT 10;

-- Only customers from Russia are of interest,
-- whose cart contains goods worth at least 1000 RUB
-- ===========================================
-- Answer
SELECT c.id, c.email, SUM(ci.amount * ci.price) as total
FROM customer as c
JOIN cart_item as ci ON c.id=ci.customer_id
WHERE c.country = 'ru'
GROUP BY c.id, c.email
HAVING total>=1000;

-- Will this index be used?
CREATE INDEX myIdx ON carts (sku, country, customer_id);

SELECT *
FROM carts
WHERE sku = 192 AND country = 'ru';
-- ===========================================
-- Answer
-- The first query can use the composite index because it filters on the leading columns (sku, country).

-- What about this query?
CREATE INDEX myIdx ON carts (sku, country, customer_id);

SELECT *
FROM carts
WHERE country = 'ru'AND customer_id = 10;
-- ===========================================
-- Answer
-- The second query cannot efficiently use the index because it skips the leading column (sku). 
-- B-tree indexes are ordered by their leftmost columns, so predicates on country and customer_id alone do not match the index structure.