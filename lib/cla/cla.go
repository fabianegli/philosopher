// Package cla (Classification) provides methods to verify if different types
// of evidences are target or decoys
package cla

import (
	"strings"

	"philosopher/lib/id"
)

// IsDecoyPSM identifies a PSM as target or Decoy based on the
// presence of the TAG string on <protein> and <alternative_proteins>
func IsDecoyPSM(p id.PeptideIdentification, tag string) bool {

	// default for TRUE (DECOY)
	// updated to FALSE
	var class bool

	if strings.HasPrefix(p.Protein, tag) {
		class = true
	} else {
		class = false
	}

	// try to find another protein, indistinguishable, that is annotate as target
	// only one evidence is enough to promote the PSM as a "no-decoy"
	if len(p.AlternativeProteins) > 1 {
		for i := range p.AlternativeProteins {
			if !strings.HasPrefix(i, tag) {
				class = false
				break
			}
		}
	}

	return class
}

// IsDecoyProtein identifies a Protein as target or Decoy based on the decoy tag
func IsDecoyProtein(p id.ProteinIdentification, tag string) bool {

	// default for TRUE (DECOY)
	// updated to FALSE
	var class bool

	if strings.HasPrefix(string(p.ProteinName), tag) {
		class = true
	} else {
		class = false
	}

	return class
}

// IsDecoy identifies a Protein as target or Decoy based on the decoy tag
func IsDecoy(name string, tag string) bool {

	// default for TRUE (DECOY)
	// updated to FALSE
	var class bool

	if strings.HasPrefix(name, tag) {
		class = true
	} else {
		class = false
	}

	return class
}

// IsDecoyIon identifies a Protein as target or Decoy based on the decoy tag
func IsDecoyIon(names map[string]int, tag string) bool {

	// default for TRUE ( DECOY)
	var class = true

	for i := range names {
		if strings.HasPrefix(i, tag) {
			class = true
		} else {
			class = false
			break
		}
	}

	return class
}
