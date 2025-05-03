locals {
  aws_vpc_vpgw_id = "vgw-0f393e2d0df8877e2"
}

resource "aws_customer_gateway" "cgw" {
  type       = "ipsec.1"
  ip_address = "34.200.180.39" # The IP address of internet routable device
  tags = {
    Name = "Cust_gw"
  }
}

resource "aws_vpn_connection" "on_prem_connect" {
  customer_gateway_id = aws_customer_gateway.cgw.id
  vpn_gateway_id      = local.aws_vpc_vpgw_id
  type                = aws_customer_gateway.cgw.type
  static_routes_only  = true # for devices that don't support BGP
}
