# Amazon Handler 

## Visão geral

Handler simples feito para facilitar o uso dos serviços da amazon em outros projetos. 

🚧 Projeto em andamento! 🚧

Atualmente o projeto suporta apenas o Amazon S3 e conta com funcionalidades como:
- Upload de objetos para buckets
- Criacação de buckets
- Exclusão de buckets e objetos
- Envio de objetos
- Listagem de objetos e buckets

## Tecnologias utilizadas:
- Go v.1.22.2
- Amazon SDK V2 para Go

## Como rodar:

### Standalone:
- ``` git clone https://github.com/ADStefano/AmazonHandler.git ```
- ``` cd AmazonHandler ```
- ``` go run main.go ```

### Como módulo:
- ``` go get github.com/ADStefano/AmazonHandler@latest ```
- ``` import "github.com/ADStefano/AmazonHandler" ```

## Exemplo de uso:
No arquivo main.go mostra como utilizar o projeto de forma simples, com algumas funções de create, delete e list sendo utilizadas.

## Teste:
O projeto conta com teste unitários e com um mock do serviço da Amazon S3, para rodar os teste utilize o comando: go test -v

## Roadmap:
- [ ] Melhorar a cobertura dos testes

- [ ] Implementar Pré signed URLs (upload/download)

## Autor:
Me chamo Ângelo P. Di Stefano, sou desenvolvedor de software backend com foco em Go, Python e Java.

Meu LinkedIn: https://www.linkedin.com/in/angelo-p-di-stefano/
