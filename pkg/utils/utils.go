package utils

import "database/sql"

func NewNullString(in string) sql.NullString {
	if in == "" {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		return sql.NullString{
			String: in,
			Valid:  true,
		}
	}
}

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// RestToBusiness из тенге в тиины
func RestToBusiness(f float64) int {
	return int(f * 100)
}

// RestToBusiness из тенге в тиины
func BusinessToRest(i int) float64 {
	if i == 0 {
		return 0
	}
	return float64(i) / 100
}

// RestToBusiness из тенге в тиины
func BusinessToRestP(i *int) *float64 {
	if i == nil {
		return nil
	}
	var out float64
	in := *i
	if in == 0 {
		out = 0
	} else {
		out = float64(in) / 100
	}
	return &out
}
