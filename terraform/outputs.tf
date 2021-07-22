output "backend_addr" {
	value = "https://${heroku_app.dayton_swing_smackdown.name}.herokuapp.com"
}

output "backend_config_vars" {
	value = local.backend_config_vars
}

output "backend_sensitive_config_vars" {
	value = local.backend_sensitive_config_vars
	sensitive = true
}

output "frontend_config_vars" {
	value = local.frontend_config
}
