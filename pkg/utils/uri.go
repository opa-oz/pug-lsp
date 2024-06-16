package utils

import (
	"errors"
	"net/url"
	"runtime"
	"strings"
)

// UriToPath converts a URI string to a file path string suitable for the current operating system.
//
// It replaces backslashes ("%5C") in the URI with forward slashes ("/"), parses the URI using url.Parse,
// and verifies that the URI scheme is "file". If the scheme is not "file", it returns an error indicating
// that the URI was not a valid file:// URI.
//
// On Windows, special handling is applied due to differences in URI parsing behavior. For URIs like
// "file:///c:/tmp/foo.md", the function ensures that the resulting path is correctly formatted for use
// as a file path in Windows. It strips an unnecessary leading '/' character and handles cases where the
// drive letter (e.g., "c:") may be in different parts of the parsed URL.
//
// Parameters:
//
//	uri string - The URI string to convert to a file path.
//
// Returns:
//
//	string - The converted file path string.
//	error - An error if there was any issue parsing the URI or if the URI is not a file:// URI.
func UriToPath(uri string) (string, error) {
	s := strings.ReplaceAll(uri, "%5C", "/")
	parsed, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	if parsed.Scheme != "file" {
		return "", errors.New("URI was not a file:// URI")
	}

	if runtime.GOOS == "windows" {
		// In Windows "file:///c:/tmp/foo.md" is parsed to "/c:/tmp/foo.md".
		// Strip the first character to get a valid path.
		if strings.Contains(parsed.Path[1:], ":") {
			// url.Parse() behaves differently with "file:///c:/..." and "file://c:/..."
			return parsed.Path[1:], nil
		} else {
			// if the windows drive is not included in Path it will be in Host
			return parsed.Host + "/" + parsed.Path[1:], nil
		}
	}
	return parsed.Path, nil
}
