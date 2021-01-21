package main

import (
	"github.com/skyhackvip/risk_engine/configs"
	"github.com/skyhackvip/risk_engine/dslparser"
	"github.com/skyhackvip/risk_engine/internal/datasource"
	"testing"
)

func TestFlow1(t *testing.T) {
	dsl := dslparser.LoadDslFromFile("flow.yaml")
	rs := dsl.Parse().Decision
	if rs == nil {
		t.Error("nil")
		return
	}
	if rs.(int) == configs.DecisionMap["reject"] {
		t.Log("ok")
	} else {
		t.Log("result error!")
	}
}

func TestFlow2(t *testing.T) {
	datasource.SetFeature("feature_1", 20)
	datasource.SetFeature("feature_2", 20)
	datasource.SetFeature("feature_3", 20)

	dsl := dslparser.LoadDslFromFile("flow.yaml")
	rs := dsl.Parse().Decision
	if rs == nil {
		t.Log("ok")
	} else {
		t.Log("result error!")
	}
}
