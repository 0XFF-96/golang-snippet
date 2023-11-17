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
variable my_ip {}

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

// 1. 
// Terraform knows 
// in which sequence the components must be created 
// 2. 
// 3. 
// 4.
resource "aws_internet_gateway" "myapp-igw" {
    vpc_id = aws_vpc.myapp-vpc.id 

    tags = {
        Name = "${var.env_prefix}-igw"
    }
}


resource "aws_route_table" "myapp-route-table" {
    vpc_id = aws_vpc.myapp-vpc.id 

    route {
        cidr_block = "0.0.0.0/0"
        gateway_id = aws_internet_gateway.myapp-igw.id 
    }

    tags = {
        Name = "${var.env_prefix}-rtb"
    }
}


resource "aws_route_table_association" "a-rtb-subnet" {
    subnet_id = aws_subnet.myapp-subnet-1.id
    route_table_id = aws_route_table.myapp-route-table.id
}

# // terraform state show -- command 
# # resource "aws_default_route_table" "main-rtb" {
# #     default_route_table_id = 
# # }


resource "aws_default_route_table" "main-rtb" {
    default_route_table_id = aws_vpc.myapp-vpc.default_route_table_id

    route {
        cidr_block = "0.0.0.0/0"
        gateway_id = aws_internet_gateway.myapp-igw.id 
    }

    tags = {
        Name = "${var.env_prefix}-rtb"
    }
}


// but still creating a new one is recommended in most cases
resource "aws_security_group" "myapp-sg" {
    name = "myapp-sg"
    vpc_id = aws_vpc.myapp-vpc.id


    // what is ingress se
    // who is allowed to access resource on port 22
    // 

    // https://whatismyipaddress.com/
    ingress {
        from_port = 22 
        to_port = 22
        protocol = "tcp"

        // you don't have to hard code here 
        // if your ip is dynamic 
        cidr_blocks = [var.my_ip]
    }

    egress {

        // nay trafic
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = []
    }

    tags = {
        Name = "${var.env_prefix}-sg"
    }
}


// Use default Security Group 
// Instead of a old one 
// 
 resource "aws_security_group" "default-sg" {
    vpc_id = aws_vpc.myapp-vpc.id
    // what is ingress se
    // who is allowed to access resource on port 22
    // 


    // https://whatismyipaddress.com/
    ingress {
        from_port = 22 
        to_port = 22
        protocol = "tcp"

        // you don't have to hard code here 
        // if your ip is dynamic 
        cidr_blocks = [var.my_ip]
    }

    egress {

        // nay trafic
        from_port = 0
        to_port = 0
        protocol = "-1"
        cidr_blocks = []
    }

    tags = {
        Name = "${var.env_prefix}-default-sg"
    }
}

// Fetch Amazon Machine Image 
// For EC2 instance 
// Amazon Linux 


resource "aws_instance" "myapp-server" {
    ami = ""
}