import "module" "common-functions" {
  source = "../../common/functions.sentinel"
}

mock "tfplan/v2" {
  module {
    source = "./mock-tfplan-v2-fail.sentinel"
  }
}

test {
  rules = {
    main = false
  }
}

param "p_required_labels" {
  value = ["billing","env","owner"]
}