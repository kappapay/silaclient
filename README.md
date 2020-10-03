# Sila Client

This project creates a native SDK for the Sila API in Golang. To learn more about Sila and how it works, please go to
their [website](https://silamoney.com) as well as read the [documentation](https://docs.silamoney.com).

## Implemented Endpoints

The current version of this library implemented the following endpoints.

### Entities

* [x] [/check_handle](https://docs.silamoney.com/docs/check_handle)
* [x] [/register](https://docs.silamoney.com/docs/register)
* [ ] [/add/\<registration-data\>](https://docs.silamoney.com/docs/addregistration-data)
* [ ] [/update/\<registration-data\>](https://docs.silamoney.com/docs/updateregistration-data)
* [ ] [/delete/\<registration-data\>](https://docs.silamoney.com/docs/deleteregistration-data)
* [ ] [/link_business_member](https://docs.silamoney.com/docs/link_business_member)
* [ ] [/unlink_business_member](https://docs.silamoney.com/docs/unlink_business_member)
* [x] [/request_kyc](https://docs.silamoney.com/docs/request_kyc)
* [x] [/check_kyc](https://docs.silamoney.com/docs/check_kyc)
* [ ] [/certify_beneficial_owner](https://docs.silamoney.com/docs/certify_beneficial_owner)
* [ ] [/certify_business](https://docs.silamoney.com/docs/certify_business)
* [x] [/get_entities](https://docs.silamoney.com/docs/get_entities)
* [x] [/get_entity](https://docs.silamoney.com/docs/get_entity)

### Accounts

* [x] [/link_account](https://docs.silamoney.com/docs/link_account)
* [x] [/plaid_sameday_auth](https://docs.silamoney.com/docs/plaid_sameday_auth)
* [x] [/get_accounts](https://docs.silamoney.com/docs/get_accounts)
* [x] [/get_account_balance](https://docs.silamoney.com/docs/get_account_balance)

### Wallets

* [x] [/register_wallet](https://docs.silamoney.com/docs/register_wallet)
* [x] [/get_wallet](https://docs.silamoney.com/docs/get_wallet)
* [x] [/get_wallets](https://docs.silamoney.com/docs/get_wallets)
* [x] [/update_wallet](https://docs.silamoney.com/docs/update_wallet)
* [x] [/get_sila_balance](https://docs.silamoney.com/docs/get_sila_balance)
* [x] [/delete_wallet](https://docs.silamoney.com/docs/delete_wallet)

### Transactions

* [x] [/issue_sila](https://docs.silamoney.com/docs/issue_sila)
* [x] [/transfer_sila](https://docs.silamoney.com/docs/transfer_sila)
* [x] [/redeem_sila](https://docs.silamoney.com/docs/redeem_sila)
* [x] [/get_transactions](https://docs.silamoney.com/docs/get_transactions)
* [x] [/cancel_transactions](https://docs.silamoney.com/docs/cancel_transaction)

### Parameters

* [x] [/get_business_types](https://docs.silamoney.com/docs/get_business_types)
* [x] [/get_business_roles](https://docs.silamoney.com/docs/get_business_roles)
* [x] [/get_naics_categories](https://docs.silamoney.com/docs/get_naics_categories)

## Usage

To use the Sila SDK, first import:

```go
import "github.com/bpancost/sila"
```

From there, create a new client by using
```go
// The the auth private key as a hex string, without the "0x" prefixed
privateKeyHex := "badba7368134dcd61c60f9b56979c09196d03f5891a20c1557b1afac0202a97c"
// The handle associated with your auth account
authHandle := "handle.silamoney.eth"
// A production client can be created by using sila.Production instead of sila.Sandbox
client, err := sila.NewClient(privateKeyHex, authHandle, sila.Sandbox)
```

With the client, it is possible to make any of the included calls. Each is a chainable series of method calls, ending
with a call to the `Do` method. For example, to get a list of entities you could use
```go
response, err := silaClient.GetEntities().
                            SetPage(1).
                            SetPerPage(20).
                            Do()
```

There are several functions within the `github.com/bpancost/sila` package which can be used to manipulate or create keys
for wallets, which will be required for certain calls.

* `GenerateNewPrivateKey() (string, error)`
    * Generates a private key for a new wallet. This should generally not be shared or shown outside of your system and
    will be used for subsequent calls related to the wallet, or as a means of identifying the user who owns the wallet.
* `GetWalletAddress(privateKeyHex string) (string, error)`
    * Gets the public address for the wallet. This is used to publicly indicate a wallet and is always visible on the
    blockchain, including any transactions.
* `GenerateWalletSignature(message []byte, walletPrivateKeyHex string) (string, error)`
    * Signs an arbitrary message using a wallet's private key as a hex string.

## Integration Tests

To use the included integration tests, create a file named `test_config.yaml` using the `test_config_sample.yaml` as a
template. Provide your auth key as a hex string, your auth handle, a unique user handle specifically for this testing,
and a private key as a hex string for the integration user's wallet which will be their main wallet. You can use a tool
like [Vanity-Eth](https://vanity-eth.tk/) to help generate a new wallet address and private key.

Most tests will complete quickly, though the transaction tests require several minute sleeps to verify they completed 
and will take around 10 minutes to complete.

If you would like to view the integration test progress via a web page, you can use `goconvey` from a terminal and then
navigate to [http://127.0.0.1:8080](http://127.0.0.1:8080).