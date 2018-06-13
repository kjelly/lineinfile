# lineinfile

Insert/remove line in file like ansible lineinfile.
Note: It only handle text. Don't support for changing file permission.

## Demo

sample file:

a.ini
```ini
[Default]
user=abc
password=abc
```

Insert `method=post` into a.ini
```
lineinfile --mode present a.ini "method=post"
```

Remove `password` from a.ini
It will remove a whole line
```
lineinfile --mode absent a.ini "password"
```

Change user value
```
lineinfile --mode replace a.ini "user=.*" --text="user=abcd"
```

Sample file b.ini

```
[Default]
user=abc
password=abc
data=full
power=false
url=http://www.example.com

[auth]
user=ccc
password=ccc

[data]
url=http://example.com
uri=http://www.example.com
driver=sql

```

change the value of user to zz in auth block
```
lineinfile --mode replace --afterpattern "\[auth\]" --beforepattern "\[data\]" --text "user=zz" b.ini "user.*"
```

add prefix `os_` to the key of user in the file
```
$ ./lineinfile --mode insertbefore --text "os_" a.ini "user"
```

add prefix `os_` to any key in auth block
```
lineinfile --mode insertbefore --afterpattern "\[auth\]" --beforepattern "\[data\]" --text "os_" a.ini ".+"
```
or
```
lineinfile --mode replace --afterpattern "\[auth\]" --beforepattern "\[data\]" --text 'os_$1' a.ini "^(.+)"
```


add prefix `os_` to any key
```
lineinfile --mode replace --afterpattern "\[.*\]" --text 'os_$1' a.ini "(.+)"
```
or
```
lineinfile --mode insertbefore --text 'os_' a.ini "^\w(.+)"
```
