resource "aws_vpc" "prod" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "default"

  tags = {
    Name = "prod"
  }
}

resource "aws_vpc" "staging" {
  cidr_block       = "10.0.0.0/16"
  instance_tenancy = "default"

  tags = {
    Name = "staging"
  }
}

output "hoge" {
  description = "value"
  value = resource.aws_vpc.staging.id
}
