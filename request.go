
package eos_contract_api_client

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