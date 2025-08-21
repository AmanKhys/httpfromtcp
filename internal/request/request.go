package request

import (
	"errors"
	"fmt"
	"io"
	rgx "regexp"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	rl, err := parseRequestLine(reader)
	if err != nil {
		return nil, err
	}
	return &Request{
		RequestLine: *rl,
	}, nil
}

func parseRequestLine(reader io.Reader) (*RequestLine, error) {

	msgbytes, err := io.ReadAll(reader)
	msg := string(msgbytes)
	if err != nil {
		return nil, err
	}

	fields := strings.Split(msg, fmt.Sprintf("\r\n"))
	requestLineFields := strings.Split(fields[0], " ")

	err = checkRequestLineFields(requestLineFields)
	if err != nil {
		return nil, err
	}

	rLine := RequestLine{
		HttpVersion:   requestLineFields[2],
		RequestTarget: requestLineFields[1],
		Method:        requestLineFields[0],
	}

	return &rLine, nil
}

func checkRequestLineFields(rl []string) error {
	var errs []error
	ok := checkRLHttpVersion(rl[0])
	if !ok {
		errs = append(errs, errors.New("http version not parsable"))
	}
	ok = checkRLRequestTarget(rl[1])
	if !ok {
		errs = append(errs, errors.New("request target not parsable"))
	}
	ok = checkRLMethod(rl[2])
	if !ok {
		errs = append(errs, errors.New("request method not parsable"))
	}

	return errors.Join(errs...)
}

func checkRLMethod(msg string) bool {
	switch msg {
	case "GET":
		return true
	case "HEAD":
		return true
	case "POST":
		return true
	case "PUT":
		return true
	case "DELETE":
		return true
	case "CONNECT":
		return true
	case "OPTION":
		return true
	case "TRACE":
		return true
	case "PATCH":
		return true
	default:
		return false
	}
}

func checkRLHttpVersion(msg string) bool {
	httpVersionCheck := rgx.MustCompile(`^HTTP/(1\.[01]|2\.0|3)$`)
	return httpVersionCheck.MatchString(msg)
}

func checkRLRequestTarget(msg string) bool {
	requestTargetCheck := rgx.MustCompile(`^/\S*$`)
	return requestTargetCheck.MatchString(msg)
}
