package sgw

func Run() {

}

/*
var load = loader.EnvLoader{
	PebbleOptions: &loader.CmdOption{
		HealthCheckURL: "https://localhost:15000/dir",
		Args:           []string{"-strict", "-config", "fixtures/pebble-config-dns.json", "-dnsserver", "localhost:8053"},
		Env:            []string{"PEBBLE_VA_NOSLEEP=1", "PEBBLE_WFE_NONCEREJECT=20"},
		Dir:            "../",
	},
	LegoOptions: []string{
		"LEGO_CA_CERTIFICATES=../fixtures/certs/pebble.minica.pem",
		"EXEC_PATH=../fixtures/update-dns.sh",
	},
	ChallSrv: &loader.CmdOption{
		Args: []string{"-http01", ":5012", "-tlsalpn01", ":5011"},
	},
}

*/
