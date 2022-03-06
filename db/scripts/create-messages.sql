CREATE TABLE IF NOT EXISTS messages (
	id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
	tenant_id uuid NOT NULL,
	user_id uuid NOT NULL,
	contents VARCHAR NOT NULL,
	created_at TIMESTAMP DEFAULT NOW(),
	FOREIGN KEY (tenant_id)
		REFERENCES tenants (id),
	FOREIGN KEY (user_id)
		REFERENCES users (id)
);
