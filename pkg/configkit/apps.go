package configkit

// Shell command: ls -l apps/ | awk  'NR>1 {printf("\"%s\", \n", $NF)}'
var Apps = []string{
	"apigw",
	"audit",
	"company",
	"default",
	"delivery",
	"infra",
	"mall",
	"notification",
	"payment",
	"promotion",
	"report",
	"security",
	"sgw",
	"strategy",
	"supply",
	"user",
	"yummy",
}
