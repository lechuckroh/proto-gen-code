package main

import (
	"bytes"
	"fmt"
	protobuf2 "github.com/lechuckroh/protogencode/internal/pkg/protobuf"
	util2 "github.com/lechuckroh/protogencode/internal/pkg/util"
	"github.com/pkg/errors"
	"log"
	"os"
	"sort"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	FlagConst         = "const"
	FlagExclude       = "exclude"
	FlagInclude       = "include"
	FlagMsg           = "msg"
	FlagMsgVarPrefix  = "msg-var-prefix"
	FlagMsgVarPostfix = "msg-var-postfix"
	FlagProto         = "proto"
	FlagRename        = "rename"
)

const VERSION = "1.0.0"

var buildDateVersion string

func GenerateAction(c *cli.Context) error {
	protoFile := c.String(FlagProto)
	constFile := c.String(FlagConst)
	msgFile := c.String(FlagMsg)
	msgVarPostfix := c.String(FlagMsgVarPostfix)
	msgVarPrefix := c.String(FlagMsgVarPrefix)
	renameParams := c.StringSlice(FlagRename)
	excludes := c.StringSlice(FlagExclude)
	includes := c.StringSlice(FlagInclude)

	// *.proto 파일 로드
	p, err := protobuf2.LoadProtoFile(protoFile)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to read %s", protoFile))
	}
	proto := protobuf2.NewProto(p, protobuf2.ProtoOption{
		Includes: includes,
		Excludes: excludes,
		NameMap:  util2.KeyValueSlicesToMap(renameParams),
	})

	// context
	ctx := &GenContext{
		Proto:         proto,
		ConstFile:     constFile,
		MsgVarPostfix: msgVarPostfix,
		MsgVarPrefix:  msgVarPrefix,
	}

	// 상수 *.ts 파일 생성
	if constFile != "" {
		constBuf := new(bytes.Buffer)
		if err := generateConstants(constBuf, ctx); err != nil {
			return err
		}
		if err := util2.WriteStringToFile(constFile, constBuf.String()); err != nil {
			return err
		}
	}

	// 메시지 *.ts 파일 생성
	if msgFile != "" {
		msgBuf := new(bytes.Buffer)
		if err := generateMessages(msgBuf, ctx); err != nil {
			return err
		}
		if err := util2.WriteStringToFile(msgFile, msgBuf.String()); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	cliApp := cli.NewApp()
	cliApp.EnableBashCompletion = true
	cliApp.Name = "protogen-ts"
	cliApp.Version = VERSION + buildDateVersion
	cliApp.Compiled = time.Now()
	cliApp.Authors = []*cli.Author{
		{
			Name:  "Lechuck Roh",
			Email: "lechuckroh@gmail.com",
		},
	}
	cliApp.Copyright = "(c) 2022 Lechuck Roh"
	cliApp.Usage = "protogen-ts"
	cliApp.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     FlagProto,
			Aliases:  []string{"p"},
			Usage:    "source proto file",
			EnvVars:  []string{"PROTOGEN_TS_PROTO"},
			Required: true,
		},
		&cli.StringFlag{
			Name:    FlagConst,
			Aliases: []string{"c"},
			Usage:   "constants ts file",
			EnvVars: []string{"PROTOGEN_TS_CONST"},
		},
		&cli.StringFlag{
			Name:    FlagMsg,
			Aliases: []string{"m"},
			Usage:   "message ts file",
			EnvVars: []string{"PROTOGEN_TS_MSG"},
		},
		&cli.StringFlag{
			Name:    FlagMsgVarPostfix,
			Usage:   "message variable name postfix",
			EnvVars: []string{"PROTOGEN_TS_MSG_VAR_POSTFIX"},
		},
		&cli.StringFlag{
			Name:    FlagMsgVarPrefix,
			Usage:   "message variable name prefix",
			EnvVars: []string{"PROTOGEN_TS_MSG_VAR_PREFIX"},
		},
		&cli.StringSliceFlag{
			Name:    FlagRename,
			Aliases: []string{"r"},
			Usage:   `rename generated types. eg) "Error=ErrorType"`,
			EnvVars: []string{"PROTOGEN_TS_RENAME"},
		},
		&cli.StringSliceFlag{
			Name:    FlagInclude,
			Aliases: []string{"i"},
			Usage:   `type to include"`,
			EnvVars: []string{"PROTOGEN_TS_INCLUDE"},
		},
		&cli.StringSliceFlag{
			Name:    FlagExclude,
			Aliases: []string{"x"},
			Usage:   `type to exclude"`,
			EnvVars: []string{"PROTOGEN_TS_EXCLUDE"},
		},
	}
	cliApp.Action = GenerateAction
	sort.Sort(cli.FlagsByName(cliApp.Flags))
	sort.Sort(cli.CommandsByName(cliApp.Commands))

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
