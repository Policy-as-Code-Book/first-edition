import "module" "common-functions" {
  source = "../../common/functions.sentinel"
}

mock "tfplan/v2" {
  module {
    source = "./mock-tfplan-v2-pass.sentinel"
  }
}

test {
  rules = {
    main = true
  }
}

param "p_required_labels" {
  value = ["billing","env","owner"]
}