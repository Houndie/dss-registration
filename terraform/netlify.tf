variable "github_token" {}
variable "netlify_token" {}

provider "github" {
	token = var.github_token
	owner = "Houndie"
}

provider "git" {}

provider "netlify" {
	token = var.netlify_token
	base_url = "https://api.netlify.com/api/v1/"
}

data "git_repository" "deploy_tag" {
	url = "https://github.com/Houndie/dss-registration"
	tag = var.deploy_version
}

resource "github_branch" "netlify_deploy" {
	repository = "dss-registration"
	branch = "netlify-deploy-${var.workspace}"
	source_sha = data.git_repository.deploy_tag.commit_sha
}

resource "netlify_deploy_key" "dayton_swing_smackdown" {}

resource "github_repository_deploy_key" "dayton_swing_smackdown" {
  title      = "netlify-deploy-${var.workspace}"
  repository = "dss-registration"
  key        = netlify_deploy_key.dayton_swing_smackdown.public_key
  read_only  = false
}

locals {
	frontend_square_data = {
		purchase_items = {
			solo_jazz = tolist(tolist(square_catalog_object.solo_jazz_variation.item_variation_data)[0].price_money)[0].amount
			team_competition = tolist(tolist(square_catalog_object.team_competition_variation.item_variation_data)[0].price_money)[0].amount
			mix_and_match = tolist(tolist(square_catalog_object.mix_and_match_variation["Leader"].item_variation_data)[0].price_money)[0].amount
			t_shirt = tolist(tolist(square_catalog_object.t_shirt_variation["Unisex Small"].item_variation_data)[0].price_money)[0].amount
			full_weekend_pass = { for tier, resource in square_catalog_object.full_weekend_pass_variation: tier => tolist(tolist(resource.item_variation_data)[0].price_money)[0].amount }
			dance_only_pass = tolist(tolist(square_catalog_object.dance_only_pass_variation["Presale"].item_variation_data)[0].price_money)[0].amount
		}
		student_discount = tolist(tolist(square_catalog_object.student_discount.discount_data)[0].amount_money)[0].amount
	}

	frontend_config = {
		GATSBY_BACKEND="https://${heroku_app.dayton_swing_smackdown.name}.herokuapp.com"
		GATSBY_FRONTEND="https://test.daytonswingsmackdown.com"
		GATSBY_CLIENT_ID=auth0_client.smackdown-website.client_id
		GATSBY_AUTH0_DOMAIN=var.auth0_domain
		GATSBY_AUTH0_AUDIENCE=auth0_resource_server.smackdown-website.identifier
		GATSBY_VERSION=var.deploy_version
		GATSBY_SQUARE_DATA=jsonencode(local.frontend_square_data)
		GATSBY_ACTIVE=var.active
	}
}

resource "netlify_site" "dayton_swing_smackdown" {
	name = "dayton-swing-smackdown-${var.workspace}"
	custom_domain = "test.daytonswingsmackdown.com"

	repo {
		repo_branch   = github_branch.netlify_deploy.branch
		command       = <<EOT
cd static && \
npm install && \
${join(" ", [for key, value in local.frontend_config: "${key}='${value}'"])} npx gatsby build
EOT
		deploy_key_id = netlify_deploy_key.dayton_swing_smackdown.id
		dir           = "static/public"
		provider      = "github"
		repo_path     = "Houndie/dss-registration"
	}
}

resource "github_repository_webhook" "netlify_push" {
	repository = "dss-registration"
	events     = ["push", "create"]

	configuration {
		content_type = "json"
		url          = "https://api.netlify.com/hooks/github"
	}

	depends_on = [netlify_site.dayton_swing_smackdown]
}
