terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }

    okta = {
      source  = "okta/okta"
      version = "~> 4.17.0"
    }
  }
}
