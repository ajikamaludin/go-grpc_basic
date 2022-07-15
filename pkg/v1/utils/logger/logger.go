package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	rotatelog "github.com/lestrrat/go-file-rotatelogs"
)

const (
	PATH      = `logs`
	REST_LOG  = `rest`
	HIF_LOG   = `hif`
	ERROR_LOG = `error`

	DATEFORMAT    = `%Y-%m-%d-%H:%M:%S`
	ROTATION_TIME = 3600 // in second
)

func StoreRestRequest(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, userid string) error {
	// create logfile
	// filepath := fmt.Sprintf("%s/%s", PATH, REST_LOG)

	// create dir if not exist
	_, err := os.Stat(PATH)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(PATH, os.ModePerm)
		} else {
			return err
		}
	}

	// file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// defer file.Close()
	// if err != nil {
	// 	return err
	// }

	file, err := rotatelog.New(
		fmt.Sprintf("%s/%s_%s.log", PATH, REST_LOG, DATEFORMAT),
		rotatelog.WithLinkName(fmt.Sprintf("%s/%s.log", PATH, REST_LOG)),
		// rotatelog.WithMaxAge(time.Second*10),
		rotatelog.WithRotationTime(time.Second*ROTATION_TIME),
	)
	if err != nil {
		return err
	}

	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	var content []byte

	if len(headers["grpcgateway-content-type"]) > 0 {
		if headers["grpcgateway-content-type"][0] == "application/json" {
			content, err = json.MarshalIndent(req, "", " ")
			if err != nil {
				return err
			}
		} else {
			content = []byte(fmt.Sprintf("%v", req))
		}
	} else {
		content = []byte(fmt.Sprintf("%v", req))
	}

	// store data
	log.SetOutput(file)

	// flags to log timestamp and lineno
	log.SetFlags(log.LstdFlags)

	// write log
	log.Println(fmt.Sprintf("User-Id: %s.%s", userid, userid))

	log.SetFlags(0)

	log.Println("--------------Rest API Request----------------")
	log.Println(fmt.Sprintf("URI: %s", info.FullMethod))

	for k, v := range headers {
		log.Println(fmt.Sprintf("%s: %s", k, v))
	}

	log.Println(fmt.Sprintf("Content: %s", content))
	log.Println("-----------------------------------------------")
	log.Println()
	log.Println()

	return nil
}

func StoreRestResponse(ctx context.Context, w http.ResponseWriter, p proto.Message) error {
	var userId string

	// create logfile
	// filepath := fmt.Sprintf("%s/%s", PATH, REST_LOG)

	// create dir if not exist
	_, err := os.Stat(PATH)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(PATH, os.ModePerm)
		} else {
			return err
		}
	}

	// file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// defer file.Close()
	// if err != nil {
	// 	return err
	// }

	file, err := rotatelog.New(
		fmt.Sprintf("%s/%s_%s.log", PATH, REST_LOG, DATEFORMAT),
		rotatelog.WithLinkName(fmt.Sprintf("%s/%s.log", PATH, REST_LOG)),
		// rotatelog.WithMaxAge(time.Second*10),
		rotatelog.WithRotationTime(time.Second*ROTATION_TIME),
	)
	if err != nil {
		return err
	}

	md, ok := runtime.ServerMetadataFromContext(ctx)
	if !ok {
		return nil
	}

	userIds := md.HeaderMD.Get("User-Id")
	if len(userIds) > 0 {
		userId = userIds[0]
	}

	content, err := json.MarshalIndent(p, "", " ")
	if err != nil {
		return err
	}

	// store data
	log.SetOutput(file)

	// flags to log timestamp and lineno
	log.SetFlags(log.LstdFlags)

	// write log
	log.Println(fmt.Sprintf("User-Id: %s.%s", userId, userId))

	log.SetFlags(0)

	log.Println("--------------Rest API Response----------------")

	statusCode := "200"
	sc := w.Header()["Grpc-Metadata-X-Http-Code"]
	if len(sc) > 0 {
		statusCode = sc[0]
	}

	log.Println(fmt.Sprintf("Status-Code: %s", statusCode))

	for k, v := range w.Header() {
		log.Println(fmt.Sprintf("%s: %s", k, v))
	}

	log.Println(fmt.Sprintf("Content: %s", string(content)))
	log.Println("-----------------------------------------------")
	log.Println()
	log.Println()

	return nil
}

