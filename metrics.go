package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/mem"
)

func getMemoryUsage() (string, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("получение данных о памяти: %v\n", err)
		return "", err
	}

	data := fmt.Sprintf("Total: %.2f GiB, Available: %.2f GiB, Used: %.1f%%\n",
		convertInGigaBytes(v.Total), convertInGigaBytes(v.Available), v.UsedPercent)

	return data, err
}

func convertInGigaBytes(num uint64) float64 {
	return float64(num) / 1024.0 / 1024.0 / 1024.0
}
