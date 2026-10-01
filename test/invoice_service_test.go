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