# QVID EST HOC (_what is this?_)

The original Latin alphabet [had only 23 letters](https://mysite.du.edu/~etuttle/classics/latalph.htm). 

There was also no lower case letters and no punctuation. The numeric system consisted of letters; zero had not yet been invented.

This program takes any modern western UTF-8 encoded text and turns it into a Latin representation.

It acts as a filter that you can use as a Unix pipe.

Examples:

    $ echo Jerusalem | latalph
    IERVSALEM
    $ echo "¡Un café, por favor!" | latalph
    VN CAFE POR FAVOR

# QVAM AEDIFICARE (_how to build?_)

Nothing fancy. [Install go](https://go.dev/) and run:

    $ go mod tidy
    $ go install

# AGENDA OPERVM (_to do_)

- Also convert Arabic numerals to latin e.g.: 5 = V, 100 = C etc.

- Write tests
