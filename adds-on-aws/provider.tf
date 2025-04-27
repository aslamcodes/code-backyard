provider "aws" {
  region = "us-east-1"
  default_tags {
    tags = {
      Owner     = "mohammedaslamm@presidio.com"
      Terraform = true
    }
  }
}

