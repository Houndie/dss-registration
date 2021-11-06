variable "auth0_domain" {}
variable "auth0_client_id" {}
variable "auth0_client_secret" {}
variable "square_access_token" {}

terraform {
	required_providers {
		auth0 = {
			source  = "alexkappa/auth0"
			version = "~> 0.19.0"
		}
		square = {
			source = "houndie/square"
			version = "0.1.1"
		}
		heroku = {
			source = "heroku/heroku"
			version = "4.2.0"
		}
		herokux = {
			source = "davidji99/herokux"
			version = "0.22.1"
		}
		github = {
			source = "integrations/github"
			version = "4.9.2"
		}
		netlify = {
			source = "aegirhealth/netlify"
			version = "0.6.12"
		}
		git = {
			source = "innovationnorway/git"
			version = "0.1.3"
		}
		aws = {
			source = "hashicorp/aws"
			version = "3.55.0"
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
	grant_types = ["authorization_code"]
	token_endpoint_auth_method = "none"

	jwt_configuration {
		alg = "RS256"
	}
}

locals {
	permissions = {
		list = {
			value = "list:registrations"
			description = "list registrations"
		}
		update = {
			value = "update:registrations"
			description = "update registrations"
		}
	}
}

resource "auth0_resource_server" "smackdown-website" {
	name = "Dayton Swing Smackdown"
	identifier = "https://dayton-swing-smackdown-testing.herokuapp.com"
	signing_alg = "RS256"
	enforce_policies = true
	token_dialect = "access_token_authz"

	dynamic "scopes" {
		for_each = [for k, v in local.permissions: v]
		
		content {
			value = scopes.value.value
			description = scopes.value.description
		}
	}

	skip_consent_for_verifiable_first_party_clients = true
}

resource "auth0_role" "admin" {
	name = "Admin"
	description = "Full Admin, can do everything"

	dynamic "permissions" {
		for_each = [for k, v in local.permissions: v.value]

		content {
			resource_server_identifier = auth0_resource_server.smackdown-website.identifier
			name = permissions.value
		}
	}
}
