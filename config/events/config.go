// Copyright 2022 Crossplane Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package events

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure adds configurations for events group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aws_cloudwatch_event_permission", func(r *config.Resource) {
		r.References["event_bus_name"] = config.Reference{
			Type:              "Bus",
			RefFieldName:      "EventBusNameRefs",
			SelectorFieldName: "EventBusNameSelector",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_cloudwatch_event_rule", func(r *config.Resource) {
		r.References["event_bus_name"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/events/v1beta1.Bus",
			RefFieldName:      "EventBusNameRefs",
			SelectorFieldName: "EventBusNameSelector",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_cloudwatch_event_target", func(r *config.Resource) {
		r.References["event_bus_name"] = config.Reference{
			Type:              "github.com/upbound/provider-aws/apis/events/v1beta1.Bus",
			RefFieldName:      "EventBusNameRefs",
			SelectorFieldName: "EventBusNameSelector",
		}
		delete(r.References, "arn")
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_cloudwatch_event_permission", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_cloudwatch_event_rule", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("aws_cloudwatch_event_target", func(r *config.Resource) {
		r.UseAsync = true
	})
}
