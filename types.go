
package eos_contract_api_client

import (
    "time"
    "strconv"
    "encoding/json"
)

type UnixTime int64

func (ts *UnixTime) UnmarshalJSON(b []byte) error {

    var i int64

    // "borrowed" from "gopkg.in/guregu/null.v4" abit.
    if err := json.Unmarshal(b, &i); err != nil {

        // If unmarshal to int64 fails, we assume that its a numeric string.
        var str string
        if err := json.Unmarshal(b, &str); err != nil {
            return err
        }

        // Then we need to parse the string into int64
        i, err = strconv.ParseInt(str, 10, 64)
        if err != nil {
            return err
        }
    }

    *ts = UnixTime(i)
    return nil
}

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

// Basic types

type Token struct {
    Contract string `json:"token_contract"`
    Symbol string   `json:"token_symbol"`
    Precision int   `json:"token_precision"`
    Amount string   `json:"amount"`
}

type Log struct {
    ID string                               `json:"log_id"`
    Name string                             `json:"name"`
    Data map[string]interface{}             `json:"data"`
    CreatedAtBlock string                   `json:"created_at_block"`
    CreatedAtTime UnixTime                  `json:"created_at_time"`
}


// Asset types

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

type ListingAsset struct {
    AssetID string                          `json:"asset_id"`
    Contract string                         `json:"contract"`
    Onwer string                            `json:"owner"`
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
    Data map[string]interface{}             `json:"data"`

    BurnedByAccount string                  `json:"burned_by_account"`
    BurnedAtBlock string                    `json:"burned_at_block"`
    BurnedAtTime string                     `json:"burned_at_time"`

    UpdatedAtBlock string                   `json:"updated_at_block"`
    UpdatedAtTime string                    `json:"updated_at_time"`

    TransferedAtBlock string                `json:"transferred_at_block"`
    TransferedAtTime string                 `json:"transferred_at_time"`

    MintedAtBlock string                    `json:"minted_at_block"`
    MintedAtTime string                     `json:"minted_at_time"`

    Sales []Sale                            `json:"sales"`
    Auctions []Auction                      `json:"actions"`
    Prices []Price                          `json:"prices"`
}

