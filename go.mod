module github.com/Fantom-foundation/go-lachesis

go 1.13

require (
	github.com/Fantom-foundation/go-opera v0.0.0-20210329030859-db15cedd849d
	github.com/Fantom-foundation/lachesis-base v0.0.0-20210323130105-a8e5ca7f15ac
	github.com/beorn7/perks v1.0.1
	github.com/cespare/cp v1.1.1
	github.com/davecgh/go-spew v1.1.1
	github.com/deckarep/golang-set v1.7.1
	github.com/docker/docker v1.13.1
	github.com/emirpasic/gods v1.12.0
	github.com/ethereum/go-ethereum v1.9.22
	github.com/evalphobia/logrus_sentry v0.8.2
	github.com/facebookgo/atomicfile v0.0.0-20151019160806-2de1f203e7d5 // indirect
	github.com/facebookgo/pidfile v0.0.0-20150612191647-f242e2999868
	github.com/fjl/memsize v0.0.0-20190710130421-bcb5799ab5e5
	github.com/golang/mock v1.4.3
	github.com/hashicorp/golang-lru v0.5.4
	github.com/mattn/go-colorable v0.1.4
	github.com/mattn/go-isatty v0.0.10
	github.com/naoina/toml v0.1.2-0.20170918210437-9fafd6967416
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.2.1
	github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4
	github.com/sirupsen/logrus v1.4.2
	github.com/status-im/keycard-go v0.0.0-20190424133014-d95853db0f48
	github.com/stretchr/testify v1.4.0
	github.com/syndtr/goleveldb v1.0.1-0.20200815110645-5c35d600f0ca
	github.com/tyler-smith/go-bip39 v1.0.2
	github.com/uber/jaeger-client-go v2.20.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gopkg.in/urfave/cli.v1 v1.22.1
)

replace gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.20.0

replace github.com/ethereum/go-ethereum => github.com/Fantom-foundation/go-ethereum v1.9.22-ftm-0.5
