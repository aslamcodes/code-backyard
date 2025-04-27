provider "aws" {
  region = "us-east-1"
  default_tags {
    Owner      = "mohammedaslamm@presidio.com"
    CreatedFor = "cba_test"
    Terraform  = true
  }
}

# TODO
provider "okta" {
  org_name       = "[ORG NAME e.g. dev-123456]"
  base_url       = "[okta.com|oktapreview.com]"
  client_id      = "[APP CLIENT_ID]"
  private_key_id = "[PRIVATE KEY ID - KID]"
  private_key    = "[PRIVATE KEY]"
  scopes         = "[COMMA,SEPARATED,SCOPE,VALUES]"
}
