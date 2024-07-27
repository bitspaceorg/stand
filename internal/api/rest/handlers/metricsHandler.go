package handlers

import (
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/bitspaceorg/STAND-FOSSHACK/internal/api/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type metricHandler struct {
}

func (m *metricHandler) metric(c *fiber.Ctx) error {
	memStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")
	hostStat, _ := host.Info()
	uptime := time.Duration(hostStat.Uptime) * time.Second
	hostname, _ := os.Hostname()
	username := os.Getenv("USER")
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":   "ok",
		"os":       runtime.GOOS,
		"cpuCount": runtime.NumCPU(),
		"memory": fiber.Map{
			"totalMemory": memStat.Total,
			"usedMemory":  memStat.Used,
		},
		"disk": fiber.Map{
			"totalDisk": ((diskStat.Total / 1024) / 1024) / 1024,
			"usedDisk":  ((diskStat.Used / 1024) / 1024) / 1024,
		},
		"upTime":   uptime.String(),
		"maxProcs": runtime.GOMAXPROCS(0),
		"hostname": hostname,
		"username": username,
	})
}

func SetupMetricRoutes(rh *rest.RestHandler) {
	app := rh.App
	m := metricHandler{}
	app.Get("/metric", m.metric)
}
