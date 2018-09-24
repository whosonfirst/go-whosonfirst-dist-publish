package publisher

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/whosonfirst/go-bindata-html-template"
	"github.com/whosonfirst/go-whosonfirst-dist"
	"github.com/whosonfirst/go-whosonfirst-dist-publish/assets/feed"
	"github.com/whosonfirst/go-whosonfirst-dist-publish/assets/html"
	"github.com/whosonfirst/go-whosonfirst-repo"
	"io"
	"io/ioutil"
	_ "log"
	_ "os"
	"time"
)

type PublishVars struct {
	Date                string
	Type                string
	Items               []*dist.Item
	BuildDate           time.Time
	DistributionName    string
	DistributionRootURL string
	DistributionBlurb   string
}

type IndexOptions struct {
	DistributionName    string
	DistributionRootURL string
	DistributionBlurb   string
}

func NewDefaultIndexOptions() (*IndexOptions, error) {

	opts := IndexOptions{
		DistributionName:    "Who's On First",
		DistributionRootURL: "https://dist.whosonfirst.org/",
		DistributionBlurb:   `Who's On First is a gazetter of all the places. Note: As of this writing "alt" (or "alternative") files are not included in any of the distributions. If you need that data you will need to clone it directly from the https://github.com/whosonfirst-data GitHub organization.`,
	}

	return &opts, nil
}

type PruneOptions struct {
	MaxDistributions int
}

func NewDefaultPruneOptions() (*PruneOptions, error) {

	opts := PruneOptions{
		MaxDistributions: 1,
	}

	return &opts, nil
}

type Publisher interface {
	Publish(io.ReadCloser, string) error
	Fetch(string) (io.ReadCloser, error)
	Prune(repo.Repo, *PruneOptions) error // most likely a string rather than a repo.Repo
	BuildIndex(repo.Repo) (map[string][]*dist.Item, error)
	IsNotFound(error) bool
}

func Index(p Publisher, r repo.Repo, opts *IndexOptions) error {

	items, err := p.BuildIndex(r)

	if err != nil {
		return err
	}

	funcs := template.FuncMap{
		"humanize_bytes": func(i int64) string { return humanize.Bytes(uint64(i)) }, // u so great Go until u r annoying this way...
		"humanize_comma": humanize.Comma,
	}

	// although it is true that all this template stuff could
	// be method-chained I find that it doesn't take long for
	// method-chaining to become inpenetrable gibberish so why
	// start now (20180807/thisisaaronland)

	// remember this is a github.com/whosonfirst/go-bindata-html-template
	// and not a plain vanilla html/template

	tpl := template.New("inventory", html.Asset)

	tpl = tpl.Funcs(funcs)

	html_tpl, err := tpl.ParseFiles("templates/html/inventory.html")

	if err != nil {
		return err
	}

	now := time.Now()

	for t, t_items := range items {

		if t == "bundle" {
			t = "bundles" // ARGGHHHHGGGHNNGNGNNNFFFPPPHPPHPTTTT
		}

		html_key := fmt.Sprintf("%s/index.html", t)
		json_key := fmt.Sprintf("%s/inventory.json", t)
		rss_key := fmt.Sprintf("%s/rss.xml", t)
		atom_key := fmt.Sprintf("%s/atom.xml", t)

		vars := PublishVars{
			Date:                now.Format(time.RFC3339),
			Type:                t,
			Items:               t_items,
			BuildDate:           now,
			DistributionName:    opts.DistributionName,
			DistributionRootURL: opts.DistributionRootURL,
			DistributionBlurb:   opts.DistributionBlurb,
		}

		// index.html

		html_fh, err := renderTemplate(html_tpl, vars, false)

		if err != nil {
			return err
		}

		err = p.Publish(html_fh, html_key)

		if err != nil {
			return err
		}

		// inventory.json

		json_b, err := json.Marshal(t_items)

		if err != nil {
			return err
		}

		json_r := bytes.NewReader(json_b)
		json_fh := ioutil.NopCloser(json_r)

		err = p.Publish(json_fh, json_key)

		if err != nil {
			return err
		}

		// feeds / start by trimming items - we should be more nuanced
		// about _how_ we do this (20180813/thisisaaronland)

		if len(t_items) > 10 {
			t_items = t_items[0:10]
		}

		// rss.xml

		rss := template.New("feed_rss_20", feed.Asset)
		rss = rss.Funcs(funcs)

		rss_tpl, err := rss.ParseFiles("templates/feed/rss_2.0.xml")

		if err != nil {
			return err
		}

		rss_fh, err := renderTemplate(rss_tpl, vars, true)

		if err != nil {
			return err
		}

		err = p.Publish(rss_fh, rss_key)

		if err != nil {
			return err
		}

		// atom.xml

		atom := template.New("feed_atom_10", feed.Asset)
		atom = atom.Funcs(funcs)

		atom_tpl, err := atom.ParseFiles("templates/feed/atom_1.0.xml")

		if err != nil {
			return err
		}

		atom_fh, err := renderTemplate(atom_tpl, vars, true)

		if err != nil {
			return err
		}

		err = p.Publish(atom_fh, atom_key)

		if err != nil {
			return err
		}

	}

	return nil
}

func renderTemplate(tpl *template.Template, vars interface{}, is_xml bool) (io.ReadCloser, error) {

	var b bytes.Buffer
	wr := bufio.NewWriter(&b)

	// because we are using html/template and there's no way to not encode '<?'
	// as '&lt?' because Go is conservative that way... (20180813/thisisaaronland)

	if is_xml {
		wr.Write([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\" ?>"))
	}

	err := tpl.Execute(wr, vars)

	if err != nil {
		return nil, err
	}

	wr.Flush()

	r := bytes.NewReader(b.Bytes())
	fh := ioutil.NopCloser(r)

	return fh, nil
}
