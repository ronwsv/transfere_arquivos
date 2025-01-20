# transfere_arquivos
Codigo em go para monitorar e garantir transferencia de arquivos entre 2 diretorios, exemplo de aplicação quando se tem algum programa que gera bilhetagem que se tem que mandar para outra pasta esse script serve para isso.
Documentação do TransferFile/main.go
Visão Geral
Este programa monitora um diretório específico em busca de arquivos com extensão .RET e os transfere automaticamente para um diretório de destino.
Constantes
startLine: 13
endLine: 16
origemPath: Caminho do diretório de origem onde os arquivos serão monitorados
destinoPath: Caminho do diretório de destino para onde os arquivos serão transferidos
Funções Principais
1. main()
startLine: 18
endLine: 73

Função principal que:
Verifica se os diretórios existem
Configura o sistema de logs
Inicia o loop infinito de monitoramento
Processa arquivos .RET encontrados
3. transferirArquivo(origem, destino string) error
   startLine: 75
    endLine: 97
Responsável por:
Abrir o arquivo de origem
Criar o arquivo no destino
Copiar o conteúdo
Retornar erro em caso de falha
4. verificarDiretorios() error
5. startLine: 99
endLine: 111
Verifica se:
O diretório de origem existe
O diretório de destino existe
Retorna erro caso algum não exista
Fluxo de Execução
Inicialização
Verifica diretórios
Configura arquivo de log
Inicia monitoramento
Loop Principal
Verifica diretório a cada 10 segundos
Filtra apenas arquivos .RET
Transfere arquivos encontrados
Remove arquivo original após transferência
Tratamento de Erros
Log de erros em arquivo
Mensagens no console
Continuação do processo mesmo após erros
Logs
O programa mantém dois tipos de output:
Console: Mensagens em tempo real
Arquivo transfer_log.txt: Histórico completo de operações
Requisitos
Go 1.23.4 ou superior
Acesso aos diretórios de rede especificados
Permissões de leitura/escrita nos diretórios
Como Executar
cd TransferFile
go mod init transferFile
go run main.go
Observações
O programa é case-insensitive para extensões (.ret, .RET, .Ret)
Aguarda 10 segundos entre verificações
Mantém logs detalhados de todas as operações
Continua executando até ser interrompido manualmente
