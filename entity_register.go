package sila

type Register struct {
	Header      *Header              `json:"header"`
	Message     string               `json:"message"`
	Address     RegistrationAddress  `json:"address,omitempty"`
	Identity    RegistrationIdentity `json:"identity,omitempty"`
	Contact     Contact              `json:"contact,omitempty"`
	CryptoEntry CryptoEntry          `json:"crypto_entry"`
	Entity      RegistrationEntity   `json:"entity"`
}

func (msg *Register) SetRef(ref string) *Register {
	msg.Header.setRef(ref)
	return msg
}

type RegistrationAddress struct {
	AddressAlias   string `json:"address_alias,omitempty"`
	StreetAddress1 string `json:"street_address_1,omitempty"`
	StreetAddress2 string `json:"street_address_2,omitempty"`
	City           string `json:"city,omitempty"`
	State          string `json:"state,omitempty"`
	Country        string `json:"country,omitempty"`
	PostalCode     string `json:"postal_code,omitempty"`
}

func (msg *Register) SetAddress(address RegistrationAddress) *Register {
	msg.Address = address
	return msg
}

type IdentityType string

const (
	Ssn IdentityType = "SSN"
	Ein              = "EIN"
)

type RegistrationIdentity struct {
	IdentityAlias IdentityType `json:"identity_alias"`
	IdentityValue string       `json:"identity_value"`
}

func (msg *Register) SetIdentity(identityType IdentityType, identityValue string) *Register {
	msg.Identity.IdentityAlias = identityType
	msg.Identity.IdentityValue = identityValue
	return msg
}

type Contact struct {
	Phone        string `json:"phone,omitempty"`
	ContactAlias string `json:"contact_alias"`
	Email        string `json:"email,omitempty"`
}

func (msg *Register) SetContact(contactAlias string, phone string, email string) *Register {
	msg.Contact.ContactAlias = contactAlias
	msg.Contact.Phone = phone
	msg.Contact.Email = email
	return msg
}

type CryptoEntry struct {
	CryptoAlias   string `json:"crypto_alias"`
	CryptoAddress string `json:"crypto_address"`
	CryptoCode    string `json:"crypto_code"`
}

func (msg *Register) SetCrypto(nickname string, address string) *Register {
	msg.CryptoEntry.CryptoAlias = nickname
	msg.CryptoEntry.CryptoAddress = address
	msg.CryptoEntry.CryptoCode = "ETH"
	return msg
}

type EntityType string

const (
	Individual EntityType = "individual"
	Business              = "business"
)

type RegistrationEntity struct {
	Type            EntityType `json:"type"`
	BirthDate       string     `json:"birthdate,omitempty"`
	FirstName       string     `json:"first_name,omitempty"`
	LastName        string     `json:"last_name,omitempty"`
	EntityName      string     `json:"entity_name,omitempty"`
	BusinessType    string     `json:"business_type,omitempty"`
	BusinessWebsite string     `json:"business_website,omitempty"`
	DoingBusinessAs string     `json:"doing_business_as,omitempty"`
	NaicsCode       int        `json:"naics_code,omitempty"`
}

func (msg *Register) SetIndividualEntity(firstName string, lastName string, birthDate string) *Register {
	msg.Entity.Type = Individual
	msg.Entity.FirstName = firstName
	msg.Entity.LastName = lastName
	msg.Entity.BirthDate = birthDate
	return msg
}

func (msg *Register) SetBusinessEntity(entityName string, businessType string, naicsCode int) *Register {
	msg.Entity.Type = Business
	msg.Entity.EntityName = entityName
	msg.Entity.BusinessType = businessType
	msg.Entity.NaicsCode = naicsCode
	return msg
}

func (msg *Register) SetBusinessWebsite(businessWebsite string) {
	msg.Entity.BusinessWebsite = businessWebsite
}

func (msg *Register) SetDoingBusinessAs(dba string) {
	msg.Entity.DoingBusinessAs = dba
}

func (msg *Register) Do() (SuccessResponse, error) {
	var responseBody SuccessResponse
	err := instance.performCall("/register", msg, &responseBody)
	return responseBody, err
}
