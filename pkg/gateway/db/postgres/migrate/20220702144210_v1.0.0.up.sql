BEGIN;
	CREATE TABLE room (
		id          SERIAL PRIMARY KEY,
		external_id uuid NOT NULL,
		name        TEXT NOT NULL
	);
COMMIT;
