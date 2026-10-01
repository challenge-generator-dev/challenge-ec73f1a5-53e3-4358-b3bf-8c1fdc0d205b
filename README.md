# Diseño y Desarrollo de Microservicios en Go

La empresa necesita implementar un sistema de microservicios en Go para gestionar la facturación de servicios de telecomunicaciones. El sistema debe manejar la creación, actualización y consulta de facturas, asegurando la consistencia de los datos y la disponibilidad del servicio. Los microservicios se comunicarán a través de gRPC y se desplegarán en un entorno de Docker y Kubernetes.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | Go Microservices |
| **Nivel** | junior-l1 |
| **Tipo** | practical |
| **Tiempo estimado** | 2 semanas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Go 1.21+, VS Code o GoLand.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Ejecuta `go build ./...`. Si no hay errores, estás listo.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Definición del Dominio y Requerimientos

**Objetivo:** Establecer los requerimientos y el dominio de los microservicios.

**Tiempo estimado:** 3 días

**Instrucciones:**

- Identificar los actores y sus interacciones en el dominio de la facturación de servicios de telecomunicaciones.
- Definir las operaciones clave que deben soportar los microservicios (crear, actualizar y consultar facturas).
- Establecer las propiedades operativas del sistema (consistencia, disponibilidad, latencia).

**Entregable:** Documento que describe el dominio, los actores, las operaciones y las propiedades operativas.

<details>
<summary>Pistas de conocimiento</summary>

- Considera las posibles fuentes de datos y sumideros en el dominio.
- Piensa en los posibles modos de falla y cómo el sistema debe manejarlos.

</details>

### Fase 2: Diseño de los Microservicios

**Objetivo:** Diseñar la arquitectura de los microservicios y sus interacciones.

**Tiempo estimado:** 5 días

**Instrucciones:**

- Diseñar los microservicios necesarios para manejar la facturación.
- Definir las interfaces gRPC para la comunicación entre microservicios.
- Establecer las reglas de consistencia y disponibilidad para las operaciones.
- Identificar los posibles puntos de falla y cómo el sistema debe manejarlos.

**Entregable:** Documento que describe la arquitectura de los microservicios, las interfaces gRPC y las reglas de consistencia y disponibilidad.

<details>
<summary>Pistas de conocimiento</summary>

- Considera la granularidad de los microservicios y cómo se relacionan entre sí.
- Piensa en las posibles estrategias para manejar la consistencia y la disponibilidad.

</details>

### Fase 3: Implementación de los Microservicios

**Objetivo:** Implementar los microservicios en Go y asegurar su correcta comunicación.

**Tiempo estimado:** 7 días

**Instrucciones:**

- Implementar los microservicios en Go siguiendo el diseño establecido.
- Configurar la comunicación gRPC entre los microservicios.
- Asegurar que los microservicios manejen correctamente las operaciones y los posibles puntos de falla.
- Verificar que los microservicios cumplan con las propiedades operativas establecidas.

**Entregable:** Código fuente de los microservicios implementados y configurados para comunicarse a través de gRPC.

<details>
<summary>Pistas de conocimiento</summary>

- Usa Docker para containerizar los microservicios.
- Considera el uso de Kubernetes para gestionar el despliegue y la escalabilidad de los microservicios.

</details>

### Fase 4: Despliegue y Pruebas en Kubernetes

**Objetivo:** Desplegar los microservicios en un entorno de Kubernetes y realizar pruebas de integración.

**Tiempo estimado:** 5 días

**Instrucciones:**

- Containerizar los microservicios usando Docker.
- Configurar y desplegar los microservicios en un entorno de Kubernetes.
- Realizar pruebas de integración para verificar la correcta comunicación y funcionamiento de los microservicios.
- Asegurar que el sistema cumpla con las propiedades operativas en el entorno de producción.

**Entregable:** Entorno de Kubernetes con los microservicios desplegados y pruebas de integración realizadas.

<details>
<summary>Pistas de conocimiento</summary>

- Usa Kubernetes para gestionar el despliegue y la escalabilidad de los microservicios.
- Realiza pruebas de carga y estrés para verificar la capacidad del sistema.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es un microservicio y cómo se relaciona con el dominio de la facturación de servicios de telecomunicaciones?
- **paraQueSirve**: ¿Para qué sirven los microservicios en este dominio y cómo mejoran el sistema?
- **comoSeUsa**: ¿Cómo se usan los microservicios para manejar la facturación de servicios de telecomunicaciones?
- **erroresComunes**: ¿Cuáles son los errores comunes que pueden ocurrir en la implementación de microservicios y cómo se pueden manejar?
- **queDecisionesImplica**: ¿Qué decisiones implica el diseño y despliegue de microservicios en este dominio?

## Criterios de Evaluacion

- Definición clara del dominio y los requerimientos de los microservicios.
- Diseño adecuado de la arquitectura de los microservicios y sus interacciones.
- Implementación correcta de los microservicios en Go y configuración de la comunicación gRPC.
- Despliegue exitoso de los microservicios en Kubernetes y realización de pruebas de integración.
- Manejo adecuado de los errores comunes y cumplimiento de las propiedades operativas.

## Como trabajar con un asistente de IA

Hay dos caminos, elegi uno:

- **AGENTS.md** (recomendado) — instrucciones nativas del repo. Abri esta carpeta con tu agente local (Claude Code, Cursor, Codex, Copilot, Gemini) y las carga solo. Sabe que archivos faltan y con que comando se verifica, y completa el scaffold escribiendo en disco.
- **PROMPT_MEJORA.md** — para copiar y pegar en un chat (claude.ai, ChatGPT). Devuelve un ZIP con el proyecto. Sirve si no tenes un agente en el IDE.

Ninguno de los dos resuelve las fases del reto: eso es tu trabajo.

## Verificacion

El proyecto esta listo para trabajar cuando este comando corre sin errores:

```bash
go build ./...
```

---

*Reto generado automaticamente por Challenge Generator - Pragma*
