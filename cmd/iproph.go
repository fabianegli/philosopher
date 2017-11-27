package cmd

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/prvst/philosopher/lib/err"
	"github.com/prvst/philosopher/lib/ext/interprophet"
	"github.com/prvst/philosopher/lib/sys"
	"github.com/spf13/cobra"
)

// iprophCmd represents the iproph command
var iprophCmd = &cobra.Command{
	Use:   "iprophet",
	Short: "MS/MS integrative analysis",
	//Long:  "Multi-level integrative analysis of shotgun proteomic data\niProphet v5.0",
	Run: func(cmd *cobra.Command, args []string) {

		if len(m.UUID) < 1 && len(m.Home) < 1 {
			e := &err.Error{Type: err.WorkspaceNotFound, Class: err.FATA}
			logrus.Fatal(e.Error())
		}

		logrus.Info("Executing InterProphet")
		var ipt = interprophet.New()

		// prepare binaries
		e := ipt.Deploy(m.OS, m.Distro)
		if e != nil {
			logrus.Fatal(e.Message)
		}

		// run
		xml, e := ipt.Run(m.InterProphet, m.Home, m.Temp, args)
		if e != nil {
			logrus.Fatal(e.Message)
		}

		m.InterProphet.InputFiles = args

		_ = xml
		//evi.IndexIdentification(xml, m.InterProphet.Decoy)

		m.Serialize()
		logrus.Info("Done")

		return
	},
}

func init() {

	if os.Args[1] == "iprophet" {

		m.Restore(sys.Meta())

		iprophCmd.Flags().IntVarP(&m.InterProphet.Threads, "threads", "", 1, "specify threads to use")
		iprophCmd.Flags().StringVarP(&m.InterProphet.Decoy, "decoy", "", "", "specify the decoy tag")
		iprophCmd.Flags().Float64VarP(&m.InterProphet.MinProb, "minProb", "", 0, "specify minimum probability of results to report")
		iprophCmd.Flags().StringVarP(&m.InterProphet.Output, "output", "", "iproph.pep.xml", "specify output name")
		iprophCmd.Flags().BoolVarP(&m.InterProphet.Length, "length", "", false, "use Peptide Length model")
		iprophCmd.Flags().BoolVarP(&m.InterProphet.Nofpkm, "nofpkm", "", false, "do not use FPKM model")
		iprophCmd.Flags().BoolVarP(&m.InterProphet.Nonss, "nonss", "", false, "do not use NSS model")
		iprophCmd.Flags().BoolVarP(&m.InterProphet.Nonse, "nonse", "", false, "do not use NSE model")
		iprophCmd.Flags().BoolVarP(&m.InterProphet.Nonrs, "nonrs", "", false, "do not use NRS model")
		iprophCmd.Flags().BoolVarP(&m.InterProphet.Nonsm, "nonsm", "", false, "do not use NSM model")
		iprophCmd.Flags().BoolVarP(&m.InterProphet.Nonsp, "nonsp", "", false, "do not use NSP model")
		iprophCmd.Flags().BoolVarP(&m.InterProphet.Sharpnse, "sharpnse", "", false, "Use more discriminating model for NSE in SWATH mode")
		iprophCmd.Flags().BoolVarP(&m.InterProphet.Nonsi, "nonsi", "", false, "do not use NSI model")
	}

	RootCmd.AddCommand(iprophCmd)
}
