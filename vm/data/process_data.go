package data

type ProcessData struct {
	PID     string `json:"pid"`
	USER    string `json:"user"`
	PR      string `json:"pr"`
	NI      string `json:"ni"`
	VIRT    string `json:"virt"`
	RES     string `json:"res"`
	SHR     string `json:"shr"`
	S       string `json:"s"`
	CPU     string `json:"cpu"`
	MEM     string `json:"mem"`
	TIME    string `json:"time"`
	COMMAND string `json:"command"`
}

func NewProcessData(pid, user, pr, ni, virt, res, shr, s, cpu, mem, time, command string) *ProcessData {
	return &ProcessData{
		PID:     pid,
		USER:    user,
		PR:      pr,
		NI:      ni,
		VIRT:    virt,
		RES:     res,
		SHR:     shr,
		S:       s,
		CPU:     cpu,
		MEM:     mem,
		TIME:    time,
		COMMAND: command,
	}
}
