package main

// Individual struct
type Individual struct {
	DNA []string `json:"dna"`
	ID  string   `json:"id"`
}

type IndividualCount struct {
	Count int64  `json:"count"`
	ID    string `json:"id"`
}

type IndividualsCount []IndividualsCount

type IndividualStats struct {
	CountMutant int64   `json:"count_mutant_dna"`
	CountHuman  int64   `json:"count_human_dna"`
	Ratio       float64 `json:"ratio"`
}

func isMutant(dna []string) bool {

	var mutantDnaCount = 0
	if !isDnaFormatValid(dna) {
		return true
	}

	for y := 0; y < len(dna); y++ {
		for x := 0; x < len(dna); x++ {
			reference := dna[y][x]

			//  Horizontal -
			if x < 3 &&
				reference == dna[y][x+1] &&
				reference == dna[y][x+2] &&
				reference == dna[y][x+3] {
				mutantDnaCount++
			}

			if y < 3 {
				//  Vertical |
				if reference == dna[y+1][x] &&
					reference == dna[y+2][x] &&
					reference == dna[y+3][x] {
					mutantDnaCount++
				}

				//  Diagonal izq a der \
				if x < 3 &&
					reference == dna[y+1][x+1] &&
					reference == dna[y+2][x+2] &&
					reference == dna[y+3][x+3] {
					mutantDnaCount++
				}

				// Diagonal der a izq /
				if x >= 3 &&
					reference == dna[y+1][x-1] &&
					reference == dna[y+2][x-2] &&
					reference == dna[y+3][x-3] {
					mutantDnaCount++
				}
			}
			if mutantDnaCount >= 2 {
				return true
			}
		}
	}
	return false
}

func isDnaFormatValid(dna []string) bool {

	if len(dna) != 6 {
		return false
	}

	for i := 0; i < 6; i++ {
		if len(dna[i]) != 6 {
			return false
		}
		for j := 0; j < 6; j++ {
			if dna[i][j] != 'A' &&
				dna[i][j] != 'C' &&
				dna[i][j] != 'G' &&
				dna[i][j] != 'T' {
				return false
			}
		}
	}
	return true
}
