import "module" "common-functions" {
  source = "./common/functions.sentinel"
}

mock "tfplan/v2" {
  module {
    source = "./mocks/mock-tfplan-v2.sentinel"
  }
}

policy "required-labels" {
  source = "./required-labels.sentinel"
  enforcement_level = "hard-mandatory"
  params = {
    "p_required_labels" = ["billing","env","owner"]
  }
}