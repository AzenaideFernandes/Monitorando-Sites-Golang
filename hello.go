package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"os"
	"time"
	//"reflect"
	"net/http"
)

const monitoramentos = 2
const delay = 5

func main() {
	// var nome string = "Douglas"
	// var versao float32 = 1.1 
	// var idade int = 22

	// var nome  = "Douglas"
	// var versao  = 1.1 
	// var idade  = 22

	//  nome  := "Douglas"
	//  versao  := 1.1 
	//  idade  := 22

	// fmt.Println("Olá Sr.", nome, " sua idade é", idade) 
	// fmt.Println("Este programa está na versão", versao)
	// fmt.Println("O tipo de variavel idade é", reflect.TypeOf(versao))	

	exibeIntroducao()

	//exibeNomes() 

	for {

	exibeMenu()
	// nome, idade := devolveNomeEIdade()
  // fmt.Println(nome, "tem", idade, "anos")

	// fmt.Println("1 - Iniciar o Monitoramento")
	// fmt.Println("2 - Exibir os Logs")
	// fmt.Println("0 - Sair do Programa")

	comando := leComando()

	//var comando int
	//fmt.Scanf("%d",&comando)
	//fmt.Scan(&comando)

	//fmt.Println("O endereço da minha varaivel é", &comando)
	//fmt.Println("O comando escolhido foi", comando)

	//exibeNomes()

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// }else if comando == 2 {
	// 	fmt.Println("Exibindo Logs...")
	// }else if comando == 0 {
	// 	fmt.Println("Saindo do programa...")
	// }else {
	// 	fmt.Println("Não conheno este comando")
	// }

	switch comando  {
	case 1:
		//fmt.Println("Monitorando...")
		inciandoMonitoramento()
	case 2:	
		fmt.Println("Exibindo Logs...")
		imprimeLogs() 
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheno este comando")
		os.Exit(-1)
	}
}
}

func exibeIntroducao() {
	fmt.Println("")
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("")
}

func exibeMenu() {
	fmt.Println("1 - Iniciar o Monitoramento")
	fmt.Println("2 - Exibir os Logs")
	fmt.Println("0 - Sair do Programa")

}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println(" ")
	fmt.Println("O comando escolhido foi", comandoLido)
	fmt.Println(" ")
	return comandoLido
}

func inciandoMonitoramento(){
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()
	fmt.Println(" ")
	// sites := []string{"https://random-status-code.herokuapp.com/",
  // "https://www.alura.com.br", "https://www.caelum.com.br"}
	// var sites [4]string
	// sites[0] = "https://random-status-code.herokuapp.com/"
  // sites[1] = "https://www.alura.com.br"
  // sites[2] = "https://www.caelum.com.br"
 
	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }

	// for i, site := range sites {
	// 	fmt.Println("Estou passa0ndo no posição", i, "do meu slice e essa posição tem o site", site)
	// }
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando", i, ":", site)
			fmt.Println(" ")
			testaSite(site)	
			fmt.Println(" ")		
		}
		time.Sleep(delay * time.Second)
		fmt.Println(" ")
	}


	fmt.Println(" ")

		//site := "https://www.alura.com.br"	
		//resp, err := http.Get(site)
		
}
func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return
	}

	fmt.Println(resp)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	}else {
		fmt.Println("Site", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {

	var sites []string

	//arquivo, err := ioutil.ReadFile("sites.txt")
	arquivo, err := os.Open("sites.txt")
	
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	//fmt.Println(string(arquivo))

	leitor := bufio.NewReader(arquivo)

	//leitor.ReadString('\n')

	for {
				linha, err := leitor.ReadString('\n')
				linha = strings.TrimSpace(linha)
				sites = append(sites, linha)
				//fmt.Println(linha)
				if err == io.EOF {
					break
				}			
			}

		arquivo.Close()
		return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	fmt.Println(arquivo)
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " _ " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()

}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Ocorreu um eroo", err)
	}
	fmt.Println(string(arquivo))
}



// func devolveNomeEIdade() (string, int)  {
// nome := "Douglas"
// idade := 24
// return nome, idade
// }

//https://golang.org/pkg/  pacotes da linguagem

// func exibeNomes() { //exemplo de slice
// 	nomes := []string{"Douglas", "Daniel", "Bernardo"}
// 	fmt.Println("O meu slice tem", len(nomes), "itens")
// 	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

// 	nomes = append(nomes, "Aparecida")
// 	fmt.Println("O meu slice tem", len(nomes), "itens")
// 	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
// }
//--------------------------------------
// func main() {
//     estados := devolveEstadosDoSudeste()
//     fmt.Println(estados)
// }

// func devolveEstadosDoSudeste() [4]string {
//     var estados [4]string
//     estados[0] = "RJ"
//     estados[1] = "SP"
//     estados[2] = "MG"
//     estados[3] = "ES"
//     return estados
// }

//---------------------------------------

// func exibeNomes() {
// 	nomes := []string{"Douglas", "Daniel", "Bernardo"}
// 	fmt.Println(nomes)
// }