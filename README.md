# lwn-newsletter

[![Build Status](https://travis-ci.org/haosdent/lwn-newsletter.svg?branch=master)](https://travis-ci.org/haosdent/lwn-newsletter)

A tool to subscribe Linux weekly archives. Just for fun. Down load it from [here](https://github.com/haosdent/lwn-newsletter/releases/).

## Config File Example

```
category=Kernel # Support Kernel, Security and Distributions
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