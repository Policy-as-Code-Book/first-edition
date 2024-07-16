mock "tfconfig" {
  module {
    source = "./mocks/mock-tfconfig.sentinel"
  }
}

mock "tfconfig/v1" {
  module {
    source = "./mocks/mock-tfconfig.sentinel"
  }
}

mock "tfconfig/v2" {
  module {
    source = "./mocks/mock-tfconfig-v2.sentinel"
  }
}

mock "tfplan" {
  module {
    source = "./mocks/mock-tfplan.sentinel"
  }
}

mock "tfplan/v1" {
  module {
    source = "./mocks/mock-tfplan.sentinel"
  }
}

mock "tfplan/v2" {
  module {
    source = "./mocks/mock-tfplan-v2.sentinel"
  }
}

mock "tfstate" {
  module {
    source = "./mocks/mock-tfstate.sentinel"
  }
}

mock "tfstate/v1" {
  module {
    source = "./mocks/mock-tfstate.sentinel"
  }
}

mock "tfstate/v2" {
  module {
    source = "./mocks/mock-tfstate-v2.sentinel"
  }
}

mock "tfrun" {
  module {
    source = "./mocks/mock-tfrun.sentinel"
  }
}

policy "eks" {
  source = "./eks.sentinel"
  enforcement_level = "hard-mandatory"
  params = {
    "p_eks_version" = ["1.26"],
    "p_region" = "us-east-1",
    "p_tf_version" = "1.5.4"
  }
}