/*
  2022 NVIDIA CORPORATION & AFFILIATES

  Licensed under the Apache License, Version 2.0 (the License);
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an AS IS BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.
*/

package state

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/Mellanox/network-operator/api/v1alpha1"
	"github.com/Mellanox/network-operator/pkg/nodeinfo"
)

var _ = Describe("MOFED state test", func() {
	var stateOfed stateOFED

	BeforeEach(func() {
		stateOfed = stateOFED{}
	})

	Context("getMofedDriverImageName", func() {
		nodeAttr := make(map[nodeinfo.AttributeType]string)
		nodeAttr[nodeinfo.AttrTypeCPUArch] = "amd64"
		nodeAttr[nodeinfo.AttrTypeOSName] = "ubuntu"
		nodeAttr[nodeinfo.AttrTypeOSVer] = "20.04"

		cr := &v1alpha1.NicClusterPolicy{
			Spec: v1alpha1.NicClusterPolicySpec{
				OFEDDriver: &v1alpha1.OFEDDriverSpec{
					ImageSpec: v1alpha1.ImageSpec{
						Image:      "mofed",
						Repository: "nvcr.io/mellanox",
					},
				},
			},
		}

		It("generates old image format", func() {
			cr.Spec.OFEDDriver.Version = "5.6-1.0.0.0"
			imageName := stateOfed.getMofedDriverImageName(cr, nodeAttr)
			Expect(imageName).To(Equal("nvcr.io/mellanox/mofed-5.6-1.0.0.0:ubuntu20.04-amd64"))
		})
		It("generates new image format", func() {
			cr.Spec.OFEDDriver.Version = "5.7-1.0.0.0"
			imageName := stateOfed.getMofedDriverImageName(cr, nodeAttr)
			Expect(imageName).To(Equal("nvcr.io/mellanox/mofed:5.7-1.0.0.0-ubuntu20.04-amd64"))
		})
		It("generates new image format double digit minor", func() {
			cr.Spec.OFEDDriver.Version = "5.10-0.0.0.1"
			imageName := stateOfed.getMofedDriverImageName(cr, nodeAttr)
			Expect(imageName).To(Equal("nvcr.io/mellanox/mofed:5.10-0.0.0.1-ubuntu20.04-amd64"))
		})
		It("return new image format in case of a bad version", func() {
			cr.Spec.OFEDDriver.Version = "1.1.1.1.1"
			imageName := stateOfed.getMofedDriverImageName(cr, nodeAttr)
			Expect(imageName).To(Equal("nvcr.io/mellanox/mofed:1.1.1.1.1-ubuntu20.04-amd64"))
		})
	})
})
