package tables

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"io"
	"net/http"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type ExampleData struct {
	ProgramName            string       `json:"program_name"`
	PolicyURL              string       `json:"policy_url"`
	PolicyURLStatus        string       `json:"policy_url_status"`
	ContactURL             string       `json:"contact_url"`
	ContactEmail           string       `json:"contact_email"`
	LaunchDate             string       `json:"launch_date"`
	OffersBounty           string       `json:"offers_bounty"`
	OffersSwag             BoolOrString `json:"offers_swag"`
	HallOfFame             string       `json:"hall_of_fame"`
	SafeHarbor             string       `json:"safe_harbor"`
	PublicDisclosure       string       `json:"public_disclosure"`
	DisclosureTimelineDays float64      `json:"disclosure_timeline_days"`
	PgpKey                 string       `json:"pgp_key"`
	Hiring                 string       `json:"hiring"`
	SecuritytxtURL         string       `json:"securitytxt_url"`
	PreferredLanguages     string       `json:"preferred_languages"`
}

// BoolOrString is a custom type to handle both boolean and string representations.
type BoolOrString struct {
	BoolValue bool
	IsSet     bool
}

func (b *BoolOrString) UnmarshalJSON(data []byte) error {
	// Handle the boolean case
	var boolValue bool
	if err := json.Unmarshal(data, &boolValue); err == nil {
		b.BoolValue = boolValue
		b.IsSet = true
		return nil
	}

	// Handle the string case
	var stringValue string
	if err := json.Unmarshal(data, &stringValue); err == nil {
		if stringValue == "true" {
			b.BoolValue = true
		} else {
			b.BoolValue = false
		}
		b.IsSet = true
		return nil
	}

	return fmt.Errorf("invalid value for BoolOrString: %s", string(data))
}

func (b *BoolOrString) ToBool() bool {
	return b.BoolValue
}

func (b *BoolOrString) IsZero() bool {
	return !b.IsSet
}

func TableJSON() *plugin.Table {
	return &plugin.Table{
		Name:        "diodb",
		Description: "diodb table provides details about entities from a JSON file.",
		List: &plugin.ListConfig{
			Hydrate: listJSON,
		},
		Columns: []*plugin.Column{
			{
				Name:        "program_name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the program.",
				Transform:   transform.FromField("ProgramName"),
			},
			{
				Name:        "policy_url",
				Type:        proto.ColumnType_STRING,
				Description: "The URL of the policy.",
				Transform:   transform.FromField("PolicyURL"),
			},
			{
				Name:        "policy_url_status",
				Type:        proto.ColumnType_STRING,
				Description: "The status of the policy URL.",
				Transform:   transform.FromField("PolicyURLStatus"),
			},
			{
				Name:        "contact_url",
				Type:        proto.ColumnType_STRING,
				Description: "The contact URL.",
				Transform:   transform.FromField("ContactURL"),
			},
			{
				Name:        "contact_email",
				Type:        proto.ColumnType_STRING,
				Description: "The contact email.",
				Transform:   transform.FromField("ContactEmail"),
			},
			{
				Name:        "launch_date",
				Type:        proto.ColumnType_STRING,
				Description: "The launch date of the program.",
				Transform:   transform.FromField("LaunchDate"),
			},
			{
				Name:        "offers_bounty",
				Type:        proto.ColumnType_STRING,
				Description: "Indicates if the program offers bounty.",
				Transform:   transform.FromField("OffersBounty"),
			},
			{
				Name:        "offers_swag",
				Type:        proto.ColumnType_BOOL,
				Description: "Indicates if the program offers swag.",
				Transform:   transform.FromField("OffersSwag.ToBool"),
			},
			{
				Name:        "hall_of_fame",
				Type:        proto.ColumnType_STRING,
				Description: "The URL of the hall of fame.",
				Transform:   transform.FromField("HallOfFame"),
			},
			{
				Name:        "safe_harbor",
				Type:        proto.ColumnType_STRING,
				Description: "The safe harbor policy.",
				Transform:   transform.FromField("SafeHarbor"),
			},
			{
				Name:        "public_disclosure",
				Type:        proto.ColumnType_STRING,
				Description: "The public disclosure policy.",
				Transform:   transform.FromField("PublicDisclosure"),
			},
			{
				Name:        "disclosure_timeline_days",
				Type:        proto.ColumnType_DOUBLE,
				Description: "The disclosure timeline in days.",
				Transform:   transform.FromField("DisclosureTimelineDays"),
			},
			{
				Name:        "pgp_key",
				Type:        proto.ColumnType_STRING,
				Description: "The URL of the PGP key.",
				Transform:   transform.FromField("PgpKey"),
			},
			{
				Name:        "hiring",
				Type:        proto.ColumnType_STRING,
				Description: "The URL of the hiring page.",
				Transform:   transform.FromField("Hiring"),
			},
			{
				Name:        "securitytxt_url",
				Type:        proto.ColumnType_STRING,
				Description: "The URL of the security.txt file.",
				Transform:   transform.FromField("SecuritytxtURL"),
			},
			{
				Name:        "preferred_languages",
				Type:        proto.ColumnType_STRING,
				Description: "The preferred languages.",
				Transform:   transform.FromField("PreferredLanguages"),
			},
		},
	}
}

func listJSON(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Define the URL of the JSON file on GitHub
	url := "https://raw.githubusercontent.com/disclose/diodb/master/program-list.json"

	// Fetch the JSON data
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JSON data: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("failed to close resp.Body")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch JSON data: received status code %d", resp.StatusCode)
	}

	// Parse the JSON data
	var data []ExampleData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON data: %v", err)
	}

	// Stream the parsed data to Steampipe
	for _, item := range data {
		d.StreamListItem(ctx, item)
	}

	return nil, nil
}
