BEGIN;
	CREATE TABLE room (
		id          SERIAL PRIMARY KEY,
		external_id uuid NOT NULL,
		name        TEXT NOT NULL
	);

	CREATE TABLE message (
		id          SERIAL PRIMARY KEY,
		external_id uuid NOT NULL,
		content     TEXT NOT NULL,
		author      TEXT NOT NULL,
		created_at   TIMESTAMP NOT NULL,
		room_id     INT NOT NULL,

		CONSTRAINT fk_room
			FOREIGN KEY (room_id)
				REFERENCES room(id)
	);
COMMIT;
