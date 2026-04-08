package main

// aqui importaremos o pacote da pasta app, e aqui só executa
import (
	"fmt"
	"linha-de-comando/app"
	"os"
	"log"
)

func main(){
	fmt.Println("Ponto de partida")

	aplicação := app.Gerar()
	// basicamente os.Args é um parametro padrão para que os comandos do OS sejam reconhecidos pela linha de comando
	// como o Run retorna um erro, eu coloco o erro na frente pra poder fazer if depois e tratar esse erro
	erro := aplicação.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}

// go build compila o aplicativo, o que ja reduz tempo de reposta