// Copyright (c) 2019 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package cooladdress

import (
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/iotexproject/go-pkgs/crypto"
	"github.com/iotexproject/iotex-address/address"
)

var (
	// ErrTimeOut indicates the address cannot found
	ErrTimeOut = errors.New("time out")
)

// Gen gen address with suffix in limit time
func Gen(args []string) (string, error) {
	num := runtime.NumCPU()
	suffix := args[0]
	limit := args[1]
	dur, err := time.ParseDuration(limit)
	if err != nil {
		return "", err
	}
	ticker := time.NewTicker(dur)
	defer ticker.Stop()
	ch := make(chan string, num)
	for i := 0; i < num; i++ {
		go gen(suffix, ch)
	}
	for {
		select {
		case ret := <-ch:
			return ret, nil
		case <-ticker.C:
			return "", ErrTimeOut
		}
	}
}

func gen(suffix string, ch chan string) {
	for {
		private, err := crypto.GenerateKey()
		if err != nil {
			log.Println("failed to create key pair", err)
		}
		addr, err := address.FromBytes(private.PublicKey().Hash())
		if err != nil {
			return
		}
		priKeyBytes := private.Bytes()
		pubKeyBytes := private.PublicKey().Bytes()
		ret := fmt.Sprintf(
			"{\"PublicKey\": \"%x\", \"PrivateKey\": \"%x\", \"Address\": \"%s\"}",
			pubKeyBytes,
			priKeyBytes, addr.String(),
		)
		if strings.EqualFold(suffix, "") || strings.HasSuffix(addr.String(), suffix) {
			ch <- ret
			break
		}
	}
}