func StoreHifRequest(url, userid string, body []byte, headers map[string]interface{}) error {
	// create logfile
	// filepath := fmt.Sprintf("%s/%s", PATH, HIF_LOG)

	// create dir if not exist
	_, err := os.Stat(PATH)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(PATH, os.ModePerm)
		} else {
			return err
		}
	}

	// file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// defer file.Close()
	// if err != nil {
	// 	return err
	// }

	file, err := rotatelog.New(
		fmt.Sprintf("%s/%s_%s.log", PATH, HIF_LOG, DATEFORMAT),
		rotatelog.WithLinkName(fmt.Sprintf("%s/%s.log", PATH, REST_LOG)),
		// rotatelog.WithMaxAge(time.Second*10),
		rotatelog.WithRotationTime(time.Second*ROTATION_TIME),
	)
	if err != nil {
		return err
	}

	var content bytes.Buffer

	if body != nil {
		if len(headers) > 0 {
			if headers["Content-Type"] == "application/json" {
				err = json.Indent(&content, body, "", "    ")
				if err != nil {
					return err
				}
			} else {
				content = *bytes.NewBuffer(body)
			}
		}
	}

	// store data
	log.SetOutput(file)

	// flags to log timestamp and lineno
	log.SetFlags(log.LstdFlags)

	// write log
	log.Println(fmt.Sprintf("User-Id: %s.%s", userid, userid))

	log.SetFlags(0)

	log.Println("--------------Hif API Request----------------")
	log.Println(fmt.Sprintf("URI: %s", url))

	for k, v := range headers {
		log.Println(fmt.Sprintf("%s: %s", k, v))
	}

	log.Println(fmt.Sprintf("Content: %s", content.String()))
	log.Println("-----------------------------------------------")
	log.Println()
	log.Println()

	return nil
}

func StoreHifResponse(url, userid string, sc int, headers http.Header, res io.Reader) error {
	// create logfile
	// filepath := fmt.Sprintf("%s/%s", PATH, HIF_LOG)

	// create dir if not exist
	_, err := os.Stat(PATH)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(PATH, os.ModePerm)
		} else {
			return err
		}
	}

	// file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// defer file.Close()
	// if err != nil {
	// 	return err
	// }
	file, err := rotatelog.New(
		fmt.Sprintf("%s/%s_%s.log", PATH, HIF_LOG, DATEFORMAT),
		rotatelog.WithLinkName(fmt.Sprintf("%s/%s.log", PATH, HIF_LOG)),
		// rotatelog.WithMaxAge(time.Second*10),
		rotatelog.WithRotationTime(time.Second*ROTATION_TIME),
	)
	if err != nil {
		return err
	}

	var content bytes.Buffer

	if res != nil {
		// read body
		buf, err := ioutil.ReadAll(res)
		if err != nil {
			return err
		}

		if len(headers.Values("Content-Type")) > 0 {
			if headers.Values("Content-Type")[0] == "application/json; charset=utf-8" {
				err = json.Indent(&content, buf, "", "    ")
				if err != nil {
					return err
				}
			} else {
				content = *bytes.NewBuffer(buf)
			}
		}
	}

	// store data
	log.SetOutput(file)

	// flags to log timestamp and lineno
	log.SetFlags(log.LstdFlags)

	// write log
	log.Println(fmt.Sprintf("User-Id: %s.%s", userid, userid))

	log.SetFlags(0)

	log.Println("--------------Hif API Response----------------")
	log.Println(fmt.Sprintf("URI: %s", url))
	log.Println(fmt.Sprintf("Status-Code: %d", sc))

	for k, v := range headers {
		log.Println(fmt.Sprintf("%s: %s", k, v))
	}

	log.Println(fmt.Sprintf("Content: %s", content.String()))
	log.Println("-----------------------------------------------")
	log.Println()
	log.Println()

	return nil
}

func StoreError(userid string, buf []byte) error {
	// create logfile
	// filepath := fmt.Sprintf("%s/%s", PATH, ERROR_LOG)

	// create dir if not exist
	_, err := os.Stat(PATH)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(PATH, os.ModePerm)
		} else {
			return err
		}
	}

	// file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// defer file.Close()
	// if err != nil {
	// 	return err
	// }
	file, err := rotatelog.New(
		fmt.Sprintf("%s/%s_%s.log", PATH, ERROR_LOG, DATEFORMAT),
		rotatelog.WithLinkName(fmt.Sprintf("%s/%s.log", PATH, ERROR_LOG)),
		// rotatelog.WithMaxAge(time.Second*10),
		rotatelog.WithRotationTime(time.Second*ROTATION_TIME),
	)
	if err != nil {
		return err
	}
	var content bytes.Buffer

	err = json.Indent(&content, buf, "", "    ")
	if err != nil {
		return err
	}

	// store data
	log.SetOutput(file)

	// flags to log timestamp and lineno
	log.SetFlags(log.LstdFlags)

	// write log
	log.Println(fmt.Sprintf("User-Id: %s.%s", userid, userid))

	log.SetFlags(0)

	log.Println("--------------Error----------------")
	log.Println(fmt.Sprintf("Content: %s", content.String()))
	log.Println("-----------------------------------------------")
	log.Println()
	log.Println()

	return nil
}
