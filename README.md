# Sila Client

This project aims to create a native SDK for Golang for the Sila API.

[Sila Documentation](https://docs.silamoney.com)

## Implemented Endpoints

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
* [ ] [/get_entities](https://docs.silamoney.com/docs/get_entities)
* [ ] [/get_entity](https://docs.silamoney.com/docs/get_entity)

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