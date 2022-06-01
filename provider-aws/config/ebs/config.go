/*
Copyright 2021 Upbound Inc.
*/

package ebs

import (
	"github.com/upbound/upjet/pkg/config"

	"github.com/upbound/official-providers/provider-aws/config/common"
)

// Configure adds configurations for ebs group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_ebs_volume", func(r *config.Resource) {
		r.Version = common.VersionV1Alpha2
		r.References = map[string]config.Reference{
			"kms_key_id": {
				Type: "github.com/upbound/official-providers/provider-aws/apis/kms/v1alpha2.Key",
			},
		}
	})
}