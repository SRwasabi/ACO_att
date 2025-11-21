package config

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type IntList []int

func (l *IntList) UnmarshalJSON(data []byte) error {
	var single int
	if err := json.Unmarshal(data, &single); err == nil {
		*l = []int{single}
		return nil
	}

	var multiple []int
	if err := json.Unmarshal(data, &multiple); err == nil {
		*l = multiple
		return nil
	}
	return fmt.Errorf("esperava int ou [int]")
}

func (l *IntList) Set(s string) error {
	parts := strings.Split(s, ",")
	var res []int
	for _, p := range parts {
		val, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return err
		}
		res = append(res, val)
	}
	*l = res
	return nil
}

func (l *IntList) String() string {
	return fmt.Sprintf("%v", *l)
}


type FloatList []float64

func (l *FloatList) UnmarshalJSON(data []byte) error {
	var single float64
	if err := json.Unmarshal(data, &single); err == nil {
		*l = []float64{single}
		return nil
	}
	var multiple []float64
	if err := json.Unmarshal(data, &multiple); err == nil {
		*l = multiple
		return nil
	}
	return fmt.Errorf("esperava float ou [float]")
}

func (l *FloatList) Set(s string) error {
	parts := strings.Split(s, ",")
	var res []float64
	for _, p := range parts {
		val, err := strconv.ParseFloat(strings.TrimSpace(p), 64)
		if err != nil {
			return err
		}
		res = append(res, val)
	}
	*l = res
	return nil
}

func (l *FloatList) String() string {
	return fmt.Sprintf("%v", *l)
}


type StringList []string

func (l *StringList) UnmarshalJSON(data []byte) error {
	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		*l = []string{single}
		return nil
	}
	var multiple []string
	if err := json.Unmarshal(data, &multiple); err == nil {
		*l = multiple
		return nil
	}
	return fmt.Errorf("esperava string ou [string]")
}

func (l *StringList) Set(s string) error {
	*l = strings.Split(s, ",")
	return nil
}

func (l *StringList) String() string {
	return fmt.Sprintf("%v", *l)
}


type BoolList []bool

func (l *BoolList) UnmarshalJSON(data []byte) error {
	var single bool
	if err := json.Unmarshal(data, &single); err == nil {
		*l = []bool{single}
		return nil
	}
	var multiple []bool
	if err := json.Unmarshal(data, &multiple); err == nil {
		*l = multiple
		return nil
	}
	return fmt.Errorf("esperava bool ou [bool]")
}

func (l *BoolList) Set(s string) error {
	parts := strings.Split(s, ",")
	var res []bool
	for _, p := range parts {
		val, err := strconv.ParseBool(strings.TrimSpace(p))
		if err != nil {
			return err
		}
		res = append(res, val)
	}
	*l = res
	return nil
}

func (l *BoolList) String() string {
	return fmt.Sprintf("%v", *l)
}