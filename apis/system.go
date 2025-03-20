package apis

import (
	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	machinecode "github.com/super-l/machine-code"
	"time"
)

func init() {

	api.Register("GET", "system/cpu-info", cpuInfo)
	api.Register("GET", "system/cpu", cpuStats)
	api.Register("GET", "system/memory", memStats)
	api.Register("GET", "system/disk", diskStats)
	api.Register("GET", "system/machine", machineInfo)
}

func memStats(ctx *gin.Context) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		api.Error(ctx, err)
		return
	}
	api.OK(ctx, stat)
}

func cpuInfo(ctx *gin.Context) {
	info, err := cpu.Info()
	if err != nil {
		api.Error(ctx, err)
		return
	}
	if len(info) == 0 {
		api.Fail(ctx, "查询失败")
		return
	}
	api.OK(ctx, info[0])
}

func cpuStats(ctx *gin.Context) {
	//times, err := cpu.Times(false)
	times, err := cpu.Percent(time.Millisecond*200, false)
	if err != nil {
		api.Error(ctx, err)
		return
	}
	if len(times) == 0 {
		api.Fail(ctx, "查询失败")
		return
	}
	api.OK(ctx, int(times[0]))
}

func diskStats(ctx *gin.Context) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	usages := make([]*disk.UsageStat, 0)
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			api.Error(ctx, err)
			return
		}
		usages = append(usages, usage)
	}
	api.OK(ctx, usages)
}

func machineInfo(ctx *gin.Context) {
	if machinecode.MachineErr != nil {
		//fmt.Println("获取机器码信息错误:" + machinecode.MachineErr.Error())
		api.Fail(ctx, "error")
		return
	}
	api.OK(ctx, &machinecode.Machine)
}
