output "backend_addr" {
	value = "https://${heroku_app.dayton_swing_smackdown.name}.herokuapp.com"
}

output "square_key" {
	value = var.square_access_token
	sensitive = true
}

output "mail_key" {
	value = var.sendinblue_token
	sensitive = true
}

output "recaptcha_key" {
	value = var.recaptcha_token
	sensitive = true
}

output "auth0_domain" {
	value = var.auth0_domain
}

output "auth0_client_id" {
	value = auth0_client.smackdown-website.client_id
}

output "auth0_audience" {
	value = auth0_resource_server.smackdown-website.identifier
}
