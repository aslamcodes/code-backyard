output "vpc_id" {
  value = module.vpc.vpc_id
}

output "instance_id" {
  value = aws_instance.web.id
}

output "vpw_id" {
  value = module.vpc.vgw_id
}
