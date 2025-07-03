# AmazonHandler - Biblioteca Go para manipulação de serviços AWS

## Visão geral

Handler simples feito para facilitar o uso dos serviços da amazon em outros projetos. 

🚧 Projeto em andamento! 🚧

Atualmente o projeto suporta apenas o Amazon S3 e conta com funcionalidades como:
- Criação de buckets
- Upload de objetos para buckets
- Exclusão de buckets e objetos
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

## Testes:
O projeto possui testes unitários, incluindo mocks do serviço da Amazon S3.
Para rodar os testes:
- ``` go test -v ```

## Roadmap:
- [x] Implementar interface do serviço da AWS S3

- [x] Implementar mocks do serviço da AWS S3

- [x] Implementar criação e exclusão de buckets

- [x] Implementar exclusão de objetos e buckets

- [x] Implementar listagem de objetos e buckets

- [ ] Implementar upload e download

- [x] Ajustar injeção de dependência no client

- [x]  Implementar errors.go com erros padronizados

- [ ]  Expandir errors.go para utilizar structs

- [ ]  Implementar arquivos de interface com verificação de implementação

- [ ] Implementar Pré signed URLs (upload/download)

- [ ] Implementar busca por prefixo ao listar objetos 

- [ ] Aumentar a cobertura dos testes unitários

- [ ] Ajustar go mod para importar como package em outros projetos

## Autor:
Me chamo Ângelo Pedersen Di Stefano, sou desenvolvedor de software com foco em backend nas linguagens Go, Python e Java.

Meu LinkedIn: https://www.linkedin.com/in/angelo-p-di-stefano/
