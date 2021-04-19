package types




type Group struct {
	Group     string   `json:"group"`
	WorkSpace []string `json:"workSpace"`
}

type Log struct {
	Up        string `json:"up"`
	Down      string `json:"down"`
	IsVisible bool   `json:"isVisible"`
}

type SpaceNdCalc struct {
	WorkSpace string
	Sum       int64
}
type Result struct {
	GroupName string
	Sum       int64
	SumGroup  []SpaceNdCalc
}
