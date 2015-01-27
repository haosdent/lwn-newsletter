# lwn-newsletter

[![Build Status](https://travis-ci.org/haosdent/lwn-newsletter.svg?branch=master)](https://travis-ci.org/haosdent/lwn-newsletter)

## Config File Example

```
receiver=xxx@gmail.com
password=xxx
server=smtp.gmail.com
port=587
```

## Usage

```
[Usage]:
         ./lwn_newsletter [config file location]
```

Add to crontab

```
5 8 * * 6 your_user_name lwn_newsletter [config file location]
```

And then check your email please. :-)