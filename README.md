# cooladdress

cooladdress is for generate address with special suffix

### Build cooladdress from code

Download the code to your desired local location (doesn't have to be under `$GOPATH/src`)
```
git clone git@github.com:iotexproject/cooladdress.git
cd cooladdress
```

If you put the project code under your `$GOPATH\src`, you will need to set up an environment variable
```
export GO111MODULE=on
set GO111MODULE=on (for windows)
```

Build the project by

```
make
```

Or

```
make all 
```

If the dependency needs to be updated, run

```
go get -u
go mod tidy
```

### Use CLI

```
./bin/addrgen gen [suffix] [timeout]

For example:
./bin/addrgen gen xy 1h

timeout such as "300ms" or "2h45m". Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"

three-letter:takes about few seconds
four-letter:takes about 150s
five-letter:takes more than 1h
```

The last 6 letters of the address is checksum,according to https://github.com/bitcoin/bips/blob/master/bip-0173.mediawiki:
The last checksum excluding "1", "b", "i", and "o"

## Contact

- Mailing list: [iotex-dev](iotex-dev@iotex.io)
- Dev Forum: [forum](https://community.iotex.io/c/research-development/protocol)
- Bugs: [issues](https://github.com/iotexproject/cooladdress/issues)

## License
This project is licensed under the [Apache License 2.0](LICENSE).
