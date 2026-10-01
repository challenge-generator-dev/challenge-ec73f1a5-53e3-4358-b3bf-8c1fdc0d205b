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