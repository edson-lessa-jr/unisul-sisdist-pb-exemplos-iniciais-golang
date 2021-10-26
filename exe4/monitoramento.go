package exe4

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

const NOME_ARQUIVO_SITES = "exe4\\lista_sites.txt"
const ARQUIVO_LOG = "exe4\\log.txt"
const QUANTIDADE_MONITORAMENTO = 5
const DELAY = 5

func MonitoramentoSites() {
Loop:
	for {
		fmt.Println("1 - Executar monitoramento")
		fmt.Println("2 - Visualizar logs")
		fmt.Println("0 - Voltar para o menu anterior")
		var comando int
		 fmt.Scan(&comando)
		switch comando {
		case 1:
			iniciarOperacao()
		case 2:
			exibirLogs()
		case 0:
			fmt.Println("Voltando...")
			break Loop
		}
	}
}

func exibirLogs() {
	fmt.Println("Exibindo logs...")
	arquivo, err := ioutil.ReadFile(ARQUIVO_LOG)
	tratamentoErros(err)
	fmt.Println(string(arquivo))
}

func iniciarOperacao() {
	var sites []string

	arquivo, err := os.Open(NOME_ARQUIVO_SITES)
	tratamentoErros(err)
	leitor := bufio.NewReader(arquivo)
	for  {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		if err == io.EOF{
			break
		} else {
			tratamentoErros(err)
		}
		fmt.Println(linha)
		sites = append(sites, linha)

	}
	arquivo.Close()
	for i := 0; i < QUANTIDADE_MONITORAMENTO; i++ {
		fmt.Println("Iniciando o monitoramento: ",i)
		for _, site := range sites{
			resposta, err := http.Get(site)
			tratamentoErros(err)
			if resposta.StatusCode ==200 {
				fmt.Printf("O site %s foi carregado com sucesso\n", site)
				registrarLog(site, true)
			} else {
				fmt.Printf("O site %s foi não foi carregado\n", site)
				registrarLog(site, false)
			}
		}
		time.Sleep(DELAY * time.Second)
		fmt.Println("....")
	}
}

func registrarLog(site string, status bool) {
	arquivo, err := os.OpenFile(ARQUIVO_LOG, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	tratamentoErros(err)
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05")+" - online: "+ strconv.FormatBool(status)+ " - "+site+"\n")
	arquivo.Close()
}

func tratamentoErros(err error) {
	if err != nil{
		fmt.Println("Ocorreu um erro ",err)
		os.Exit(-1)
	}
}