package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 5
const delay = 5

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		// Fluxo com IF
		// if comando == 1 {
		// 	fmt.Println("Monitorando...")
		// } else if comando == 2 {
		// 	fmt.Println("Exibindo Logs...")
		// } else if comando == 0 {
		// 	fmt.Println("Saindo...")
		// } else {
		// 	fmt.Println("Comando incorreto.")
		// }

		// Fluxo com SWITCH
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Comando incorreto.")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Marcos"
	versao := 1.1

	fmt.Println("Ola sr.", nome)
	fmt.Println("Este programa esta na versao", versao)

}

func exibeMenu() {
	fmt.Println("\n1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	// fmt.Println("Endereco de comando", &comando)
	fmt.Println("Comando:", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com.br", "https://www.caelum.com.br"}
	sites := leSitesDosArquivo()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println()
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println("Testando site:", site)
	if resp != nil {
		if resp.StatusCode == 200 {
			fmt.Println("Site: [", site, "] carregado com sucesso!")
			registraLog(site, true)
		} else {
			fmt.Println("Site: [", site, "] com problema. Status Code:", resp.StatusCode)
			registraLog(site, false)
		}
	} else {
		fmt.Println("Site: [", site, "] com problema. Erro:", err)
		registraLog(site, false)
	}
}

func leSitesDosArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return nil
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')

		sites = append(sites, strings.TrimSpace(linha))
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
		return
	}

	arquivo.WriteString(time.Now().Format("2006-01-02 15:04:05.000000") + " | " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		return
	}

	fmt.Println("\nExibindo Logs...")
	fmt.Println(string(arquivo))
}
