package delivery

import (
	"ohlc-data-api/api/service"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Test_dataController_Retrieve(t *testing.T) {
	type fields struct {
		svc    service.DataService
		logger *zap.Logger
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dataController{
				svc:    tt.fields.svc,
				logger: tt.fields.logger,
			}
			d.Retrieve(tt.args.ctx)
		})
	}
}
