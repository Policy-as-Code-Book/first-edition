# Variable Overrides

# AWS
region = "us-east-2"

# EKS Cluster
cluster_name           = "tf-test"
cluster_version        = "1.26"
cluster_log_types      = ["audit", "api", "authenticator", "controllerManager", "scheduler"]
# secrets_encryption_kms = "arn:aws:kms:us-east-2:123456789012:key/<KEY_ID>>"

# Networking
vpc_cidr = "10.0.0.0/16"
vpc_id   = "vpc-..."
private_subnet_ids = [
  "subnet-...",
  "subnet-..."
]
public_subnet_ids = [
  "subnet-...",
  "subnet-...",
  "subnet-..."
]

#MNGs
mng_instance_types = ["m5.large"]
mng_al2_name       = "tf-test-al2"
# arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy
mng_policies = ["arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
  "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
  "arn:aws:iam::aws:policy/AmazonSSMManagedInstanceCore"]
#   ,
#   "arn:aws:iam::123456789012:policy/KmsKeyUserSsmOps",
# "arn:aws:iam::123456789012:policy/s3-jimmyray-eks-read"]
mng_labels = {
  owner   = "jimmyray"
  env     = "dev"
  billing = "lob-cc"
}
mng_tags = {
  Terraform = "true"
}
ami_type_al2 = "AL2_x86_64"
# ami_id_al2 = "ami-..." #1.23
# ami_id_ubuntu = "ami-..." #1.23


