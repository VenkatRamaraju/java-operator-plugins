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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/pflag"
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/plugin"
)

var _ = Describe("v1", func() {
	successInitSubcommand := &initSubcommand{
		domain:      "testDomain",
		projectName: "test-123",
	}

	failureInitSubcommand := &initSubcommand{
		domain:      "testDomain",
		projectName: "?&fail&?",
		commandName: "failureTest",
	}

	Describe("UpdateMetadata", func() {
		testCliMetadata := plugin.CLIMetadata{CommandName: "TestCommand"}
		testSubcommandMetadata := plugin.SubcommandMetadata{}

		It("Check command name inequality pre function call", func() {
			Expect(failureInitSubcommand.commandName).NotTo(Equal(testCliMetadata.CommandName))
		})

		successInitSubcommand.UpdateMetadata(testCliMetadata, &testSubcommandMetadata)

		It("Check command name equality post function call", func() {
			Expect(successInitSubcommand.commandName, testCliMetadata.CommandName)
		})
	})

	Describe("BindFlags", func() {
		flagTest := pflag.NewFlagSet("testFlag", -1)
		// Need clarification on what to set ErrorHandlingNumber as

		successInitSubcommand.BindFlags(flagTest)

		It("should set SortFlags to false", func() {
			Expect(flagTest.SortFlags).To(BeFalse())
		})

		It("should set domain to my.domain", func() {
			Expect(successInitSubcommand.domain, "my.domain")
		})

		It("should set projectName to an empty string", func() {
			Expect(successInitSubcommand.projectName, "")
		})

		It("should set group to an empty string", func() {
			Expect(successInitSubcommand.group, "")
		})

		It("should set version to an empty string", func() {
			Expect(successInitSubcommand.version, "")
		})

		It("should set kind to an empty string", func() {
			Expect(successInitSubcommand.kind, "")
		})
	})

	Describe("InjectConfig", func() {
		testConfig, _ := config.New(config.Version{Number: 3})
		// Need clarification on what to set Version number as

		It("should error", func() {
			Expect(failureInitSubcommand.InjectConfig(testConfig)).To(HaveOccurred())
		})

		It("should return nil", func() {
			Expect(successInitSubcommand.InjectConfig(testConfig)).To(BeNil())
		})
	})

	Describe("Validate", func() {
		It("should return nil", func() {
			Expect(successInitSubcommand.Validate()).To(BeNil())
		})
	})

	Describe("PostScaffold", func() {
		It("should return nil", func() {
			Expect(successInitSubcommand.PostScaffold()).To(BeNil())
		})
	})
})
