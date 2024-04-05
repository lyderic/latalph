# Quid ?

The original Latin alphabet [had only 23 letters](https://mysite.du.edu/~etuttle/classics/latalph.htm). 

There was also no lower case letters and no punctuation.

This program takes any modern western UTF-8 encoded text and turns it into a Latin representation.

It acts as a filter.

Examples:

    $ echo Jerusalem | latalph
    IERVSALEM
    $ echo "¡Un café, por favor!" | latalph
    VN CAFE POR FAVOR

# To do

Also convert Arabic numerals to latin e.g.: 5 = V, 100 = C etc.
