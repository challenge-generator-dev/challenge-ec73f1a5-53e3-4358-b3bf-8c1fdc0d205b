# AGENTS.md

Instrucciones para el agente de IA que abra este repositorio (Claude Code, Cursor, Codex, Copilot, Gemini). Se cargan solas: no hay que pegar nada en ningun chat.

## Que es este repositorio

Es el codigo base de un reto de aprendizaje de Pragma: **Diseño y Desarrollo de Microservicios en Go**.

| | |
|---|---|
| Tema | Go Microservices |
| Nivel | junior-l1 |
| Chapter | Generico |
| Especialidad | Inferido del contexto |
| Stack | Go / gRPC |
| Patron arquitectonico | microservicio reactivo con capas estándar |
| Tiempo estimado | 2 semanas |

## Receta del stack

Esqueleto obligatorio:

- `el manifiesto de dependencias del stack, en la raiz`
- `el punto de entrada del stack`
- `la capa de interfaz (controller, handler o equivalente)`
- `las capas que el patron arquitectonico declarado exija, como directorios reales`

Dependencias:

- google.golang.org/grpc 1.64.0
- google.golang.org/protobuf 1.33.0
- github.com/golang/protobuf/protoc-gen-go 1.5.3
- github.com/stretchr/testify 1.9.0
- golang.org/x/net 0.23.0

## Tu tarea

Dejar este proyecto en estado **verificable**: que el comando de verificacion corra sin errores. Escribi los archivos en disco, en este repositorio. No generes ZIPs ni archivos adjuntos.

En orden:

1. Corre `go build ./...` y mira que falla.
2. Completa lo que falte de la lista de abajo: manifiesto de dependencias, punto de entrada, capa de interfaz y las capas del patron declarado.
3. Arregla SOLO los errores que impiden compilar o arrancar.
4. Volve a correr `go build ./...` hasta que pase.
5. Pará ahí.

## Regla dura: las fases son trabajo del humano

**PROHIBIDO implementar los entregables de las fases.** El valor del reto esta en que la persona los resuelva. Tu trabajo es que tenga un proyecto que arranca; el hueco pedagogico se queda como esta.

No resuelvas nada de esto:

- **Fase 1 — Definición del Dominio y Requerimientos**: Documento que describe el dominio, los actores, las operaciones y las propiedades operativas.
- **Fase 2 — Diseño de los Microservicios**: Documento que describe la arquitectura de los microservicios, las interfaces gRPC y las reglas de consistencia y disponibilidad.
- **Fase 3 — Implementación de los Microservicios**: Código fuente de los microservicios implementados y configurados para comunicarse a través de gRPC.
- **Fase 4 — Despliegue y Pruebas en Kubernetes**: Entorno de Kubernetes con los microservicios desplegados y pruebas de integración realizadas.

Distincion operativa:

- **Arreglar** (si): import faltante, tipo que no existe, dependencia sin declarar, error de sintaxis, archivo referenciado que no existe.
- **No tocar** (no): logica de negocio incompleta, validaciones ausentes, secretos hardcodeados, APIs deprecadas que funcionan, concurrencia insegura, patrones mejorables. Eso es lo que la persona tiene que encontrar.

## Superficie de practica (NO completes)

Estos archivos SON el ejercicio de la persona. No los implementes; deja stubs. No toques la logica que el reto pide completar.

- [ ] `Dockerfile` — El topic pide contenedores/orquestacion: este archivo es el ejercicio, no scaffolding.
- [ ] `docker-compose.yml` — El topic pide contenedores/orquestacion: este archivo es el ejercicio, no scaffolding.
- [ ] `deployments/k8s/invoice-service-deployment.yaml` — El topic pide contenedores/orquestacion: este archivo es el ejercicio, no scaffolding.
- [ ] `deployments/k8s/invoice-service-service.yaml` — El topic pide contenedores/orquestacion: este archivo es el ejercicio, no scaffolding.

## Lo que falta y tenes que completar

No se detectaron huecos: estan los archivos declarados, el boilerplate del stack y ninguna referencia quedo colgando. Igual corre el comando de verificacion — que los archivos existan no garantiza que compilen.

### Presentes (15)

- `go.mod`
- `cmd/invoice-service/main.go`
- `internal/domain/invoice.go`
- `internal/infrastructure/repository.go`
- `internal/interfaces/grpc/invoice_handler.go`
- `internal/application/invoice_service.go`
- `pkg/proto/invoice.proto`
- `pkg/proto/invoice.pb.go`
- `Dockerfile`
- `.dockerignore`
- `docker-compose.yml`
- `README.md`
- `deployments/k8s/invoice-service-deployment.yaml`
- `deployments/k8s/invoice-service-service.yaml`
- `test/invoice_service_test.go`

### Capas del patron declarado

Cada una tiene que existir como directorio real con al menos un archivo. Codigo plano en la raiz no satisface el patron.

- `cmd/invoice-service`
- `internal/domain`
- `internal/application`
- `internal/infrastructure`
- `internal/interfaces/grpc`
- `pkg/proto`
- `deployments/k8s`
- `test`

## Verificacion

```bash
go build ./...
```

El comando tiene que pasar SIN implementar los archivos de la superficie de practica: solo andamiaje.

Ese comando pasando es la definicion de "terminado" para vos.

## Convenciones que tenes que respetar

- Un solo ecosistema: no declares librerias de otro lenguaje ni mezcles gestores de paquetes.
- Toda libreria que uses tiene que estar declarada en el manifiesto de dependencias.
- Todo import declarado tiene que usarse; todo tipo usado tiene que existir o venir de una dependencia declarada.
- El patron es **microservicio reactivo con capas estándar**: los contratos (interfaces, puertos) los define la capa interna y los implementa la externa, nunca al revés.
- Los archivos que crees llevan implementacion real, no stubs: sin `TODO`, sin cuerpos vacios, sin `// getters y setters`.

## Contexto del candidato

Sirve para calibrar el nivel del codigo, no para resolver las fases.

- Brecha que el reto ataca: Diseñar microservicios en Go con gRPC, Docker y Kubernetes

---

*Generado por Challenge Generator — Pragma. `README.md` tiene el enunciado completo del reto para la persona. `PROMPT_MEJORA.md` es la variante para pegar en un chat, si se prefiere ese flujo.*
