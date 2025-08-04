# atomic-go

The official Atomic Go client library for interacting with the Atomic API.

## Installation

```bash
go get github.com/libatomic/atomic-go
```

## Quick Start

```go
package main

import (
    "context"
    "log"
    
    "github.com/libatomic/atomic-go"
)

func main() {
    // Create a client with your access token
    client := atomic.New(
        atomic.WithHost("api.atomic.com"),
        atomic.WithToken("your-access-token"),
    )
    
    // Get a user
    user, err := client.UserGet(context.Background(), &atomic.UserGetInput{
        UserID: atomic.String("user-id"),
    })
    if err != nil {
        log.Fatal(err)
    }
    
    log.Printf("User: %s", user.Email)
}
```

## Authentication

The Atomic API uses OAuth2 for authentication. You can configure authentication in several ways:

### Access Token Authentication

```go
client := atomic.New(
    atomic.WithHost("api.atomic.com"),
    atomic.WithToken("your-access-token"),
)
```

### Client Credentials Flow

```go
client := atomic.New(
    atomic.WithHost("api.atomic.com"),
    atomic.WithClientCredentials("client-id", "client-secret", "instance:read", "instance:write"),
)
```

### Custom HTTP Client

```go
customClient := &http.Client{
    Timeout: 30 * time.Second,
}

client := atomic.New(
    atomic.WithHost("api.atomic.com"),
    atomic.WithToken("your-access-token"),
    atomic.WithHTTPClient(customClient),
)
```

## API Endpoints

The atomic-go library provides access to all major Atomic API endpoints:

### Users

Manage user accounts and profiles.

```go
// Get a user
user, err := client.UserGet(ctx, &atomic.UserGetInput{
    UserID: atomic.String("user-id"),
})

// Create a user
user, err := client.UserCreate(ctx, &atomic.UserCreateInput{
    Email:    atomic.String("user@example.com"),
    Password: atomic.String("secure-password"),
    Name:     atomic.String("John Doe"),
})

// Update a user
user, err := client.UserUpdate(ctx, &atomic.UserUpdateInput{
    UserID: atomic.String("user-id"),
    Name:   atomic.String("Jane Doe"),
})

// Delete a user
err := client.UserDelete(ctx, &atomic.UserDeleteInput{
    UserID: atomic.String("user-id"),
})

// List users
users, err := client.UserList(ctx, &atomic.UserListInput{
    Limit: atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Applications

Manage OAuth applications.

```go
// Create an application
app, err := client.ApplicationCreate(ctx, &atomic.ApplicationCreateInput{
    Name:        atomic.String("My App"),
    Description: atomic.String("My application description"),
    RedirectURI: atomic.String("https://myapp.com/callback"),
})

// Get an application
app, err := client.ApplicationGet(ctx, &atomic.ApplicationGetInput{
    ApplicationID: atomic.String("app-id"),
})

// Update an application
app, err := client.ApplicationUpdate(ctx, &atomic.ApplicationUpdateInput{
    ApplicationID: atomic.String("app-id"),
    Name:          atomic.String("Updated App Name"),
})

// Delete an application
err := client.ApplicationDelete(ctx, &atomic.ApplicationDeleteInput{
    ApplicationID: atomic.String("app-id"),
})

