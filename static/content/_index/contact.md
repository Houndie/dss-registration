+++
fragment = "contact"
#disabled = true
date = "2017-09-10"
weight = 1100
#background = "light"
form_name = "defaultContact"

title = "Have questions?"
subtitle  = "*Contact us!*"

# PostURL can be used with backends such as mailout from caddy
email = "info@daytonswingsmackdown.com"
button = "Send Button" # defaults to theme default
#netlify = false

# Optional google captcha
#[recaptcha]
#  sitekey = ""

[message]
  #success = "" # defaults to theme default
  #error = "" # defaults to theme default

# Only defined fields are shown in contact form
[fields.name]
  text = "Your Name *"

[fields.email]
  text = "Your Email *"

[fields.message]
  text = "Your Message *"
+++
