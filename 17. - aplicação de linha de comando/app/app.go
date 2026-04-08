package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// maiusculo, pra poder ser exportado
// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		// comando 1
		{
			// isso aqui basicamente permite no terminal algo como go run main.go ip --host amazon.com.br
			Name:   "ip",
			Usage:  "Busca IPS de endereços na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		// segundo comando
		{
			// go run main.go servidores --host amazon.com.br
			Name:   "servidores",
			Usage:  "Busca o nome dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

// função para buscar ips
func buscarIps(c *cli.Context) {
	// vai pegar o valor da flag host do terminal e jogar nessa linha de comando
	host := c.String("host")

	// usando o pacote net pra passar o host pro metodo LookupIP pra ele, e ele vai retornar um slice de ips do host e um erro
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro) //pausa a execução
	}

	for _, ip := range ips { // _ é pra ocultar o indice
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {

	host := c.String("host")

	servidores, erro := net.LookupNS(host) //nome do servidor
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)//pra imprimir em string
	}

}
