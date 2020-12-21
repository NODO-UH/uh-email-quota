package quota

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var (
	ErrCommandError = errors.New("error running internal command")
)

type QuotaInfo struct {
	Value int64
	Limit int64
}

func ParseQuotaInfo(line_str string) (*QuotaInfo, error) {
	lines := strings.Split(line_str, "\n")
	if len(lines) < 2 {
		log.Println("ERROR: unexpected output format")
		return nil, ErrCommandError
	}
	// Check header
	if lines[0] != "Quota name\tType\tValue\tLimit\t%" {
		log.Println("ERROR: unexpected output format")
		return nil, ErrCommandError
	}

	// Parse info
	words := strings.Split(lines[1], "\t")
	if len(words) < 5 {
		log.Println("ERROR: unexpected output format")
		return nil, ErrCommandError
	}

	value, err := strconv.ParseInt(words[2], 10, 64)
	if err != nil {
		log.Println("ERROR: unexpected output format in Value field")
		return nil, ErrCommandError
	}
	limnit, err := strconv.ParseInt(words[3], 10, 64)
	if err != nil {
		log.Println("ERROR: unexpected output format in Limit field")
		return nil, ErrCommandError
	}
	return &QuotaInfo{
		Value: value,
		Limit: limnit,
	}, nil
}

func GetUserQuota(user_name string) (*QuotaInfo, error) {
	cmd := exec.Command("doveadm", "-f", "tab", "quota", "get", "-u", user_name)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		log.Println(err.Error())
		return nil, ErrCommandError
	}
	out_str := string(out)
	return ParseQuotaInfo(out_str)
}
