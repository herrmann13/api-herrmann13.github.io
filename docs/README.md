# Backend Portfolio — Documentação de Arquitetura

## Objetivo

Backend de portfólio pessoal em **Go**, com foco em:

- **Observabilidade**: logs estruturados, métricas e tracing distribuído
- **Visualização de requisições**: rastreamento completo do ciclo de vida de cada request
- **Escalabilidade**: adicionar novos projetos/módulos sem reestruturar o sistema
- **Deploy simples e gratuito**: free tier do Render + GitHub Pages

---

## Arquitetura

### Visão Geral

```
┌─────────────────────┐     ┌──────────────────┐     ┌───────────────┐
│   GitHub Pages      │────▶│   Render (Go)    │────▶│  PostgreSQL   │
│   (Frontend)        │     │   (Backend API)  │     │  (Supabase/   │
│                     │     │                  │     │   Railway)    │
└─────────────────────┘     └──────────────────┘     └───────────────┘
```

### Frontend (GitHub Pages)

- Visualização de traces e spans por `trace_id`
- Filtro e exibição de logs por `trace_id`
- Dashboards de métricas
- Gatilhos para simulações e testes de carga
- Navegação entre projetos do portfólio

### Backend (Render — Go)

- Monolito simples, sem microsserviços
- API gateway única para todo o sistema
- Geração de `trace_id` por requisição
- Criação e gerenciamento de traces (spans)
- Emissão de logs estruturados
- Coleta e armazenamento de métricas
- Serve dados para os dashboards do frontend
- Rotas organizadas por projeto/módulo

### Banco de Dados (PostgreSQL)

Armazena tudo: logs, traces e métricas em tabelas relacionais.

---

## Observabilidade

### Logs

- Logs estruturados armazenados diretamente no PostgreSQL
- Sem collector externo, sem fila de mensagens

| Campo       | Descrição                     |
|-------------|-------------------------------|
| timestamp   | Momento do evento             |
| service     | Nome do serviço/módulo        |
| level       | INFO, WARN, ERROR, DEBUG      |
| message     | Conteúdo textual do log       |
| trace_id    | Identificador da trace        |

### Tracing

- Tracing distribuído **manual**, implementado em Go
- Cada request gera um `trace_id` que identifica todo o ciclo de vida
- Cada etapa do processamento gera um **span** associado ao trace
- Spans armazenados em tabela no PostgreSQL

**Conceito:** `trace` = request completo, `span` = passo dentro do trace.

### Métricas

- Contadores e timings simples, armazenados no PostgreSQL
- Atualizados a cada requisição pelo próprio backend

| Métrica               | Tipo      |
|-----------------------|-----------|
| request_count         | counter   |
| request_duration_ms   | timing    |
| error_count           | counter   |

---

## Modelo de Correlação

A chave central de correlação é o `trace_id`:

```
trace_id ──▶ logs (associados ao trace)
         ──▶ spans (pertencem ao trace)
         ──▶ métricas (opcionalmente etiquetadas com trace_id)
```

Isso permite reconstruir o fluxo completo de execução de qualquer requisição no frontend.

---

## Infraestrutura

| Componente       | Provedor              |
|------------------|-----------------------|
| Frontend         | GitHub Pages          |
| Backend          | Render (free tier)    |
| Banco de dados   | Supabase ou Railway   |

- Sem NGINX, sem gateway externo, sem fila de mensagens
- Deploy simples: monolito Go + banco externo

---

## Decisões de Design

| Decisão                                       | Motivo                                   |
|-----------------------------------------------|------------------------------------------|
| Monolito, não microsserviços                  | Simplicidade, custo zero                 |
| Sem fila de mensagens / event bus             | Complexidade desnecessária               |
| Sem NGINX ou gateway                          | Backend único serve tudo                 |
| Sem collector de logs separado                | Logs direto no PostgreSQL                |
| PostgreSQL como sistema de registro           | Única fonte de verdade                   |
| Tracing manual (sem OpenTelemetry)            | Fins educacionais, controle total        |
| Design modular com rotas por projeto          | Adicionar projetos sem reestruturar      |
| Zero dependências externas (stdlib apenas)    | Aprendizado, simplicidade                |

---

## Estrutura do Projeto

```
cmd/server/main.go          # ponto de entrada
internal/
  system/                    # health check, uptime
  projects/                  # listagem de projetos
  observability/             # tracing, logging, métricas (planejado)
  database/                  # conexão e migrações (planejado)
docs/                        # documentação
```

---

## Endpoints Atuais

| Método | Rota        | Descrição                     |
|--------|-------------|-------------------------------|
| GET    | `/health`   | Status do serviço + uptime    |
| GET    | `/projects` | Lista de projetos do portfólio|

---

## Próximos Passos

1. Conexão com PostgreSQL
2. Geração de `trace_id` por request (middleware)
3. Logs estruturados com `slog` (correlacionados ao trace)
4. Implementação manual de tracing (spans)
5. Coleta de métricas (request count, duration, errors)
6. Endpoints para consulta de logs, traces e métricas
7. Frontend com dashboards (GitHub Pages)
8. Deploy no Render

---

## Stack

- **Linguagem:** Go 1.26+
- **Banco:** PostgreSQL
- **Hosting Backend:** Render
- **Hosting Frontend:** GitHub Pages
