# Diter

`diter` is a date iteration command.

## Installation

    $ go get github.com/ojima-h/diter

Or download the script:

    $ curl -L https://github.com/ojima-h/diter/releases/download/v1.0.0/diter-1.0.0.linux-amd64 > diter
    $ chmod +x diter

## Usage

*Basic usage:*

    $ diter 2016-01-01 2016-01-03
    2016-01-01
    2016-01-02
    2016-01-03

*Run command:*

    $ diter 2016-01-01 2016-01-03 -- echo date={}
    date=2016-01-01
    date=2016-01-02
    date=2016-01-03

*Reverse order:*

    $ diter 2016-01-03 2016-01-01 -- echo date={}
    date=2016-01-03
    date=2016-01-02
    date=2016-01-01

*Date format:*

    $ diter 2016-01-01 2016-01-03 -F %Y/%m/%d
    2016/01/01
    2016/01/02
    2016/01/03

*Date format in command:*

    $ diter 2016-01-01 2016-01-03 -- echo date=%Y/%m/%d
    date=2016/01/01
    date=2016/01/02
    date=2016/01/03

Refer to [Date#strftime](https://docs.ruby-lang.org/en/2.1.0/Date.html#method-i-strftime) for possible date formats.

*Parallel execution (with xargs):*

    $ diter 2016-01-01 2016-01-03 | xargs -I{} -P2 -t echo date={}
    echo date=2016-01-01
    echo date=2016-01-02
    date=2016-01-01
    date=2016-01-02
    echo date=2016-01-03
    date=2016-01-03

*Loop:*

    $ for lang in en ja fr; do
    >   for date in `diter 2016-01-01 2016-01-03`; do
    >      echo lang=$lang date=$date
    >   done
    > done
    lang=en date=2016-01-01
    lang=en date=2016-01-02
    lang=en date=2016-01-03
    lang=ja date=2016-01-01
    lang=ja date=2016-01-02
    lang=ja date=2016-01-03
    lang=fr date=2016-01-01
    lang=fr date=2016-01-02
    lang=fr date=2016-01-03

*Filter by the day of week:*

    $ diter 2016-01-01 2016-01-31 --wday=1
    2016-01-04
    2016-01-11
    2016-01-18
    2016-01-25

*Filter by the day of month:*

    $ diter 2016-01-01 2016-04-01 --mday=15
    2016-01-15
    2016-02-15
    2016-03-15

A negative integer is assumed to be relative to the end of month:

    $ diter 2016-01-01 2016-04-01 --mday=-1
    2016-01-31
    2016-02-29
    2016-03-31
