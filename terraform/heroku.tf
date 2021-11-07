variable "heroku_api_key" {}
variable "deploy_version" {}
variable "workspace" {}
variable "sendinblue_token" {}
variable "recaptcha_token" {}
variable "active" {}

provider "heroku" {
	api_key = var.heroku_api_key
}

provider "herokux" {
	api_key = var.heroku_api_key
}

locals {
	backend_square_variations = {
		"solo_jazz": square_catalog_object.solo_jazz_variation
		"team_competition": square_catalog_object.team_competition_variation
		"dance_only_pass": square_catalog_object.dance_only_pass_variation["Presale"]
	}
	backend_square_set_variations = {
		"full_weekend_pass": square_catalog_object.full_weekend_pass_variation
		"mix_and_match": square_catalog_object.mix_and_match_variation
		"t_shirt": square_catalog_object.t_shirt_variation
	}

	backend_square_purchase_items = {
		for purchase_item, purchase_item_value in local.backend_square_variations: purchase_item => {
			id = purchase_item_value.id
			price = tolist(tolist(purchase_item_value.item_variation_data)[0].price_money)[0].amount
		}
	}

	backend_square_set_purchase_items =  {
		for purchase_item, purchase_item_value in local.backend_square_set_variations: purchase_item => {
			for k, v in purchase_item_value: k => {
				id = v.id
				price = tolist(tolist(v.item_variation_data)[0].price_money)[0].amount
			}
		}
	}

	backend_square_data =  {
		purchase_items = merge(local.backend_square_purchase_items, local.backend_square_set_purchase_items)
		discounts = {
			student = {
				id = square_catalog_object.student_discount.id
				amount = tolist(tolist(square_catalog_object.student_discount.discount_data)[0].amount_money)[0].amount
				discount_type = tolist(square_catalog_object.student_discount.discount_data)[0].discount_type
				applied_to = "Full Weekend"
			}
		}
	}

	backend_config_vars = {
		DSS_FRONTEND = "https://test.daytonswingsmackdown.com"
		DSS_AUTHENDPOINT = "https://${var.auth0_domain}"
		DSS_ENVIRONMENT = "development"
		DSS_VERSION = var.deploy_version
		DSS_SQUAREDATA = jsonencode(local.backend_square_data)
		DSS_AWS_ACCESSKEY = aws_iam_access_key.backend.id
		DSS_AWS_VAXBUCKET = aws_s3_bucket.vax.bucket
		DSS_PERMISSIONS_LIST = local.permissions.list.value
		DSS_PERMISSIONS_UPDATE = local.permissions.update.value
		DSS_ACTIVE = var.active
	}

	backend_sensitive_config_vars = {
		DSS_SQUAREKEY = var.square_access_token
		DSS_MAILKEY = var.sendinblue_token
		DSS_REAPTCHAKEY = var.recaptcha_token
		DSS_AWS_SECRETKEY = aws_iam_access_key.backend.secret
	}
}

resource "heroku_app" "dayton_swing_smackdown" {
	name   = "dayton-swing-smackdown-${var.workspace}"
	region = "us"

	config_vars = local.backend_config_vars

	sensitive_config_vars = local.backend_sensitive_config_vars
}

data "herokux_registry_image" "dayton_swing_smackdown" {
	app_id = heroku_app.dayton_swing_smackdown.uuid
	process_type = "web"
	docker_tag = var.deploy_version
}

resource "herokux_app_container_release" "dayton_swing_smackdown" {
	app_id = heroku_app.dayton_swing_smackdown.uuid
	image_id = data.herokux_registry_image.dayton_swing_smackdown.digest
	process_type = data.herokux_registry_image.dayton_swing_smackdown.process_type
}

resource "heroku_formation" "dayton_swing_smackdown" {
	app = heroku_app.dayton_swing_smackdown.name
	type = data.herokux_registry_image.dayton_swing_smackdown.process_type
	quantity = 1
	size = "free"

	depends_on = [herokux_app_container_release.dayton_swing_smackdown]
}

resource "heroku_addon" "database" {
  app  = heroku_app.dayton_swing_smackdown.name
  plan = "heroku-postgresql:hobby-dev"
}
