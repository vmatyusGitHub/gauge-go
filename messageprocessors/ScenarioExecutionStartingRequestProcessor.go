package messageprocessors

import (
	"fmt"

	m "github.com/getgauge-contrib/gauge-go/gauge_messages"
	t "github.com/getgauge-contrib/gauge-go/testsuit"
)

type ScenarioExecutionStartingRequestProcessor struct{}

func (r *ScenarioExecutionStartingRequestProcessor) Process(msg *m.Message, context *t.GaugeContext) *m.Message {
	tags := msg.GetScenarioExecutionStartingRequest().GetCurrentExecutionInfo().GetCurrentScenario().GetTags()
	fmt.Println("==> Current Scenario Tags:", tags)
	fmt.Println("==> Current Spec Tags:", msg.GetScenarioExecutionStartingRequest().GetCurrentExecutionInfo().GetCurrentSpec().GetTags())
	specTags := msg.GetScenarioExecutionStartingRequest().GetCurrentExecutionInfo().GetCurrentSpec().GetTags()
	tags = append(tags, specTags[:]...)
	fmt.Println("==> Merged Tags:", tags)
	hooks := context.GetHooks(t.BEFORESCENARIO, tags)
	exInfo := msg.GetScenarioExecutionStartingRequest().GetCurrentExecutionInfo()

	return executeHooks(hooks, msg, exInfo)
}
