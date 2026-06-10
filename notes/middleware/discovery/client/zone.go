package client


// Zone zone scheduler info.
type Zone struct {
	Src string           `json:"src"`
	Dst map[string]int64 `json:"dst"`
}


