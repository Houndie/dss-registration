ALTER TABLE registrations
	ADD COLUMN pass_manually_paid boolean NOT NULL DEFAULT false,
	ADD COLUMN mix_and_match_manually_paid boolean NOT NULL DEFAULT false,
	ADD COLUMN solo_jazz_manually_paid boolean NOT NULL DEFAULT false,
	ADD COLUMN team_competition_manually_paid boolean NOT NULL DEFAULT false,
	ADD COLUMN tshirt_manually_paid boolean NOT NULL DEFAULT false;
