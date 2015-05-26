package funcs

import (
	"github.com/my-open-falcon/common/model"
)

func AgentMetrics() []*model.MetricValue {
	return []*model.MetricValue{GaugeValue("agent.alive", 1)}
}
