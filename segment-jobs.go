package main

import (
	"os"
	"os/exec"
	"strings"
	"strconv"
	"fmt"
)

func segmentJobs(p *powerline) {
	nJobs := -1

	ppid := os.Getppid()
	out, _ := exec.Command("ps", "-a", "-f", "-o", "ppid").Output()
	processes := strings.Split(string(out), "\n")
	for _, processPpidStr := range processes {
		processPpid, _ := strconv.ParseInt(strings.TrimSpace(processPpidStr), 10, 64)
		if int(processPpid) == ppid {
			nJobs++
		}
	}

	if nJobs > 0 {
		p.appendSegment(segment{
			content: fmt.Sprintf(" %d ", nJobs),
			foreground: p.theme.JobsFg,
			background: p.theme.JobsBg,
		})
	}
}
