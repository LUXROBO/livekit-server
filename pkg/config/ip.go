package config

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/pion/stun"
	"github.com/pkg/errors"
)

func (conf *Config) determineIP() (string, error) {
	if conf.RTC.UseExternalIP {
		stunServers := conf.RTC.StunServers
		if len(stunServers) == 0 {
			stunServers = DefaultStunServers
		}
		var err error
		for i := 0; i < 3; i++ {
			var ip string
			ip, err = GetExternalIP(stunServers)
			if err == nil {
				return ip, nil
			} else {
				time.Sleep(500 * time.Millisecond)
			}
		}
		return "", errors.Errorf("could not resolve external IP: %v", err)
	}

	// use local ip instead
	return GetLocalIPAddress()
}

func GetLocalIPAddress() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	// handle err
	var loopBack string
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP.To4()
			case *net.IPAddr:
				ip = v.IP.To4()
			default:
				continue
			}
			if ip == nil {
				continue
			}
			if ip.IsLoopback() {
				loopBack = ip.String()
			} else {
				return ip.String(), nil
			}
		}
	}

	if loopBack != "" {
		return loopBack, nil
	}
	return "", fmt.Errorf("could not find local IP address")
}

func GetExternalIP(stunServers []string) (string, error) {
	if len(stunServers) == 0 {
		return "", errors.New("STUN servers are required but not defined")
	}
	c, err := stun.Dial("udp4", stunServers[0])
	if err != nil {
		return "", err
	}
	defer c.Close()

	message, err := stun.Build(stun.TransactionID, stun.BindingRequest)
	if err != nil {
		return "", err
	}

	var stunErr error
	// sufficiently large buffer to not block it
	ipChan := make(chan string, 20)
	err = c.Start(message, func(res stun.Event) {
		if res.Error != nil {
			stunErr = res.Error
			return
		}

		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err != nil {
			stunErr = err
			return
		}
		ip := xorAddr.IP.To4()
		if ip != nil {
			ipChan <- ip.String()
		}
	})
	if err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	select {
	case nodeIP := <-ipChan:
		return nodeIP, nil
	case <-ctx.Done():
		msg := "could not determine public IP"
		if stunErr != nil {
			return "", errors.Wrap(stunErr, msg)
		} else {
			return "", fmt.Errorf(msg)
		}
	}
}
