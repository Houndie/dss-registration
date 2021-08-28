variable "aws_access_key" {}
variable "aws_secret_key" {}

provider "aws" {
	region = "us-east-2"
	access_key = var.aws_access_key
	secret_key = var.aws_secret_key
}

resource "aws_s3_bucket" "vax" {
	bucket = "dayton-swing-smackdown-testing-vaccines"
	acl = "private"

	cors_rule {
		allowed_headers = ["*"]
		allowed_methods = ["PUT", "POST", "GET"]
		allowed_origins = ["https://test.daytonswingsmackdown.com", "http://localhost:8081"]
		expose_headers  = ["ETag"]
		max_age_seconds = 3000
	}
}

resource "aws_iam_user" "backend" {
	name = "backend"
}

resource "aws_iam_policy" "backend_vax" {
	name = "backend_vax"
	
	policy = jsonencode({
		Version = "2012-10-17"
		Statement = [
			{
				Resource = "${aws_s3_bucket.vax.arn}/*"
				Effect = "Allow"
				Action = ["s3:GetObject", "s3:PutObject"]
			}
		]
	})
}

resource "aws_iam_user_policy_attachment" "backend_vax" {
	user = aws_iam_user.backend.name
	policy_arn = aws_iam_policy.backend_vax.arn
}

resource "aws_iam_access_key" "backend" {
	user = aws_iam_user.backend.name
}