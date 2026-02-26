package cmd

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"

	"github.com/darknamer/black-vault-lib"
	"github.com/spf13/cobra"
)

const maxPortTries = 20 // ลองได้สูงสุด 20 พอร์ตถ้าพอร์ตเริ่มต้นถูกใช้

var (
	servePort int
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run gRPC server for GUI (default :50051, try next port if in use)",
	Long:  "Starts the BlackVault gRPC server. If --port is in use, tries port+1, port+2, ... and writes the chosen port to ~/.blackvault/grpc_port for the GUI.",
	RunE:  runServe,
}

func init() {
	serveCmd.Flags().IntVar(&servePort, "port", 50051, "gRPC listen port (try next ports if in use)")
}

func runServe(cmd *cobra.Command, args []string) error {
	portFilePath := blackvault.GrpcPortFilePath()
	_ = os.MkdirAll(filepath.Dir(portFilePath), 0755)

	var lis net.Listener
	var chosenPort int
	for try := 0; try < maxPortTries; try++ {
		p := servePort + try
		addr := ":" + strconv.Itoa(p)
		var err error
		lis, err = net.Listen("tcp", addr)
		if err == nil {
			chosenPort = p
			break
		}
		if try == maxPortTries-1 {
			return fmt.Errorf("could not bind any port from %d to %d: %w", servePort, servePort+try, err)
		}
	}
	defer lis.Close()

	if err := os.WriteFile(portFilePath, []byte(strconv.Itoa(chosenPort)), 0600); err != nil {
		return fmt.Errorf("write port file: %w", err)
	}
	defer os.Remove(portFilePath)

	addr := ":" + strconv.Itoa(chosenPort)
	if chosenPort != servePort {
		fmt.Printf("Port %d in use, listening on %s instead\n", servePort, addr)
	}
	fmt.Printf("gRPC server placeholder listening on %s (port file: %s)\n", addr, portFilePath)
	fmt.Println("To enable full gRPC: install protoc, run 'make proto', then rebuild.")
	select {} // รออยู่ที่นี่; ในเวอร์ชันจริงใช้ s.Serve(lis)
}
