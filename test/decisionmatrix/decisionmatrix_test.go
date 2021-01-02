package main

import (
	//	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/dslparser"
	"github.com/skyhackvip/risk_engine/internal"
	"testing"
)

func TestDecisionMatrix(t *testing.T) {
	internal.SetFeature("model_1", 85)
	internal.SetFeature("model_2", 180)
	dsl := dslparser.LoadDslFromFile("decisionmatrix.yaml")
	rs, err := dsl.ParseDecisionMatrix(dsl.DecisionMatrixs[0])
	if err != nil {
		t.Error(err)
	}
	if rs == "C" {
		t.Log("result is ", rs)
	} else {
		t.Error("result error,expert C, result is ", rs)
	}
}
