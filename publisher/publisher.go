package publisher

import (
	"github.com/whosonfirst/go-whosonfirst-dist"
	"github.com/whosonfirst/go-whosonfirst-repo"
	"io"
)

type Publisher interface {
	Publish(io.ReadCloser, string) error
	Fetch(string) (io.ReadCloser, error)
	Prune(repo.Repo) error // most likely a string rather than a repo.Repo
	Index(repo.Repo) error // DEPRECATED... ALMOST
	BuildIndex(repo.Repo) (map[string][]*dist.Item, error)
	IsNotFound(error) bool
}

// THIS WILL BE THE NEW NEW...BUT NOT YET
// (20180810/thisisaaronland)

/*

func Index(p Publisher, r repo.Repo) error {

	items, err := p.BuildIndex(r)

	if err != nil {
		return err
	}

	// although it is true that all this template stuff could
	// be method-chained I find that it doesn't take long for
	// method-chaining to become inpenetrable gibberish so why
	// start now (20180807/thisisaaronland)

	// remember this is a github.com/whosonfirst/go-bindata-html-template
	// and not a plain vanilla html/template

	tpl := template.New("inventory", html.Asset)

	funcs := template.FuncMap{
		"humanize_bytes": func(i int64) string { return humanize.Bytes(uint64(i)) }, // u so great Go until u r annoying this way...
		"humanize_comma": humanize.Comma,
	}

	tpl = tpl.Funcs(funcs)

	html_tpl, err := tpl.ParseFiles("templates/html/inventory.html")

	if err != nil {
		return err
	}

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

		// please rename to something more generic than HTMLVars
		// because we're also going to (eventually) pass it to the
		// feed templates

		vars := HTMLVars{
			Date:  now.Format(time.RFC3339),
			Type:  t,
			Items: t_items,
		}

		// index.html

		var html_b bytes.Buffer
		html_wr := bufio.NewWriter(&html_b)

		err = html_tpl.Execute(html_wr, vars)

		if err != nil {
			return err
		}

		html_r := bytes.NewReader(html_b.Bytes())
		html_fh := ioutil.NopCloser(html_r)

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

		// rss.xml - please make this work
		// atom.xml - please make this work
	}

	return nil
}

*/
