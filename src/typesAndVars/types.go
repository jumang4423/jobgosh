package types

// group stracture
type Group struct {
	Group     string   `json:"group"`
	WorkSpace []string `json:"workSpace"`
}

// each foldeer log datas stracture
type Log struct {
	Up        string `json:"up"`
	Down      string `json:"down"`
	IsVisible bool   `json:"isVisible"`
}

// calculation struct
type SpaceNdCalc struct {
	WorkSpace string
	Sum       int64
	IsUp      bool
}

// calculation struct 2
type Result struct {
	GroupName string
	Sum       int64
	SumGroup  []SpaceNdCalc
}
