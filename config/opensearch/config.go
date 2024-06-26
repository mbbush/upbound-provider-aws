// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: CC0-1.0

package opensearch

import (
	"github.com/crossplane/upjet/pkg/config"

	"github.com/upbound/provider-aws/config/common"
)

// Configure adds configurations for the opensearch group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_opensearch_domain", func(r *config.Resource) {
		config.MoveToStatus(r.TerraformResource, "access_policies")
		r.References["encrypt_at_rest.kms_key_id"] = config.Reference{
			// its KMS key ARN in AWS API
			TerraformName: "aws_kms_key",
			Extractor:     common.PathARNExtractor,
		}

		r.References["vpc_options.security_group_ids"] = config.Reference{
			TerraformName:     "aws_security_group",
			RefFieldName:      "SecurityGroupIDRefs",
			SelectorFieldName: "SecurityGroupIDSelector",
		}

		r.References["vpc_options.subnet_ids"] = config.Reference{
			TerraformName:     "aws_subnet",
			RefFieldName:      "SubnetIDRefs",
			SelectorFieldName: "SubnetIDSelector",
		}

		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_opensearch_domain_policy", func(r *config.Resource) {
		r.References["domain_name"] = config.Reference{
			TerraformName: "aws_opensearch_domain",
		}
		r.UseAsync = true
	})

	p.AddResourceConfigurator("aws_opensearch_domain_saml_options", func(r *config.Resource) {
		r.UseAsync = true
	})
}
