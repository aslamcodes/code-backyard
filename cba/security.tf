module "appstream_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "5.3.0"

  name        = "appstream_sg"
  description = ""
  vpc_id      = module.vpc.id

  ingress_with_cidr_blocks = [
    {
      protocol    = "-1"
      cidr_blocks = "0.0.0.0/0"
    }
  ]
}
