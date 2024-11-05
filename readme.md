
# Projeto Korp - HTTP Server em Golang com Nginx Reverse Proxy

## Objetivo do Projeto
Este projeto tem como objetivo avaliar habilidades com Docker, programação, redes, servidores e automação em um ambiente Linux. O projeto envolve a criação de um serviço HTTP em Golang, a configuração de um container Docker, a criação de uma rede Docker e a configuração de um proxy reverso utilizando o Nginx.

## Instruções para Execução do Projeto

### 1. Instalar o Go (Golang) no Linux
Para instalar a versão mais recente do Go, siga os passos abaixo:

```bash
# Baixar a última versão do Go (verifique o site oficial para a versão mais recente)
wget https://golang.org/dl/go1.20.linux-amd64.tar.gz

# Remover qualquer versão antiga do Go
sudo rm -rf /usr/local/go

# Extrair o arquivo tar.gz para o diretório /usr/local
sudo tar -C /usr/local -xzf go1.20.linux-amd64.tar.gz

# Adicionar o Go ao seu PATH
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
source ~/.profile

# Verifique a instalação
go version
```

### 2. Instalar o Docker no Linux
Para instalar o Docker, use os seguintes comandos:

```bash
# Atualizar os pacotes
sudo apt update

# Instalar o Docker
sudo apt install -y docker.io

# Iniciar e habilitar o serviço Docker
sudo systemctl start docker
sudo systemctl enable docker

# Adicionar seu usuário ao grupo Docker (opcional)
sudo usermod -aG docker $USER

# Verifique a instalação
docker --version
```

### 3. Instalar o Docker Compose
Para instalar o Docker Compose, siga estes passos:

```bash
# Baixar a última versão do Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/latest/download/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

# Aplicar permissões executáveis
sudo chmod +x /usr/local/bin/docker-compose

# Verifique a instalação
docker-compose --version
```

### 4. Clonar o Repositório do Projeto
Clone o repositório do projeto do GitHub:

```bash
git clone <URL_DO_REPOSITORIO>
cd <NOME_DIRETORIO_PROJETO>
```

### 5. Executar o Projeto
Construa e inicie os containers usando o Docker Compose:

```bash
docker-compose up --build
```

### 6. Testar a Configuração
Após os containers estarem em execução, teste a configuração usando `curl`:

```bash
curl http://localhost:80
```

Você deve receber uma resposta em JSON do serviço Golang.

## Conclusão
Este projeto permite a prática de várias habilidades importantes na criação e gerenciamento de serviços em containers, além de fortalecer o entendimento sobre a comunicação entre serviços e o uso de proxies reversos.
