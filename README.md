## 🔗 GoURL - Encurtador de URL Simples
Este projeto é uma aplicação de backend RESTful construída em **Go (Golang)** que funciona como um encurtador de URLs simples. Ele demonstra a capacidade da linguagem Go para construir serviços web eficientes e lidar com requisições HTTP.

---

## ✨ Visão Geral
*O GoURL recebe URLs longas e as converte em URLs curtas e únicas. Ao acessar uma URL curta, o serviço redireciona para a URL longa original. Para simplificar, todas as URLs encurtadas são armazenadas apenas na memória enquanto o serviço está ativo. Isso significa que, ao reiniciar o aplicativo, os dados serão perdidos.*

---

## 🚀 Funcionalidades
**Encurtar URL:** Aceita uma URL longa via requisição **POST** e retorna uma URL curta gerada de forma única.

**Redirecionar URL:** Ao receber uma URL curta, o serviço redireciona automaticamente o usuário para a URL longa correspondente.

Armazenamento em Memória: Mapeamento de URLs curtas para longas é mantido em memória RAM.

---

## 🛠️ Tecnologias Utilizadas
- `Go (Golang): Linguagem de programação principal para o backend.`

- `net/http: Pacote padrão do Go para construir servidores web e lidar com requisições HTTP.`

- `sync: Pacote padrão do Go para controle de concorrência (usado para proteger o mapa de URLs com um sync.Mutex).`

- `encoding/json: Pacote padrão do Go para serialização e desserialização de dados JSON.`

- `github.com/teris-io/shortid: Biblioteca externa para gerar códigos curtos alfanuméricos únicos.`

---

## ⚙️ Como Rodar Localmente
**Siga os passos abaixo para configurar e executar a aplicação GoURL em sua máquina local.**

*Pré-requisitos:*
Go 1.22+: Baixe e instale a versão mais recente do Go em go.dev/dl/.

**Verifique a instalação com: go version**

- Passos:
*Clone o Repositório:*

- `git clone https://github.com/Wozniak7/Go_Project.git`
- `cd Go_Project`

**(Ajuste o cd se você moveu o projeto Go para um subdiretório dentro de Go_Project)**

*Baixe as Dependências:*

- `go mod tidy`

*Este comando lerá seu go.mod e main.go, baixará as dependências necessárias e atualizará go.sum.*

**Execute o Servidor Go:**

- `go run main.go`

*Você deverá ver uma mensagem no terminal indicando que o servidor iniciou na porta :8080:* `Servidor GoURL iniciado na porta :8080`

**Testando a API Localmente:**
Com o servidor GoURL rodando, você pode testá-lo usando curl (ou ferramentas como Postman/Insomnia):

*Encurtar uma URL (POST):*

- `curl -X POST -H "Content-Type: application/json" -d '{"long_url": "https://www.google.com/sua/pagina/muito/longa/e/complexa"}' http://localhost:8080/shorten`

*Você receberá uma resposta JSON com a URL curta, por exemplo:*

- `{"short_url":"http://localhost:8080/abcde"}`

**Redirecionar uma URL (GET):**
Abra seu navegador e cole a short_url que você recebeu (ex: http://localhost:8080/abcde). Você deverá ser redirecionado para a URL longa original.

---

## ☁️ Deploy no Render.com
Este serviço GoURL pode ser facilmente implantado em plataformas como o **Render.com.**

*Configurações no Render:*
Tipo de Serviço: **Web Service**

Runtime: Go (detectado automaticamente)

Build Command: go build -o server .

Start Command: ./server

Port: Certifique-se de que seu main.go utiliza a variável de ambiente PORT fornecida pelo Render:

// Trecho do main.go
port := os.Getenv("PORT")
if port == "" {
    port = "8080" // Fallback para desenvolvimento local
}
log.Fatal(http.ListenAndServe(":"+port, nil))

*(Lembre-se de adicionar import "os" no seu main.go).*

**CORS:** A aplicação Go já inclui cabeçalhos CORS (Access-Control-Allow-Origin: *) para permitir o acesso de frontends em diferentes origens, o que é útil para desenvolvimento. Para produção, considere restringir o Origin aos domínios específicos do seu frontend.

---

## 🌐 Frontend Web (Opcional)
Para uma demonstração completa do seu encurtador de URL, você pode criar uma interface web simples em HTML/CSS/JavaScript. O arquivo index.html na raiz deste projeto é um exemplo de frontend que se comunica com esta API GoURL.

Se a API GoURL estiver deployada, edite o index.html e substitua http://localhost:8080 pela URL pública do seu serviço Render para conectar o frontend ao backend deployado.

## 🤝 Contribuições
Sinta-se à vontade para explorar o código, sugerir melhorias ou reportar problemas.