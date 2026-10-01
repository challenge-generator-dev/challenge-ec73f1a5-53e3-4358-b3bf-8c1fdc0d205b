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