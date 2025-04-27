module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "5.21.0"
  name    = "dc-vpc"
  cidr    = "172.16.0.0/16"

  azs             = ["us-east-1a", "us-east-1b", "us-east-1c"]
  private_subnets = ["172.16.1.0/24", "172.16.2.0/24", "172.16.3.0/24"]
  public_subnets  = ["172.16.101.0/24", "172.16.102.0/24", "172.16.103.0/24"]

  enable_nat_gateway  = true
  enable_vpn_gateway  = true
  enable_dhcp_options = true

  dhcp_options_domain_name         = local.domain_name
  dhcp_options_domain_name_servers = [aws_instance.web.private_ip]
}
