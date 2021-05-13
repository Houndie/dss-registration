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
	tag = "netlify-test-deploy-tag"
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

resource "netlify_site" "dayton_swing_smackdown" {
	name = "dayton-swing-smackdown-${var.workspace}"

	repo {
		repo_branch   = github_branch.netlify_deploy.branch
		command       = <<EOT
cd static && \
npm install && \
GATSBY_BACKEND=https://${heroku_app.dayton_swing_smackdown.name}.herokuapp.com \
GATSBY_FRONTEND=https://test.daytonswingsmackdown.com \
GATSBY_CLIENT_ID=${auth0_client.smackdown-website.client_id} \
GATSBY_AUTH0_DOMAIN=${var.auth0_domain}
GATSBY_AUTH_AUDIENCE=${auth0_resource_server.smackdown-website.identifier} \
npx gatsby build
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
