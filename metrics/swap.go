package metrics

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	. "github.com/mlabouardy/cloudwatch/services"
	"github.com/shirou/gopsutil/mem"
)

type Swap struct{}

func (d Swap) Collect(instanceId string, c CloudWatchService) {
	swapMetrics, err := mem.SwapMemory()
	if err != nil {
		log.Fatal(err)
	}

	swapUtilizationData := constructMetricDatum("SwapUtilization", swapMetrics.UsedPercent, cloudwatch.StandardUnitPercent, instanceId)
	c.Publish(swapUtilizationData, "CustomMetrics")

	swapUsedData := constructMetricDatum("SwapUsed", float64(swapMetrics.Used), cloudwatch.StandardUnitBytes, instanceId)
	c.Publish(swapUsedData, "CustomMetrics")

	swapFreeData := constructMetricDatum("SwapFree", float64(swapMetrics.Free), cloudwatch.StandardUnitBytes, instanceId)
	c.Publish(swapFreeData, "CustomMetrics")

	log.Printf("Swap - Utilization:%v%% Used:%v Free:%v\n", swapMetrics.UsedPercent, swapMetrics.Used, swapMetrics.Free)
}
