// Copyright 2018 Authors of Cilium
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

package RuntimeTest

import (
	"fmt"

	. "github.com/cilium/cilium/test/ginkgo-ext"
	"github.com/cilium/cilium/test/helpers"

	. "github.com/onsi/gomega"
)

// ExpectPolicyEnforcementUpdated sets PolicyEnforcement mode on the provided vm
// running Cilium to policyEnforcementType, and waits for endpoints on vm to
// be in 'ready' state. It asserts whether the configuration update was successful,
// and if the endpoints are ready.
func ExpectPolicyEnforcementUpdated(vm *helpers.SSHMeta, policyEnforcementType string) {
	By(fmt.Sprintf("Setting PolicyEnforcement=%s", policyEnforcementType))
	areEndpointsReady := vm.SetPolicyEnforcementAndWait(policyEnforcementType)
	ExpectWithOffset(1, areEndpointsReady).Should(BeTrue(), "unable to set PolicyEnforcement=%s", policyEnforcementType)
}
