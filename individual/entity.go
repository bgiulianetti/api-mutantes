package individual

// Individual struct
type Individual struct {
	DNA []string `json:"dna"`
	ID  string   `json:"id"`
}

// DTO ...
type DTO struct {
	ID string `json:"id"`
}

// Count ...
type Count struct {
	Count int64  `json:"count"`
	ID    string `json:"id"`
}

// IndividualsCount ...
type IndividualsCount []IndividualsCount

// Stats ...
type Stats struct {
	CountMutant float64 `json:"count_mutant_dna"`
	CountHuman  float64 `json:"count_human_dna"`
	Ratio       float64 `json:"ratio"`
}
