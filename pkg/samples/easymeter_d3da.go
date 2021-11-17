package samples

var (
	Easymeter_D3DA = "/ESY5Q3DA1024 V3.03\r\n" +
		"\r\n" +
		"1-0:0.0.0*255(111940207)\r\n" +
		"1-0:1.8.0*255(00000000.7801269*kWh)\r\n" +
		"1-0:2.8.0*255(00007041.2292533*kWh)\r\n" +
		"1-0:21.7.255*255(-000073.69*W)\r\n" +
		"1-0:41.7.255*255(000000.00*W)\r\n" +
		"1-0:61.7.255*255(-000068.61*W)\r\n" +
		"1-0:1.7.255*255(-000142.30*W)\r\n" +
		"1-0:96.5.5*255(A2)\r\n" +
		"0-0:96.1.255*255(1ESY1127000167)\r\n" +
		"!\r\n"
)

func init() {
	All["easymeter_d3da"] = Easymeter_D3DA
}
