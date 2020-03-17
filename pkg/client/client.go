package client

const (
	// Dev is environment type for development.
	Dev = "dev"
	// Prod is environment type for production.
	Prod = "prod"
	// DevDomainURL is domain URL for development environment.
	DevDomainURL = "https://oropay-dev.e-mas.com"
	// ProdDomainURL is domain URL for production environment.
	ProdDomainURL = "https://api.e-mas.com"
	// DefaultEnvironment is default environment type value.
	DefaultEnvironment = "dev"
	// DefaultDebug is default debug boolean value.
	DefaultDebug = true
)

// Client is e-mas SSO client.
type Client struct {
	// ClientName is client name for using e-mas SSO.
	// Required when requesting client id and secret.
	ClientName string

	// ClientId and ClientSecret is client's credential data
	// that will be used for authorization. Need to be requested
	// for using client for the first time.
	ClientId     string
	ClientSecret string

	// Token is access token granted after login. Token will
	// be used in query params for most of e-mas sso endpoints.
	Token string

	// DomainURL is e-mas API domain URL depends on the environment.
	DomainURL string
	// Environment is environment type of the client.
	// Will affect the e-mas domain URL when initializing client.
	Environment string
	// Debug is boolean for debugging.
	Debug bool
}

// DefaultClient is default client config.
var DefaultClient = Client{
	Environment: DefaultEnvironment,
	Debug:       DefaultDebug,
}

// New to create a new e-mas SSO client.
func New(clientName string, clientCredentials ...string) *Client {
	// Use default client as base client.
	c := DefaultClient
	c.ClientName = clientName

	// Assign client credentials if provided.
	if len(clientCredentials) == 2 {
		c.ClientId = clientCredentials[0]
		c.ClientSecret = clientCredentials[1]
	}

	c.init()
	return &c
}

// init to initiate field values in client.
func (c *Client) init() {
	c.DomainURL = DevDomainURL
	if c.Environment == Prod {
		c.DomainURL = ProdDomainURL
	}
}

// SetClientName to set client name.
func (c *Client) SetClientName(clientName string) *Client {
	c.ClientName = clientName
	return c
}

// SetClientCredential to set client id and secret.
func (c *Client) SetClientCredential(clientId, clientSecret string) *Client {
	c.ClientId = clientId
	c.ClientSecret = clientSecret
	return c
}

// SetEnv to set client environment. This will set domain
// URL automatically depends on the environment. Don't
// call this function after SetDomain() or it will overwrite
// the domain URL.
func (c *Client) SetEnv(env string) *Client {
	// Handle invalid environment type.
	if env != Dev && env != Prod {
		env = Dev
	}
	c.Environment = env
	c.init()
	return c
}

// SetDomain to set e-mas API domain URL.
// Can be used if e-mas API is in localhost.
func (c *Client) SetDomain(domainURL string) *Client {
	c.DomainURL = domainURL
	return c
}

// SetDebug to set client debug boolean value.
func (c *Client) SetDebug(debug bool) *Client {
	c.Debug = debug
	return c
}

// SetToken to set client access token.
func (c *Client) SetToken(token string) *Client {
	c.Token = token
	return c
}
