# Prompt para Mejorar el Codigo Base

Copia y pega el contenido del bloque de abajo en un asistente de IA (Claude, ChatGPT)
para obtener un ZIP con el proyecto completo y arrancable.

Si preferis trabajar en tu editor con un agente local (Claude Code, Cursor, Copilot), usa `AGENTS.md` en vez de este archivo: dice lo mismo pero para que escriba los archivos en disco.

## Las dos reglas que no se negocian

1. **Completa el boilerplate.** Todo lo que el proyecto necesita para compilar y arrancar: manifiesto de dependencias, punto de entrada, configuracion, capa de interfaz, y las capas del patron arquitectonico declarado. Eso es andamiaje y es tu trabajo.
2. **NO resuelvas el reto.** Los entregables de las fases son el trabajo de la persona. El hueco pedagogico se deja como esta: el proyecto arranca, pero lo que el reto pide implementar NO esta implementado.

Dicho de otra forma: si algo impide compilar, arreglalo. Si algo es logica de negocio incompleta, validaciones ausentes, un secreto hardcodeado o un patron mejorable, dejalo exactamente como esta — es lo que la persona tiene que encontrar.

## Superficie de practica — NO resuelvas

Estos archivos SON el ejercicio de la persona. No los implementes; deja stubs.

- `Dockerfile` — El topic pide contenedores/orquestacion: este archivo es el ejercicio, no scaffolding.
- `docker-compose.yml` — El topic pide contenedores/orquestacion: este archivo es el ejercicio, no scaffolding.
- `deployments/k8s/invoice-service-deployment.yaml` — El topic pide contenedores/orquestacion: este archivo es el ejercicio, no scaffolding.
- `deployments/k8s/invoice-service-service.yaml` — El topic pide contenedores/orquestacion: este archivo es el ejercicio, no scaffolding.

## Como saber que terminaste

```bash
go build ./...
```

Ese comando corriendo sin errores es la definicion de "listo".

---

