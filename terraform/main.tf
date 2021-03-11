variable "auth0_domain" {}
variable "auth0_client_id" {}
variable "auth0_client_secret" {}

terraform {
	required_providers {
		auth0 = {
			source  = "alexkappa/auth0"
			version = "~> 0.19.0"
		}
	}
	backend "remote" {
		organization = "daytonswingsmackdown"
		workspaces {
			name = "testing"
		}
	}
}

provider "auth0" {
	domain        = var.auth0_domain
	client_id     = var.auth0_client_id
	client_secret = var.auth0_client_secret
}

resource "auth0_client" "smackdown-website" {
	name                = "Dayton Swing Smackdown"
	description         = "Dayton Swing Smackdown"
	app_type            = "spa"
	callbacks           = ["http://localhost:8081"]
	allowed_logout_urls = ["http://localhost:8081"]
	web_origins         = ["http://localhost:8081"]
	oidc_conformant = true

	jwt_configuration {
		alg = "RS256"
	}
}

resource "auth0_resource_server" "smackdown-website" {
	name = "Dayton Swing Smackdown"
	identifier = "https://dayton-swing-smackdown-testing.herokuapp.com"
	signing_alg = "RS256"
	enforce_policies = true
	token_dialect = "access_token_authz"

	scopes {
		value = "list:discounts"
		description = "list discounts"
	}
}
