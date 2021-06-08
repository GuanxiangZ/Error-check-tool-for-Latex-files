package config

import (
	"github.com/sony/sonyflake"
	"sync"
	"time"
)

var (
	fileconfig *FileConfig
	serverconfig *ServerConfig
	fileonce sync.Once
	serveronce sync.Once
)

type FileConfig struct {
	Input    string
	Output   string
	Outdir   string
	Language string
	Addr     string
}


func GetFileCmdConfig() *FileConfig {
	fileonce.Do(func() {
		fileconfig = &FileConfig{}
	})

	return fileconfig
}

type ServerConfig struct {
	ParserUrl string
	Port uint32
	LogDir string
	flake *sonyflake.Sonyflake
}

func GetServeCmdConfig() *ServerConfig {
	serveronce.Do(func() {
		serverconfig = &ServerConfig{}
		serverconfig.flake = sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: time.Now(),
		})
	})
	return serverconfig
}

func(sc *ServerConfig) NextID() uint64 {
	id, err := sc.flake.NextID()
	if err != nil {
		sc.flake = sonyflake.NewSonyflake(sonyflake.Settings{
			StartTime: time.Now(),
		})
		id = sc.NextID()
	}
	return id
}
