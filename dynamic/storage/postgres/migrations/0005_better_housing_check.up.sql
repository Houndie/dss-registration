ALTER TABLE registrations DROP CONSTRAINT registrations_check4;

ALTER TABLE registrations ADD CONSTRAINT registrations_check4 CHECK ((housing = 'Provide' AND require_housing_pet_allergies = '' AND require_housing_details = '') OR (housing = 'Require' AND provide_housing_pets = '' AND provide_housing_quantity = 0 AND provide_housing_details = '') OR (housing = 'No Housing' AND provide_housing_pets = '' AND provide_housing_quantity = 0 AND provide_housing_details = '' AND require_housing_pet_allergies = '' AND require_housing_details = ''));
