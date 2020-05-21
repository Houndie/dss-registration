CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;

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

CREATE TABLE admins(
	id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
	created_at timestamp with time zone NOT NULL DEFAULT now(),
	user_id text NOT NULL
);

CREATE TYPE pass_type_type AS ENUM ('Full Weekend', 'Dance Only', 'No Pass');
CREATE TYPE full_weekend_level_type AS ENUM ('Level 1', 'Level 2', 'Level 3');
CREATE TYPE full_weekend_tier_type AS ENUM ('Tier 1', 'Tier 2', 'Tier 3', 'Tier 4', 'Tier 5');
CREATE TYPE mix_and_match_role_type AS ENUM ('Leader', 'Follower');
CREATE TYPE tshirt_style_type AS ENUM ('Unisex S', 'Unisex M', 'Unisex L', 'Unisex XL', 'Unisex 2XL', 'Unisex 3XL', 'Bella S', 'Bella M', 'Bella L', 'Bella XL', 'Bella 2XL');
CREATE TYPE housing_type AS ENUM ('Provide', 'Require', 'No Housing');

CREATE TABLE registrations(
	id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
	created_at timestamp with time zone NOT NULL DEFAULT now(),
	first_name text NOT NULL,
	last_name text NOT NULL,
	street_address text NOT NULL DEFAULT '',
	city text NOT NULL DEFAULT '',
	state text NOT NULL DEFAULT '',
	zip_code text NOT NULL DEFAULT '',
	email text NOT NULL DEFAULT '',
	home_scene text NOT NULL DEFAULT '',
	is_student boolean NOT NULL DEFAULT FALSE,
	pass_type pass_type_type NOT NULL DEFAULT 'No Pass',
	full_weekend_level full_weekend_level_type,
	full_weekend_tier full_weekend_tier_type,
	mix_and_match boolean NOT NULL DEFAULT false,
	mix_and_match_role mix_and_match_role_type,
	solo_jazz boolean NOT NULL DEFAULT false,
	team_competition boolean NOT NULL DEFAULT false,
	team_competition_name text NOT NULL DEFAULT '',
	tshirt boolean NOT NULL DEFAULT false,
	tshirt_style tshirt_style_type,
	housing housing_type NOT NULL DEFAULT 'No Housing',
	provide_housing_pets text NOT NULL DEFAULT '',
	provide_housing_quantity int NOT NULL DEFAULT 0,
	provide_housing_details text NOT NULL DEFAULT '',
	require_housing_pet_allergies text NOT NULL DEFAULT '',
	require_housing_details text NOT NULL DEFAULT '',
	user_id text NOT NULL DEFAULT '',
	order_ids text[] NOT NULL DEFAULT '{}',
	discount_codes text[] NOT NULL DEFAULT '{}',
	enabled boolean NOT NULL DEFAULT true,
	CHECK ((pass_type = 'Full Weekend' AND full_weekend_level IS NOT NULL AND full_weekend_tier IS NOT NULL) OR (pass_type != 'Full Weekend' AND full_weekend_level IS NULL AND full_weekend_tier IS NULL)),
	CHECK ((mix_and_match AND mix_and_match_role IS NOT NULL) OR (NOT mix_and_match AND mix_and_match_role IS NULL)),
	CHECK ((team_competition AND team_competition_name != '') OR (NOT team_competition AND team_competition_name = '')),
	CHECK ((tshirt AND tshirt_style IS NOT NULL) OR (NOT tshirt AND tshirt_style IS NULL)),
	CHECK ((housing = 'Provide' AND provide_housing_pets != '' AND provide_housing_quantity != 0 AND provide_housing_details != '' AND require_housing_pet_allergies = '' AND require_housing_details = '') OR (housing = 'Require' AND provide_housing_pets = '' AND provide_housing_quantity = 0 AND provide_housing_details = '' AND require_housing_pet_allergies != '' AND require_housing_details != '') OR (housing = 'No Housing' AND provide_housing_pets = '' AND provide_housing_quantity = 0 AND provide_housing_details = '' AND require_housing_pet_allergies = '' AND require_housing_details = ''))
);
