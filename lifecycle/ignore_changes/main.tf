provider "aws" {
  region = "us-west-2"
}

module "s3_bucket" {
  source      = "./s3"
  bucket_name = "lifecycle-poc-bucket-aslam-trying"
  region      = "us-west-2"
}

output "s3_bucket_arn" {
  value = module.s3_bucket.bucket_arn
}

terraform {
  required_version = ">= 1.0.0"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0" # Adjust to the latest version as needed
    }
  }
}
