package cli

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/cs23m001/cipherc2/client/constants"
	"github.com/cs23m001/cipherc2/protobuf/clientpb"
	"github.com/cs23m001/cipherc2/server/assets"
	"github.com/cs23m001/cipherc2/server/c2"
	"github.com/cs23m001/cipherc2/server/certs"
	"github.com/cs23m001/cipherc2/server/console"
	"github.com/cs23m001/cipherc2/server/cryptography"
	"github.com/cs23m001/cipherc2/server/daemon"
	"github.com/cs23m001/cipherc2/server/db"
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Force start server in daemon mode",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		force, err := cmd.Flags().GetBool(forceFlagStr)
		if err != nil {
			fmt.Printf("Failed to parse --%s flag %s\n", forceFlagStr, err)
			return
		}
		lhost, err := cmd.Flags().GetString(lhostFlagStr)
		if err != nil {
			fmt.Printf("Failed to parse --%s flag %s\n", lhostFlagStr, err)
			return
		}
		lport, err := cmd.Flags().GetUint16(lportFlagStr)
		if err != nil {
			fmt.Printf("Failed to parse --%s flag %s\n", lportFlagStr, err)
			return
		}

		tailscale, err := cmd.Flags().GetBool(tailscaleFlagStr)
		if err != nil {
			fmt.Printf("Failed to parse --%s flag %s\n", tailscaleFlagStr, err)
			return
		}

		appDir := assets.GetRootAppDir()
		logFile := initConsoleLogging(appDir)
		defer logFile.Close()

		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic:\n%s", debug.Stack())
				fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
				os.Exit(99)
			}
		}()

		assets.Setup(force, false)
		c2.SetupDefaultC2Profiles()
		certs.SetupCAs()
		certs.SetupWGKeys()
		cryptography.AgeServerKeyPair()
		cryptography.MinisignServerPrivateKey()

		listenerJobs, err := db.ListenerJobs()
		if err != nil {
			fmt.Println(err)
		}

		err = StartPersistentJobs(listenerJobs)
		if err != nil {
			fmt.Println(err)
		}

		daemon.Start(lhost, uint16(lport), tailscale)
	},
}

func StartPersistentJobs(listenerJobs []*clientpb.ListenerJob) error {
	if len(listenerJobs) > 0 {
		// StartPersistentJobs - Start persistent jobs
		for _, j := range listenerJobs {
			listenerJob, err := db.ListenerByJobID(j.JobID)
			if err != nil {
				return err
			}
			switch j.Type {
			case constants.HttpStr:
				job, err := c2.StartHTTPListenerJob(listenerJob.HTTPConf)
				if err != nil {
					return err
				}
				j.JobID = uint32(job.ID)
			case constants.HttpsStr:
				job, err := c2.StartHTTPListenerJob(listenerJob.HTTPConf)
				if err != nil {
					return err
				}
				j.JobID = uint32(job.ID)
			case constants.MtlsStr:
				job, err := c2.StartMTLSListenerJob(listenerJob.MTLSConf)
				if err != nil {
					return err
				}
				j.JobID = uint32(job.ID)
			case constants.WGStr:
				job, err := c2.StartWGListenerJob(listenerJob.WGConf)
				if err != nil {
					return err
				}
				j.JobID = uint32(job.ID)
			case constants.DnsStr:
				job, err := c2.StartDNSListenerJob(listenerJob.DNSConf)
				if err != nil {
					return err
				}
				j.JobID = uint32(job.ID)
			case constants.MultiplayerModeStr:
				id, err := console.JobStartClientListener(listenerJob.MultiConf)
				if err != nil {
					return err
				}
				j.JobID = uint32(id)
			}
			db.UpdateHTTPC2Listener(j)
		}
	}

	return nil
}
