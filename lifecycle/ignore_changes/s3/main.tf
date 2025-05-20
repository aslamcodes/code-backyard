terraform {
  required_version = ">= 1.0.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0" # Adjust to the latest version as needed
    }
  }
}

provider "aws" {
  region = var.region
}

variable "bucket_name" {
  description = "The name of the S3 bucket"
  type        = string
}

variable "region" {
  description = "The AWS region to create the S3 bucket in"
  type        = string
}

resource "aws_s3_bucket" "this" {
  bucket = var.bucket_name

  lifecycle {
    ignore_changes = [bucket]
  }

}

output "bucket_arn" {
  value = aws_s3_bucket.this.arn
}
