package main

// pacote fmt para formatação (entrada e saída)
// pacote reflect para usar o typeOf
// pacote os para interação com o OS e para trabalhar com arquivos
// pacote io/util para manipulação de arquivos
// pacote bufio (buffer IO) para manipulação de dados de arquivo
// pacote net/http para fazer requisições web
// pacote time para função sleep
// strconv converter para string (boolean)

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// https://pkg.go.dev/std

// constantes

const monitoramentos = 3
const delay = 5 //segundos!

func main() {

	//leSitesdoArquivo()
	//registraLog("site-falso", true)
	// exibeIntroducao()

	/*
		if comando == 1 {
			fmt.Println("Monitorando...")
		}else if comando == 2 {
			fmt.Println("Exibindo logs...")
		}else if comando == 0 {
			fmt.Println("Até mais.")
		}else{
			fmt.Println("Comando desconhecido")
		}
	*/

	// não tem WHILE no GO, para isso, usa-se um FOR sem parâmetros
	for {

		exibeMenu()

		// já chama via função
		// ou comando := leComando()

		switch leComando() {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Até mais!")
			os.Exit(0)
		default:
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}

	}

}

// funções que retornam mais de um valor
func devolveNomeEIdade() (string, int) {
	nome := "Juliano Vieira"
	idade := 36
	return nome, idade
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// array/vetor
	// var sites [4]string

	// slices é um array melhorado abstraído, sem precisar definir o tamanho
	// sites := []string{"https://www.unifique.com.br", "https://random-status-code.herokuapp.com", "https://www.google.com"}

	// vamos ler os sites de um arquivo!

	sites := leSitesdoArquivo()

	//site := "https://www.unifique.com.br"
	//site := "https://random-status-code.herokuapp.com"

	//for i := 0; i < len(sites); i++ { ... ou com RANGE

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site:", i)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("\naguardando...")
	}
}

func leSitesdoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	// imprimir direto
	// casting para converter o array de bytes para string ao usar o ioutil.ReadFil
	// fmt.Println(string(arquivo))

	// ler linha a linha com buffer IO
	leitor := bufio.NewReader(arquivo)
	// aqui vai a limitação, o separador! com aspas simples pois é um byte
	// aqui vai ler a PRIMEIRA linha apenas

	// for até chegar no err EndOfFile!
	for {
		linha, err := leitor.ReadString('\n')
		// remover a quebra de linha e espaços no final das strings
		// nota importante, atribuição simples usa só =, o := é para a declaração, apenas
		linha = strings.TrimSpace(linha)
		// adicionar as string no array (slice) de sites usa-se o append
		sites = append(sites, linha)
		//fmt.Println(linha)
		// break para sair do for caso o erro for igual ao EOF, tipo um DO..WHILE
		if err == io.EOF {
			break
		}
	}

	// fechar o arquivo depois de usar!
	arquivo.Close()

	return sites
}

func imprimeLogs() {
	fmt.Println("Exibindo logs...")

	// esse readFile já fecha o arquivo automaticamente depois que termina de ler
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Deu um erro: ", err)
	}
	// string() pode converter um string de bytes para string
	fmt.Println(string(arquivo))

}

func registraLog(site string, status bool) {

	// abrir arquivo, mas com mais métodos, como de criar arquivo inexistente, tem argumentos adicionais, como Read Write and Create e permissões
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Deu erro:", err)
	}

	// https://go.dev/src/time/format.go
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	//fmt.Println(arquivo)
	arquivo.Close()

}

func testaSite(site string) {
	// funções que retornam mais de um valor, ver função 'devolveNomeEIdade'
	// usar _ no lugar da variável que se deseja ignorar no retorno da função que retorna múltiplos valores
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro na função http.Get:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "está ONLINE")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está OFFLINE, status:", resp.StatusCode)
		registraLog(site, false)
	}

}

func exibeMenu() {
	fmt.Println("")
	fmt.Println("Menu")
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	// %d para formatar a entrada como um decimal inteiro
	// & indica que é o endereço da variável (ponteiro), se eu usar só 'comando', ela vai representar o valor de 'comando', que será 0
	// fmt.Scanf("%d", &comando)
	// ou, sem o %d
	fmt.Print("Informe uma opção:")
	fmt.Scan(&comandoLido)
	return comandoLido

}

func exibeIntroducao() {
	// só precisa tipar se não houver atribuição
	// var nome = "Juliano" << funciona
	// var também é opcional, mas precisa usar := como atribuição
	// nome := "Juliano"

	var nome string = "Juliano"
	var versao float32 = 1.1
	var idade int = 36
	fmt.Println("Olá, senhor", nome, "sua idade é", idade, "anos")
	fmt.Println("Programa está na versão", versao)

	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
}

// go build arquivo.go - para buildar o EXE
// go run arquivo.go - para compilar e já executar!
