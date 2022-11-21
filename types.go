
package eos_contract_api_client

import (
    "time"
    null "gopkg.in/guregu/null.v4"
)

type UnixTime int64

func (ts UnixTime) Time() time.Time {
    return time.Unix(int64(ts) / 1000, int64(ts) % 1000).UTC()
}

// HTTP

type HTTPResponse struct {
    HTTPStatusCode int
}

func (resp *HTTPResponse) IsError() bool {
    return resp.HTTPStatusCode == 0 || resp.HTTPStatusCode > 399
}

type APIError struct {
    Success null.Bool    `json:"success"`
    Message null.String  `json:"message"`
}

type APIResponse struct {
    HTTPResponse
    Success bool        `json:"success"`
    QueryTime UnixTime  `json:"query_time"`
}

// Health

type ChainHealth struct {
    Status      string      `json:"status"`
    HeadBlock   int64       `json:"head_block"`
    HeadTime    UnixTime    `json:"head_time"`
}

type RedisHealth struct {
    Status string `json:"status"`
}

type PostgresHealth struct {
    Status string                       `json:"status"`
    Readers []map[string]interface{}    `json:"readers"`
}

type HealthData struct {
    Version string          `json:"version"`
    Postgres PostgresHealth `json:"postgres"`
    Redis RedisHealth       `json:"redis"`
    Chain ChainHealth       `json:"chain"`
}

type Health struct {
    APIResponse
    Data HealthData
}


// Assets request/response

type AssetResponse struct {
    APIResponse
    Data Asset
}

type AssetsRequestParams struct {
    CollectionName string           `qs:"collection_name,omitempty"`
    CollectionBlacklist []string    `qs:"collection_blacklist,omitempty"`
    CollectionWhitelist []string    `qs:"collection_whitelist,omitempty"`
    SchemaName string               `qs:"schema_name,omitempty"`
    TemplateID int                  `qs:"template_id,omitempty"`
    TemplateWhitelist []int         `qs:"template_whitelist,omitempty"`
    TemplateBlacklist []int         `qs:"template_blacklist,omitempty"`
    Owner string                    `qs:"owner,omitempty"`
    Match string                    `qs:"match,omitempty"`
    MatchImmutableName string       `qs:"match_immutable_name,omitempty"`
    MatchMutableName string         `qs:"match_mutable_name,omitempty"`
    HideTemplatesByAccounts string  `qs:"hide_templates_by_accounts,omitempty"`

    IsTransferable bool             `qs:"is_transferable,omitempty"`
    IsBurnable bool                 `qs:"is_burnable,omitempty"`
    Burned bool                     `qs:"burned,omitempty"`
    OnlyDuplicatedTemplates bool    `qs:"only_duplicated_templates,omitempty"`
    HasBackedTokens bool            `qs:"has_backend_tokens,omitempty"`
    HideOffers bool                 `qs:"hide_offers,omitempty"`

    LowerBound string               `qs:"lower_bound,omitempty"`
    UpperBound string               `qs:"upper_bound,omitempty"`

    Before int                      `qs:"before,omitempty"`
    After int                       `qs:"after,omitempty"`

    Limit int                       `qs:"limit,omitempty"`
    Order string                    `qs:"order,omitempty"`
    Sort string                     `qs:"sort,omitempty"`
}

type AssetsResponse struct {
    APIResponse
    Data []Asset
}

// Token Type
type Token struct {
    Contract string `json:"token_contract"`
    Symbol string   `json:"token_symbol"`
    Precision int   `json:"token_precision"`
    Amount string   `json:"amount"`
}

// Asset type

type Asset struct {
    ID string                               `json:"asset_id"`
    Contract string                         `json:"contract"`
    Owner string                            `json:"owner"`
    Name string                             `json:"name"`
    IsTransferable bool                     `json:"is_transferable"`
    IsBurnable bool                         `json:"is_burnable"`
    TemplateMint string                     `json:"template_mint"`
    Collection Collection                   `json:"collection"`
    Schema Schema                           `json:"schema"`
    Template Template                       `json:"template"`
    BackedTokens []Token                    `json:"backed_tokens"`
    ImmutableData map[string]interface{}    `json:"immutable_data"`
    MutableData map[string]interface{}      `json:"mutable_data"`

    BurnedByAccount string                  `json:"burned_by_account"`
    BurnedAtBlock string                    `json:"burned_at_block"`
    BurnedAtTime string                     `json:"burned_at_time"`

    UpdatedAtBlock string                   `json:"updated_at_block"`
    UpdatedAtTime string                    `json:"updated_at_time"`

    TransferedAtBlock string                `json:"transferred_at_block"`
    TransferedAtTime string                 `json:"transferred_at_time"`

    MintedAtBlock string                    `json:"minted_at_block"`
    MintedAtTime string                     `json:"minted_at_time"`
}

// Schema type

type SchemaFormat struct {
    Name string `json:"name"`
    Type string `json:"type"`
}

type Schema struct {
    Name string             `json:"schema_name"`
    Contract string         `json:"contract"`
    Format []SchemaFormat   `json:"format"`
    CreatedAtBlock string   `json:"created_at_block"`
    CreatedAtTime string    `json:"created_at_time"`
}

// Collection type

type Collection struct {
    CollectionName string       `json:"collection_name"`
    Contract string             `json:"contract"`
    Name string                 `json:"name"`
    Author string               `json:"author"`
    AllowNotify bool            `json:"allow_notify"`
    AuthorizedAccounts []string `json:"authorized_accounts"`
    NotifyAccounts []string     `json:"notify_accounts"`
    MarketFee float64           `json:"market_fee"`
    CreatedAtBlock string       `json:"created_at_block"`
    CreatedAtTime string        `json:"created_at_time"`
}

type Template struct {
    ID string                               `json:"template_id"`
    Contract string                         `json:"contract"`
    MaxSupply string                        `json:"max_supply"`
    IssuedSupply string                     `json:"issued_supply"`
    IsTransferable bool                     `json:"is_transferable"`
    IsBurnable bool                         `json:"is_burnable"`
    ImmutableData map[string]interface{}    `json:"immutable_data"`
    CreatedAtBlock string                   `json:"created_at_block"`
    CreatedAtTime string                    `json:"created_at_time"`
}
