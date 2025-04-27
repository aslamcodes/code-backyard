resource "aws_appstream_stack" "cba_stack" {
  name         = "aslam cba test stack"
  description  = "stack description"
  display_name = "stack display name"


  user_settings {
    action     = "CLIPBOARD_COPY_FROM_LOCAL_DEVICE"
    permission = "ENABLED"
  }
  user_settings {
    action     = "CLIPBOARD_COPY_TO_LOCAL_DEVICE"
    permission = "ENABLED"
  }
  # user_settings {
  #   action     = "DOMAIN_PASSWORD_SIGNIN"
  #   permission = "ENABLED"
  # }
  # user_settings {
  #   action     = "DOMAIN_SMART_CARD_SIGNIN"
  #   permission = "DISABLED"
  # }
  user_settings {
    action     = "FILE_DOWNLOAD"
    permission = "ENABLED"
  }
  user_settings {
    action     = "FILE_UPLOAD"
    permission = "ENABLED"
  }
  user_settings {
    action     = "PRINTING_TO_LOCAL_DEVICE"
    permission = "ENABLED"
  }

  application_settings {
    enabled        = true
    settings_group = "SettingsGroup"
  }

  tags = {
    Owner = "Aslam"
  }
}

resource "aws_appstream_fleet" "cba_fleet" {
  name = "cba-fleet"

  compute_capacity {
    desired_instances = 1
  }

  description                        = "cba fleet"
  idle_disconnect_timeout_in_seconds = 60
  display_name                       = "cba-fleet"
  enable_default_internet_access     = false
  fleet_type                         = "ON_DEMAND"
  image_name                         = "wiki-firefox"
  instance_type                      = "stream.standard.large"
  max_user_duration_in_seconds       = 600

  vpc_config {
    subnet_ids = module.vpc.public_subnets
  }
}

resource "aws_appstream_fleet_stack_association" "stack_association" {
  fleet_name = aws_appstream_fleet.cba_fleet.name
  stack_name = aws_appstream_stack.cba_stack.name
}
