package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/prvst/philosopher/lib/err"
	"github.com/prvst/philosopher/lib/ext/peptideprophet"
	"github.com/prvst/philosopher/lib/meta"
	"github.com/prvst/philosopher/lib/sys"
	"github.com/spf13/cobra"
)

var pep peptideprophet.PeptideProphet

// peprophCmd represents the peproph command
var peprophCmd = &cobra.Command{
	Use:   "peptideprophet",
	Short: "Peptide assignment validation",
	//Long:  "Statistical validation of peptide assignments for MS/MS Proteomics data\nPeptidProphet v5.0",
	Run: func(cmd *cobra.Command, args []string) {

		var m meta.Data
		m.Restore(sys.Meta())
		if len(m.UUID) < 1 && len(m.Home) < 1 {
			e := &err.Error{Type: err.WorkspaceNotFound, Class: err.FATA}
			logrus.Fatal(e.Error())
		}

		if len(pep.Database) < 1 {
			logrus.Fatal("You need to provide a protein database")
		}

		// deploy the binaries
		err := pep.Deploy()
		if err != nil {
			logrus.Fatal(err)
		}

		// run
		err = pep.Run(args)
		if err != nil {
			logrus.Fatal(err)
		}

		logrus.Info("Done")

		return
	},
}

func init() {

	pep = peptideprophet.New()

	peprophCmd.Flags().BoolVarP(&pep.Exclude, "exclude", "", false, "exclude deltaCn*, Mascot*, and Comet* results from results (default Penalize * results)")
	peprophCmd.Flags().BoolVarP(&pep.Leave, "leave", "", false, "leave alone deltaCn*, Mascot*, and Comet* results from results (default Penalize * results)")
	peprophCmd.Flags().BoolVarP(&pep.Icat, "icat", "", false, "apply ICAT model (default Autodetect ICAT)")
	peprophCmd.Flags().BoolVarP(&pep.Noicat, "noicat", "", false, "do no apply ICAT model (default Autodetect ICAT)")
	peprophCmd.Flags().BoolVarP(&pep.Zero, "zero", "", false, "report results with minimum probability 0")
	peprophCmd.Flags().BoolVarP(&pep.Accmass, "accmass", "", false, "use Accurate Mass model binning")
	peprophCmd.Flags().IntVarP(&pep.Clevel, "clevel", "", 0, "set Conservative Level in neg_stdev from the neg_mean, low numbers are less conservative, high numbers are more conservative")
	peprophCmd.Flags().BoolVarP(&pep.Ppm, "ppm", "", false, "use PPM mass error instead of Daltons for mass modeling")
	peprophCmd.Flags().BoolVarP(&pep.Nomass, "nomass", "", false, "disable mass model")
	peprophCmd.Flags().Float64VarP(&pep.Masswidth, "masswidth", "", 5.0, "model mass width")
	peprophCmd.Flags().BoolVarP(&pep.Pi, "pi", "", false, "enable peptide pI model")
	peprophCmd.Flags().IntVarP(&pep.Minpintt, "minpintt", "", 2, "minimum number of NTT in a peptide used for positive pI model")
	peprophCmd.Flags().Float64VarP(&pep.Minpiprob, "minpiprob", "", 0.9, "minimum probability after first pass of a peptide used for positive pI model")
	peprophCmd.Flags().BoolVarP(&pep.Rt, "rt", "", false, "enable peptide RT model")
	peprophCmd.Flags().Float64VarP(&pep.Minrtprob, "minrtprob", "", 0.9, "minimum probability after first pass of a peptide used for positive RT model")
	peprophCmd.Flags().IntVarP(&pep.Minrtntt, "minrtntt", "", 2, "minimum number of NTT in a peptide used for positive RT model")
	peprophCmd.Flags().BoolVarP(&pep.Glyc, "glyc", "", false, "enable peptide Glyco motif model")
	peprophCmd.Flags().BoolVarP(&pep.Phospho, "phospho", "", false, "enable peptide Phospho motif model")
	peprophCmd.Flags().BoolVarP(&pep.Maldi, "maldi", "", false, "enable MALDI mode")
	peprophCmd.Flags().BoolVarP(&pep.Instrwarn, "instrwarn", "", false, "warn and continue if combined data was generated by different instrument models")
	peprophCmd.Flags().Float64VarP(&pep.Minprob, "minprob", "", 0.05, "report results with minimum probability")
	peprophCmd.Flags().StringVarP(&pep.Decoy, "decoy", "", "", "semi-supervised mode, protein name prefix to identify Decoy entries")
	peprophCmd.Flags().BoolVarP(&pep.Decoyprobs, "decoyprobs", "", false, "compute possible non-zero probabilities for Decoy entries on the last iteration")
	peprophCmd.Flags().BoolVarP(&pep.Nontt, "nontt", "", false, "disable NTT enzymatic termini model")
	peprophCmd.Flags().BoolVarP(&pep.Nonmc, "nonmc", "", false, "disable NMC missed cleavage model")
	peprophCmd.Flags().BoolVarP(&pep.Expectscore, "expectscore", "", false, "use expectation value as the only contributor to the f-value for modeling")
	peprophCmd.Flags().BoolVarP(&pep.Nonparam, "nonparam", "", false, "use semi-parametric modeling, must be used in conjunction with --decoy option")
	peprophCmd.Flags().BoolVarP(&pep.Neggamma, "neggamma", "", false, "use Gamma distribution to model the negative hits")
	peprophCmd.Flags().BoolVarP(&pep.Forcedistr, "forcedistr", "", false, "bypass quality control checks, report model despite bad modelling")
	peprophCmd.Flags().BoolVarP(&pep.Optimizefval, "optimizefval", "", false, "(SpectraST only) optimize f-value function f(dot,delta) using PCA")
	peprophCmd.Flags().IntVarP(&pep.MinPepLen, "minpeplen", "", 7, "minimum peptide length not rejected")
	peprophCmd.Flags().StringVarP(&pep.Output, "output", "", "interact", "Output name prefix")
	peprophCmd.Flags().BoolVarP(&pep.Combine, "combine", "", false, "combine the results from PeptideProphet into a single result file")
	peprophCmd.Flags().StringVarP(&pep.Database, "database", "", "", "path to the database")
	//peprophCmd.Flags().BoolVarP(&pep.Perfectlib, "perfectlib", "", false, "multiply by SpectraST library probability")
	//peprophCmd.Flags().StringVarP(&pep.Rtcat, "rtcat", "", "", "enable peptide RT model, use <rtcatalog_file> peptide RTs when available as the theoretical value")
	//peprophCmd.Flags().Uint8VarP(&pep.Ignorechg, "ignorechg", "", 0, "can be used multiple times to specify all charge states to exclude from modeling")

	RootCmd.AddCommand(peprophCmd)
}
