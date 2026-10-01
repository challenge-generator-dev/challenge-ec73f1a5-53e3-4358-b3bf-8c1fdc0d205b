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