```
## Briefing del reto (autoridad)
Este bloque manda sobre los archivos adjuntos. El stack y el rol salen de AQUÍ, no de un topic genérico ni de markdown placeholder.

### Contexto técnico original
Diseñar microservicios en Go con gRPC, Docker y Kubernetes

### Reto
- Tema: Go Microservices
- Seniority: junior-l1
- Tipo: practical
- Título: Diseño y Desarrollo de Microservicios en Go
- Tiempo estimado: 2 semanas

### Fases (trabajo del HUMANO — PROHIBIDO completarlas)
No implementes estos entregables. Dejalos como hueco pedagógico. El asistente solo materializa el proyecto arrancable para que el participante pueda trabajar.
- Fase 1: Definición del Dominio y Requerimientos — objetivo: Establecer los requerimientos y el dominio de los microservicios. — entregable (NO resolver): Documento que describe el dominio, los actores, las operaciones y las propiedades operativas.
- Fase 2: Diseño de los Microservicios — objetivo: Diseñar la arquitectura de los microservicios y sus interacciones. — entregable (NO resolver): Documento que describe la arquitectura de los microservicios, las interfaces gRPC y las reglas de consistencia y disponibilidad.
- Fase 3: Implementación de los Microservicios — objetivo: Implementar los microservicios en Go y asegurar su correcta comunicación. — entregable (NO resolver): Código fuente de los microservicios implementados y configurados para comunicarse a través de gRPC.
- Fase 4: Despliegue y Pruebas en Kubernetes — objetivo: Desplegar los microservicios en un entorno de Kubernetes y realizar pruebas de integración. — entregable (NO resolver): Entorno de Kubernetes con los microservicios desplegados y pruebas de integración realizadas.

Eres un asistente experto en análisis, corrección y generación de archivos de cualquier tipo:
código fuente, documentación, hojas de cálculo, documentos Word, configuraciones, entre otros.
Voy a enviarte una cadena de texto que contiene uno o más archivos. Cada archivo está delimitado por un marcador con el siguiente formato:
// === ARCHIVO: ruta/del/archivo.extension ===
o también puede aparecer como:
## === ARCHIVO: ruta/del/archivo.extension ===
Lo que sigue al marcador puede ser:

El contenido real del archivo (código, texto, YAML, etc.)
Una descripción en lenguaje natural de lo que debe contener el archivo


TU TAREA
PASO 0 — ¿Esto es un proyecto o una carcasa?
Antes de extraer archivos, leé el Briefing (si está) y diagnosticá el adjunto.

Es CARCASA si ocurre CUALQUIERA de estas:
- No hay manifiesto de dependencias del stack del briefing (manifest.json de VTEX IO / package.json / pom.xml / build.gradle / requirements.txt / go.mod / *.tf / *.csproj, según corresponda)
- Hay un "binario" que en realidad es un comentario ("no puede ser mostrado como texto plano", placeholder .fig/.docx vacío)
- Los markdowns ya completan entregables de fases posteriores ("se implementó fade-in", lista de áreas ya resuelta)

Si es CARCASA:
- MATERIALIZÁ un proyecto que arranca en el stack del briefing (VTEX IO Store Framework, Angular, Terraform, pytest, Nest, etc.). Incluí manifiesto, punto de entrada y capa de interfaz reales.
- NO copies los markdowns de "solución" como si fueran el producto. Son ruido de generación.
- NO resuelvas las fases del briefing (están marcadas PROHIBIDO). Dejá el hueco pedagógico: el flujo existe, las microinteracciones/calidad/infra que el reto pide NO están hechas.
- Después seguí al PASO 5 (ZIP).

Si es un proyecto REAL (manifiesto + código que compila o arranca):
- Seguí PASO 1 en adelante. 🔴 compilación sí. 🟡 pedagógico no.

PASO 1 — Detección y extracción
Identifica todos los archivos presentes en la cadena. Para cada archivo extrae:

Su ruta completa (ej: src/main/java/com/pragma/Service.java)
Su contenido o descripción

PASO 2 — Clasificación por tipo
Clasifica cada archivo en una de estas categorías:
A) Código fuente (Java, Python, TypeScript, JavaScript, Kotlin, etc.)
B) Configuración / documentación (YAML, properties, Markdown, JSON, txt, etc.)
C) Excel (.xlsx, .xls, .csv)
D) Word (.docx, .doc)
E) Otro tipo de archivo binario o especial
PASO 3 — Clasificación de errores en código fuente

Objetivo prioritario: que el proyecto compile. No corrijas flujo de negocio ni lógica funcional.

Antes de modificar cualquier archivo de código fuente, clasifica cada problema encontrado en una de estas dos categorías:
🔴 ERROR DE COMPILACIÓN — corregir siempre
Son errores que impiden que el proyecto arranque, sin valor pedagógico:

Import faltante o incorrecto
Clase, método o variable referenciada que no existe en ningún archivo del proyecto
Error de sintaxis
Anotación con atributos inválidos
Dependencia ausente en pom.xml, package.json, etc.
Archivo referenciado que no existe y debe ser creado con implementación mínima

→ CORREGIR estos errores.
🟡 PROBLEMA FUNCIONAL O DE CALIDAD — preservar siempre
Son problemas que no impiden compilar. Pueden ser intencionales para el aprendizaje:

Clave secreta hardcodeada ("secret", "password123")
API deprecada que funciona pero tiene reemplazo moderno
Lógica de negocio incorrecta o incompleta
Código redundante o de baja legibilidad
Falta de validaciones en flujo de negocio
Patrones de diseño incorrectos pero funcionales
Concurrencia no segura
Configuración funcional pero no óptima

→ PRESERVAR tal cual. No corregir, no mejorar, no comentar.
PASO 4 — Procesamiento según tipo de archivo
Tipo A — Código fuente
Aplica únicamente las correcciones clasificadas como 🔴 ERROR DE COMPILACIÓN.
No alteres ningún elemento clasificado como 🟡 PROBLEMA FUNCIONAL O DE CALIDAD.
Si falta un archivo referenciado, créalo con la implementación mínima necesaria para compilar.
Tipo B — Configuración / documentación
Extrae el contenido tal cual, sin modificaciones salvo errores evidentes de sintaxis
(ej: YAML mal indentado).
Tipo C — Excel (.xlsx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un archivo Excel funcional con:

Fila de encabezados en negrita con color de fondo distintivo
Columnas con ancho ajustado al contenido
Tipos de dato correctos por columna
Validaciones si la descripción lo indica
Hojas nombradas descriptivamente si hay más de una
Filas de ejemplo si no hay datos reales

Tipo D — Word (.docx)
Si viene con contenido real, genera el archivo respetando ese contenido.
Si viene con descripción en lenguaje natural, genera un documento Word funcional con:

Estilos de título (Título 1, Título 2) para jerarquía de secciones
Fuente legible (Calibri o equivalente), tamaño 11-12pt para cuerpo
Márgenes estándar
Tabla de contenido si tiene múltiples secciones
Tablas con encabezados en negrita si aplica

Tipo E — Otro
Genera el archivo con el contenido o estructura más apropiada según la descripción.
PASO 5 — Exportación en ZIP
Empaqueta todos los archivos en un único archivo ZIP descargable respetando exactamente
la estructura de rutas indicada por los marcadores.
El ZIP debe incluir:

Archivos de código con únicamente los errores de compilación corregidos
Archivos de configuración y documentación sin cambios
Archivos nuevos creados para resolver dependencias de compilación faltantes
Archivos Excel y Word generados desde descripción

IMPORTANTE: El ZIP debe estar listo para descargar al finalizar. No preguntes si el usuario
quiere generarlo. Simplemente genera el archivo y proporciona el enlace de descarga; No debes desplegar en el chat el resumen de lo que arreglaste al Zip, solo entregalo.

REGLAS IMPORTANTES

No omitas ningún archivo aunque no tenga errores ni modificaciones
Respeta los nombres y rutas exactas indicadas por los marcadores
Si un archivo no tiene marcador claro, infiere el nombre desde su contenido
Si la cadena contiene solo documentación, placeholders o binarios fake, NO la reproduzcas:
aplicá PASO 0 (materializar el proyecto del briefing). Reproducir la carcasa es un fallo.
No agregues texto después del enlace de descarga del ZIP
No preguntes si el usuario quiere el ZIP: simplemente generalo siempre
Si detectas que falta un archivo de configuración necesario para compilar
(pom.xml, package.json, requirements.txt, build.gradle, etc.), créalo e inclúyelo
inferiendo su contenido desde los imports y frameworks detectados en el código
Nunca corrijas problemas 🟡 aunque parezcan obvios o fáciles de mejorar.
El participante que recibirá este proyecto los debe encontrar y resolver él mismo.


INPUT
Aquí está la cadena con los archivos:

// === ARCHIVO: go.mod ===
module invoice-service

go 1.22

require (
	github.com/golang/protobuf/protoc-gen-go v1.5.3
	github.com/stretchr/testify v1.9.0
	golang.org/x/net v0.23.0
	google.golang.org/grpc v1.64.0
	google.golang.org/protobuf v1.33.0
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
)

// === ARCHIVO: cmd/invoice-service/main.go ===
package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	"invoice-service/internal/application"
	"invoice-service/pkg/proto"
)

const (
	grpcPort = ":50051"
)

func main() {
	// Configurar el contexto para manejo de señales
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Canal para capturar señales del sistema
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Inicializar el servidor gRPC
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Crear el servidor gRPC con opciones de timeout
	grpcServer := grpc.NewServer(
		grpc.ConnectionTimeout(5 * time.Second),
		grpc.MaxRecvMsgSize(1024*1024*5), // 5MB
		grpc.MaxSendMsgSize(1024*1024*5), // 5MB
	)

	// Inicializar el servicio de aplicación
	invoiceService := application.NewInvoiceService()

	// Registrar el servicio gRPC
	invoiceServer := &invoiceGRPCServer{invoiceService: invoiceService}
	proto.RegisterInvoiceServiceServer(grpcServer, invoiceServer)

	// Registrar el servicio de health check
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("invoice.InvoiceService", grpc_health_v1.HealthCheckResponse_SERVING)

	// Registrar reflection para debugging
	reflection.Register(grpcServer)

	// Canal para esperar la finalización del servidor
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		log.Printf("Starting gRPC server on %s", grpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Esperar señal de terminación
	select {
	case sig := <-sigChan:
		log.Printf("Received signal: %v", sig)
		grpcServer.GracefulStop()
	case <-ctx.Done():
		log.Println("Context cancelled, shutting down")
		grpcServer.GracefulStop()
	}

	wg.Wait()
}

// invoiceGRPCServer implementa el servidor gRPC para el servicio de facturas.
type invoiceGRPCServer struct {
	proto.UnimplementedInvoiceServiceServer
	invoiceService *application.InvoiceService
}

// CreateInvoice maneja la creación de una nueva factura.
func (s *invoiceGRPCServer) CreateInvoice(ctx context.Context, req *proto.CreateInvoiceRequest) (*proto.InvoiceResponse, error) {
	if req == nil || req.Invoice == nil {
		return nil, fmt.Errorf("invalid request: invoice data is required")
	}

	invoice := &application.Invoice{
		ID:          req.Invoice.Id,
		CustomerID: req.Invoice.CustomerId,
		Amount:     req.Invoice.Amount,
		IssueDate:  req.Invoice.IssueDate.AsTime(),
		DueDate:    req.Invoice.DueDate.AsTime(),
		Status:     req.Invoice.Status,
	}

	createdInvoice, err := s.invoiceService.CreateInvoice(ctx, invoice)
	if err != nil {
		return nil, fmt.Errorf("failed to create invoice: %w", err)
	}

	return &proto.InvoiceResponse{
		Invoice: &proto.Invoice{
			Id:         createdInvoice.ID,
			CustomerId: createdInvoice.CustomerID,
			Amount:     createdInvoice.Amount,
			IssueDate:  protoTimestamp(createdInvoice.IssueDate),
			DueDate:    protoTimestamp(createdInvoice.DueDate),
			Status:     createdInvoice.Status,
		},
	}, nil
}

// GetInvoice recupera una factura por su ID.
func (s *invoiceGRPCServer) GetInvoice(ctx context.Context, req *proto.GetInvoiceRequest) (*proto.InvoiceResponse, error) {
	if req == nil || req.Id == "" {
		return nil, fmt.Errorf("invalid request: invoice ID is required")
	}

	invoice, err := s.invoiceService.GetInvoice(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get invoice: %w", err)
	}

	return &proto.InvoiceResponse{
		Invoice: &proto.Invoice{
			Id:         invoice.ID,
			CustomerId: invoice.CustomerID,
			Amount:     invoice.Amount,
			IssueDate:  protoTimestamp(invoice.IssueDate),
			DueDate:    protoTimestamp(invoice.DueDate),
			Status:     invoice.Status,
		},
	}, nil
}

// UpdateInvoice actualiza una factura existente.
func (s *invoiceGRPCServer) UpdateInvoice(ctx context.Context, req *proto.UpdateInvoiceRequest) (*proto.InvoiceResponse, error) {
	if req == nil || req.Invoice == nil {
		return nil, fmt.Errorf("invalid request: invoice data is required")
	}

	invoice := &application.Invoice{
		ID:          req.Invoice.Id,
		CustomerID: req.Invoice.CustomerId,
		Amount:     req.Invoice.Amount,
		IssueDate:  req.Invoice.IssueDate.AsTime(),
		DueDate:    req.Invoice.DueDate.AsTime(),
		Status:     req.Invoice.Status,
	}

	updatedInvoice, err := s.invoiceService.UpdateInvoice(ctx, invoice)
	if err != nil {
		return nil, fmt.Errorf("failed to update invoice: %w", err)
	}

	return &proto.InvoiceResponse{
		Invoice: &proto.Invoice{
			Id:         updatedInvoice.ID,
			CustomerId: updatedInvoice.CustomerID,
			Amount:     updatedInvoice.Amount,
			IssueDate:  protoTimestamp(updatedInvoice.IssueDate),
			DueDate:    protoTimestamp(updatedInvoice.DueDate),
			Status:     updatedInvoice.Status,
		},
	}, nil
}

// protoTimestamp convierte un time.Time a un *timestamppb.Timestamp.
func protoTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

// === ARCHIVO: internal/domain/invoice.go ===
package domain

import (
	"errors"
	"time"
)

// Invoice representa una factura en el dominio de telecomunicaciones.
type Invoice struct {
	ID          string
	CustomerID  string
	Amount      float64
	DueDate     time.Time
	IssueDate   time.Time
	Status      InvoiceStatus
	Description string
}

// InvoiceStatus define los estados posibles de una factura.
type InvoiceStatus string

const (
	StatusPending   InvoiceStatus = "pending"
	StatusPaid      InvoiceStatus = "paid"
	StatusOverdue   InvoiceStatus = "overdue"
	StatusCancelled InvoiceStatus = "cancelled"
)

// InvoiceRepository define el contrato para la persistencia de facturas.
type InvoiceRepository interface {
	Save(invoice *Invoice) error
	FindByID(id string) (*Invoice, error)
	Update(invoice *Invoice) error
}

// NewInvoice crea una nueva instancia de Invoice con validación básica.
func NewInvoice(id, customerID string, amount float64, dueDate, issueDate time.Time, description string) (*Invoice, error) {
	if id == "" {
		return nil, errors.New("invoice ID cannot be empty")
	}
	if customerID == "" {
		return nil, errors.New("customer ID cannot be empty")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be positive")
	}
	if dueDate.Before(issueDate) {
		return nil, errors.New("due date cannot be before issue date")
	}

	return &Invoice{
		ID:          id,
		CustomerID:  customerID,
		Amount:      amount,
		DueDate:     dueDate,
		IssueDate:   issueDate,
		Status:      StatusPending,
		Description: description,
	}, nil
}

// Validate verifica que la factura cumpla con las reglas de negocio.
func (i *Invoice) Validate() error {
	if i.Amount <= 0 {
		return errors.New("amount must be positive")
	}
	if i.DueDate.Before(i.IssueDate) {
		return errors.New("due date cannot be before issue date")
	}
	if i.Status == "" {
		return errors.New("status cannot be empty")
	}
	return nil
}

// UpdateStatus actualiza el estado de la factura segun reglas de negocio.
func (i *Invoice) UpdateStatus(newStatus InvoiceStatus) error {
	switch i.Status {
	case StatusPaid:
		return errors.New("cannot change status of a paid invoice")
	case StatusCancelled:
		return errors.New("cannot change status of a cancelled invoice")
	}

	switch newStatus {
	case StatusPaid, StatusOverdue, StatusCancelled:
		i.Status = newStatus
		return nil
	default:
		return errors.New("invalid invoice status")
	}
}

// === ARCHIVO: internal/infrastructure/repository.go ===
package infrastructure

import (
	"errors"
	"invoice-service/internal/domain"
	"sync"
)

// InvoiceRepositoryInMemory es una implementación en memoria de InvoiceRepository.
type InvoiceRepositoryInMemory struct {
	mu       sync.RWMutex
	invoices map[string]*domain.Invoice
}

// NewInvoiceRepositoryInMemory crea una nueva instancia de InvoiceRepositoryInMemory.
func NewInvoiceRepositoryInMemory() *InvoiceRepositoryInMemory {
	return &InvoiceRepositoryInMemory{
		invoices: make(map[string]*domain.Invoice),
	}
}

// Save guarda una factura en el repositorio.
func (r *InvoiceRepositoryInMemory) Save(invoice *domain.Invoice) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.invoices[invoice.ID]; exists {
		return errors.New("invoice already exists")
	}

	r.invoices[invoice.ID] = invoice
	return nil
}

// FindByID busca una factura por su ID.
func (r *InvoiceRepositoryInMemory) FindByID(id string) (*domain.Invoice, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	invoice, exists := r.invoices[id]
	if !exists {
		return nil, errors.New("invoice not found")
	}
	return invoice, nil
}

// Update actualiza una factura existente en el repositorio.
func (r *InvoiceRepositoryInMemory) Update(invoice *domain.Invoice) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.invoices[invoice.ID]; !exists {
		return errors.New("invoice not found")
	}

	r.invoices[invoice.ID] = invoice
	return nil
}

// === ARCHIVO: internal/interfaces/grpc/invoice_handler.go ===
package grpc

import (
	"context"
	"errors"
	"invoice-service/internal/domain"
	"invoice-service/pkg/proto"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// InvoiceHandler maneja las solicitudes gRPC para operaciones con facturas.
type InvoiceHandler struct {
	proto.UnimplementedInvoiceServiceServer
	service domain.InvoiceService
}

// NewInvoiceHandler crea una nueva instancia de InvoiceHandler.
func NewInvoiceHandler(service domain.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{
		service: service,
	}
}

// CreateInvoice crea una nueva factura.
func (h *InvoiceHandler) CreateInvoice(ctx context.Context, req *proto.CreateInvoiceRequest) (*proto.InvoiceResponse, error) {
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request cannot be nil")
	}

	invoice, err := domain.NewInvoice(
		req.GetId(),
		req.GetCustomerId(),
		req.GetAmount(),
		req.GetDueDate().AsTime(),
		req.GetIssueDate().AsTime(),
		req.GetDescription(),
	)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid invoice data: %v", err)
	}

	if err := h.service.CreateInvoice(ctx, invoice); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create invoice: %v", err)
	}

	return h.toInvoiceResponse(invoice), nil
}

// GetInvoice recupera una factura por su ID.
func (h *InvoiceHandler) GetInvoice(ctx context.Context, req *proto.GetInvoiceRequest) (*proto.InvoiceResponse, error) {
	if req == nil || req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invoice ID cannot be empty")
	}

	invoice, err := h.service.GetInvoice(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, errors.New("invoice not found")) {
			return nil, status.Errorf(codes.NotFound, "invoice not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get invoice: %v", err)
	}

	return h.toInvoiceResponse(invoice), nil
}

// UpdateInvoice actualiza una factura existente.
func (h *InvoiceHandler) UpdateInvoice(ctx context.Context, req *proto.UpdateInvoiceRequest) (*proto.InvoiceResponse, error) {
	if req == nil || req.GetId() == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invoice ID cannot be empty")
	}

	invoice, err := h.service.GetInvoice(ctx, req.GetId())
	if err != nil {
		if errors.Is(err, errors.New("invoice not found")) {
			return nil, status.Errorf(codes.NotFound, "invoice not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get invoice: %v", err)
	}

	if req.Amount != nil {
		invoice.Amount = req.GetAmount()
	}
	if req.DueDate != nil {
		invoice.DueDate = req.GetDueDate().AsTime()
	}
	if req.Description != nil {
		invoice.Description = req.GetDescription()
	}

	if err := invoice.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid invoice data: %v", err)
	}

	if err := h.service.UpdateInvoice(ctx, invoice); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update invoice: %v", err)
	}

	return h.toInvoiceResponse(invoice), nil
}

// toInvoiceResponse convierte una entidad Invoice a InvoiceResponse.
func (h *InvoiceHandler) toInvoiceResponse(invoice *domain.Invoice) *proto.InvoiceResponse {
	return &proto.InvoiceResponse{
		Id:          invoice.ID,
		CustomerId:  invoice.CustomerID,
		Amount:      invoice.Amount,
		DueDate:     timestamppb.New(invoice.DueDate),
		IssueDate:   timestamppb.New(invoice.IssueDate),
		Status:      string(invoice.Status),
		Description: invoice.Description,
	}
}


// === ARCHIVO: internal/application/invoice_service.go ===
package application

import (
	"context"
	"fmt"
	"time"

	"invoice-service/internal/domain"
	"invoice-service/pkg/proto"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type InvoiceService interface {
	CreateInvoice(ctx context.Context, customerID string, amount float64, dueDate, issueDate time.Time, description string) (*domain.Invoice, error)
	GetInvoice(ctx context.Context, id string) (*domain.Invoice, error)
	UpdateInvoice(ctx context.Context, id string, amount float64, status domain.InvoiceStatus) (*domain.Invoice, error)
}

type invoiceService struct {
	repository domain.InvoiceRepository
}

func NewInvoiceService(repository domain.InvoiceRepository) InvoiceService {
	return &invoiceService{
		repository: repository,
	}
}

func (s *invoiceService) CreateInvoice(ctx context.Context, customerID string, amount float64, dueDate, issueDate time.Time, description string) (*domain.Invoice, error) {
	if customerID == "" {
		return nil, fmt.Errorf("el ID del cliente es requerido")
	}

	if amount <= 0 {
		return nil, fmt.Errorf("el monto debe ser mayor a cero")
	}

	if dueDate.Before(issueDate) {
		return nil, fmt.Errorf("la fecha de vencimiento no puede ser anterior a la fecha de emisión")
	}

	invoice, err := domain.NewInvoice(
		generateInvoiceID(),
		customerID,
		amount,
		dueDate,
		issueDate,
		description,
	)
	if err != nil {
		return nil, fmt.Errorf("error al crear la factura: %w", err)
	}

	if err := s.repository.Save(invoice); err != nil {
		return nil, fmt.Errorf("error al persistir la factura: %w", err)
	}

	return invoice, nil
}

func (s *invoiceService) GetInvoice(ctx context.Context, id string) (*domain.Invoice, error) {
	if id == "" {
		return nil, fmt.Errorf("el ID de la factura es requerido")
	}

	invoice, err := s.repository.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("error al buscar la factura: %w", err)
	}

	if invoice == nil {
		return nil, fmt.Errorf("factura no encontrada con ID: %s", id)
	}

	return invoice, nil
}

func (s *invoiceService) UpdateInvoice(ctx context.Context, id string, amount float64, status domain.InvoiceStatus) (*domain.Invoice, error) {
	if id == "" {
		return nil, fmt.Errorf("el ID de la factura es requerido")
	}

	invoice, err := s.repository.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("error al buscar la factura para actualizar: %w", err)
	}

	if invoice == nil {
		return nil, fmt.Errorf("factura no encontrada con ID: %s", id)
	}

	if amount > 0 {
		invoice.Amount = amount
	}

	if status != "" {
		if err := invoice.UpdateStatus(status); err != nil {
			return nil, fmt.Errorf("error al actualizar el estado de la factura: %w", err)
		}
	}

	if err := invoice.Validate(); err != nil {
		return nil, fmt.Errorf("la factura no es válida después de la actualización: %w", err)
	}

	if err := s.repository.Update(invoice); err != nil {
		return nil, fmt.Errorf("error al actualizar la factura: %w", err)
	}

	return invoice, nil
}

func generateInvoiceID() string {
	return fmt.Sprintf("INV-%d", time.Now().UnixNano())
}

func ToProtoInvoice(invoice *domain.Invoice) *proto.InvoiceResponse {
	return &proto.InvoiceResponse{
		Id:          invoice.ID,
		CustomerId:  invoice.CustomerID,
		Amount:      invoice.Amount,
		Status:      string(invoice.Status),
		DueDate:     timestamppb.New(invoice.DueDate),
		IssueDate:   timestamppb.New(invoice.IssueDate),
		Description: invoice.Description,
	}
}

func ToProtoInvoiceList(invoices []*domain.Invoice) []*proto.InvoiceResponse {
	responses := make([]*proto.InvoiceResponse, len(invoices))
	for i, inv := range invoices {
		responses[i] = ToProtoInvoice(inv)
	}
	return responses
}

// === ARCHIVO: pkg/proto/invoice.proto ===
syntax = "proto3";

package proto;

option go_package = "invoice-service/pkg/proto";

import "google/protobuf/timestamp.proto";

message CreateInvoiceRequest {
  string customer_id = 1;
  double amount = 2;
  google.protobuf.Timestamp due_date = 3;
  google.protobuf.Timestamp issue_date = 4;
  string description = 5;
}

message CreateInvoiceResponse {
  InvoiceResponse invoice = 1;
  string error = 2;
}

message GetInvoiceRequest {
  string id = 1;
}

message GetInvoiceResponse {
  InvoiceResponse invoice = 1;
  string error = 2;
}

message UpdateInvoiceRequest {
  string id = 1;
  double amount = 2;
  string status = 3;
}

message UpdateInvoiceResponse {
  InvoiceResponse invoice = 1;
  string error = 2;
}

message InvoiceResponse {
  string id = 1;
  string customer_id = 2;
  double amount = 3;
  string status = 4;
  google.protobuf.Timestamp due_date = 5;
  google.protobuf.Timestamp issue_date = 6;
  string description = 7;
}

message HealthCheckRequest {
  string service = 1;
}

message HealthCheckResponse {
  enum ServingStatus {
    UNKNOWN = 0;
    SERVING = 1;
    NOT_SERVING = 2;
  }
  ServingStatus status = 1;
}

service InvoiceService {
  rpc CreateInvoice(CreateInvoiceRequest) returns (CreateInvoiceResponse);
  rpc GetInvoice(GetInvoiceRequest) returns (GetInvoiceResponse);
  rpc UpdateInvoice(UpdateInvoiceRequest) returns (UpdateInvoiceResponse);
}

service Health {
  rpc Check(HealthCheckRequest) returns (HealthCheckResponse);
  rpc Watch(HealthCheckRequest) returns (stream HealthCheckResponse);
}

// === ARCHIVO: pkg/proto/invoice.pb.go ===
package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	constpref "google.golang.org/genproto/googleapis/type/date"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	File_invoice_proto_name          = "invoice.proto"
	File_invoice_proto_rawDesc       = "\n\rinvoice.proto\x12\x05proto\x1a\x1fgoogle/protobuf/timestamp.proto" +
		"\xa0\x01\x03\xa8\x01\x03\xe8\x01\x03\x12\x8b\x01\n\x17CreateInvoiceRequest\x12\x12\n\x0bcustomer_id" +
		"\x18\x01 \x01(\tR\x0bcustomerId\x12\n\n\x06amount\x18\x02 \x01(\x01R\x06amount\x12\x31\n\x08due_date\x18\x03 " +
		"\x01(\x0b2\x1a.google.protobuf.TimestampR\x07dueDate\x12\x33\n\nissue_date\x18\x04 " +
		"\x01(\x0b2\x1a.google.protobuf.TimestampR\tissueDate\x12\x14\n\x0bdescription\x18\x05 " +
		"\x01(\tR\x0bdescription\x12\x8d\x01\n\x18CreateInvoiceResponse\x12-\n\x07invoice\x18\x01 " +
		"\x01(\x0b2\x14.proto.InvoiceResponseR\x07invoice\x12\x10\n\x05error\x18\x02 \x01(\tR\x05error" +
		"\x12\x80\x01\n\x14GetInvoiceRequest\x12\n\n\x02id\x18\x01 \x01(\tR\x02id\x12\x81\x01\n\x15GetInvoiceResponse" +
		"\x12-\n\x07invoice\x18\x01 \x01(\x0b2\x14.proto.InvoiceResponseR\x07invoice\x12\x10\n\x05error\x18\x02 " +
		"\x01(\tR\x05error\x12\x8a\x01\n\x16UpdateInvoiceRequest\x12\n\n\x02id\x18\x01 \x01(\tR\x02id\x12\n\n\x06amount" +
		"\x18\x02 \x01(\x01R\x06amount\x12\x10\n\x06status\x18\x03 \x01(\tR\x06status\x12\x8b\x01\n\x17UpdateInvoiceResponse" +
		"\x12-\n\x07invoice\x18\x01 \x01(\x0b2\x14.proto.InvoiceResponseR\x07invoice\x12\x10\n\x05error\x18\x02 " +
		"\x01(\tR\x05error\x12\x8a\x01\n\x0eInvoiceResponse\x12\n\n\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n\x0bcustomer_id" +
		"\x18\x02 \x01(\tR\x0bcustomerId\x12\n\n\x06amount\x18\x03 \x01(\x01R\x06amount\x12\x10\n\x06status\x18\x04 " +
		"\x01(\tR\x06status\x12\x31\n\x08due_date\x18\x05 \x01(\x0b2\x1a.google.protobuf.TimestampR\x07dueDate" +
		"\x12\x33\n\nissue_date\x18\x06 \x01(\x0b2\x1a.google.protobuf.TimestampR\tissueDate\x12\x14\n\x0bdescription" +
		"\x18\x07 \x01(\tR\x0bdescription\x12\x80\x01\n\x11HealthCheckRequest\x12\x14\n\x07service\x18\x01 " +
		"\x01(\tR\x07service\x12\x8d\x01\n\x12HealthCheckResponse\x12-\n\x06status\x18\x01 " +
		"\x01(\x0e2\x1d.proto.HealthCheckResponse.ServingStatusR\x06status" +
		"\xe8\x01\x02\x12\n\n\x0eInvoiceService\x12\x46\n\x0eCreateInvoice\x12\x1a.proto.CreateInvoiceRequest" +
		"(\x0b2\x1b.proto.CreateInvoiceResponseR\x0eCreateInvoice\x12\x41\n\nGetInvoice\x12\x17.proto.GetInvoiceRequest" +
		"(\x0b2\x18.proto.GetInvoiceResponseR\nGetInvoice\x12\x43\n\x0cUpdateInvoice\x12\x19.proto.UpdateInvoiceRequest" +
		"(\x0b2\x1a.proto.UpdateInvoiceResponseR\x0cUpdateInvoice\x12\x1e\n\x05Health\x12\x14.proto.HealthCheckRequest" +
		"(\x0b2\x15.proto.HealthCheckResponseR\x05Health\x12\x1e\n\x05Watch\x12\x14.proto.HealthCheckRequest" +
		"\x12\x15.proto.HealthCheckResponseB\xca\x01\x1bGo package invoice-service/pkg/protob\x06proto3"
	File_invoice_proto_reflectDescFileInfo = []string{
		"invoice.proto",
	}
)

var File_invoice_proto protoreflect.FileDescriptor

func init() { protoimpl.RegisterFile(File_invoice_proto_name, File_invoice_proto_rawDesc, File_invoice_proto_reflectDescFileInfo) }

var File_invoice_proto_rawDescOnce sync.Once

func File_invoice_proto_rawDescGZIP() []byte {
	File_invoice_proto_rawDescOnce.Do(func() {
		File_invoice_proto_rawDesc = []byte{
			0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcd, 0x6e, 0xe3, 0x36,
			0x10, 0xc6, 0x5e, 0x40, 0x12, 0x8e, 0x9c, 0x84, 0x4a, 0x93, 0x4e, 0x92, 0x4b, 0x93, 0x52, 0x24,
			0x12, 0x65, 0x25, 0x24, 0x65, 0x25, 0xb2, 0x4a, 0x89, 0x52, 0x48, 0xa5, 0x28, 0x4b, 0xa5, 0x6c,
			0x28, 0x4a, 0x81, 0xd4, 0x6a, 0x28, 0x94, 0x0a, 0x52, 0x5c, 0x90, 0x6a, 0x90, 0x02, 0x52, 0x5c,
			0x90, 0x6a, 0x90, 0x02, 0x52, 0x5c, 0x90, 0x6a, 0x90, 0x02, 0x52, 0x5c, 0x90, 0x6a, 0x90, 0x02,
			0x52, 0x5c, 0x90, 0x6a, 0x90, 0x02, 0x52, 0x5c, 0x90, 0x6a, 0x90, 0x82, 0x50, 0x2d, 0x25, 0x4a,
			0x4a, 0x6d, 0xf1, 0x44, 0x60, 0x71, 0xe5, 0x97, 0xdc, 0x23, 0xb1, 0x89, 0x5c, 0x22, 0x24, 0x95,
			0x4a, 0x32, 0xb9, 0x40, 0x90, 0x90, 0x92, 0x4a, 0x32, 0xb9, 0x40, 0x90, 0x92, 0x4a, 0x32, 0xb9,
			0x40, 0x90, 0x92, 0x4a, 0x32, 0xb9, 0x40, 0x90, 0x92, 0x4a, 0x32, 0xb9, 0x40, 0x90, 0x92, 0x4a,
			0x32, 0xb9, 0x40, 0x90, 0x92, 0x4a, 0x32, 0xb9, 0x40, 0x90, 0x92, 0x4a, 0x32, 0xb9, 0x40, 0x90,
			0x92, 0x4a, 0x12, 0x23, 0x21, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81,
			0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64,
			0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25,
			0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81,
			0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64,
			0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25, 0x95, 0x64, 0x72, 0x81, 0x20, 0x25,
			0x95, 0xa4, 0x48, 0x52, 0x52, 0x49, 0x93, 0x03, 0x04, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23,
			0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23,
			0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23,
			0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23,
			0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23,
			0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23, 0x23,
			0xa3, 0xfd, 0x0e, 0x6e, 0x80, 0x8c, 0x8c, 0x4e, 0x40, 0x22, 0x25, 0x0e, 0x2e, 0x46, 0x4a, 0x9c,
			0x80, 0x9c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		}
	})
	return File_invoice_proto_rawDesc
}

type CreateInvoiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields unknownFields

	CustomerID  string                 `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Amount      float64                `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	DueDate     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	IssueDate   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	Description string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateInvoiceRequest) Reset() {
	*x = CreateInvoiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateInvoiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateInvoiceRequest) ProtoMessage() {}

func (x *CreateInvoiceRequest) GetCustomerID() string {
	if x != nil {
		return x.CustomerID
	}
	return ""
}

func (x *CreateInvoiceRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateInvoiceRequest) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

func (x *CreateInvoiceRequest) GetIssueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.IssueDate
	}
	return nil
}

func (x *CreateInvoiceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type InvoiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields unknownFields

	Id          string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CustomerId  string                  `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
	Amount      float64                 `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Status      string                  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	DueDate     *timestamppb.Timestamp  `protobuf:"bytes,5,opt,name=due_date,json=dueDate,proto3" json:"due_date,omitempty"`
	IssueDate   *timestamppb.Timestamp  `protobuf:"bytes,6,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	Description string                  `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *InvoiceResponse) Reset() {
	*x = InvoiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_invoice_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvoiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvoiceResponse) ProtoMessage() {}

