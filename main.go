package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	origemPath  = `\\192.168.10.01\exemplo\exemplo\teste\`
	destinoPath = `\\192.168.10.02\bancos\exemplo\RETORNO`
)

func main() {
	if err := verificarDiretorios(); err != nil {
		log.Fatal("Erro ao verificar diretórios:", err)
	}

	// Configurar log para arquivo
	logFile, err := os.OpenFile("transfer_log.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Erro ao criar arquivo de log:", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	fmt.Println("Monitorando arquivos .RET...")
	log.Println("Iniciando monitoramento de arquivos .RET")

	for {
		// Verificar diretório de origem
		arquivos, err := os.ReadDir(origemPath)
		if err != nil {
			log.Printf("Erro ao ler diretório de origem: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		// Se encontrar arquivos, transferir apenas .RET
		if len(arquivos) > 0 {
			for _, arquivo := range arquivos {
				if !arquivo.IsDir() && strings.ToUpper(filepath.Ext(arquivo.Name())) == ".RET" {
					origem := filepath.Join(origemPath, arquivo.Name())
					destino := filepath.Join(destinoPath, arquivo.Name())

					err := transferirArquivo(origem, destino)
					if err != nil {
						log.Printf("Erro ao transferir arquivo %s: %v", arquivo.Name(), err)
						continue
					}

					// Remover arquivo original após transferência bem-sucedida
					err = os.Remove(origem)
					if err != nil {
						log.Printf("Erro ao remover arquivo original %s: %v", arquivo.Name(), err)
						continue
					}

					mensagem := fmt.Sprintf("Arquivo %s transferido com sucesso", arquivo.Name())
					log.Println(mensagem)
					fmt.Println(mensagem)
				}
			}
		}

		// Aguardar antes de verificar novamente
		time.Sleep(10 * time.Second)
	}
}

func transferirArquivo(origem, destino string) error {
	// Abrir arquivo de origem
	sourceFile, err := os.Open(origem)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo de origem: %v", err)
	}
	defer sourceFile.Close()

	// Criar arquivo de destino
	destinoFile, err := os.Create(destino)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de destino: %v", err)
	}
	defer destinoFile.Close()

	// Copiar conteúdo
	_, err = io.Copy(destinoFile, sourceFile)
	if err != nil {
		return fmt.Errorf("erro ao copiar arquivo: %v", err)
	}

	return nil
}

func verificarDiretorios() error {
	// Verificar diretório de origem
	if _, err := os.Stat(origemPath); os.IsNotExist(err) {
		return fmt.Errorf("diretório de origem não existe: %v", err)
	}

	// Verificar diretório de destino
	if _, err := os.Stat(destinoPath); os.IsNotExist(err) {
		return fmt.Errorf("diretório de destino não existe: %v", err)
	}

	return nil
}
