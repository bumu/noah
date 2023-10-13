package configkit

// Shell command: ls -l apps/ | awk  'NR>1 {printf("\"%s\", ", $NF)}'
var Apps = []string{
	"apigw",
	"audit",
	"delivery",
	"mall",
	"notification",
	"payment",
	"promotion",
	"sgw",
	"strategy",
	"supply",
	"user",
	"yummy",
}
