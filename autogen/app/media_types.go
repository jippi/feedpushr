// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "feedpushr": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/ncarlier/feedpushr/v3/design
// --out=/home/fr23972/workspace/fe/feedpushr/autogen
// --version=v1.4.3

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// The search result (default view)
//
// Identifier: application/vnd.feedpushr.explore.v2+json; view=default
type ExploreResponse struct {
	// Feed description
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// URL of the feed website
	HTMLURL string `form:"htmlUrl" json:"htmlUrl" yaml:"htmlUrl" xml:"htmlUrl"`
	// Feed title
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
	// URL of the XML feed
	XMLURL string `form:"xmlUrl" json:"xmlUrl" yaml:"xmlUrl" xml:"xmlUrl"`
}

// Validate validates the ExploreResponse media type instance.
func (mt *ExploreResponse) Validate() (err error) {
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.XMLURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "xmlUrl"))
	}
	if mt.HTMLURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "htmlUrl"))
	}
	return
}

// ExploreResponseCollection is the media type for an array of ExploreResponse (default view)
//
// Identifier: application/vnd.feedpushr.explore.v2+json; type=collection; view=default
type ExploreResponseCollection []*ExploreResponse

// Validate validates the ExploreResponseCollection media type instance.
func (mt ExploreResponseCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A RSS feed (default view)
//
// Identifier: application/vnd.feedpushr.feed.v2+json; view=default
type FeedResponse struct {
	// Date of creation
	Cdate time.Time `form:"cdate" json:"cdate" yaml:"cdate" xml:"cdate"`
	// Number of consecutive aggregation errors
	ErrorCount *int `form:"errorCount,omitempty" json:"errorCount,omitempty" yaml:"errorCount,omitempty" xml:"errorCount,omitempty"`
	// Last aggregation error
	ErrorMsg *string `form:"errorMsg,omitempty" json:"errorMsg,omitempty" yaml:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// URL of the feed website
	HTMLURL *string `form:"htmlUrl,omitempty" json:"htmlUrl,omitempty" yaml:"htmlUrl,omitempty" xml:"htmlUrl,omitempty"`
	// URL of the PubSubHubbud hub
	HubURL *string `form:"hubUrl,omitempty" json:"hubUrl,omitempty" yaml:"hubUrl,omitempty" xml:"hubUrl,omitempty"`
	// ID of feed (MD5 of the xmlUrl)
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Last aggregation pass
	LastCheck *time.Time `form:"lastCheck,omitempty" json:"lastCheck,omitempty" yaml:"lastCheck,omitempty" xml:"lastCheck,omitempty"`
	// Date of modification
	Mdate time.Time `form:"mdate" json:"mdate" yaml:"mdate" xml:"mdate"`
	// Total number of processed items
	NbProcessedItems *int `form:"nbProcessedItems,omitempty" json:"nbProcessedItems,omitempty" yaml:"nbProcessedItems,omitempty" xml:"nbProcessedItems,omitempty"`
	// Next aggregation pass
	NextCheck *time.Time `form:"nextCheck,omitempty" json:"nextCheck,omitempty" yaml:"nextCheck,omitempty" xml:"nextCheck,omitempty"`
	// Aggregation status
	Status *string `form:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty" xml:"status,omitempty"`
	// List of tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" yaml:"tags,omitempty" xml:"tags,omitempty"`
	// Title of the Feed
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
	// URL of the XML feed
	XMLURL string `form:"xmlUrl" json:"xmlUrl" yaml:"xmlUrl" xml:"xmlUrl"`
}

// Validate validates the FeedResponse media type instance.
func (mt *FeedResponse) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.XMLURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "xmlUrl"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	if mt.Status != nil {
		if !(*mt.Status == "running" || *mt.Status == "stopped") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.status`, *mt.Status, []interface{}{"running", "stopped"}))
		}
	}
	return
}

// A RSS feed (link view)
//
// Identifier: application/vnd.feedpushr.feed.v2+json; view=link
type FeedResponseLink struct {
	// ID of feed (MD5 of the xmlUrl)
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// URL of the XML feed
	XMLURL string `form:"xmlUrl" json:"xmlUrl" yaml:"xmlUrl" xml:"xmlUrl"`
}

// Validate validates the FeedResponseLink media type instance.
func (mt *FeedResponseLink) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.XMLURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "xmlUrl"))
	}
	return
}

// A RSS feed (tiny view)
//
// Identifier: application/vnd.feedpushr.feed.v2+json; view=tiny
type FeedResponseTiny struct {
	// Date of creation
	Cdate time.Time `form:"cdate" json:"cdate" yaml:"cdate" xml:"cdate"`
	// ID of feed (MD5 of the xmlUrl)
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// List of tags
	Tags []string `form:"tags,omitempty" json:"tags,omitempty" yaml:"tags,omitempty" xml:"tags,omitempty"`
	// Title of the Feed
	Title string `form:"title" json:"title" yaml:"title" xml:"title"`
	// URL of the XML feed
	XMLURL string `form:"xmlUrl" json:"xmlUrl" yaml:"xmlUrl" xml:"xmlUrl"`
}

// Validate validates the FeedResponseTiny media type instance.
func (mt *FeedResponseTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.XMLURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "xmlUrl"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}

	return
}

// FeedResponseCollection is the media type for an array of FeedResponse (default view)
//
// Identifier: application/vnd.feedpushr.feed.v2+json; type=collection; view=default
type FeedResponseCollection []*FeedResponse

// Validate validates the FeedResponseCollection media type instance.
func (mt FeedResponseCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// FeedResponseCollection is the media type for an array of FeedResponse (link view)
//
// Identifier: application/vnd.feedpushr.feed.v2+json; type=collection; view=link
type FeedResponseLinkCollection []*FeedResponseLink

// Validate validates the FeedResponseLinkCollection media type instance.
func (mt FeedResponseLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// FeedResponseCollection is the media type for an array of FeedResponse (tiny view)
//
// Identifier: application/vnd.feedpushr.feed.v2+json; type=collection; view=tiny
type FeedResponseTinyCollection []*FeedResponseTiny

// Validate validates the FeedResponseTinyCollection media type instance.
func (mt FeedResponseTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A pagignated list of feeds (default view)
//
// Identifier: application/vnd.feedpushr.feeds-page.v2+json; view=default
type FeedsPageResponse struct {
	// Current page number
	Current int `form:"current" json:"current" yaml:"current" xml:"current"`
	// List of feeds
	Data FeedResponseCollection `form:"data" json:"data" yaml:"data" xml:"data"`
	// Max number of feeds by page
	Size int `form:"size" json:"size" yaml:"size" xml:"size"`
	// Total number of feeds
	Total int `form:"total" json:"total" yaml:"total" xml:"total"`
}

// Validate validates the FeedsPageResponse media type instance.
func (mt *FeedsPageResponse) Validate() (err error) {

	if mt.Data == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "data"))
	}
	if err2 := mt.Data.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// The filter specification (default view)
//
// Identifier: application/vnd.feedpushr.filter-spec.v2+json; view=default
type FilterSpecResponse struct {
	// Description of the filter
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// Name of the filter
	Name  string             `form:"name" json:"name" yaml:"name" xml:"name"`
	Props PropSpecCollection `form:"props" json:"props" yaml:"props" xml:"props"`
}

// Validate validates the FilterSpecResponse media type instance.
func (mt *FilterSpecResponse) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.Props == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "props"))
	}
	if err2 := mt.Props.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// FilterSpecResponseCollection is the media type for an array of FilterSpecResponse (default view)
//
// Identifier: application/vnd.feedpushr.filter-spec.v2+json; type=collection; view=default
type FilterSpecResponseCollection []*FilterSpecResponse

// Validate validates the FilterSpecResponseCollection media type instance.
func (mt FilterSpecResponseCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// A filter (default view)
//
// Identifier: application/vnd.feedpushr.filter.v2+json; view=default
type FilterResponse struct {
	// Alias of the filter
	Alias string `form:"alias" json:"alias" yaml:"alias" xml:"alias"`
	// Conditional expression of the filter
	Condition string `form:"condition" json:"condition" yaml:"condition" xml:"condition"`
	// Description of the filter
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// Status
	Enabled bool `form:"enabled" json:"enabled" yaml:"enabled" xml:"enabled"`
	// ID of the filter
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Name of the filter
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Number of error
	NbError int `form:"nbError" json:"nbError" yaml:"nbError" xml:"nbError"`
	// Number of success
	NbSuccess int `form:"nbSuccess" json:"nbSuccess" yaml:"nbSuccess" xml:"nbSuccess"`
	// Filter properties
	Props map[string]interface{} `form:"props" json:"props" yaml:"props" xml:"props"`
}

// Validate validates the FilterResponse media type instance.
func (mt *FilterResponse) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Alias == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "alias"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.Condition == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "condition"))
	}
	if mt.Props == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "props"))
	}
	return
}

// FilterResponseCollection is the media type for an array of FilterResponse (default view)
//
// Identifier: application/vnd.feedpushr.filter.v2+json; type=collection; view=default
type FilterResponseCollection []*FilterResponse

// Validate validates the FilterResponseCollection media type instance.
func (mt FilterResponseCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// HAL link (default view)
//
// Identifier: application/vnd.feedpushr.hal-links.v2+json; view=default
type HALLink struct {
	// Link's destination
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
}

// Validate validates the HALLink media type instance.
func (mt *HALLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// API info (default view)
//
// Identifier: application/vnd.feedpushr.info.v2+json; view=default
type Info struct {
	// HAL links
	Links map[string]*HALLink `form:"_links" json:"_links" yaml:"_links" xml:"_links"`
	// Service description
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// Service name
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Service version
	Version string `form:"version" json:"version" yaml:"version" xml:"version"`
}

// Validate validates the Info media type instance.
func (mt *Info) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}
	if mt.Links == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "_links"))
	}
	for _, e := range mt.Links {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// OPMLImportJobResponse media type (default view)
//
// Identifier: application/vnd.feedpushr.ompl-import-job.v2+json; view=default
type OPMLImportJobResponse struct {
	// ID of the import job
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
}

// Validate validates the OPMLImportJobResponse media type instance.
func (mt *OPMLImportJobResponse) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	return
}

// The output channel specification (default view)
//
// Identifier: application/vnd.feedpushr.output-spec.v2+json; view=default
type OutputSpecResponse struct {
	// Description of the output channel
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// Name of the output channel
	Name  string             `form:"name" json:"name" yaml:"name" xml:"name"`
	Props PropSpecCollection `form:"props" json:"props" yaml:"props" xml:"props"`
}

// Validate validates the OutputSpecResponse media type instance.
func (mt *OutputSpecResponse) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.Props == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "props"))
	}
	if err2 := mt.Props.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// OutputSpecResponseCollection is the media type for an array of OutputSpecResponse (default view)
//
// Identifier: application/vnd.feedpushr.output-spec.v2+json; type=collection; view=default
type OutputSpecResponseCollection []*OutputSpecResponse

// Validate validates the OutputSpecResponseCollection media type instance.
func (mt OutputSpecResponseCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// The output channel (default view)
//
// Identifier: application/vnd.feedpushr.output.v2+json; view=default
type OutputResponse struct {
	// Alias of the output channel
	Alias string `form:"alias" json:"alias" yaml:"alias" xml:"alias"`
	// Conditional expression of the filter
	Condition string `form:"condition" json:"condition" yaml:"condition" xml:"condition"`
	// Description of the output channel
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// Status
	Enabled bool `form:"enabled" json:"enabled" yaml:"enabled" xml:"enabled"`
	// Filters
	Filters FilterResponseCollection `form:"filters,omitempty" json:"filters,omitempty" yaml:"filters,omitempty" xml:"filters,omitempty"`
	// ID of the output
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Name of the output channel
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Number of error
	NbError int `form:"nbError" json:"nbError" yaml:"nbError" xml:"nbError"`
	// Number of success
	NbSuccess int `form:"nbSuccess" json:"nbSuccess" yaml:"nbSuccess" xml:"nbSuccess"`
	// Output channel properties
	Props map[string]interface{} `form:"props" json:"props" yaml:"props" xml:"props"`
}

// Validate validates the OutputResponse media type instance.
func (mt *OutputResponse) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Alias == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "alias"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.Condition == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "condition"))
	}
	if mt.Props == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "props"))
	}
	if err2 := mt.Filters.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// OutputResponseCollection is the media type for an array of OutputResponse (default view)
//
// Identifier: application/vnd.feedpushr.output.v2+json; type=collection; view=default
type OutputResponseCollection []*OutputResponse

// Validate validates the OutputResponseCollection media type instance.
func (mt OutputResponseCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// The specification of a property (default view)
//
// Identifier: application/vnd.feedpushr.prop-spec.v2+json; view=default
type PropSpec struct {
	// Description of the output channel
	Desc string `form:"desc" json:"desc" yaml:"desc" xml:"desc"`
	// Name of the property
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Property options
	Options map[string]string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
	// Property type ('text', 'url', ...)
	Type string `form:"type" json:"type" yaml:"type" xml:"type"`
}

// Validate validates the PropSpec media type instance.
func (mt *PropSpec) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Desc == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "desc"))
	}
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	return
}

// PropSpecCollection is the media type for an array of PropSpec (default view)
//
// Identifier: application/vnd.feedpushr.prop-spec.v2+json; type=collection; view=default
type PropSpecCollection []*PropSpec

// Validate validates the PropSpecCollection media type instance.
func (mt PropSpecCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}
