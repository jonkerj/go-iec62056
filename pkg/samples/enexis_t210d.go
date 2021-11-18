package samples

var (
	Enexis_T210D = "/Ene5\\T210-D ESMR5.0\r\n" +
		"\r\n" +
		"1-3:0.2.8(50)\r\n" +
		"0-0:1.0.0(190816125118S)\r\n" +
		"0-0:96.1.1(4530303438303030303335383536323139)\r\n" +
		"1-0:1.8.1(000001.200*kWh)\r\n" +
		"1-0:1.8.2(000001.766*kWh)\r\n" +
		"1-0:2.8.1(000000.092*kWh)\r\n" +
		"1-0:2.8.2(000000.000*kWh)\r\n" +
		"0-0:96.14.0(0002)\r\n" +
		"1-0:1.7.0(00.479*kW)\r\n" +
		"1-0:2.7.0(00.000*kW)\r\n" +
		"0-0:96.7.21(00007)\r\n" +
		"0-0:96.7.9(00003)\r\n" +
		"1-0:99.97.0(0)(0-0:96.7.19)\r\n" +
		"1-0:32.32.0(00002)\r\n" +
		"1-0:52.32.0(00002)\r\n" +
		"1-0:72.32.0(00002)\r\n" +
		"1-0:32.36.0(00000)\r\n" +
		"1-0:52.36.0(00000)\r\n" +
		"1-0:72.36.0(00000)\r\n" +
		"0-0:96.13.0()\r\n" +
		"1-0:32.7.0(241.0*V)\r\n" +
		"1-0:52.7.0(243.0*V)\r\n" +
		"1-0:72.7.0(242.0*V)\r\n" +
		"1-0:31.7.0(000*A)\r\n" +
		"1-0:51.7.0(000*A)\r\n" +
		"1-0:71.7.0(001*A)\r\n" +
		"1-0:21.7.0(00.186*kW)\r\n" +
		"1-0:41.7.0(00.012*kW)\r\n" +
		"1-0:61.7.0(00.281*kW)\r\n" +
		"1-0:22.7.0(00.000*kW)\r\n" +
		"1-0:42.7.0(00.000*kW)\r\n" +
		"1-0:62.7.0(00.000*kW)\r\n" +
		"0-1:24.1.0(003)\r\n" +
		"0-1:96.1.0(0000000000000000000000000000000000)\r\n" +
		"0-1:24.2.1(632525252525S)(00000.000)\r\n" +
		"!162E\r\n"
)

func init() {
	All["enexis_t210d"] = Enexis_T210D
}