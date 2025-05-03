output "vpc_id" {
  value = module.vpc.vpc_id
}

output "instance_id" {
  value = aws_instance.web.id
}

output "cgw_device_ip" {
  value = aws_eip.cgw_ip.public_ip
}
