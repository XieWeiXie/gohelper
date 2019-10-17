package go_helper

import (
	"bytes"
	"fmt"
	"strings"
	"syscall"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)
const (
	NORMAL   = "normal"
	WARNING  = "warning"
	CRITICAL = "critical" // default: 95%
)

type Disk struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}

var DefaultDisk = Disk{}

func NewDisk(path string) *Disk {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return nil
	}
	all := fs.Blocks * uint64(fs.Bsize)
	free := fs.Bfree * uint64(fs.Bsize)
	used := all - free
	return &Disk{
		All:  all,
		Used: used,
		Free: free,
	}
}

func (D Disk) GetTotal(unit string) uint64 {
	unit = strings.ToUpper(unit)
	switch unit {
	case "B":
		return D.All
	case "KB":
		return D.All / KB
	case "MB":
		return D.All / MB
	default:
		return D.All / GB
	}
}
func (D Disk) GetUsed(unit string) uint64 {
	unit = strings.ToUpper(unit)
	switch unit {
	case "B":
		return D.Used
	case "KB":
		return D.Used / KB
	case "MB":
		return D.Used / MB
	default:
		return D.Used / GB
	}
}
func (D Disk) GetFree(unit string) uint64 {
	return D.GetTotal(unit) - D.GetUsed(unit)
}

func (D Disk) PercentForFree() float64 {
	return float64(D.Free) * 100 / float64(D.All)
}
func (D Disk) PercentForUsed() float64 {
	return float64(D.Used) * 100 / float64(D.All)
}

func (D Disk) Status(threshold int) string {
	percent := D.PercentForUsed()
	if percent >= 95 {
		return CRITICAL
	} else if percent >= float64(threshold) {
		return WARNING
	} else {
		return NORMAL
	}
}

func (D Disk) String() string {
	by := bytes.Buffer{}
	by.WriteString("{")
	by.WriteString(fmt.Sprintf(`"all": "%dGB", "used": "%dGB", "free": "%dGB", "usedpercent": "%f%%"`,
		D.GetTotal("GB"), D.GetUsed("GB"), D.GetFree("GB"), D.PercentForUsed()))
	by.WriteString("}")
	return by.String()

}
