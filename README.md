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
    - Abrir uma janela do terminal na pasta do programa e executar o comando:
        1. docker-compose up

# Como testar o funcionanmento
    - http://localhost:9000/alunos/status
    - http://localhost:9000/exercicios/status
    - http://localhost:9000/provas/status
    - http://localhost:9000/provas/5/35.329.394-5