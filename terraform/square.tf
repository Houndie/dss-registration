provider "square" {
	access_token = var.square_access_token
}

resource "square_catalog_object" "team_competition" {
	type = "ITEM"

	item_data {
		name = "Team Competition"
	}
}

resource "square_catalog_object" "team_competition_variation" {
	type = "ITEM_VARIATION"

	item_variation_data {
		item_id = square_catalog_object.team_competition.id
		name = "Regular"
		pricing_type = "FIXED_PRICING"

		price_money {
			amount = 5500
			currency = "USD"
		}
	}
}

locals {
	mix_and_match_variants = ["Leader", "Follower"]
}

resource "square_catalog_object" "mix_and_match" {
	type = "ITEM"
	
	item_data {
		name = "Mix and Match"
	}
}

resource "square_catalog_object" "mix_and_match_variation" {
	for_each = toset(local.mix_and_match_variants)

	type = "ITEM_VARIATION"

	item_variation_data {
		item_id = square_catalog_object.mix_and_match.id
		name = each.value
		pricing_type = "FIXED_PRICING"

		price_money {
			amount = 500
			currency = "USD"
		}
	}
}

resource "square_catalog_object" "solo_jazz" {
	type = "ITEM"
	
	item_data {
		name = "Solo Jazz"
	}
}

resource "square_catalog_object" "solo_jazz_variation" {
	type = "ITEM_VARIATION"

	item_variation_data {
		item_id = square_catalog_object.solo_jazz.id
		name = "Regular"
		pricing_type = "FIXED_PRICING"

		price_money {
			amount = 500
			currency = "USD"
		}
	}
}

resource "square_catalog_object" "dance_only_pass" {
	type = "ITEM"
	
	item_data {
		name = "Dance Only"
	}
}

resource "square_catalog_object" "dance_only_pass_variation" {
	type = "ITEM_VARIATION"

	item_variation_data {
		item_id = square_catalog_object.dance_only_pass.id
		name = "Regular"
		pricing_type = "FIXED_PRICING"

		price_money {
			amount = 4500
			currency = "USD"
		}
	}
}

locals {
	t_shirt_variants = ["Unisex Small", "Unisex Medium", "Unisex Large", "Unisex XL", "Unisex 2XL", "Unisex 3XL", "Bella Small", "Bella Medium", "Bella Large", "Bella XL", "Bella 2XL"]
}

resource "square_catalog_object" "t_shirt" {
	type = "ITEM"
	
	item_data {
		name = "T-Shirt"
	}
}

resource "square_catalog_object" "t_shirt_variation" {
	for_each = toset(local.t_shirt_variants)

	type = "ITEM_VARIATION"

	item_variation_data {
		item_id = square_catalog_object.t_shirt.id
		name = each.value
		pricing_type = "FIXED_PRICING"

		price_money {
			amount = 1500
			currency = "USD"
		}
	}
}

locals {
	full_weekend_variants = {
		"Tier 1": 5500,
		"Tier 2": 6500,
		"Tier 3": 7500,
		"Tier 4": 8500,
		"Tier 5": 9500,
	}
}

resource "square_catalog_object" "full_weekend_pass" {
	type = "ITEM"
	
	item_data {
		name = "Full Weekend Pass"
	}
}

resource "square_catalog_object" "full_weekend_pass_variation" {
	for_each = local.full_weekend_variants

	type = "ITEM_VARIATION"

	item_variation_data {
		item_id = square_catalog_object.full_weekend_pass.id
		name = each.key
		pricing_type = "FIXED_PRICING"
		track_inventory = true

		price_money {
			amount = each.value
			currency = "USD"
		}
	}
}

resource "square_catalog_object" "student_discount" {
	type = "DISCOUNT"
	
	discount_data {
		name = "Student Discount"
		discount_type = "FIXED_AMOUNT"
		amount_money {
			amount = 500
			currency = "USD"	
		}
	}
}
