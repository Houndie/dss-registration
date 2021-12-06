variable "aws_access_key" {}
variable "aws_secret_key" {}

provider "aws" {
	region = "us-east-2"
	access_key = var.aws_access_key
	secret_key = var.aws_secret_key
}

resource "aws_kms_key" "backend_vax" {
  description             = "This key is used to encrypt bucket objects for ${var.workspace} vaccines"
  deletion_window_in_days = 7
}

resource "aws_s3_bucket" "vax" {
	bucket = "dayton-swing-smackdown-${var.workspace}-vaccines"
	acl = "private"

	cors_rule {
		allowed_headers = ["*"]
		allowed_methods = ["PUT", "POST", "GET"]
		allowed_origins = local.sites
		expose_headers  = ["ETag"]
		max_age_seconds = 3000
	}

	server_side_encryption_configuration {
		rule {
			apply_server_side_encryption_by_default {
				kms_master_key_id = aws_kms_key.backend_vax.arn
				sse_algorithm     = "aws:kms"
			}
			bucket_key_enabled = true
		}
	}
}

resource "aws_iam_user" "backend" {
	name = "backend-${var.workspace}"
}

resource "aws_iam_policy" "backend_vax" {
	name = "backend_vax-${var.workspace}"
	
	policy = jsonencode({
		Version = "2012-10-17"
		Statement = [
			{
				Resource = "${aws_s3_bucket.vax.arn}/*"
				Effect = "Allow"
				Action = ["s3:GetObject", "s3:PutObject", "s3:GetObjectAcl", "s3:DeleteObject"]
			}
		]
	})
}

resource "aws_iam_policy" "backend_vax_kms" {
	name = "backend_vax-${var.workspace}-kms"
	
	policy = jsonencode({
		Version = "2012-10-17"
		Statement = [
			{
				Resource = aws_kms_key.backend_vax.arn
				Effect = "Allow"
				Action = ["kms:*"]
			}
		]
	})
}

resource "aws_iam_user_policy_attachment" "backend_vax" {
	user = aws_iam_user.backend.name
	policy_arn = aws_iam_policy.backend_vax.arn
}

resource "aws_iam_user_policy_attachment" "backend_vax_kms" {
	user = aws_iam_user.backend.name
	policy_arn = aws_iam_policy.backend_vax_kms.arn
}

resource "aws_iam_access_key" "backend" {
	user = aws_iam_user.backend.name
}
