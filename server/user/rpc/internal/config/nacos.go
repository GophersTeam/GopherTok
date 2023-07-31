package config

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/zeromicro/go-zero/core/conf"
	"log"
	"sync"
)

var (
	configClient config_client.IConfigClient
	once         sync.Once
)

type NacosConf struct {
	Host        string
	Port        uint64
	NamespaceId string
	Group       string
	DataId      string
}

func (nc *NacosConf) LoadConfig(c *Config) {
	clientConfig := constant.ClientConfig{
		NamespaceId:         nc.NamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "../../../tmp/nacos/log",
		CacheDir:            "../../../tmp/nacos/cache",
	}
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: nc.Host,
			Port:   nc.Port,
		},
	}

	once.Do(func() {
		var err error
		configClient, err = clients.NewConfigClient(
			vo.NacosClientParam{
				ClientConfig:  &clientConfig,
				ServerConfigs: serverConfigs,
			},
		)
		if err != nil {
			log.Fatalf("Create nacos config client failed: %v", err)
		}
		content, err := configClient.GetConfig(vo.ConfigParam{
			DataId: nc.DataId,
			Group:  nc.Group,
		})
		if err != nil {
			log.Fatalf("Get config from nacos failed: %v", err)
		}
		err = conf.LoadFromYamlBytes([]byte(content), c)
		if err != nil {
			log.Fatalf("Load config from nacos failed: %v", err)
		}
	})
}

func (nc *NacosConf) ListenConfig(f func(namespace, group, dataId, data string)) {
	if configClient == nil {
		log.Fatalf("Please load config first")
	}
	err := configClient.ListenConfig(vo.ConfigParam{
		DataId:   nc.DataId,
		Group:    nc.Group,
		OnChange: f,
	})
	if err != nil {
		log.Fatalf("Listen config from nacos failed: %v", err)
	}
}
