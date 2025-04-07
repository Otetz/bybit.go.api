package bybit_connector

import (
	"context"
	"net/http"

	"github.com/wuhewuhe/bybit.go.api/handlers"
)

// GetOrdersRealtime get Open & Closed orders
func (s *BybitClientRequest) GetOrdersRealtime(ctx context.Context, opts ...RequestOption) (res *ServerResponse, err error) {
	if err = handlers.ValidateParams(s.params); err != nil {
		return nil, err
	}
	r := &request{
		method:   http.MethodGet,
		endpoint: "/v5/order/realtime",
		secType:  secTypeSigned,
	}
	data := SendRequest(ctx, opts, r, s, err)
	return GetServerResponse(err, data)
}
