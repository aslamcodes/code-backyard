resource "aws_security_group" "domain_sg" {
  name        = "domain_sg"
  description = "Security group for Domain Controller"
  vpc_id      = module.vpc.vpc_id

  tags = {
    Name = "domain_sg"
  }
}

resource "aws_vpc_security_group_ingress_rule" "kerberos" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = 88
  ip_protocol       = "tcp"
  to_port           = 88
}

resource "aws_vpc_security_group_ingress_rule" "ms_rpc" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = 135
  ip_protocol       = "tcp"
  to_port           = 135
}

resource "aws_vpc_security_group_ingress_rule" "ldap" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = 389
  ip_protocol       = "tcp"
  to_port           = 389
}

resource "aws_vpc_security_group_ingress_rule" "ldaps" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = 636
  ip_protocol       = "tcp"
  to_port           = 636
}

resource "aws_vpc_security_group_ingress_rule" "global_catalog" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = 3268
  ip_protocol       = "tcp"
  to_port           = 3268
}

resource "aws_vpc_security_group_ingress_rule" "dns_tcp" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = 53
  ip_protocol       = "tcp"
  to_port           = 53
}

resource "aws_vpc_security_group_ingress_rule" "dns_udp" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = 53
  ip_protocol       = "udp"
  to_port           = 53
}

resource "aws_vpc_security_group_ingress_rule" "netbios_name_service" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = 137
  ip_protocol       = "udp"
  to_port           = 137
}

resource "aws_vpc_security_group_ingress_rule" "netbios_datagram_service" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = 138
  ip_protocol       = "udp"
  to_port           = 138
}

resource "aws_vpc_security_group_ingress_rule" "netbios_session_service" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = 139
  ip_protocol       = "tcp"
  to_port           = 139
}

resource "aws_vpc_security_group_ingress_rule" "microsoft_ds" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = 445
  ip_protocol       = "tcp"
  to_port           = 445
}

resource "aws_vpc_security_group_ingress_rule" "dynamic_rpc" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = 1024
  ip_protocol       = "tcp"
  to_port           = 65535
}

resource "aws_vpc_security_group_ingress_rule" "icmp_ping" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = var.domain_cidr
  from_port         = -1
  ip_protocol       = "icmp"
  to_port           = -1
}

resource "aws_vpc_security_group_egress_rule" "all_outbound" {
  security_group_id = aws_security_group.domain_sg.id
  cidr_ipv4         = "0.0.0.0/0"
  ip_protocol       = "-1"
}
