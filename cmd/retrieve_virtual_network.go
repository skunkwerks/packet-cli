// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// retrieveVirtualNetworksCmd represents the retrieveVirtualNetworks command
var retrieveVirtualNetworksCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a list of virtual networks for a single project.",
	Long: `Example:

packet virtual-network get -p [project_UUID]

	`,
	Run: func(cmd *cobra.Command, args []string) {
		vnets, _, err := PacknGo.ProjectVirtualNetworks.List(projectID, nil)
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		data := make([][]string, len(vnets.VirtualNetworks))

		for i, n := range vnets.VirtualNetworks {
			data[i] = []string{n.ID, n.Description, strconv.Itoa(n.VXLAN), n.FacilityCode, n.CreatedAt}
		}
		header := []string{"ID", "Description", "VXLAN", "Facility", "Created"}

		output(vnets, header, &data)
	},
}

func init() {
	retrieveVirtualNetworksCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "UUID of the project")
	retrieveVirtualNetworksCmd.MarkFlagRequired("project-id")
	retrieveVirtualNetworksCmd.PersistentFlags().BoolVarP(&isJSON, "json", "j", false, "JSON output")
	retrieveVirtualNetworksCmd.PersistentFlags().BoolVarP(&isYaml, "yaml", "y", false, "YAML output")
}
