# Examples

Please feel free to contribute your own examples on github

## Nginx

### HTTP Status

Parse error codes and graph in a histogram

```sh
$ rare h -m "\" (\d+)" -e "{1}" -z -x testdata/*

404                 5,557,374  [66.4%] ||||||||||||||||||||||||||||||||||||||||||||||||||
200                 2,564,984  [30.6%] |||||||||||||||||||||||
400                 243,282    [ 2.9%] ||
405                 5,708      [ 0.1%]
408                 1,397      [ 0.0%]
Matched: 8,373,328 / 8,383,717 (Groups: 8)
```

### Extracting Page Sizes

Page sizes, ignoring 0-sized pages

```sh
$ rare h -m "\" (\d+) (\d+)" -e "{bytesize {bucket {2} 1024}}" -i "{lt {2} 1024}" -z -x testdata/*

234 KB              3,602      [14.6%] ||||||||||||||||||||||||||||||||||||||||||||||||||
149 KB              2,107      [ 8.5%] |||||||||||||||||||||||||||||
193 KB              1,519      [ 6.2%] |||||||||||||||||||||
192 KB              1,470      [ 6.0%] ||||||||||||||||||||
191 KB              1,421      [ 5.8%] |||||||||||||||||||
Matched: 24,693 / 8,383,717 (Groups: 96) (Ignored: 8,348,635)
```

### Table of URLs to HTTP Status

Know how your URLs are responding by their http statuses

```sh
$ rare t -m "\"(\w+) (.+).*\" (\d+) (\d+)" -e "{$ {3} {substr {2} 0 20}}" -z testdata/*

                     404                  200                  400
/ HTTP/1.1           0                    127,624              5,681
/ HTTP/1.0           0                    5,222                0
/test.php HTTP/1.1   3,241                0                    0
/1.php HTTP/1.1      2,508                0                    0
/qq.php HTTP/1.1     1,908                0                    0
/index.php HTTP/1.1  1,776                0                    0
/shell.php HTTP/1.1  1,750                0                    0
/cmd.php HTTP/1.1    1,588                0                    0
/x.php HTTP/1.1      1,573                0                    0
/log.php HTTP/1.1    1,261                0                    0
/confg.php HTTP/1.1  1,253                0                    0
/api.php HTTP/1.1    1,241                0                    0
/ss.php HTTP/1.1     1,233                0                    0
/mirror/distros/vlit 0                    1,122                0
/robots.txt HTTP/1.1 1,056                0                    0
/vendor/phpunit/phpu 1,055                0                    0
/aaa.php HTTP/1.1    954                  0                    0
/hell.php HTTP/1.1   948                  0                    0
/z.php HTTP/1.1      948                  0                    0
Matched: 465,348 / 470,163 (R: 2396; C: 8)
```