package samples

import (
	_ "embed"
)

var All map[string]*[]byte = make(map[string]*[]byte)

//go:embed easymeter_d3da
var Easymeter_D3DA []byte

//go:embed enexis_t210d
var Enexis_T210D []byte

//go:embed hager_ehz161wa
var Hager_eHZ161WA []byte

//go:embed iskra_am550t_1011
var IskraAM550_1011 []byte

//go:embed iskra_am550t_1012
var IskraAM550_1012 []byte

//go:embed iskra_mt382_1000_dsmr3
var IskraMT382_1000_DSMRv3 []byte

//go:embed iskra_mt382_1000_dsmr5
var IskraMT382_1000_DSMRv5 []byte

//go:embed iskra_mt382_1003
var IskraMT382_1003 []byte

//go:embed iskra_mt382_1004
var IskraMT382_1004 []byte

//go:embed kaifa_ma105
var Kaifa_MA105 []byte

//go:embed kamstrup_162
var Kamstrup_162 []byte

//go:embed kamstrup_mc66
var Kamstrup_Multical66 []byte

//go:embed landis_gir_zcf100
var LandisGirZCF100 []byte

//go:embed ziv_5cta3
var Ziv_5CTA3 []byte

func init() {
	All["easymeter_d3da"] = &Easymeter_D3DA
	All["enexis_t210d"] = &Enexis_T210D
	All["hager_ehz161wa"] = &Hager_eHZ161WA
	All["iskra_am550t_1011"] = &IskraAM550_1011
	All["iskra_am550t_1012"] = &IskraAM550_1012
	All["iskra_mt382_1000_dsmr3"] = &IskraMT382_1000_DSMRv3
	All["iskra_mt382_1000_dsmr5"] = &IskraMT382_1000_DSMRv5
	All["iskra_mt382_1003"] = &IskraMT382_1003
	All["iskra_mt382_1004"] = &IskraMT382_1004
	All["kaifa_ma105"] = &Kaifa_MA105
	All["kamstrup_162"] = &Kamstrup_162
	All["kamstrup_mc66"] = &Kamstrup_Multical66
	All["landis_gir_zcf100"] = &LandisGirZCF100
	All["ziv_5cta3"] = &Ziv_5CTA3
}
