// Copyright 2019 Thorsten Kukuk
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

package kubeadm

import (
	"strings"

	"github.com/thkukuk/kubic-control/pkg/tools"
)

func ListNodes() (bool, string, []string) {

	// Get list of all worker nodes:
        success, message := tools.ExecuteCmd("salt", "-G", "kubicd:kubic-worker-node", "grains.get",  "kubic-worker-node")
        if success != true {
		return success, message, nil
        }
        message = strings.TrimSuffix(message, "\n")
        nodelist := strings.Split (strings.Replace(message, ":", "", -1), "\n")

	return true, "", nodelist
}