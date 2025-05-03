data "aws_ami" "ubuntu" {
  most_recent = true

  filter {
    name   = "name"
    values = ["ubuntu/images/hvm-ssd-gp3/ubuntu-noble-24.04-amd64-server-20250305"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["099720109477"]
}


resource "aws_instance" "customer_gateway_device" {
  ami                    = data.aws_ami.ubuntu.id
  key_name               = data.aws_key_pair.aslam_task.key_name
  subnet_id              = module.vpc.public_subnets[0]
  user_data              = file("${path.module}/scripts/strongSwan.sh")
  vpc_security_group_ids = [aws_security_group.cgw_device.id]
  instance_type          = "t3.medium"
  source_dest_check      = false

  root_block_device {
    volume_size           = 8
    volume_type           = "gp3"
    delete_on_termination = true
  }
}

resource "aws_eip" "cgw_ip" {
  domain = "vpc"
}

resource "aws_eip_association" "eip_assoc" {
  instance_id   = aws_instance.customer_gateway_device.id
  allocation_id = aws_eip.cgw_ip.id
}
