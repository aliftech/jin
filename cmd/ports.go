package main

import (
	"strconv"
	"strings"

	"github.com/aliftech/jin/internal/bootstrap"

	"github.com/spf13/cobra"
)

var (
	portJSON bool
	portList string // e.g., "80,443,8080"
)

var portsCmd = &cobra.Command{
	Use:   "ports <host>",
	Short: "Scan for open ports on a host",
	Long:  `Performs a TCP connect scan on common ports (or custom list) of a given host.`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		host := args[0]

		// Parse custom ports if provided
		var customPorts []int
		if portList != "" {
			parts := strings.Split(portList, ",")
			for _, p := range parts {
				portNum, err := strconv.Atoi(strings.TrimSpace(p))
				if err != nil {
					return err
				}
				customPorts = append(customPorts, portNum)
			}
		}

		handler := bootstrap.NewPortScanHandler()
		return handler.HandleScan(host, portJSON, customPorts)
	},
}

func init() {
	portsCmd.Flags().BoolVarP(&portJSON, "json", "j", false, "Output as JSON")
	portsCmd.Flags().StringVarP(&portList, "ports", "p", "", "Custom ports to scan (comma-separated)")
}
