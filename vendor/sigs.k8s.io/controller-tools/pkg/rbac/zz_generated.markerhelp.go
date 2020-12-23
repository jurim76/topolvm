// +build !ignore_autogenerated

/*
Copyright2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by helpgen. DO NOT EDIT.

package rbac

import (
	"sigs.k8s.io/controller-tools/pkg/markers"
)

func (Generator) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "",
		DetailedHelp: markers.DetailedHelp{
			Summary: "generates ClusterRole objects.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{
			"RoleName": markers.DetailedHelp{
				Summary: "sets the name of the generated ClusterRole.",
				Details: "",
			},
		},
	}
}

func (Rule) Help() *markers.DefinitionHelp {
	return &markers.DefinitionHelp{
		Category: "RBAC",
		DetailedHelp: markers.DetailedHelp{
			Summary: "specifies an RBAC rule to all access to some resources or non-resource URLs.",
			Details: "",
		},
		FieldHelp: map[string]markers.DetailedHelp{
			"Groups": markers.DetailedHelp{
				Summary: "specifies the API groups that this rule encompasses.",
				Details: "",
			},
			"Resources": markers.DetailedHelp{
				Summary: "specifies the API resources that this rule encompasses.",
				Details: "",
			},
			"ResourceNames": markers.DetailedHelp{
				Summary: "specifies the names of the API resources that this rule encompasses. ",
				Details: "Create requests cannot be restricted by resourcename, as the object's name is not known at authorization time.",
			},
			"Verbs": markers.DetailedHelp{
				Summary: "specifies the (lowercase) kubernetes API verbs that this rule encompasses.",
				Details: "",
			},
			"URLs": markers.DetailedHelp{
				Summary: "URL specifies the non-resource URLs that this rule encompasses.",
				Details: "",
			},
			"Namespace": markers.DetailedHelp{
				Summary: "specifies the scope of the Rule. If not set, the Rule belongs to the generated ClusterRole. If set, the Rule belongs to a Role, whose namespace is specified by this field.",
				Details: "",
			},
		},
	}
}