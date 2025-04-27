module "domain_sg" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "5.3.0"

  name        = "appstream_sg"
  description = ""
  vpc_id      = module.vpc.vpc_id

  ingress_with_cidr_blocks = [
    { from_port = 88, to_port = 88, protocol = "tcp", cidr_blocks = local.dc_ingress, description = "Kerberos Authentication" },
    { from_port = 135, to_port = 135, protocol = "tcp", cidr_blocks = local.dc_ingress, description = "MS RPC Endpoint Mapper" },
    { from_port = 389, to_port = 389, protocol = "tcp", cidr_blocks = local.dc_ingress, description = "LDAP" },
    { from_port = 636, to_port = 636, protocol = "tcp", cidr_blocks = local.dc_ingress, description = "LDAPS" },
    { from_port = 3268, to_port = 3268, protocol = "tcp", cidr_blocks = local.dc_ingress, description = "Global Catalog" },
    { from_port = 53, to_port = 53, protocol = "tcp", cidr_blocks = local.dc_ingress, description = "DNS (TCP)" },
    { from_port = 53, to_port = 53, protocol = "udp", cidr_blocks = local.dc_ingress, description = "DNS (UDP)" },
    { from_port = 137, to_port = 137, protocol = "udp", cidr_blocks = local.dc_ingress, description = "NetBIOS Name Service" },
    { from_port = 138, to_port = 138, protocol = "udp", cidr_blocks = local.dc_ingress, description = "NetBIOS Datagram Service" },
    { from_port = 139, to_port = 139, protocol = "tcp", cidr_blocks = local.dc_ingress, description = "NetBIOS Session Service" },
    { from_port = 445, to_port = 445, protocol = "tcp", cidr_blocks = local.dc_ingress, description = "Microsoft-DS" },
    { from_port = 1024, to_port = 65535, protocol = "tcp", cidr_blocks = local.dc_ingress, description = "Dynamic RPC" },
    { from_port = -1, to_port = -1, protocol = "icmp", cidr_blocks = local.dc_ingress, description = "ICMP (Ping)" }
  ]

  egress_with_cidr_blocks = [
    {
      protocol    = "-1"
      cidr_blocks = "0.0.0.0/0"
    }
  ]
}
