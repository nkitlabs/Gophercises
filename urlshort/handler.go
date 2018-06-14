package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dst, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dst, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// BuildMapByYaml will parse the provided YAML and then return
// the map of mapHandle (path: url)
//
// YAML is expected to be in the format:
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having invalid YAML data.
func BuildMapByYaml(yamlBytes []byte) (map[string]string, error) {
	pathsToUrls := map[string]string{}

	var pathURLs []pathURL

	if err := yaml.Unmarshal(yamlBytes, &pathURLs); err != nil {
		return nil, err
	}

	for _, pathURL := range pathURLs {
		pathsToUrls[pathURL.Path] = pathURL.URL
	}

	return pathsToUrls, nil

}
