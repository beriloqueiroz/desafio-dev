# Desafio backend

## Desenho da solução

![img.png](img.png)

## todo
- fazer retry no send queue
- fazer next app

## Introdução
A solução tem como objetivo enviar mensagens com informações de previsão do tempo 
para uma aplicação web sendo possível o agendamento desse envio, e o conteúdo da mensagem depende 
da localização do usuário. O usuário poderá também ser cadastrado, ativado e desativado.

A autenticação e autorização não está contemplada no projeto. Considerou-se que é feita por outro serviço, utilizando talvez a camada de infra e etc.  

## app_containers
É o conjunto de aplicativos internos da solução. Os blocos contidos são containers. Esses blocos chamaremos neste documento de módulos

## cache_sync_service
Atualiza cache de tempos em tempos com informações das localidades existentes no mgmt.
### casos de uso
- sync
  - job que roda de tempos em tempos. Este consulta as localidades diferentes no mgmt, e popula no cache. Considera-se rodar este job todos os dias.
- get_msg
  - possibilita captura da mensagem com base num valor de 'nome cidade - UF', Usado pelo core, este uso é feito utilizando injeção do usecase, poderia ter sido feito via http, mas não foi a decisão.

## core
Responsável pela gestão de usuários e agendamentos. De forma que de tempos em tempos é verificado no mgmt se há schedule a ser acionado, inserindo assim as notificações com as mensagens populadas, dos usuários ativos na fila.
### casos de uso
- insert user
- activate user
- deactivate user
- create schedule notification
- delete schedule notification
- sync schedule
  - verifica agendamentos que devem ser acionados e envia para as filas cadastradas, hoje somente com a web_queue, as notificações para os usuários ativos.

## web_worker
Ler da ewb_queue constantemente e envia as notificações para o web conforme localização e email.
### casos de uso
- sync_send_notifications

## app_worker (pode ser implementado futuramente)
Ler da web_queue constantemente e envia as notificações para o app conforme localização e email.
### casos de uso (padrão para workers)
- sync_send_notifications

## estrutua e organização

Decidiu-se utilizar alguns conceitos de clean architecture e domínios ricos. 
Os módulos estão contidos na pasta internal, e seus entrypoints na pasta cmd, uma outra possibilidade seria tornar esses módulos em projetos isolados.
- Dentro de cada módulo há:
  - entity
    - contém as regras de négocio, representa o domínio.
  - infra
    - contém as implementações e integrações com meio externo, banco de dados, api e etc.
  - usecase
    - representa intensão do usuário, o usuário pode ser outro módulo ou sistema.
### pastas
```tree
├── cmd
│   ├── cache_sync_service
│   │   └── factories
│   ├── core
│   ├── web_client_test
│   └── web_worker
├── configs
├── internal
│   ├── cache_sync_service
│   │   ├── entity
│   │   ├── infra
│   │   │   └── implements
│   │   └── usecase
│   │       └── interfaces
│   ├── core
│   │   ├── entity
│   │   ├── infra
│   │   │   ├── database
│   │   │   │   └── migrations
│   │   │   ├── implements
│   │   │   └── web
│   │   └── usecase
│   │       └── interfaces
│   └── web_worker
│       ├── entity
│       ├── infra
│       │   └── implements
│       └── usecase
│           └── interfaces
└── pkg

```

### entrypoints
```tree
./cmd
├── cache_sync_service
│   ├── factories
│   │   └── factories.go
│   └── main.go
├── core
│   └── main.go
└── web_worker
    └── main.go
```

### exemplo de módulo
```tree
./internal/web_worker
├── entity
│   └── notification.go
├── infra
│   └── implements
│       ├── coastal_cities.go
│       ├── web_kafka_repository.implements.go
│       └── web_rest_service.implements.go
└── usecase
    ├── interfaces
    │   ├── notification_queue_repository.interface.go
    │   └── web_service.interface.go
    └── sync_notifications.usecase.go

```
O conteúdo da pasta implements são as implementações das interfaces em /usecase/interfaces.

## Resiliência e Escalabilidade

- CircuitBreak
  - Implementados na neste projeto nas implementações nos módulos na pasta infra/implements nos casos que há request para serviço de terceiros.
- Containers
  - Containerização que pode ser escalada com replicaset e deployments usando kubernetes


## Tecnologias e serviço de terceiros

- cache: Redis
- message_api: CPTEC API REST
- web_queue: Kafka
- mgmt: postgres
- notification_web_worker: API REST [WEBSERVER - POST]
- web: web app test 
  - [//]: # (  - todo)

## Testando

- rodando
```bash
docker compose up -d
```
em caso de problemas de permissão ao criar o volumes usar:
```bash
sudo chmod -R g+rwX ./.docker/kafka
```
- gestão de usuário e agendamento

Usar uma extensão do vscode Name: Thunder Client: https://marketplace.visualstudio.com/items?itemName=rangav.vscode-thunder-client, que é algo parecido com postman.
Pode importar o arquivo: thunder-collection_desafio-dev-back.json

- Swagger
  - http://localhost:8080/swagger/# desafio-dev
