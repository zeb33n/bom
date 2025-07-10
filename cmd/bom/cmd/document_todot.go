/*
Copyright 2023 The Kubernetes Authors.

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

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"sigs.k8s.io/bom/pkg/spdx"
)

func AddToDot(parent *cobra.Command) {
	toDotOpts := &spdx.ToDotOptions{}
	toDotCmd := &cobra.Command{
		PersistentPreRunE: initLogging,
		Short:             "dump SPDX as a dot file",
		Long:              `TODO!`,
		Use:               "todot SPDX_FILE|URL",
		SilenceUsage:      true,
		SilenceErrors:     true,
		RunE: func(_ *cobra.Command, args []string) error {
			if len(args) == 0 {
				args = append(args, "")
			}
			doc, err := spdx.OpenDoc(args[0])
			if err != nil {
				return fmt.Errorf("opening doc: %w", err)
			}
			fmt.Println(doc.ToDot(toDotOpts))
			return nil
		},
	}
	toDotCmd.PersistentFlags().StringVarP(
		&toDotOpts.Find,
		"find",
		"f",
		"",
		"Find node in DAG",
	)
	toDotCmd.PersistentFlags().IntVarP(
		&toDotOpts.Recursion,
		"depth",
		"d",
		-1,
		"Depth to traverse",
	)
	toDotCmd.PersistentFlags().StringVarP(
		&toDotOpts.SubGraphRoot,
		"subgraph",
		"s",
		"",
		"SPXID of the root node for the subgraph",
	)
	parent.AddCommand(toDotCmd)
}