type AssetSale struct {
    ID string                   `json:"sale_id"`
    MarketContract string       `json:"market_contract"`
    AuctionID string            `json:"auction_id"`
    BuyOfferID string           `json:"buyoffer_id"`
    Price string                `json:"price"`
    TokenSymbol string          `json:"token_symbol"`
    TokenPrecision int64        `json:"token_precision"`
    TokenContract string        `json:"token_contract"`
    Seller string               `json:"seller"`
    Buyer string                `json:"buyer"`
    BlockTime UnixTime          `json:"block_time"`
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

// Schema types

type Schema struct {
    Name string             `json:"schema_name"`
    Contract string         `json:"contract"`
    Format []SchemaFormat   `json:"format"`
    CreatedAtBlock string   `json:"created_at_block"`
    CreatedAtTime string    `json:"created_at_time"`
}

type SchemaFormat struct {
    Name string `json:"name"`
    Type string `json:"type"`
}

// Template types

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

// Offer types

type Offer struct {
    ID string                   `json:"offer_id"`
    Contract string             `json:"contract"`
    Sender string               `json:"sender_name"`
    Recipient string            `json:"recipient_name"`
    Memo string                 `json:"memo"`
    State int64                 `json:"state"`
    IsSenderContract bool       `json:"is_sender_contract"`
    IsRecipientContract bool    `json:"is_recipient_contract"`
    SenderAssets []Asset        `json:"sender_assets"`
    RecipientAssets []Asset     `json:"recipient_assets"`
    UpdatedAtBlock string       `json:"updated_at_block"`
    UpdatedAtTime UnixTime      `json:"updated_at_time"`
    CreatedAtBlock string       `json:"created_at_block"`
    CreatedAtTime UnixTime      `Json:"created_at_time"`
}

type ListingOffer struct {
    ID string                           `json:"offer_id"`
    Contract string                     `json:"contract"`
    Sender string                       `json:"sender_name"`
    Recipient string                    `json:"recipient_name"`
    Memo string                         `json:"memo"`
    State int64                         `json:"state"`
    IsSenderContract bool               `json:"is_sender_contract"`
    IsRecipientContract bool            `json:"is_recipient_contract"`
    SenderAssets []ListingAsset         `json:"sender_assets"`
    RecipientAssets []ListingAsset      `json:"recipient_assets"`
    UpdatedAtBlock string               `json:"updated_at_block"`
    UpdatedAtTime UnixTime              `json:"updated_at_time"`
    CreatedAtBlock string               `json:"created_at_block"`
    CreatedAtTime UnixTime              `Json:"created_at_time"`
}

type BuyOffer struct {
    ID string                   `json:"buyoffer_id"`
    MarketContract string       `json:"market_contract"`
    AssetsContract string       `json:"assets_contract"`
    Seller string               `json:"seller"`
    Buyer string                `json:"buyer"`
    Price Token                 `json:"price"`
    Assets []Asset              `json:"assets"`
    MakerMarketplace string     `json:"maker_marketplace,omitempty"`
    TakerMarketplace string     `json:"taker_marketplace,omitempty"`
    Collection Collection       `json:"collection"`
    State int64                 `json:"state"`
    Memo string                 `json:"memo"`
    DeclineMemo string          `json:"decline_memo"`
    UpdatedAtBlock string       `json:"updated_at_block"`
    UpdatedAtTime UnixTime      `json:"updated_at_time"`
    CreatedAtBlock string       `json:"created_at_block"`
    CreatedAtTime UnixTime      `Json:"created_at_time"`
}

// Transfer types

type Transfer struct {
    ID string                   `json:"transfer_id"`
    Contract string             `json:"contract"`
    Sender string               `json:"sender_name"`
    Recipient string            `json:"recipient_name"`
    Memo string                 `json:"memo"`
    Assets []Asset              `json:"assets"`
    CreatedAtBlock string       `json:"created_at_block"`
    CreatedAtTime UnixTime      `Json:"created_at_time"`
}

type ListingTransfer struct {
    ID string                   `json:"transfer_id"`
    Contract string             `json:"contract"`
    Sender string               `json:"sender_name"`
    Recipient string            `json:"recipient_name"`
    Memo string                 `json:"memo"`
    Assets []ListingAsset       `json:"assets"`
    CreatedAtBlock string       `json:"created_at_block"`
    CreatedAtTime UnixTime      `Json:"created_at_time"`
}

// Sale types

type Sale struct {
    ID string                   `json:"sales_id"`
    MarketContract string       `json:"market_contract"`
    AsssetsContract string      `json:"assets_contract"`
    Seller string               `json:"seller"`
    Buyer string                `json:"buyer"`
    OfferID string              `json:"offer_id"`
    Price Price                 `json:"price"`
    ListingPrice int64          `json:"listing_price"`
    ListingSymbol string        `json:"listing_symbol"`
    Assets []Asset              `json:"assets"`
    MakerMarketplace string     `json:"maker_marketplace,omitempty"`
    TakerMarketplace string     `json:"taker_marketplace,omitempty"`
    Collection Collection       `json:"collection"`
    State int64                 `json:"state"`
    UpdatedAtBlock string       `json:"updated_at_block"`
    UpdatedAtTime UnixTime      `json:"updated_at_time"`
    CreatedAtBlock string       `json:"created_at_block"`
    CreatedAtTime UnixTime      `Json:"created_at_time"`
}

// Action types

type Auction struct {
    ID string               `json:"action_id"`
    MarketContract string   `json:"market_contract"`
}

// Marketplace types

type MarketPlace struct {
    Name string                 `json:"marketplace_name"`
    Creator string              `json:"creator"`
    CreatedAtBlock string       `json:"created_at_block"`
    CreatedAtTime UnixTime      `Json:"created_at_time"`
}

// Price types

type Price struct {
    Average	string          `json:"average"`
    MarketContract string   `json:"market_contract"`
    Max	string              `json:"max"`
    Median string           `json:"median"`
    Min	string              `json:"min"`
    Sales string            `json:"sales"`
    SuggestedAverage string `json:"suggested_average"`
    SuggestedMedian	string  `json:"suggested_median"`
    Token PriceToken        `json:"token"`
}

type PriceToken struct {
    Contract string     `json:"token_contract"`
    Symbol string       `json:"token_symbol"`
    Precision int       `json:"token_precision"`
}

// Link types

type Link struct {
    ID string               `json:"link_id"`
    ToolsContract string    `json:"tools_contract"`
    AssetsContract string   `json:"assets_contract"`
    Creator string          `json:"creator"`
    Claimer string          `json:"claimer,omitempty"`
    State int64             `json:"state"`
    PublicKey string        `json:"public_key"`
    Memo string             `json:"memo"`
    TxID string             `json:"txid"`
    Assets []Asset          `json:"assets"`
    CreatedAtBlock string   `json:"created_at_block"`
    CreatedAtTime UnixTime  `Json:"created_at_time"`
}
