package test

import (
	"context"
	"testing"

	"github.com/bentenison/microservice/business/sdk/dbtest"
)

func BenchmarkExecution(b *testing.B) {
	db, err := dbtest.New()
	if err != nil {
		b.Fatal(err)
	}
	ctx := context.WithValue(context.TODO(), "tracectx", "1212345678890")
	for i := 0; i < b.N; i++ {
		res, err := db.BusDomain.ExecBus.ExecuteCode(ctx, "../static/code_0f047a28-797f-463e-97d8-381d412c27f2.py", "python")
		if err != nil {
			db.Log.Errorc(ctx, "error in executing code", map[string]interface{}{
				"error": err.Error(),
			})
		}
		db.Log.Infoc(ctx, "result from exec is", map[string]interface{}{
			"res": res,
		})
	}

}
