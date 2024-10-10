package handler

type OrdersGrpcHandler struct {

	//service injection
	//unimplemented UnimplementedOrderServiceServer
}

func NewGrpcOrdersService() {
	gRPCHandler := &OrdersGrpcHandler{}

	// register the OrderServiceServer
}
