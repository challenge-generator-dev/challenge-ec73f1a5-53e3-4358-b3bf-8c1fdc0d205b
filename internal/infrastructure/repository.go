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