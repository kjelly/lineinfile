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
