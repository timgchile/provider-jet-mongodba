/*
Copyright 2021 The Crossplane Authors.

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

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	"github.com/timgchile/provider-jet-mongodba/config/database"
	"github.com/timgchile/provider-jet-mongodba/config/mongodba"
	"github.com/timgchile/provider-jet-mongodba/config/project"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	resourcePrefix = "mongodba"
	modulePath     = "github.com/timgchile/provider-jet-mongodba"
)

//go:embed schema.json
var providerSchema string

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(DefaultResource(
			groupKindOverrides(),
			identifierAssignedByMongoDBAtlas(),
			commonReferences(),
		)),
		tjconfig.WithSkipList([]string{"mongodbatlas_encryption_at_rest", "mongodbatlas_teams"}),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		// add custom config functions
		mongodba.Configure,
		database.Configure,
		project.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// DefaultResource returns a DefaultResoruceFn that makes sure the original
// DefaultResource call is made with given options here.
func DefaultResource(opts ...tjconfig.ResourceOption) tjconfig.DefaultResourceFn {
	return func(name string, terraformResource *schema.Resource, orgOpts ...tjconfig.ResourceOption) *tjconfig.Resource {
		return tjconfig.DefaultResource(name, terraformResource, append(orgOpts, opts...)...)
	}
}
