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
variable instance_type{}
variable public_key_location{}

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

        // 可能是漏了这个配置
        // 1. 导致没有运行所有规则出站
        cidr_blocks = ["0.0.0.0/0"]
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

        // 这条规则好像没有生效果
        // any trafic
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


// Query date from amazon !~
data "aws_ami" "latest-amazon-linux-image" {
    most_recent = true 
    owners = ["amazon"]
    filter {
        name = "name"
        values = ["amzn2-ami-hvm-*-x86_64-gp2"]
    }

    filter {
        name = "virtualization-type"
        values = ["hvm"]
    }
}

// output 命令用于 debug 的
output "aws_ami_id" {
    value = data.aws_ami.latest-amazon-linux-image
}

# // set ami dynamically 
resource "aws_instance" "myapp-server" {
    ami = data.aws_ami.latest-amazon-linux-image.id 
    instance_type = var.instance_type

    subnet_id = aws_subnet.myapp-subnet-1.id
    vpc_security_group_ids = [aws_security_group.default-sg.id]
    availability_zone = var.avail_zone
    
    associate_public_ip_address = true 
    // key_name = "server-key-pair"
    key_name = aws_key_pair.ssh-key.key_name

    # user_data = <<EOF 
    #                 #!/bin/bash
    #                 sudo yum update -y && sudo yum install -y docker 
    #                 sudo systemctl start docker 
    #                 sudo usermod -aG docker ec2-user 
    #                 docker run -p 8080:80 nginx 
    #             EOF

    user_data = <<EOF
#!/bin/bash
echo "Hello, World!"
EOF

  # Other arguments or blocks...

    tags = {
        Name = "${var.env_prefix}-server"
    }
}

// create EC2 type 
// 1. create key pair 
//    AWS rejects ssh request, if permission not set correctly !~ 
// 2. 


// private ip address: ec2-user@ip-10-0-10-130 ~
// 
// Do Not check this into your git repo
// a key pair must already exists locally on your machine
resource "aws_key_pair" "ssh-key" {
    key_name = "server-key"
    
   // public_key = "${file(var.public_key_location)}"
   public_key = file(var.public_key_location)
}


//  aws_instance.myapp-server must be replaced
// instance need to be replace

output "ec2_public_ip" {
    value = aws_instance.myapp-server.public_ip
}

//  ssh -i ~/.ssh/ssh-github-v2 ec2-user@35.180.251.173
// Automate as many as possible 
// 1. you may forget components, when it's time to clean up 
// 2. environment replication 
// 3. you have to document everything or remember ~ 
// 
// 

// P23 
// Configure EC2 server to run entryoint script 
// to start Docker contianer 
// 

// 1. EC2 Server running!
// 2. Networking configured!

// we want to automate this configuration too!
// 
// 
// 
// 


// Terraform 本身不提供直接的命令来检查 EC2 实例是否能够连接外网。
// 但是，你可以使用 Terraform 输出（output）来显示相关的配置信息，然后手动检查这些配置是否满足连接外网的条件。
# output "route_table" {
#   value = aws_route_table.main.id
# }

# output "internet_gateway" {
#   value = aws_internet_gateway.main.id
# }

# output "security_group" {
#   value = aws_security_group.main.id
# }

# output "network_acl" {
#   value = aws_network_acl.main.id
# }

