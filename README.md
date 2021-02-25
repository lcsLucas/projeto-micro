# Projeto-micro Golang
    Estrutura simples de microsserviços em Golang utilizando a ferramenta Go-kit

# Tecnologias utilizadas
    - Go
        - Go kit
        - Gorilla/mux
    - gRPC
        - protobuf
    - Docker
        - Dockerfile
    - Postgresql
    - Mongodb

# Pré-requisitos
    - Docker instalado, junto com o docker-compose

# Como rodar o projeto
    - Abrir uma janela do terminal na pasta do programa e executar os seguintes comandos (um por vez):
        1. docker build --no-cache -t dependencies -f ./Dockerfile.dependencies .
        2. docker-compose up --force-recreate --no-deps --build

# Como testar o funcionanmento
    - http://localhost:9000/alunos/status
    - http://localhost:9000/exercicios/status
    - http://localhost:9000/provas/status
    - http://localhost:9000/provas/5/35.329.394-5