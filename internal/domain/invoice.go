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