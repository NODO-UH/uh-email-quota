package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type QuotaInfo struct {
	Value      int
	Limit      int
	Percentage float64
}

func (qi *QuotaInfo) String() string {
	return fmt.Sprintf("%d/%d - %f\n", qi.Value, qi.Limit, qi.Percentage)
}

func ParseQuotaInfo(line_str string) (*QuotaInfo, error) {
	lines := strings.Split(line_str, "\n")
	if len(lines) < 2 {
		return nil, errors.New("unexpected output format")
	}
	// Check header
	if lines[0] != "Quota name\tType\tValue\tLimit\t%" {
		return nil, errors.New("unexpected output format")
	}

	// Parse info
	words := strings.Split(lines[1], "\t")
	if len(words) < 5 {
		return nil, errors.New("unexpected output format")
	}

	value, err := strconv.Atoi(words[2])
	if err != nil {
		return nil, errors.New("unexpected output format in Value field")
	}
	limnit, err := strconv.Atoi(words[3])
	if err != nil {
		return nil, errors.New("unexpected output format in Limit field")
	}
	percentage, err := strconv.ParseFloat(words[4], 64)
	if err != nil {
		return nil, errors.New("unexpected output format in % field")
	}
	return &QuotaInfo{
		Value:      value,
		Limit:      limnit,
		Percentage: percentage,
	}, nil
}

func get_user_quota(user_name string) (*QuotaInfo, error) {
	cmd := exec.Command("doveadm", "-f", "tab", "quota", "get", "-u", user_name)
	cmd.Stderr = os.Stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	out_str := string(out)
	return ParseQuotaInfo(out_str)
}

func main() {
	user := os.Args[1]
	quota_info, err := get_user_quota(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(quota_info)
}
