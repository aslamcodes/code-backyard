data "aws_key_pair" "aslam_task" {
  key_name = "aslamtask"
}

data "aws_ami" "windows" {
  most_recent = true

  filter {
    name   = "name"
    values = ["Windows_Server-2019-English-Full-Base-*"]
  }

  filter {
    name   = "virtualization-type"
    values = ["hvm"]
  }

  owners = ["801119661308"]
}

resource "aws_instance" "web" {
  ami           = data.aws_ami.windows.id
  instance_type = "t3.medium"
  user_data = templatefile("${path.module}/scripts/dc.ps1", {
    DomainName       = var.domain_name
    SafeModePassword = var.safe_mode_password
    NetBiosName      = var.netbios_name
  })
  subnet_id              = module.vpc.public_subnets[0]
  vpc_security_group_ids = [aws_security_group.domain_sg.id]
  key_name               = data.aws_key_pair.aslam_task.key_name

  root_block_device {
    volume_size           = 50
    volume_type           = "gp3"
    delete_on_termination = true
  }
}
