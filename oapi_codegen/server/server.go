package server

import (
	"microservice-apis-go/oapi_codegen/generated/api"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var orders []api.GetOrderSchema

type Server struct{}

func NewServer() api.ServerInterface {
	return Server{}
}
func (Server) GetOrders(ctx *gin.Context, params api.GetOrdersParams) {
	resp := orders
	ctx.JSON(http.StatusOK, resp)
}
func (Server) CreateOrder(ctx *gin.Context) {
	var order api.CreateOrderSchema
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resp := api.GetOrderSchema{
		Id:      uuid.New(),
		Created: time.Now(),
		Order:   order.Order,
		Status:  api.Created,
	}
	ctx.JSON(http.StatusCreated, resp)

	orders = append(orders, resp)
}
func (Server) GetOrder(ctx *gin.Context, id uuid.UUID) {
}

func (Server) CancelOrder(ctx *gin.Context, id uuid.UUID) {
}
func (Server) DeleteOrder(ctx *gin.Context, id uuid.UUID) {
}

func (Server) PayOrder(ctx *gin.Context, id uuid.UUID) {
}
func (Server) UpdateOrder(ctx *gin.Context, id uuid.UUID) {
}
