resource "aws_subnet" "myapp-subnet-1" {
    vpc_id = var.vpc_id
    cidr_block = var.subnet_cidr_block
    availability_zone = var.avail_zone
    tags = {
        Name: "${var.env_prefix}-subnet-1"
    }
}

resource "aws_internet_gateway" "myapp-igw" {
    vpc_id = var.vpc_id 

    tags = {
        Name = "${var.env_prefix}-igw"
    }
}


resource "aws_route_table" "myapp-route-table" {
    vpc_id = var.vpc_id 

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
    route_table_id = var.default_route_table_id
}


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
    vpc_id = var.vpc_id


    // what is ingress se
    // who is allowed to access resource on port 22
    // 

    // https://whatismyipaddress.com/
    ingress {
        from_port = 22 
        to_port = 22
        protocol = "tcp"

        cidr_blocks = [var.my_ip]
    }

    egress {
        from_port = 0
        to_port = 0
        protocol = "-1"

        cidr_blocks = ["0.0.0.0/0"]
    }

    tags = {
        Name = "${var.env_prefix}-sg"
    }
}
