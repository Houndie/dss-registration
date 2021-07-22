CREATE TABLE discount_bundles(
	id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
	created_at timestamp with time zone NOT NULL DEFAULT now(),
	code text NOT NULL UNIQUE
);

CREATE TYPE applied_to_type AS ENUM ('Full Weekend','Dance Only','Mix And Match','Solo Jazz','Team Competition','TShirt');

CREATE TABLE discounts(
	id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
	created_at timestamp with time zone NOT NULL DEFAULT now(),
	discount_bundle_id uuid NOT NULL REFERENCES discount_bundles(id) ON DELETE CASCADE,
	applied_to applied_to_type NOT NULL,
	name text NOT NULL
);