// List applications
apps, err := client.ApplicationList(ctx, &atomic.ApplicationListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Access Tokens

Manage OAuth access tokens.

```go
// Create an access token for a user
token, err := client.AccessTokenCreate(ctx, &atomic.AccessTokenCreateInput{
    UserID: atomic.String("user-id"),
    Scopes: []string{"read", "write"},
})

// Create an access token for an application
token, err := client.AccessTokenCreate(ctx, &atomic.AccessTokenCreateInput{
    ApplicationID: atomic.String("app-id"),
    Scopes:        []string{"read", "write"},
})

// Get an access token
token, err := client.AccessTokenGet(ctx, &atomic.AccessTokenGetInput{
    AccessTokenID: atomic.String("token-id"),
})

// Revoke an access token
err := client.AccessTokenRevoke(ctx, &atomic.AccessTokenRevokeInput{
    AccessTokenID: atomic.String("token-id"),
})

// Delete an access token
err := client.AccessTokenDelete(ctx, &atomic.AccessTokenDeleteInput{
    AccessTokenID: atomic.String("token-id"),
})
```

### Articles

Manage content articles.

```go
// Get an article
article, err := client.ArticleGet(ctx, &atomic.ArticleGetInput{
    ArticleID: atomic.String("article-id"),
})

// Create an article
article, err := client.ArticleCreate(ctx, &atomic.ArticleCreateInput{
    Title:   atomic.String("My Article"),
    Content: atomic.String("Article content..."),
    Status:  atomic.String("draft"),
})

// Update an article
article, err := client.ArticleUpdate(ctx, &atomic.ArticleUpdateInput{
    ArticleID: atomic.String("article-id"),
    Title:     atomic.String("Updated Title"),
    Status:    atomic.String("published"),
})

// Delete an article
err := client.ArticleDelete(ctx, &atomic.ArticleDeleteInput{
    ArticleID: atomic.String("article-id"),
})

// List articles
articles, err := client.ArticleList(ctx, &atomic.ArticleListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Assets

Manage media assets.

```go
// Get an asset
asset, err := client.AssetGet(ctx, &atomic.AssetGetInput{
    AssetID: atomic.String("asset-id"),
})

// Create an asset
asset, err := client.AssetCreate(ctx, &atomic.AssetCreateInput{
    Name: atomic.String("image.jpg"),
    Type: atomic.String("image"),
    URL:  atomic.String("https://example.com/image.jpg"),
})

// Update an asset
asset, err := client.AssetUpdate(ctx, &atomic.AssetUpdateInput{
    AssetID: atomic.String("asset-id"),
    Name:    atomic.String("updated-name.jpg"),
})

// Delete an asset
err := client.AssetDelete(ctx, &atomic.AssetDeleteInput{
    AssetID: atomic.String("asset-id"),
})

// List assets
assets, err := client.AssetList(ctx, &atomic.AssetListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Templates

Manage email and notification templates.

```go
// Get a template
template, err := client.TemplateGet(ctx, &atomic.TemplateGetInput{
    TemplateID: atomic.String("template-id"),
})

// Create a template
template, err := client.TemplateCreate(ctx, &atomic.TemplateCreateInput{
    Name:    atomic.String("Welcome Email"),
    Subject: atomic.String("Welcome to our platform"),
    Body:    atomic.String("Hello {{name}}, welcome!"),
    Type:    atomic.String("email"),
})

// Update a template
template, err := client.TemplateUpdate(ctx, &atomic.TemplateUpdateInput{
    TemplateID: atomic.String("template-id"),
    Subject:    atomic.String("Updated Subject"),
})

// Delete a template
err := client.TemplateDelete(ctx, &atomic.TemplateDeleteInput{
    TemplateID: atomic.String("template-id"),
})

// List templates
templates, err := client.TemplateList(ctx, &atomic.TemplateListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Subscriptions

Manage user subscriptions and billing.

```go
// Get a subscription
subscription, err := client.SubscriptionGet(ctx, &atomic.SubscriptionGetInput{
    SubscriptionID: atomic.String("subscription-id"),
})

// Create a subscription
subscription, err := client.SubscriptionCreate(ctx, &atomic.SubscriptionCreateInput{
    UserID:  atomic.String("user-id"),
    PlanID:  atomic.String("plan-id"),
    Status:  atomic.String("active"),
})

// Update a subscription
subscription, err := client.SubscriptionUpdate(ctx, &atomic.SubscriptionUpdateInput{
    SubscriptionID: atomic.String("subscription-id"),
    Status:         atomic.String("cancelled"),
})

// Delete a subscription
err := client.SubscriptionDelete(ctx, &atomic.SubscriptionDeleteInput{
    SubscriptionID: atomic.String("subscription-id"),
})

// List subscriptions
subscriptions, err := client.SubscriptionList(ctx, &atomic.SubscriptionListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Plans

Manage subscription plans and pricing.

```go
// Get a plan
plan, err := client.PlanGet(ctx, &atomic.PlanGetInput{
    PlanID: atomic.String("plan-id"),
})

// Create a plan
plan, err := client.PlanCreate(ctx, &atomic.PlanCreateInput{
    Name:        atomic.String("Pro Plan"),
    Description: atomic.String("Professional features"),
    Price:       atomic.Float64(29.99),
    Interval:    atomic.String("monthly"),
})

// Update a plan
plan, err := client.PlanUpdate(ctx, &atomic.PlanUpdateInput{
    PlanID:      atomic.String("plan-id"),
    Name:        atomic.String("Updated Plan Name"),
    Price:       atomic.Float64(39.99),
})

// Delete a plan
err := client.PlanDelete(ctx, &atomic.PlanDeleteInput{
    PlanID: atomic.String("plan-id"),
})

// List plans
plans, err := client.PlanList(ctx, &atomic.PlanListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Prices

Manage pricing tiers and billing.

```go
// Get a price
price, err := client.PriceGet(ctx, &atomic.PriceGetInput{
    PriceID: atomic.String("price-id"),
})

// Create a price
price, err := client.PriceCreate(ctx, &atomic.PriceCreateInput{
    PlanID:  atomic.String("plan-id"),
    Amount:  atomic.Float64(2999), // $29.99 in cents
    Currency: atomic.String("usd"),
    Interval: atomic.String("monthly"),
})

// Update a price
price, err := client.PriceUpdate(ctx, &atomic.PriceUpdateInput{
    PriceID: atomic.String("price-id"),
    Amount:  atomic.Float64(3999), // $39.99 in cents
})

// Delete a price
err := client.PriceDelete(ctx, &atomic.PriceDeleteInput{
    PriceID: atomic.String("price-id"),
})

// List prices
prices, err := client.PriceList(ctx, &atomic.PriceListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Audiences

Manage user audiences and segments.

```go
// Get an audience
audience, err := client.AudienceGet(ctx, &atomic.AudienceGetInput{
    AudienceID: atomic.String("audience-id"),
})

// Create an audience
audience, err := client.AudienceCreate(ctx, &atomic.AudienceCreateInput{
    Name:        atomic.String("Premium Users"),
    Description: atomic.String("Users with premium subscriptions"),
    Filter:      atomic.String("subscription.status = 'active'"),
})

// Update an audience
audience, err := client.AudienceUpdate(ctx, &atomic.AudienceUpdateInput{
    AudienceID:  atomic.String("audience-id"),
    Name:        atomic.String("Updated Audience Name"),
    Description: atomic.String("Updated description"),
})

// Delete an audience
err := client.AudienceDelete(ctx, &atomic.AudienceDeleteInput{
    AudienceID: atomic.String("audience-id"),
})

// List audiences
audiences, err := client.AudienceList(ctx, &atomic.AudienceListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Publishers

Manage content publishers and channels.

```go
// Get a publisher
publisher, err := client.PublisherGet(ctx, &atomic.PublisherGetInput{
    PublisherID: atomic.String("publisher-id"),
})

// Create a publisher
publisher, err := client.PublisherCreate(ctx, &atomic.PublisherCreateInput{
    Name:        atomic.String("My Publisher"),
    Description: atomic.String("My publishing channel"),
    URL:         atomic.String("https://mypublisher.com"),
})

// Update a publisher
publisher, err := client.PublisherUpdate(ctx, &atomic.PublisherUpdateInput{
    PublisherID: atomic.String("publisher-id"),
    Name:        atomic.String("Updated Publisher Name"),
})

// Delete a publisher
err := client.PublisherDelete(ctx, &atomic.PublisherDeleteInput{
    PublisherID: atomic.String("publisher-id"),
})

// List publishers
publishers, err := client.PublisherList(ctx, &atomic.PublisherListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Distributions

Manage content distribution and syndication.

```go
// Get a distribution
distribution, err := client.DistributionGet(ctx, &atomic.DistributionGetInput{
    DistributionID: atomic.String("distribution-id"),
})

// Create a distribution
distribution, err := client.DistributionCreate(ctx, &atomic.DistributionCreateInput{
    ArticleID:   atomic.String("article-id"),
    PublisherID: atomic.String("publisher-id"),
    Status:      atomic.String("scheduled"),
})

// Update a distribution
distribution, err := client.DistributionUpdate(ctx, &atomic.DistributionUpdateInput{
    DistributionID: atomic.String("distribution-id"),
    Status:         atomic.String("published"),
})

// Delete a distribution
err := client.DistributionDelete(ctx, &atomic.DistributionDeleteInput{
    DistributionID: atomic.String("distribution-id"),
})

// List distributions
distributions, err := client.DistributionList(ctx, &atomic.DistributionListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Instances

Manage multi-tenant instances.

```go
// Get an instance
instance, err := client.InstanceGet(ctx, &atomic.InstanceGetInput{
    InstanceID: atomic.String("instance-id"),
})

// Create an instance
instance, err := client.InstanceCreate(ctx, &atomic.InstanceCreateInput{
    Name:        atomic.String("My Instance"),
    Description: atomic.String("My tenant instance"),
    Domain:      atomic.String("myinstance.atomic.com"),
})

// Update an instance
instance, err := client.InstanceUpdate(ctx, &atomic.InstanceUpdateInput{
    InstanceID:  atomic.String("instance-id"),
    Name:        atomic.String("Updated Instance Name"),
    Description: atomic.String("Updated description"),
})

// Delete an instance
err := client.InstanceDelete(ctx, &atomic.InstanceDeleteInput{
    InstanceID: atomic.String("instance-id"),
})

// List instances
instances, err := client.InstanceList(ctx, &atomic.InstanceListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Options

Manage configuration options and settings.

```go
// Get an option
option, err := client.OptionGet(ctx, &atomic.OptionGetInput{
    Name: atomic.String("site_name"),
})

// Update an option
option, err := client.OptionUpdate(ctx, &atomic.OptionUpdateInput{
    Name:  atomic.String("site_name"),
    Value: atomic.String("My Site"),
})

// Remove an option
err := client.OptionRemove(ctx, &atomic.OptionRemoveInput{
    Name: atomic.String("site_name"),
})

// List options
options, err := client.OptionList(ctx, &atomic.OptionListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})
```

### Communication

Send emails and SMS messages.

```go
// Send an email
messages, err := client.SendMail(ctx, &atomic.SendMailInput{
    To:      []string{"user@example.com"},
    Subject: atomic.String("Welcome!"),
    Body:    atomic.String("Welcome to our platform!"),
    TemplateID: atomic.String("template-id"), // Optional
})

// Send an SMS
messages, err := client.SendSMS(ctx, &atomic.SendSMSInput{
    To:   []string{"+1234567890"},
    Body: atomic.String("Your verification code is 123456"),
})
```

## Error Handling

The library provides detailed error information:

```go
user, err := client.UserGet(ctx, &atomic.UserGetInput{
    UserID: atomic.String("invalid-id"),
})
if err != nil {
    if apiErr, ok := err.(atomic.Error); ok {
        log.Printf("API Error: %s (Code: %s)", apiErr.Message, apiErr.Code)
    } else {
        log.Printf("Network Error: %v", err)
    }
}
```

## Response Handling

All API responses are wrapped in a `Resource` type that provides access to the response data and metadata:

```go
user, err := client.UserGet(ctx, &atomic.UserGetInput{
    UserID: atomic.String("user-id"),
})
if err != nil {
    log.Fatal(err)
}

// Access the user data
fmt.Printf("User: %s\n", user.Email)

// Access response metadata (if available)
if user.LastResponse != nil {
    fmt.Printf("Status: %s\n", user.LastResponse.Status)
    fmt.Printf("Headers: %v\n", user.LastResponse.Headers)
}
```

## Pagination

For list endpoints, use the `Limit` and `Offset` parameters for pagination:

```go
// Get first page
users, err := client.UserList(ctx, &atomic.UserListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(0),
})

// Get second page
users, err := client.UserList(ctx, &atomic.UserListInput{
    Limit:  atomic.Int(10),
    Offset: atomic.Int(10),
})
```

## Context Support

All API methods accept a `context.Context` for cancellation and timeouts:

```go
ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
defer cancel()

user, err := client.UserGet(ctx, &atomic.UserGetInput{
    UserID: atomic.String("user-id"),
})
```

## Instance Support

For multi-tenant applications, you can specify an instance ID in the context:

```go
ctx := atomic.ContextWithInstance(context.Background(), "instance-id")

user, err := client.UserGet(ctx, &atomic.UserGetInput{
    UserID: atomic.String("user-id"),
})
```

## Dependencies

The library depends on the following packages:

- `github.com/libatomic/atomic` - Core Atomic types and models
- `github.com/google/go-querystring` - Query string encoding
- `golang.org/x/oauth2` - OAuth2 authentication
- `github.com/go-ozzo/ozzo-validation/v4` - Input validation

## License

This library is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.


