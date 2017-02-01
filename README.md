#deptest

Output of glide-report:
$ glide-report
[WARN]	Disclaimer, this report is to help highlight things to consider. It is
[WARN]	alpha software and the rules are still under consideration.
[INFO]	Reading glide.yaml file to understand configured versions and ranges
[INFO]	Reading glide.lock file to understand pinned revisions
[INFO]	Fetching dependency data, this may take a moment...
DEBUG dep.Remote in fetchMissingDeps: https://golang.org/x/crypto
Report on .

------------------------------------------------------------------------------
Direct Imports
------------------------------------------------------------------------------

Analysis of golang.org/x/crypto:
● Dependency does not provide Semantic Version releases
[ERROR]	Unable to get cached repo data: open /Users/dmitris/.glide/cache/info/https-go.googlesource.com-crypto.json: no such file or directory