func (x *InvoiceResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InvoiceResponse) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *InvoiceResponse) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *InvoiceResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *InvoiceResponse) GetDueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DueDate
	}
	return nil
}

func (x *InvoiceResponse) GetIssueDate() *timestamppb.Timestamp {
	if x != nil {
		return x.IssueDate
	}
	return nil
}

func (x *InvoiceResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var file_invoice_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_invoice_proto_goTypes = []interface{}{
	(*CreateInvoiceRequest)(nil),
	(*CreateInvoiceResponse)(nil),
	(*GetInvoiceRequest)(nil),
	(*GetInvoiceResponse)(nil),
	(*UpdateInvoiceRequest)(nil),
	(*UpdateInvoiceResponse)(nil),
	(*InvoiceResponse)(nil),
	(*HealthCheckRequest)(nil),
	(*HealthCheckResponse)(nil),
}

var file_invoice_proto_rawDesc = protoimpl.RawDesc{
	TX:     rawDesc,
	GoTypes: file_invoice_proto_goTypes,
	MessageInfos: file_invoice_proto_msgTypes,
}

func init() { proto.RegisterFile(File_invoice_proto_name, file_invoice_proto_rawDesc) }

var File_invoice_proto protoreflect.FileDescriptor = file_invoice_proto_rawDesc


// === ARCHIVO: Dockerfile ===
# Etapa de compilación
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Instalar herramientas necesarias para compilación
RUN apk add --no-cache git make protobuf

# Copiar go.mod y go.sum primero para aprovechar cache
COPY go.mod go.sum ./
RUN go mod download

# Copiar código fuente
COPY . .

# Generar código protobuf
RUN protoc --go_out=. --go-grpc_out=. pkg/proto/invoice.proto

# Compilar el binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /invoice-service ./cmd/invoice-service

# Etapa final - imagen minima
FROM alpine:3.19

WORKDIR /app

# Instalar certificadas necesarias para Alpine
RUN apk add --no-cache ca-certificates openssl

# Copiar el binary desde la etapa de build
COPY --from=builder /invoice-service /app/invoice-service

# Copiar configuración si existe
COPY config/ /app/config/

# Exponer puerto del servicio gRPC
EXPOSE 50051

# Health check para Kubernetes
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD grpc_health_probe -addr=localhost:50051 || exit 1

# Instalar grpc_health_probe
RUN apk add --no-cache grpc-tools

# Ejecutar el servicio
CMD ["/app/invoice-service"]

// === ARCHIVO: .dockerignore ===
# Archivos de control de versiones
.git
.gitignore

# Documentación
README.md
docs/
documentation/
*.md

# Archivos de configuración local
.env
.env.local
.env.*.local

# Archivos de IDE y herramientas
.idea/
.vscode/
*.swp
*.swo
*~

# Archivos de build y compilación
/bin/
/obj/
/out/
dist/
build/

# Tests y coverage
*_test.go
*.test
coverage.txt
coverage.html

# Archivos de dependencias descargadas
vendor/

# Archivos temporales
tmp/
temp/
*.log

# Kubernetes y Docker (se manejan fuera del contexto de build)
deployments/
k8s/
docker-compose*.yml
Dockerfile

# Archivos de desarrollo local
Makefile
*.sum

// === ARCHIVO: docker-compose.yml ===
version: '3.8'

services:
  invoice-service:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: invoice-service
    ports:
      - "50051:50051"
      - "9090:9090"
    environment:
      - GRPC_PORT=50051
      - HTTP_PORT=9090
      - LOG_LEVEL=info
      - DB_TYPE=inmemory
    networks:
      - invoice-network
    restart: unless-stopped
    healthcheck:
      test: ["CMD", "grpc_health_probe", "-addr=localhost:50051"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 10s
    volumes:
      - ./config:/app/config:ro

  # Base de datos PostgreSQL para producción
  postgres:
    image: postgres:15-alpine
    container_name: invoice-postgres
    environment:
      POSTGRES_DB: invoice_db
      POSTGRES_USER: invoice_user
      POSTGRES_PASSWORD: invoice_pass
    ports:
      - "5432:5432"
    volumes:
      - postgres-data:/var/lib/postgresql/data
    networks:
      - invoice-network
    restart: unless-stopped

  # Redis para caché
  redis:
    image: redis:7-alpine
    container_name: invoice-redis
    ports:
      - "6379:6379"
    networks:
      - invoice-network
    restart: unless-stopped
    command: redis-server --appendonly yes
    volumes:
      - redis-data:/data

  # Prometheus para métricas
  prometheus:
    image: prom/prometheus:latest
    container_name: invoice-prometheus
    ports:
      - "9091:9090"
    volumes:
      - ./monitoring/prometheus.yml:/etc/prometheus/prometheus.yml:ro
    networks:
      - invoice-network
    restart: unless-stopped

  # Grafana para visualización de métricas
  grafana:
    image: grafana/grafana:latest
    container_name: invoice-grafana
    ports:
      - "3000:3000"
    environment:
      - GF_SECURITY_ADMIN_PASSWORD=admin
      - GF_USERS_ALLOW_SIGN_UP=false
    volumes:
      - grafana-data:/var/lib/grafana
      - ./monitoring/dashboards:/etc/grafana/provisioning/dashboards:ro
    networks:
      - invoice-network
    restart: unless-stopped
    depends_on:
      - prometheus

networks:
  invoice-network:
    driver: bridge

volumes:
  postgres-data:
  redis-data:
  grafana-data:


// === ARCHIVO: README.md ===
# Invoice Service - Microservicio de Facturación

Microservicio en Go para la gestión de facturas de servicios de telecomunicaciones. Implementa comunicación gRPC y está diseñado para desplegarse en contenedores Docker y clústeres Kubernetes.

## Requisitos Previos

- Go 1.22 o superior
- Docker y Docker Compose
- Kubernetes (minikube, kind o un clúster cloud)
- protoc (Protocol Buffers compiler)

## Estructura del Proyecto

```
invoice-service/
├── cmd/                          # Punto de entrada del servicio
│   └── invoice-service/main.go
├── internal/                     # Código interno del dominio
│   ├── domain/                   # Entidades y lógica de dominio
│   │   └── invoice.go
│   ├── application/              # Servicios de aplicación
│   │   └── invoice_service.go
│   ├── infrastructure/           # Implementaciones de persistencia
│   │   └── repository.go
│   └── interfaces/               # Adaptadores de entrada
│       └── grpc/
│           └── invoice_handler.go
├── pkg/proto/                    # Definiciones protobuf generadas
│   ├── invoice.proto
│   └── invoice.pb.go
├── deployments/                  # Configuraciones de despliegue
│   ├── docker/
│   │   ├── Dockerfile
│   │   └── docker-compose.yml
│   └── k8s/
│       ├── invoice-service-deployment.yaml
│       └── invoice-service-service.yaml
├── test/                         # Pruebas unitarias y de integración
│   └── invoice_service_test.go
├── go.mod                        # Dependencias Go
└── README.md
```

## Compilación y Ejecución Local

### Compilar el servicio

```bash
go build -o bin/invoice-service ./cmd/invoice-service
```

### Ejecutar el servicio

```bash
./bin/invoice-service
```

El servicio arrancará en el puerto 50051 por defecto.

### Ejecutar pruebas

```bash
go test -v ./...
```

## Docker

### Construir imagen

```bash
docker build -t invoice-service:latest .
```

### Ejecutar con Docker Compose

```bash
docker-compose up -d
```

## Kubernetes

### Requisitos

- kubectl configurado
- Acceso a un clúster Kubernetes

### Desplegar en Kubernetes

```bash
kubectl apply -f deployments/k8s/invoice-service-deployment.yaml
kubectl apply -f deployments/k8s/invoice-service-service.yaml
```

### Verificar despliegue

```bash
kubectl get pods -l app=invoice-service
kubectl get services invoice-service
```

### Ver logs

```bash
kubectl logs -l app=invoice-service
```

## gRPC API

El servicio expone los siguientes métodos:

- **CreateInvoice**: Crear una nueva factura
- **GetInvoice**: Consultar una factura por ID
- **UpdateInvoice**: Actualizar el estado de una factura

### Definición de proto

Ver `pkg/proto/invoice.proto` para la definición completa de la API.

## Configuración

El servicio configurable mediante variables de entorno:

- `PORT`: Puerto del servidor gRPC (default: 50051)
- `LOG_LEVEL`: Nivel de logging (debug, info, warn, error)

## Decisiones de Diseño

- **Patrón arquitectónico**: Capas estándar de microservicio (dominio, aplicación, infraestructura, interfaces)
- **Comunicación**: gRPC para comunicación síncrona entre servicios
- **Persistencia**: Repositorio en memoria para desarrollo, extensible a base de datos real
- **Salud**: Implementa health checks de Kubernetes mediante gRPC health protocol

## Contribuir

1. Fork del repositorio
2. Crear branch de feature (`git checkout -b feature/nueva-funcionalidad`)
3. Commit de cambios (`git commit -m 'Agrega nueva funcionalidad'`)
4. Push al branch (`git push origin feature/nueva-funcionalidad`)
5. Crear Pull Request

## Licencia

MIT License - consultar archivo LICENSE para más detalles.

// === ARCHIVO: deployments/k8s/invoice-service-deployment.yaml ===
apiVersion: apps/v1
kind: Deployment
metadata:
  name: invoice-service
  labels:
    app: invoice-service
    version: v1
spec:
  replicas: 3
  selector:
    matchLabels:
      app: invoice-service
  template:
    metadata:
      labels:
        app: invoice-service
        version: v1
    spec:
      containers:
      - name: invoice-service
        image: invoice-service:latest
        imagePullPolicy: IfNotPresent
        ports:
        - containerPort: 50051
          name: grpc
          protocol: TCP
        env:
        - name: PORT
          value: "50051"
        - name: LOG_LEVEL
          value: "info"
        resources:
          requests:
            memory: "128Mi"
            cpu: "100m"
          limits:
            memory: "256Mi"
            cpu: "500m"
        livenessProbe:
          grpc:
            port: 50051
            service: ""
          initialDelaySeconds: 10
          periodSeconds: 10
          timeoutSeconds: 5
          failureThreshold: 3
        readinessProbe:
          grpc:
            port: 50051
            service: ""
          initialDelaySeconds: 5
          periodSeconds: 5
          timeoutSeconds: 3
          failureThreshold: 2
// === ARCHIVO: deployments/k8s/invoice-service-service.yaml ===
apiVersion: v1
kind: Service
metadata:
  name: invoice-service
  labels:
    app: invoice-service
spec:
  type: ClusterIP
  ports:
  - port: 50051
    targetPort: 50051
    protocol: TCP
    name: grpc
  selector:
    app: invoice-service
---
apiVersion: v1
kind: Service
metadata:
  name: invoice-service-headless
  labels:
    app: invoice-service
spec:
  clusterIP: None
  ports:
  - port: 50051
    targetPort: 50051
    protocol: TCP
    name: grpc
  selector:
    app: invoice-service


// === ARCHIVO: test/invoice_service_test.go ===
package test

import (
	"context"
	"net"
	"testing"
	"time"

	"invoice-service/cmd"
	"invoice-service/internal/domain"
	"invoice-service/internal/infrastructure"
	"invoice-service/internal/interfaces/grpc"
	"invoice-service/pkg/proto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	testPort       = "127.0.0.1:0"
	testTimeout    = 5 * time.Second
)

type testServer struct {
	ln           net.Listener
	grpcServer   *grpc.Server
	client       proto.InvoiceServiceClient
	closeFunc    func()
}

func setupTestServer(t *testing.T) *testServer {
	ln, err := net.Listen("tcp", testPort)
	require.NoError(t, err, "Error al iniciar listener para pruebas")

	repository := infrastructure.NewInvoiceRepositoryInMemory()
	service := domain.NewInvoiceService(repository)
	handler := grpc.NewInvoiceHandler(service)

	grpcServer := grpc.NewServer()
	proto.RegisterInvoiceServiceServer(grpcServer, handler)

	go func() {
		if err := grpcServer.Serve(ln); err != nil {
			t.Logf("Servidor gRPC detenido: %v", err)
		}
	}()

	conn, err := grpc.NewClient(
		ln.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err, "Error al conectar con el servidor de pruebas")

	client := proto.NewInvoiceServiceClient(conn)

	return &testServer{
		ln:         ln,
		grpcServer: grpcServer,
		client:     client,
		closeFunc: func() {
			conn.Close()
			grpcServer.GracefulStop()
		},
	}
}

func TestCreateInvoice_Success(t *testing.T) {
	server := setupTestServer(t)
	defer server.closeFunc()

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	dueDate := time.Now().Add(30 * 24 * time.Hour)
	req := &proto.CreateInvoiceRequest{
		CustomerId:   "CUST-001",
		Amount:       1500.50,
		DueDate:      timestamppb.New(dueDate),
		Description:  "Servicio de telecomunicación mensual",
	}

	resp, err := server.client.CreateInvoice(ctx, req)
	require.NoError(t, err, "CreateInvoice debería成功了")

	assert.NotEmpty(t, resp.InvoiceId, "El ID de factura no debería estar vacío")
	assert.Equal(t, "CUST-001", resp.CustomerId)
	assert.Equal(t, 1500.50, resp.Amount)
	assert.Equal(t, string(domain.InvoiceStatusPending), resp.Status)
}

func TestCreateInvoice_InvalidAmount(t *testing.T) {
	server := setupTestServer(t)
	defer server.closeFunc()

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	req := &proto.CreateInvoiceRequest{
		CustomerId:   "CUST-002",
		Amount:       -100.0,
		DueDate:      timestamppb.New(time.Now().Add(30 * 24 * time.Hour)),
		Description:  "Monto inválido",
	}

	resp, err := server.client.CreateInvoice(ctx, req)
	assert.Error(t, err, "Debería fallar con monto negativo")
	assert.Nil(t, resp)
}

func TestGetInvoice_Exists(t *testing.T) {
	server := setupTestServer(t)
	defer server.closeFunc()

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	createReq := &proto.CreateInvoiceRequest{
		CustomerId:   "CUST-003",
		Amount:       2500.00,
		DueDate:      timestamppb.New(time.Now().Add(15 * 24 * time.Hour)),
		Description:  "Factura para recuperación",
	}

	createResp, err := server.client.CreateInvoice(ctx, createReq)
	require.NoError(t, err, "La creación debería funcionar")

	getReq := &proto.GetInvoiceRequest{
		InvoiceId: createResp.InvoiceId,
	}

	getResp, err := server.client.GetInvoice(ctx, getReq)
	require.NoError(t, err, "GetInvoice debería encontrar la factura")

	assert.Equal(t, createResp.InvoiceId, getResp.InvoiceId)
	assert.Equal(t, "CUST-003", getResp.CustomerId)
	assert.Equal(t, 2500.00, getResp.Amount)
}

func TestGetInvoice_NotFound(t *testing.T) {
	server := setupTestServer(t)
	defer server.closeFunc()

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	req := &proto.GetInvoiceRequest{
		InvoiceId: "NONEXISTENT-ID",
	}

	resp, err := server.client.GetInvoice(ctx, req)
	assert.Error(t, err, "Debería fallar al buscar una factura inexistente")
	assert.Nil(t, resp)
}

func TestUpdateInvoice_StatusChange(t *testing.T) {
	server := setupTestServer(t)
	defer server.closeFunc()

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	createReq := &proto.CreateInvoiceRequest{
		CustomerId:   "CUST-004",
		Amount:       3000.75,
		DueDate:      timestamppb.New(time.Now().Add(45 * 24 * time.Hour)),
		Description:  "Factura para actualizar estado",
	}

	createResp, err := server.client.CreateInvoice(ctx, createReq)
	require.NoError(t, err)

	updateReq := &proto.UpdateInvoiceRequest{
		InvoiceId: createResp.InvoiceId,
		Status:    string(domain.InvoiceStatusPaid),
	}

	updateResp, err := server.client.UpdateInvoice(ctx, updateReq)
	require.NoError(t, err, "UpdateInvoice debería actualizar el estado")

	assert.Equal(t, string(domain.InvoiceStatusPaid), updateResp.Status)
}

func TestUpdateInvoice_InvalidStatus(t *testing.T) {
	server := setupTestServer(t)
	defer server.closeFunc()

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	createReq := &proto.CreateInvoiceRequest{
		CustomerId:   "CUST-005",
		Amount:       500.00,
		DueDate:      timestamppb.New(time.Now().Add(10 * 24 * time.Hour)),
		Description:  "Factura para probar estado inválido",
	}

	createResp, err := server.client.CreateInvoice(ctx, createReq)
	require.NoError(t, err)

	updateReq := &proto.UpdateInvoiceRequest{
		InvoiceId: createResp.InvoiceId,
		Status:    "INVALID_STATUS",
	}

	resp, err := server.client.UpdateInvoice(ctx, updateReq)
	assert.Error(t, err, "Debería fallar con estado inválido")
	assert.Nil(t, resp)
}

func TestCreateInvoice_MultipleInvoices(t *testing.T) {
	server := setupTestServer(t)
	defer server.closeFunc()

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	invoicesCreated := make(map[string]bool)

	for i := 0; i < 5; i++ {
		req := &proto.CreateInvoiceRequest{
			CustomerId:   "CUST-MULTI",
			Amount:       float64((i + 1) * 100),
			DueDate:      timestamppb.New(time.Now().Add(time.Duration(i+1) * 24 * time.Hour)),
			Description:  "Factura múltiple número " + string(rune('1'+i)),
		}

		resp, err := server.client.CreateInvoice(ctx, req)
		require.NoError(t, err, "Debería crear la factura %d", i+1)
		invoicesCreated[resp.InvoiceId] = true
	}

	assert.Equal(t, 5, len(invoicesCreated), "Deberían crearse 5 facturas")
}

func TestInvoiceConcurrency(t *testing.T) {
	server := setupTestServer(t)
	defer server.closeFunc()

	ctx, cancel := context.WithTimeout(context.Background(), testTimeout)
	defer cancel()

	resultChan := make(chan *proto.InvoiceResponse, 10)
	errChan := make(chan error, 10)

	for i := 0; i < 10; i++ {
		go func(index int) {
			req := &proto.CreateInvoiceRequest{
				CustomerId:   "CUST-CONCURRENT",
				Amount:       float64(index * 10),
				DueDate:      timestamppb.New(time.Now().Add(20 * 24 * time.Hour)),
				Description:  "Factura concurrente",
			}

			resp, err := server.client.CreateInvoice(ctx, req)
			if err != nil {
				errChan <- err
				return
			}
			resultChan <- resp
		}(i)
	}

	count := 0
	errors := 0

	select {
	case resp := <-resultChan:
		count++
		assert.NotEmpty(t, resp.InvoiceId)
	case err := <-errChan:
		errors++
		t.Logf("Error concurrente: %v", err)
	case <-time.After(testTimeout):
	}

	for {
		select {
		case resp := <-resultChan:
			count++
			assert.NotEmpty(t, resp.InvoiceId)
		case err := <-errChan:
			errors++
			t.Logf("Error concurrente: %v", err)
		default:
			goto done
		}
	}

done:
	assert.Equal(t, 10, count+errors, "Deberían procesarse 10 solicitudes")
}

func TestContextCancellation(t *testing.T) {
	server := setupTestServer(t)
	defer server.closeFunc()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	req := &proto.CreateInvoiceRequest{
		CustomerId:   "CUST-CANCEL",
		Amount:       100.00,
		DueDate:      timestamppb.New(time.Now().Add(5 * 24 * time.Hour)),
		Description:  "Factura con contexto cancelado",
	}

	resp, err := server.client.CreateInvoice(ctx, req)
	assert.Error(t, err, "Debería fallar con contexto cancelado")
	assert.Nil(t, resp)
}

```
