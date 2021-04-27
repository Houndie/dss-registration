variable "heroku_api_key" {}
variable "deploy_version" {}

provider "heroku" {
	api_key = var.heroku_api_key
}

provider "herokux" {
	api_key = var.heroku_api_key
}

resource "heroku_app" "dayton_swing_smackdown" {
  name   = "dayton-swing-smackdown-testing"
  region = "us"
}
