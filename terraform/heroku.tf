variable "heroku_api_key" {}
variable "deploy_version" {}
variable "workspace" {}
variable "sendinblue_token" {}
variable "recaptcha_token" {}

provider "heroku" {
	api_key = var.heroku_api_key
}

provider "herokux" {
	api_key = var.heroku_api_key
}

resource "heroku_app" "dayton_swing_smackdown" {
	name   = "dayton-swing-smackdown-${var.workspace}"
	region = "us"

	config_vars = {
		DSS_FRONTEND = "https://test.daytonswingsmackdown.com"
		DSS_AUTH0ENDPOINT = var.auth0_domain
		DSS_ENVIRONMENT = "development"
		DSS_VERSION = var.deploy_version
   }

	sensitive_config_vars = {
		DSS_SQUAREKEY = var.square_access_token
		DSS_MAILKEY = var.sendinblue_token
		DSS_REAPTCHAKEY = var.recaptcha_token
   }
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

output "backend_addr" {
	value = "https://${heroku_app.dayton_swing_smackdown.name}.herokuapp.com"
}
