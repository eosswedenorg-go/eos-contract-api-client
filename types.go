
package eos_contract_api_client

import (
    "time"
)

type UnixTime int64

func (ts UnixTime) Time() time.Time {
    return time.Unix(int64(ts) / 1000, int64(ts) % 1000).UTC()
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
