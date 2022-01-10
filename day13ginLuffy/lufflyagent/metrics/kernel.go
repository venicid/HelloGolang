package metrics

import (
	"fmt"
	"log"
	"lufflyagent/models"
	"lufflyagent/nux"
	"lufflyagent/settings"
)

func KernelMetrics() (L []*models.MetricValue) {

	maxFiles, err := nux.KernelMaxFiles()
	if err != nil {
		log.Println(err)
		return
	}
	tags := fmt.Sprintf("__IP=%s", settings.IP())
	L = append(L, models.GaugeValue("kernel.maxfiles", maxFiles, tags))

	allocateFiles, err := nux.KernelAllocateFiles()
	if err != nil {
		log.Println(err)
		return
	}

	L = append(L, models.GaugeValue("kernel.files.allocated", allocateFiles, tags))
	return
}
