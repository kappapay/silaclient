module github.com/kappapay/silaclient/examples/register_individual

go 1.19

replace github.com/kappapay/silaclient => ../..

require (
	github.com/ardanlabs/conf/v3 v3.1.3
	github.com/joho/godotenv v1.4.0
	github.com/kappapay/silaclient v0.2.5
)

require (
	github.com/btcsuite/btcd v0.0.0-20171128150713-2e60448ffcc6 // indirect
	github.com/ethereum/go-ethereum v1.9.21 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/sys v0.0.0-20200824131525-c12d262b63d8 // indirect
)
