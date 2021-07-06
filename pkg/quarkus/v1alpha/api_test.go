// Copyright 2021 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"
	"sigs.k8s.io/kubebuilder/v3/pkg/model/resource"
)

// Remaining functions - InjectResource, Scaffold

var _ = Describe("v1", func() {
	testAPISubcommand := &createAPISubcommand{}
	testResource := resource.Resource{}

	Describe("UpdateResource", func() {
		testAPIOptions := &createAPIOptions{}
		testAPIOptions.UpdateResource(&testResource)

		It("verify that resource path is an empty string", func() {
			Expect(testResource.Path, "")
		})

		It("verify that resource controller is false", func() {
			Expect(testResource.Controller).To(BeFalse())
		})
	})

	Describe("InjectResource", func() {
		err := testAPISubcommand.InjectResource(&testResource)

		It("verify that the subcommand resource was set", func() {
			Expect(testAPISubcommand.resource, testResource)
		})

		// It("should return nil", func() {
		// 	fmt.Println("-------")
		// 	fmt.Println(err)
		// 	fmt.Println("-------")
		// 	Expect(err).To(BeNil())
		// })
	})

	Describe("BindFlags", func() {
		flagTest := pflag.NewFlagSet("testFlag", -1)
		// Need clarification on what to set ErrorHandlingNumber as

		testAPISubcommand.BindFlags(flagTest)

		It("should set SortFlags to false", func() {
			Expect(flagTest.SortFlags).To(BeFalse())
		})

		It("should set CRDVersion to v1", func() {
			Expect(testAPISubcommand.options.CRDVersion, "v1")
		})

		It("should set Namespaced to true", func() {
			Expect(testAPISubcommand.options.Namespaced).To(BeTrue())
		})
	})

	Describe("InjectConfig", func() {
		testConfig, _ := config.New(config.Version{Number: 3})
		// Need clarification on what to set Version number as

		err := testAPISubcommand.InjectConfig(testConfig)

		It("should set config", func() {
			Expect(testAPISubcommand.config, testConfig)
		})

		It("should return nil", func() {
			Expect(err).To(BeNil())
		})
	})

	Describe("Run", func() {
		testFileSystem := machinery.Filesystem{}
		It("should return nil", func() {
			Expect(testAPISubcommand.Run(testFileSystem)).To(BeNil())
		})
	})

	Describe("Validate", func() {
		It("should return nil", func() {
			Expect(testAPISubcommand.Validate()).To(BeNil())
		})
	})

	Describe("PostScaffold", func() {
		It("should return nil", func() {
			Expect(testAPISubcommand.PostScaffold()).To(BeNil())
		})
	})
})
