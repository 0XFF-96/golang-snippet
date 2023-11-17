# terraform {
#   required_providers {
#     aws = {
#       source  = "hashicorp/aws"
#       version = "~> 5.0"
#     }
#   }
# }

provider "aws" {
    region = "eu-west-3"
}

variable vpc_cidr_block {}
variable avail_zone {}
variable subnet_cidr_block {}
variable env_prefix {}

resource "aws_vpc" "myapp-vpc" {
    // cidr_block = "10.0.0.0/16"
    cidr_block = var.vpc_cidr_block
    tags = {
        Name: "${var.env_prefix}-vpc"
    }
}

resource "aws_subnet" "myapp-subnet-1" {
    // create a resource based on something 
    // that haven't been created yet.
    vpc_id = aws_vpc.myapp-vpc.id 
    // cidr_block = "10.0.10.0/24"
    cidr_block = var.subnet_cidr_block
    availability_zone = var.avail_zone
    tags = {
        Name: "${var.env_prefix}-subnet-1"
    }
}
