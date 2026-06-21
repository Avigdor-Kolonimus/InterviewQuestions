CREATE TABLE holds (
    id         TEXT PRIMARY KEY,
    account_id BIGINT NOT NULL REFERENCES accounts(id),
    amount     BIGINT NOT NULL
);

CREATE TABLE accounts (
    id      BIGSERIAL PRIMARY KEY,
    amount  BIGINT NOT NULL -- account balance excluding held funds
);

-- POST /accounts/:id/hold?amount=10

BEGIN;

SELECT amount
FROM accounts
WHERE id = :id;

-- Check that the account has sufficient funds
-- Calculate the new balance

UPDATE accounts
SET amount = 90
WHERE id = :id;

INSERT INTO holds (...);

COMMIT;ROLLBACK;

-- What problems can occur with this implementation?
-- How would you fix them?

-- =================================================================================================================================
-- ============================================================ Answer =============================================================
-- =================================================================================================================================

-- 1. Problem: Transaction conflicts (race conditions)
-- If two requests call POST /accounts/:id/hold simultaneously,
-- both transactions may read the same value from accounts.amount.
-- This can lead to a negative balance or inconsistent data.

-- Solution: Use row-level locking
-- Use SELECT ... FOR UPDATE to lock the row until the transaction completes.

BEGIN;

-- Lock the row to prevent modifications by other transactions
SELECT amount
FROM accounts
WHERE id = :id
FOR UPDATE;

-- Update the balance and create a hold
UPDATE accounts
SET amount = amount - 10
WHERE id = :id;

INSERT INTO holds (id, account_id, amount)
VALUES (:uuid, :id, 10);

COMMIT;


-- 2. Problem: Missing data integrity validation
-- The query does not verify whether the account has sufficient funds
-- to create the hold. This may result in a negative balance.

-- Solution: Add business logic validation
-- The available balance should be checked before updating accounts.amount.

-- Check available funds
IF (amount - 10) < 0 THEN
    ROLLBACK;
    RETURN 'Insufficient funds';
END IF;


-- 3. Problem: Potential deadlocks
-- If multiple transactions access different rows in accounts and holds
-- concurrently, deadlocks may occur.

-- Solution: Enforce a consistent operation order
-- Ensure that operations are always executed in the same order:
-- SELECT, then UPDATE, then INSERT.


-- Final version
-- With the improvements above, the code could look like this:

BEGIN;

-- Lock the row for modification
SELECT amount
FROM accounts
WHERE id = :id
FOR UPDATE;

-- Check available funds
IF (amount - :hold_amount) < 0 THEN
    ROLLBACK;
    RETURN 'Insufficient funds';
END IF;

-- Update the balance and create a hold
UPDATE accounts
SET amount = amount - :hold_amount
WHERE id = :id;

INSERT INTO holds (id, account_id, amount)
VALUES (gen_random_uuid(), :id, :hold_amount);

COMMIT;