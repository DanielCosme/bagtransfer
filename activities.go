package app

import (
	"context"
	"errors"
	"fmt"
)

func GreetActivity(ctx context.Context, data Data) (string, error) {
	var res string
	switch data.Name {
	case "err":
		return "", errors.New("unidentified error")
	default:
		res = fmt.Sprintf("Hello, %s", data.Name)
		fmt.Println(res)
	}
	return res, nil
}

func GreetActivity2(ctx context.Context, data Data) (string, error) {
	var res string
	switch data.Name {
	case "err":
		return "", errors.New("unidentified error")
	default:
		res = fmt.Sprintf("Aloha, %s", data.Name)
		fmt.Println(res)
	}
	return res, nil
}
